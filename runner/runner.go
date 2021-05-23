package runner

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

// Runner is used to run tasks
type Runner struct {
	taskNames              []string
	project                *config.Project
	taskStatus             map[string]*TaskStatus
	taskStatusMutex        *sync.RWMutex
	cmdInfo                map[string]*CmdInfo
	cmdInfoMutex           *sync.RWMutex
	killed                 bool
	killedMutex            *sync.RWMutex
	done                   bool
	doneMutex              *sync.RWMutex
	statusInterval         time.Duration
	startTimeMutex         *sync.RWMutex
	startTime              time.Time
	spaces                 string
	surpressWaitError      bool
	surpressWaitErrorMutex *sync.RWMutex
	logger                 output.Logger
	decoration             *output.Decoration
	autoTerminate          bool
	autoTerminateDelay     time.Duration
}

// NewRunner create new runner
func NewRunner(
	logger output.Logger, decoration *output.Decoration, project *config.Project, taskNames []string,
	statusIntervalStr string, autoTerminate bool, autoTerminateDelayStr string,
) (runner *Runner, err error) {
	if !project.IsInitialized {
		return &Runner{}, fmt.Errorf("cannot create runner because project was not initialized")
	}
	if err = project.ValidateByTaskNames(taskNames); err != nil {
		return &Runner{}, err
	}
	if project.GetAutoTerminate(taskNames) {
		autoTerminate = true
	}
	statusInterval, err := time.ParseDuration(statusIntervalStr)
	if err != nil {
		return &Runner{}, fmt.Errorf("cannot parse statusInterval '%s': %s", statusIntervalStr, err)
	}
	autoTerminateDelayInterval, err := time.ParseDuration(autoTerminateDelayStr)
	if err != nil {
		return &Runner{}, fmt.Errorf("cannot parse autoTerminateDelay '%s': %s", autoTerminateDelayStr, err)
	}
	return &Runner{
		taskNames:              taskNames,
		project:                project,
		taskStatus:             map[string]*TaskStatus{},
		taskStatusMutex:        &sync.RWMutex{},
		cmdInfo:                map[string]*CmdInfo{},
		cmdInfoMutex:           &sync.RWMutex{},
		killed:                 false,
		killedMutex:            &sync.RWMutex{},
		done:                   false,
		doneMutex:              &sync.RWMutex{},
		statusInterval:         statusInterval,
		startTimeMutex:         &sync.RWMutex{},
		spaces:                 fmt.Sprintf("%s %s", decoration.Empty, decoration.Empty),
		surpressWaitError:      false,
		surpressWaitErrorMutex: &sync.RWMutex{},
		logger:                 logger,
		decoration:             decoration,
		autoTerminate:          autoTerminate,
		autoTerminateDelay:     autoTerminateDelayInterval,
	}, nil
}

// Run Tasks
func (r *Runner) Run() (err error) {
	r.startTime = time.Now()
	r.showStatus()
	ch := make(chan error)
	go r.handleTerminationSignal(ch)
	go r.run(ch)
	go r.waitAnyProcessError(ch)
	go r.showStatusByInterval()
	err = <-ch
	if err == nil && r.getKilledSignal() {
		r.showStatus()
		return fmt.Errorf("terminated")
	}
	if !r.getKilledSignal() {
		r.Terminate()
	}
	r.showStatus()
	return err
}

// Terminate all processes
func (r *Runner) Terminate() {
	r.logger.DPrintfError("Terminating\n")
	r.setKilledSignal()
	// kill unfinished commands
	r.cmdInfoMutex.Lock()
	killedCh := map[string]chan error{}
	for label, cmdInfo := range r.cmdInfo {
		killedCh[label] = make(chan error)
		cmd := cmdInfo.Cmd
		r.logger.DPrintfKill("Kill %s (PID=%d)\n", label, cmd.Process.Pid)
		go r.killByPid(-cmd.Process.Pid, killedCh[label])
	}
	for label := range r.cmdInfo {
		if err := <-killedCh[label]; err != nil {
			r.logger.Println(r.spaces, err)
		}
		delete(r.cmdInfo, label)
	}
	r.cmdInfoMutex.Unlock()
}

func (r *Runner) showStatusByInterval() {
	for {
		r.sleep(r.statusInterval)
		if r.getKilledSignal() {
			return
		}
		r.showStatus()
	}
}

func (r *Runner) waitAnyProcessError(ch chan error) {
	seen := map[string]bool{}
	for {
		r.sleep(10 * time.Millisecond)
		if r.getKilledSignal() {
			ch <- fmt.Errorf("terminated")
			return
		}
		r.cmdInfoMutex.Lock()
		for label, cmdInfo := range r.cmdInfo {
			if _, exist := seen[label]; exist || !cmdInfo.IsProcess {
				continue
			}
			seen[label] = true
			currentLabel := label
			currentCmd := cmdInfo.Cmd
			currentTaskName := cmdInfo.TaskName
			go func() {
				err := currentCmd.Wait()
				if err != nil {
					if !r.getKilledSignal() && !r.getSurpressWaitErrorSignal() {
						r.logger.DPrintfError("%s exited:\n%s\n\n%s\n", currentLabel, r.sprintfCmdArgs(currentCmd), err)
					} else {
						r.logger.DPrintfError("%s exited: %s\n", currentLabel, err)
					}
					r.unregisterCmd(currentLabel)
					r.setSurpressWaitErrorSignal()
					ch <- err
					return
				}
				if !r.isTaskFinished(currentTaskName) {
					if !r.getKilledSignal() && !r.getSurpressWaitErrorSignal() {
						r.logger.DPrintfError("%s stopped before ready:\n%s\n", currentLabel, r.sprintfCmdArgs(currentCmd))
					} else {
						r.logger.DPrintfError("%s stopped before ready\n", currentLabel)
					}
					r.unregisterCmd(currentLabel)
					r.setSurpressWaitErrorSignal()
					ch <- fmt.Errorf("%s stopped before ready", currentLabel)
					return
				}
				r.unregisterCmd(currentLabel)
				r.logger.DPrintfError("%s exited without any error message\n", currentLabel)
			}()
		}
		r.cmdInfoMutex.Unlock()
	}
}

func (r *Runner) handleTerminationSignal(ch chan error) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	sig := <-signalChannel
	errorMsg := ""
	switch sig {
	case os.Interrupt:
		errorMsg = "Receiving SIGINT"
		break
	case syscall.SIGTERM:
		errorMsg = "Receiving SIGTERM"
		break
	default:
		errorMsg = "Receiving termination signal"
	}
	r.logger.Println()
	r.logger.DPrintfError("%s\n", errorMsg)
	ch <- fmt.Errorf(errorMsg)
}

func (r *Runner) setSurpressWaitErrorSignal() {
	r.surpressWaitErrorMutex.Lock()
	r.surpressWaitError = true
	r.surpressWaitErrorMutex.Unlock()
}

func (r *Runner) getSurpressWaitErrorSignal() (isSurpressWaitError bool) {
	r.surpressWaitErrorMutex.RLock()
	isSurpressWaitError = r.surpressWaitError
	r.surpressWaitErrorMutex.RUnlock()
	return isSurpressWaitError
}

func (r *Runner) setDoneSignal() {
	r.doneMutex.Lock()
	r.done = true
	r.doneMutex.Unlock()
}

func (r *Runner) getDoneSignal() (isDone bool) {
	r.doneMutex.RLock()
	isDone = r.done
	r.doneMutex.RUnlock()
	return isDone
}

func (r *Runner) setKilledSignal() {
	r.killedMutex.Lock()
	r.killed = true
	r.killedMutex.Unlock()
}

func (r *Runner) getKilledSignal() (isKilled bool) {
	r.killedMutex.RLock()
	isKilled = r.killed
	r.killedMutex.RUnlock()
	return isKilled
}

func (r *Runner) run(ch chan error) {
	if err := r.runTaskByNames(r.taskNames); err != nil {
		ch <- err
		return
	}
	r.setDoneSignal()
	r.showStatus()
	d := r.decoration
	r.logger.DPrintfSuccess("%s\n", strings.Repeat(d.Success, 11))
	r.logger.DPrintfSuccess("%s%sJob Complete!!! %s%s\n", d.Bold, d.Green, strings.Repeat(d.Success, 3), d.Normal)
	if r.autoTerminate {
		r.sleep(100 * time.Millisecond)
		r.sleep(r.autoTerminateDelay)
		ch <- nil
		return
	}
	// wait until no cmd left
	for {
		r.sleep(10 * time.Millisecond)
		if r.getKilledSignal() {
			ch <- fmt.Errorf("terminated")
			return
		}
		processExist := false
		r.cmdInfoMutex.RLock()
		for range r.cmdInfo {
			processExist = true
			break
		}
		r.cmdInfoMutex.RUnlock()
		if !processExist {
			ch <- nil
			return
		}
	}
}

func (r *Runner) runTaskByNames(taskNames []string) (err error) {
	tasks := []*config.Task{}
	for _, taskName := range taskNames {
		task := r.project.Tasks[taskName]
		tasks = append(tasks, task)
	}
	ch := make(chan error)
	for _, task := range tasks {
		go r.runTask(task, ch)
	}
	for index := 0; index < len(tasks); index++ {
		err = <-ch
		if err != nil {
			return err
		}
	}
	return err
}

func (r *Runner) runTask(task *config.Task, ch chan error) {
	if !r.registerTask(task.GetName()) {
		ch <- r.waitTaskFinished(task.GetName())
		return
	}
	if err := r.runTaskByNames(task.GetDependencies()); err != nil {
		ch <- err
		return
	}
	startLogDone := make(chan error)
	startCmd, startExist, startErr := task.GetStartCmd(startLogDone)
	if !startExist {
		r.logger.DPrintfSuccess("Reach %s '%s' wrapper\n", r.decoration.Icon(task.Icon), task.GetName())
		r.finishTask(task.GetName(), nil)
		ch <- nil
		return
	}
	if startErr != nil {
		ch <- startErr
		return
	}
	checkLogDone := make(chan error)
	checkCmd, checkExist, checkErr := task.GetCheckCmd(checkLogDone)
	if !checkExist {
		err := r.runCommandTask(task, startCmd, startLogDone)
		r.finishTask(task.GetName(), err)
		ch <- err
		return
	}
	if checkErr != nil {
		ch <- checkErr
		return
	}
	err := r.runServiceTask(task, startCmd, checkCmd, checkLogDone)
	r.finishTask(task.GetName(), err)
	ch <- err
}

func (r *Runner) runCommandTask(task *config.Task, startCmd *exec.Cmd, startLogDone chan error) (err error) {
	r.logger.DPrintfStarted("Run %s '%s' command on %s\n", r.decoration.Icon(task.Icon), task.GetName(), startCmd.Dir)
	startStdinPipe, err := startCmd.StdinPipe()
	if err == nil {
		err = startCmd.Start()
	}
	if err != nil {
		if !r.getKilledSignal() && !r.getSurpressWaitErrorSignal() {
			r.logger.DPrintfError("Error running command %s '%s':\n%s\n%s\n", r.decoration.Icon(task.Icon), task.GetName(), r.sprintfCmdArgs(startCmd), err)
		} else {
			r.logger.DPrintfError("Error running command %s '%s': %s\n", r.decoration.Icon(task.Icon), task.GetName(), err)
		}
		return err
	}
	startCmdLabel := fmt.Sprintf("%s '%s' command", r.decoration.Icon(task.Icon), task.GetName())
	r.registerCommandCmd(startCmdLabel, task.GetName(), startCmd, startStdinPipe)
	err = r.waitTaskCmd(task, startCmd, startCmdLabel, startLogDone)
	r.unregisterCmd(startCmdLabel)
	return err
}

func (r *Runner) runServiceTask(task *config.Task, startCmd *exec.Cmd, checkCmd *exec.Cmd, checkLogDone chan error) (err error) {
	if err = r.runStartServiceTask(task, startCmd); err != nil {
		return err
	}
	err = r.runCheckServiceTask(task, checkCmd, checkLogDone)
	return err
}

func (r *Runner) runStartServiceTask(task *config.Task, startCmd *exec.Cmd) (err error) {
	r.logger.DPrintfStarted("Run %s '%s' service on %s\n", r.decoration.Icon(task.Icon), task.GetName(), startCmd.Dir)
	startStdinPipe, err := startCmd.StdinPipe()
	if err == nil {
		err = startCmd.Start()
	}
	if err != nil {
		if !r.getKilledSignal() && !r.getSurpressWaitErrorSignal() {
			r.logger.DPrintfError("Error starting service %s '%s':\n%s\n%s\n", r.decoration.Icon(task.Icon), task.GetName(), r.sprintfCmdArgs(startCmd), err)
		} else {
			r.logger.DPrintfError("Error starting service %s '%s': %s\n", r.decoration.Icon(task.Icon), task.GetName(), err)
		}
		return err
	}
	startCmdLabel := fmt.Sprintf("%s '%s' service", r.decoration.Icon(task.Icon), task.GetName())
	r.registerProcessCmd(startCmdLabel, task.GetName(), startCmd, startStdinPipe)
	return err
}

func (r *Runner) runCheckServiceTask(task *config.Task, checkCmd *exec.Cmd, checkLogDone chan error) (err error) {
	r.logger.DPrintfStarted("Check %s '%s' readiness on %s\n", r.decoration.Icon(task.Icon), task.GetName(), checkCmd.Dir)
	checkStdinPipe, err := checkCmd.StdinPipe()
	if err == nil {
		err = checkCmd.Start()
	}
	if err != nil {
		if !r.getKilledSignal() && !r.getSurpressWaitErrorSignal() {
			r.logger.DPrintfError("Error checking service %s '%s' readiness:\n%s\n\n%s", r.decoration.Icon(task.Icon), task.GetName(), r.sprintfCmdArgs(checkCmd), err)
		} else {
			r.logger.DPrintfError("Error checking service %s '%s' readiness: %s\n", r.decoration.Icon(task.Icon), task.GetName(), err)
		}
		return err
	}
	checkCmdLabel := fmt.Sprintf("%s '%s' readiness check", r.decoration.Icon(task.Icon), task.GetName())
	r.registerCommandCmd(checkCmdLabel, task.GetName(), checkCmd, checkStdinPipe)
	err = r.waitTaskCmd(task, checkCmd, checkCmdLabel, checkLogDone)
	r.unregisterCmd(checkCmdLabel)
	return err
}

func (r *Runner) waitTaskCmd(task *config.Task, cmd *exec.Cmd, cmdLabel string, logDone chan error) (err error) {
	executed := false
	ch := make(chan error)
	go func() {
		waitErr := cmd.Wait()
		if waitErr != nil {
			if !r.getKilledSignal() && !r.getSurpressWaitErrorSignal() {
				r.logger.DPrintfError("Error running %s:\n%s\n%s\n", cmdLabel, r.sprintfCmdArgs(cmd), waitErr)
			} else {
				r.logger.DPrintfError("Error running %s: %s\n", cmdLabel, waitErr)
			}
			ch <- waitErr
			return
		}
		executed = true
		r.logger.DPrintfSuccess("Successfully running %s\n", cmdLabel)
		if logErr := <-logDone; logErr != nil {
			r.logger.DPrintfError("Error logging %s: %s\n", cmdLabel, logErr)
		} else {
			r.logger.DPrintfSuccess("Successfully logging %s\n", cmdLabel)
		}
		ch <- nil
		return
	}()
	go func() {
		r.sleep(task.GetTimeoutDuration())
		if executed {
			return
		}
		timeoutMessage := fmt.Sprintf("Getting timeout while running %s", cmdLabel)
		r.logger.DPrintfError("%s\n", timeoutMessage)
		ch <- fmt.Errorf(timeoutMessage)
		return
	}()
	err = <-ch
	return err
}

func (r *Runner) registerCommandCmd(label, taskName string, cmd *exec.Cmd, stdinPipe io.WriteCloser) {
	r.registerCmd(label, taskName, cmd, stdinPipe, false)
}

func (r *Runner) registerProcessCmd(label, taskName string, cmd *exec.Cmd, stdinPipe io.WriteCloser) {
	r.registerCmd(label, taskName, cmd, stdinPipe, true)
}

func (r *Runner) registerCmd(label, taskName string, cmd *exec.Cmd, stdinPipe io.WriteCloser, isProcess bool) {
	r.cmdInfoMutex.Lock()
	r.cmdInfo[label] = &CmdInfo{
		Cmd:       cmd,
		IsProcess: isProcess,
		StdInPipe: stdinPipe,
		TaskName:  taskName,
	}
	r.cmdInfoMutex.Unlock()
}

func (r *Runner) unregisterCmd(label string) {
	r.cmdInfoMutex.Lock()
	delete(r.cmdInfo, label)
	r.cmdInfoMutex.Unlock()
}

func (r *Runner) registerTask(taskName string) (success bool) {
	r.taskStatusMutex.Lock()
	_, isStarted := r.taskStatus[taskName]
	if isStarted {
		success = false
	} else {
		r.taskStatus[taskName] = NewTaskStatus()
		success = true
	}
	r.taskStatusMutex.Unlock()
	return success
}

func (r *Runner) finishTask(taskName string, err error) {
	r.taskStatusMutex.Lock()
	r.taskStatus[taskName].Finish(err)
	r.taskStatusMutex.Unlock()
}

func (r *Runner) isTaskFinished(taskName string) (isFinished bool) {
	r.taskStatusMutex.RLock()
	isFinished = r.taskStatus[taskName].Finished
	r.taskStatusMutex.RUnlock()
	return isFinished
}

func (r *Runner) isTaskError(taskName string) (err error) {
	r.taskStatusMutex.RLock()
	err = r.taskStatus[taskName].Error
	r.taskStatusMutex.RUnlock()
	return err
}

func (r *Runner) waitTaskFinished(taskName string) (err error) {
	for {
		r.sleep(10 * time.Millisecond)
		if r.isTaskFinished(taskName) {
			return r.isTaskError(taskName)
		}
		if r.getKilledSignal() {
			return fmt.Errorf("terminated")
		}
	}
}

func (r *Runner) sprintfCmdArgs(cmd *exec.Cmd) (output string) {
	d := r.decoration
	formattedArgs := []string{}
	for _, arg := range cmd.Args {
		rows := strings.Split(arg, "\n")
		for index, row := range rows {
			prefix := "  "
			if index == 0 {
				prefix = "* "
			}
			if len(rows) > 1 {
				prefix += fmt.Sprintf("%s%4d |%s ", d.Yellow, index+1, d.NoColor)
			}
			row = fmt.Sprintf("%s%s %s%s%s%s", r.spaces, d.Empty, d.Faint, prefix, row, d.Normal)
			rows[index] = row
		}
		formattedArg := strings.Join(rows, "\n")
		formattedArgs = append(formattedArgs, formattedArg)
	}
	output = strings.Join(formattedArgs, "\n")
	return output
}

func (r *Runner) sleep(duration time.Duration) {
	done := make(chan bool)
	ticker := time.NewTimer(duration)
	go func() {
		<-ticker.C
		ticker.Stop()
		done <- true
	}()
	<-done
}

func (r *Runner) getProcessRow(label string, cmd *exec.Cmd) string {
	d := r.decoration
	return fmt.Sprintf("%s* (PID=%d) %s%s", d.Faint, cmd.Process.Pid, label, d.Normal)
}

func (r *Runner) showStatus() {
	d := r.decoration
	descriptionPrefix := r.spaces + d.Empty + d.Empty
	processPrefix := r.spaces + r.spaces + " "
	processRows := []string{}
	r.cmdInfoMutex.Lock()
	for label, cmdInfo := range r.cmdInfo {
		cmd := cmdInfo.Cmd
		processRow := r.getProcessRow(label, cmd)
		processRows = append(processRows, processRow)
	}
	r.cmdInfoMutex.Unlock()
	statusCaption := r.getStatusCaption()
	r.startTimeMutex.RLock()
	elapsedTime := time.Since(r.startTime)
	elapsedTimeCaption := fmt.Sprintf("%s%sElapsed Time: %s%s\n", descriptionPrefix, d.Faint, elapsedTime, d.Normal)
	r.startTimeMutex.RUnlock()
	currentTime := time.Now()
	currentTimeString := currentTime.Format("15:04:05")
	currentTimeCaption := fmt.Sprintf("%s%sCurrent Time: %s%s\n", descriptionPrefix, d.Faint, currentTimeString, d.Normal)
	activeProcessLabel := ""
	processCaption := ""
	if len(processRows) > 0 {
		activeProcessLabel = fmt.Sprintf("%s%sActive Process:%s\n", descriptionPrefix, d.Faint, d.Normal)
		processCaption = processPrefix + strings.Join(processRows, "\n"+processPrefix) + "\n"
	}
	r.logger.DPrintfInspect("%s%s%s%s%s", statusCaption, elapsedTimeCaption, currentTimeCaption, activeProcessLabel, processCaption)
}

func (r *Runner) getStatusCaption() (statusCaption string) {
	d := r.decoration
	if killed := r.getKilledSignal(); killed {
		return fmt.Sprintf("%sJob Ended...%s\n", d.Bold, d.Normal)
	}
	if done := r.getDoneSignal(); done {
		return fmt.Sprintf("%s%sJob Running...%s\n", d.Bold, d.Green, d.Normal)
	}
	return fmt.Sprintf("%sJob Starting...%s\n", d.Bold, d.Normal)
}

func (r *Runner) killByPid(pid int, ch chan error) {
	r.sleep(100 * time.Millisecond)
	err := syscall.Kill(pid, syscall.SIGINT)
	r.sleep(100 * time.Millisecond)
	syscall.Kill(pid, syscall.SIGTERM)
	r.sleep(100 * time.Millisecond)
	syscall.Kill(pid, syscall.SIGKILL)
	ch <- err
}
