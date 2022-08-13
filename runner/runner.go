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

	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

// Runner is used to run tasks
type Runner struct {
	taskNames              []string
	project                *dsl.Project
	taskStatus             map[string]*TaskStatus
	taskStatusMutex        *sync.RWMutex
	cmdInfo                map[string]*CmdInfo
	cmdInfoMutex           *sync.RWMutex
	killed                 bool
	killedMutex            *sync.RWMutex
	done                   bool
	doneMutex              *sync.RWMutex
	statusTimeInterval     time.Duration
	statusLineInterval     int
	startTimeMutex         *sync.RWMutex
	startTime              time.Time
	spaces                 string
	surpressWaitError      bool
	surpressWaitErrorMutex *sync.RWMutex
	logger                 output.Logger
	recordLogger           output.RecordLogger
	decoration             *output.Decoration
	autoTerminate          bool
	autoTerminateDelay     time.Duration
}

// NewRunner create new runner
func NewRunner(logger output.Logger, recordLogger output.RecordLogger, project *dsl.Project, taskNames []string, statusTImeIntervalStr string, statusLineInterval int, autoTerminate bool, autoTerminateDelayStr string) (runner *Runner, err error) {
	if !project.IsInitialized {
		return &Runner{}, fmt.Errorf("cannot create runner because project was not initialized")
	}
	if err = project.ValidateByTaskNames(taskNames); err != nil {
		return &Runner{}, err
	}
	if project.GetAutoTerminate(taskNames) {
		autoTerminate = true
	}
	statusTImeInterval, err := time.ParseDuration(statusTImeIntervalStr)
	if err != nil {
		return &Runner{}, fmt.Errorf("cannot parse statusInterval '%s': %s", statusTImeIntervalStr, err)
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
		statusTimeInterval:     statusTImeInterval,
		statusLineInterval:     statusLineInterval,
		startTimeMutex:         &sync.RWMutex{},
		spaces:                 fmt.Sprintf("%s %s", project.Decoration.EmptyIcon, project.Decoration.EmptyIcon),
		surpressWaitError:      false,
		surpressWaitErrorMutex: &sync.RWMutex{},
		logger:                 logger,
		recordLogger:           recordLogger,
		decoration:             project.Decoration,
		autoTerminate:          autoTerminate,
		autoTerminateDelay:     autoTerminateDelayInterval,
	}, nil
}

// Run Tasks
func (r *Runner) Run() (err error) {
	r.startTime = time.Now()
	r.showStatus()
	go r.logStderr()
	go r.logStdout()
	go r.logStderrRow()
	go r.logStdoutRow()
	ch := make(chan error)
	go r.handleTerminationSignal(ch)
	go r.run(ch)
	go r.waitLongRunningCmd(ch)
	go r.showStatusByInterval()
	err = <-ch
	r.waitOutputWg(50*time.Millisecond, 2)
	if err == nil && r.getKilledSignal() {
		r.showStatus()
		return fmt.Errorf("Terminated")
	}
	if !r.getKilledSignal() {
		r.waitOutputWg(50*time.Millisecond, 4)
		r.Terminate()
	}
	r.showStatus()
	return err
}

func (r *Runner) waitOutputWg(sleepDuration time.Duration, maxWaitAttempt int) {
	for waitAttempt := 0; waitAttempt < maxWaitAttempt; waitAttempt++ {
		r.sleep(sleepDuration)
		r.project.OutputWg.Wait()
	}
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

func (r *Runner) logStdout() {
	lineCounter := 0
	for {
		content := <-r.project.StdoutChan
		r.logger.DPrintf(content)
		r.project.OutputWgMutex.Lock()
		r.project.OutputWg.Done()
		r.project.OutputWgMutex.Unlock()
		if r.statusLineInterval < 1 {
			continue
		}
		lineCounter++
		if lineCounter >= r.statusLineInterval {
			lineCounter = 0
			r.showStatus()
		}
	}
}

func (r *Runner) logStderr() {
	for {
		content := <-r.project.StderrChan
		r.logger.DPrintfError(content)
		r.project.ProcessOutputWg()
	}
}

func (r *Runner) logStdoutRow() {
	for {
		content := <-r.project.StdoutRecordChan
		r.recordLogger.Log(content...)
		r.project.ProcessOutputWg()
	}
}

func (r *Runner) logStderrRow() {
	for {
		content := <-r.project.StderrRecordChan
		r.recordLogger.Log(content...)
		r.project.ProcessOutputWg()
	}
}

func (r *Runner) showStatusByInterval() {
	for {
		r.sleep(r.statusTimeInterval)
		if r.getKilledSignal() {
			return
		}
		r.showStatus()
	}
}

func (r *Runner) waitLongRunningCmd(ch chan error) {
	seen := map[string]bool{}
	for {
		r.sleep(50 * time.Millisecond)
		if r.getKilledSignal() {
			ch <- fmt.Errorf("Terminated")
			return
		}
		r.cmdInfoMutex.Lock()
		for label, cmdInfo := range r.cmdInfo {
			if _, exist := seen[label]; exist || !cmdInfo.IsProcess {
				continue
			}
			seen[label] = true
			currentLabel := label
			currentTask := cmdInfo.Task
			currentTaskName := currentTask.GetName()
			currentCmdMaker := cmdInfo.CmdMaker
			currentMaxRetry := cmdInfo.MaxRetry
			currentRetryDelayDuration := cmdInfo.RetryDelayDuration
			isInfiniteRetry := currentMaxRetry < 1 // 0 or less means infinite
			currentCmd := cmdInfo.Cmd
			go func(currentCmd *exec.Cmd) {
				err := currentCmd.Wait()
				if err != nil {
					r.logger.DPrintfError("Error running %s (Attempt: %d/%d)\n", currentLabel, 1, currentMaxRetry)
					r.unregisterCmd(currentLabel)
					for attempt := 2; isInfiniteRetry || attempt <= currentMaxRetry; attempt++ {
						isLastAttempt := !isInfiniteRetry && attempt == currentMaxRetry
						currentCmd, err = currentCmdMaker()
						if err != nil {
							ch <- err
							return
						}
						stdinPipe, err := currentCmd.StdinPipe()
						if err != nil {
							ch <- err
							return
						}
						r.registerCmd(currentLabel, currentTask, currentCmdMaker, currentCmd, stdinPipe, currentMaxRetry, currentRetryDelayDuration, true)
						r.logger.DPrintfStarted("Running %s on %s (Attempt: %d/%d)\n", currentLabel, currentCmd.Dir, attempt, currentMaxRetry)
						err = currentCmd.Start()
						if err != nil {
							r.logger.DPrintfError("Failed running %s (Attempt: %d/%d)\n", currentLabel, attempt, currentMaxRetry)
							r.unregisterCmd(currentLabel)
							continue
						}
						err = currentCmd.Wait()
						// no error, quit loop
						if err == nil {
							r.logger.DPrintfSuccess("Successfully running %s (Attempt: %d/%d)\n", currentLabel, attempt, currentMaxRetry)
							break
						}
						// any error
						if !r.getKilledSignal() && !r.getSurpressWaitErrorSignal() {
							r.logger.DPrintfError("Error running %s (Attempt: %d/%d):\n%s\n%s\n", currentLabel, attempt, currentMaxRetry, r.sprintfCmdArgs(currentCmd), err)
						} else {
							r.logger.DPrintfError("Error running %s (Attempt: %d/%d): %s\n", currentLabel, attempt, currentMaxRetry, err)
						}
						// going to retry
						if !isLastAttempt {
							r.unregisterCmd(currentLabel)
							r.sleep(currentRetryDelayDuration)
						}
					}
				}
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
				ch <- fmt.Errorf("%s exited without any error message", currentLabel)
			}(currentCmd)
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
	case syscall.SIGTERM:
		errorMsg = "Receiving SIGTERM"
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
	r.logger.DPrintfSuccess("%s\n", strings.Repeat(d.SuccessIcon, 11))
	r.logger.DPrintfSuccess("%s%sJob Complete!!! %s%s\n", d.Bold, d.Green, strings.Repeat(d.SuccessIcon, 3), d.Normal)
	if r.autoTerminate {
		r.sleep(r.autoTerminateDelay)
		ch <- nil
		return
	}
	// wait until no cmd left
	for {
		r.sleep(100 * time.Millisecond)
		if r.getKilledSignal() {
			ch <- fmt.Errorf("Terminated")
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
	tasks := []*dsl.Task{}
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

func (r *Runner) runTask(task *dsl.Task, ch chan error) {
	if !r.registerTask(task.GetName()) {
		ch <- r.waitTaskFinished(task.GetName())
		return
	}
	if err := r.runTaskByNames(task.GetDependencies()); err != nil {
		ch <- err
		return
	}
	if !task.IsHavingStartCmd() {
		// wrapper task
		r.logger.DPrintfSuccess("Reach %s '%s' wrapper\n", r.decoration.Icon(task.Icon), task.GetName())
		r.markTaskFinished(task.GetName(), nil)
		ch <- nil
		return
	}
	if !task.IsHavingCheckCmd() {
		// simple task
		err := r.runSimpleTask(task)
		r.markTaskFinished(task.GetName(), err)
		ch <- err
		return
	}
	// long running task
	err := r.runLongRunningTask(task)
	r.markTaskFinished(task.GetName(), err)
	ch <- err
}

func (r *Runner) runSimpleTask(task *dsl.Task) (err error) {
	maxStartRetry := task.GetMaxStartRetry()
	startRetryDelayDuration := task.GetStartRetryDelayDuration()
	startCmdLabel := fmt.Sprintf("%s '%s' runner", r.decoration.Icon(task.Icon), task.GetName())
	err = r.waitSimpleCmd(startCmdLabel, task, task.GetStartCmd, maxStartRetry, startRetryDelayDuration)
	r.unregisterCmd(startCmdLabel)
	return err
}

func (r *Runner) runLongRunningTask(task *dsl.Task) (err error) {
	if err = r.startLongRunningTask(task); err != nil {
		return err
	}
	err = r.checkLongRunningTask(task)
	return err
}

func (r *Runner) startLongRunningTask(task *dsl.Task) (err error) {
	maxStartRetry := task.GetMaxStartRetry()
	startRetryDelayDuration := task.GetStartRetryDelayDuration()
	startCmdLabel := fmt.Sprintf("%s '%s' starter", r.decoration.Icon(task.Icon), task.GetName())
	startCmdMaker := task.GetStartCmd
	startCmd, err := startCmdMaker()
	if err != nil {
		return err
	}
	startStdinPipe, err := startCmd.StdinPipe()
	if err != nil {
		return err
	}
	r.logger.DPrintfStarted("Running %s on %s (Attempt: %d)\n", startCmdLabel, startCmd.Dir, 1)
	r.registerCmd(startCmdLabel, task, startCmdMaker, startCmd, startStdinPipe, maxStartRetry, startRetryDelayDuration, true)
	startCmd.Start()
	return nil
}

func (r *Runner) checkLongRunningTask(task *dsl.Task) (err error) {
	maxCheckRetry := task.GetMaxCheckRetry()
	checkRetryDelayDuration := task.GetCheckRetryDelayDuration()
	checkCmdLabel := fmt.Sprintf("%s '%s' readiness checker", r.decoration.Icon(task.Icon), task.GetName())
	err = r.waitSimpleCmd(checkCmdLabel, task, task.GetCheckCmd, maxCheckRetry, checkRetryDelayDuration)
	r.unregisterCmd(checkCmdLabel)
	return err
}

func (r *Runner) waitSimpleCmd(cmdLabel string, task *dsl.Task, cmdMaker func() (*exec.Cmd, error), maxRetry int, retryDelayDuration time.Duration) (err error) {
	isInfiniteRetry := maxRetry < 1 // 0 or less means infinite
	executed := false
	ch := make(chan error)
	// run task
	go func() {
		var cmdErr error
		var cmd *exec.Cmd
		var stdinPipe io.WriteCloser
		for attempt := 1; isInfiniteRetry || attempt <= maxRetry; attempt++ {
			isLastAttempt := !isInfiniteRetry && attempt == maxRetry
			cmd, cmdErr = cmdMaker()
			if cmdErr != nil {
				ch <- cmdErr
				return
			}
			stdinPipe, cmdErr = cmd.StdinPipe()
			if cmdErr != nil {
				ch <- cmdErr
				return
			}
			r.logger.DPrintfStarted("Running %s on %s (Attempt: %d/%d)\n", cmdLabel, cmd.Dir, attempt, maxRetry)
			r.registerCmd(cmdLabel, task, cmdMaker, cmd, stdinPipe, maxRetry, retryDelayDuration, false)
			cmdErr = cmd.Start()
			if cmdErr != nil {
				r.logger.DPrintfError("Failed running %s (Attempt: %d/%d)\n", cmdLabel, attempt, maxRetry)
				r.unregisterCmd(cmdLabel)
				continue
			}
			cmdErr = cmd.Wait()
			// no error, quit this function
			if cmdErr == nil {
				r.logger.DPrintfSuccess("Successfully running %s (Attempt: %d/%d)\n", cmdLabel, attempt, maxRetry)
				ch <- nil
				return
			}
			// any error
			if !r.getKilledSignal() && !r.getSurpressWaitErrorSignal() {
				r.logger.DPrintfError("Error running %s (Attempt: %d/%d):\n%s\n%s\n", cmdLabel, attempt, maxRetry, r.sprintfCmdArgs(cmd), cmdErr)
			} else {
				r.logger.DPrintfError("Error running %s (Attempt: %d/%d): %s\n", cmdLabel, attempt, maxRetry, cmdErr)
			}
			// going to retry
			if !isLastAttempt {
				r.unregisterCmd(cmdLabel)
				r.sleep(retryDelayDuration)
			}
		}
		ch <- cmdErr
	}()
	// checking timeout
	go func() {
		r.sleep(task.GetTimeoutDuration())
		if executed {
			return
		}
		timeoutMessage := fmt.Sprintf("Getting timeout while running %s", cmdLabel)
		r.logger.DPrintfError("%s\n", timeoutMessage)
		ch <- fmt.Errorf(timeoutMessage)
	}()
	err = <-ch
	return err
}

func (r *Runner) registerCmd(label string, task *dsl.Task, cmdMaker func() (*exec.Cmd, error), cmd *exec.Cmd, stdinPipe io.WriteCloser, maxRetry int, retryDelayDuration time.Duration, isProcess bool) {
	r.cmdInfoMutex.Lock()
	r.cmdInfo[label] = &CmdInfo{
		Cmd:                cmd,
		IsProcess:          isProcess,
		StdInPipe:          stdinPipe,
		Task:               task,
		CmdMaker:           cmdMaker,
		MaxRetry:           maxRetry,
		RetryDelayDuration: retryDelayDuration,
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

func (r *Runner) markTaskFinished(taskName string, err error) {
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

func (r *Runner) getTaskError(taskName string) (err error) {
	r.taskStatusMutex.RLock()
	err = r.taskStatus[taskName].Error
	r.taskStatusMutex.RUnlock()
	return err
}

func (r *Runner) waitTaskFinished(taskName string) (err error) {
	for {
		r.sleep(100 * time.Millisecond)
		if r.isTaskFinished(taskName) {
			return r.getTaskError(taskName)
		}
		if r.getKilledSignal() {
			return fmt.Errorf("Terminated")
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
			row = strings.ReplaceAll(row, "\x1b", "\\x1b")
			row = fmt.Sprintf("%s%s %s%s%s%s", r.spaces, d.EmptyIcon, d.Faint, prefix, row, d.Normal)
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
	descriptionPrefix := r.spaces + d.EmptyIcon + d.EmptyIcon
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
	var err error
	if _, findErr := os.FindProcess(int(pid)); findErr == nil {
		r.sleep(300 * time.Millisecond)
		if _, findErr := os.FindProcess(int(pid)); findErr == nil {
			err = syscall.Kill(pid, syscall.SIGINT)
		}
	}
	if _, findErr := os.FindProcess(int(pid)); findErr == nil {
		r.sleep(300 * time.Millisecond)
		if _, findErr := os.FindProcess(int(pid)); findErr == nil {
			syscall.Kill(pid, syscall.SIGTERM)
		}
	}
	if _, findErr := os.FindProcess(int(pid)); findErr == nil {
		r.sleep(300 * time.Millisecond)
		if _, findErr := os.FindProcess(int(pid)); findErr == nil {
			syscall.Kill(pid, syscall.SIGKILL)
		}
	}
	ch <- err
}
