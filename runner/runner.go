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
	"github.com/state-alchemists/zaruba/logger"
)

// TaskStatus represent task status
type TaskStatus struct {
	Finished bool
	Error    error
}

// NewTaskStatus create new task status
func NewTaskStatus() (ts *TaskStatus) {
	return &TaskStatus{
		Finished: false,
		Error:    nil,
	}
}

// Finish task status
func (ts *TaskStatus) Finish(err error) {
	ts.Finished = true
	ts.Error = err
}

// CmdInfo represent information of Cmd
type CmdInfo struct {
	Cmd       *exec.Cmd
	IsProcess bool
	StdInPipe io.WriteCloser
	TaskName  string
}

// Runner is used to run tasks
type Runner struct {
	TaskNames       []string
	Project         *config.Project
	TaskStatus      map[string]*TaskStatus
	TaskStatusMutex *sync.RWMutex
	CmdInfo         map[string]*CmdInfo
	CmdInfoMutex    *sync.RWMutex
	Killed          bool
	KilledMutex     *sync.RWMutex
	Done            bool
	DoneMutex       *sync.RWMutex
	StatusInterval  time.Duration
	StartTimeMutex  *sync.RWMutex
	StartTime       time.Time
	Spaces          string
}

// NewRunner create new runner
func NewRunner(project *config.Project, taskNames []string, statusInterval time.Duration) (runner *Runner, err error) {
	if !project.IsInitialized {
		return &Runner{}, fmt.Errorf("Cannot create runner because project was not initialize")
	}
	return &Runner{
		TaskNames:       taskNames,
		Project:         project,
		TaskStatus:      map[string]*TaskStatus{},
		TaskStatusMutex: &sync.RWMutex{},
		CmdInfo:         map[string]*CmdInfo{},
		CmdInfoMutex:    &sync.RWMutex{},
		Killed:          false,
		KilledMutex:     &sync.RWMutex{},
		Done:            false,
		DoneMutex:       &sync.RWMutex{},
		StatusInterval:  statusInterval,
		Spaces:          "     ",
		StartTimeMutex:  &sync.RWMutex{},
	}, nil
}

// Run Tasks
func (r *Runner) Run() (err error) {
	r.StartTime = time.Now()
	r.showStatus()
	ch := make(chan error)
	go r.handleTerminationSignal(ch)
	go r.run(ch)
	go r.waitAnyProcessError(ch)
	go r.showStatusByInterval()
	go r.readInput()
	err = <-ch
	if err == nil && r.getKilledSignal() {
		r.showStatus()
		return fmt.Errorf("Terminated")
	}
	if !r.getKilledSignal() {
		r.Terminate()
	}
	r.showStatus()
	return err
}

// Terminate all processes
func (r *Runner) Terminate() {
	logger.PrintfError("Terminating\n")
	r.setKilledSignal()
	// kill unfinished commands
	r.CmdInfoMutex.Lock()
	killedCh := map[string]chan error{}
	for label, cmdInfo := range r.CmdInfo {
		killedCh[label] = make(chan error)
		cmd := cmdInfo.Cmd
		logger.PrintfKill("Kill %s (PID=%d)\n", label, cmd.Process.Pid)
		go r.killByPid(-cmd.Process.Pid, killedCh[label])
	}
	for label := range r.CmdInfo {
		if err := <-killedCh[label]; err != nil {
			fmt.Println(r.Spaces, err)
		}
		delete(r.CmdInfo, label)
	}
	r.CmdInfoMutex.Unlock()
}

func (r *Runner) readInput() {
	d := logger.NewDecoration()
	for {
		r.sleep(1 * time.Microsecond)
		input := ""
		fmt.Scanf("%s", &input)
		if input == "" {
			continue
		}
		r.CmdInfoMutex.Lock()
		cmdCount := len(r.CmdInfo)
		for label, cmdInfo := range r.CmdInfo {
			redirect := false
			if cmdCount > 1 {
				logger.Printf("%sDo you want to send the input to `%s`? (Y/n)%s\n", d.Bold, label, d.Normal)
				confirmInput := ""
				fmt.Scanf("%s", &confirmInput)
				redirect = confirmInput == "Y" || confirmInput == "y"
			}
			if cmdCount == 1 || redirect {
				io.WriteString(cmdInfo.StdInPipe, input+"\n")
			}
		}
		r.CmdInfoMutex.Unlock()
	}
}

func (r *Runner) showStatusByInterval() {
	for {
		r.sleep(r.StatusInterval)
		if r.getKilledSignal() {
			return
		}
		r.showStatus()
	}
}

func (r *Runner) waitAnyProcessError(ch chan error) {
	seen := map[string]bool{}
	for {
		r.sleep(1 * time.Microsecond)
		if r.getKilledSignal() {
			ch <- fmt.Errorf("Terminated")
			return
		}
		r.CmdInfoMutex.Lock()
		for label, cmdInfo := range r.CmdInfo {
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
					if !r.getKilledSignal() {
						logger.PrintfError("%s exited with error:\n%s\n", currentLabel, r.sprintfCmdArgs(currentCmd))
					}
					fmt.Println(err)
					r.unregisterCmd(currentLabel)
					ch <- err
					return
				}
				if !r.isTaskFinished(currentTaskName) {
					logger.PrintfError("%s stopped before ready:\n%s\n", currentLabel, r.sprintfCmdArgs(currentCmd))
					r.unregisterCmd(currentLabel)
					ch <- fmt.Errorf("%s stopped before ready", currentLabel)
					return
				}
				r.unregisterCmd(currentLabel)
				logger.PrintfError("%s exited without any error message\n", currentLabel)
			}()
		}
		r.CmdInfoMutex.Unlock()
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
	fmt.Println()
	logger.PrintfError("%s\n", errorMsg)
	ch <- fmt.Errorf(errorMsg)
}

func (r *Runner) setDoneSignal() {
	r.DoneMutex.Lock()
	r.Done = true
	r.DoneMutex.Unlock()
}

func (r *Runner) getDoneSignal() (isDone bool) {
	r.DoneMutex.RLock()
	isDone = r.Done
	r.DoneMutex.RUnlock()
	return isDone
}

func (r *Runner) setKilledSignal() {
	r.KilledMutex.Lock()
	r.Killed = true
	r.KilledMutex.Unlock()
}

func (r *Runner) getKilledSignal() (isKilled bool) {
	r.KilledMutex.RLock()
	isKilled = r.Killed
	r.KilledMutex.RUnlock()
	return isKilled
}

func (r *Runner) run(ch chan error) {
	if err := r.runTaskByNames(r.TaskNames); err != nil {
		ch <- err
		return
	}
	r.setDoneSignal()
	r.showStatus()
	d := logger.NewDecoration()
	logger.PrintfSuccess("%s%s🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉%s\n", d.Bold, d.Green, d.Normal)
	logger.PrintfSuccess("%s%sJob Complete!!! 🎉🎉🎉%s\n", d.Bold, d.Green, d.Normal)
	autostop, autostopDefined := r.Project.GetValue("autostop"), r.Project.IsValueExist("autostop")
	if autostopDefined {
		if autostop != "" && autostop != "true" {
			autostopDuration, parseErr := time.ParseDuration(autostop)
			if parseErr != nil {
				ch <- parseErr
				logger.PrintfError("Cannot parse autostop duration %s", autostop)
				return
			}
			r.sleep(autostopDuration)
		}
		ch <- nil
		return
	}
	// wait until no cmd left
	for {
		r.sleep(1 * time.Microsecond)
		if r.getKilledSignal() {
			ch <- fmt.Errorf("Terminated")
			return
		}
		processExist := false
		r.CmdInfoMutex.RLock()
		for range r.CmdInfo {
			processExist = true
			break
		}
		r.CmdInfoMutex.RUnlock()
		if !processExist {
			ch <- nil
			return
		}
	}
}

func (r *Runner) runTaskByNames(taskNames []string) (err error) {
	tasks := []*config.Task{}
	for _, taskName := range taskNames {
		task, exists := r.Project.Tasks[taskName]
		if !exists {
			return fmt.Errorf("Task %s is not exist", taskName)
		}
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
		logger.PrintfSuccess("Reach %s '%s' wrapper\n", task.Icon, task.GetName())
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
	logger.PrintfStarted("Run %s '%s' command on %s\n", task.Icon, task.GetName(), startCmd.Dir)
	startStdinPipe, err := startCmd.StdinPipe()
	if err == nil {
		err = startCmd.Start()
	}
	if err != nil {
		logger.PrintfError("Error running command %s '%s':\n%s\n", task.Icon, task.GetName(), r.sprintfCmdArgs(startCmd))
		fmt.Println(r.Spaces, err)
		return err
	}
	startCmdLabel := fmt.Sprintf("%s '%s' command", task.Icon, task.GetName())
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
	logger.PrintfStarted("Run %s '%s' service on %s\n", task.Icon, task.GetName(), startCmd.Dir)
	startStdinPipe, err := startCmd.StdinPipe()
	if err == nil {
		err = startCmd.Start()
	}
	if err != nil {
		logger.PrintfError("Error running service %s '%s':\n%s\n", task.Icon, task.GetName(), r.sprintfCmdArgs(startCmd))
		fmt.Println(r.Spaces, err)
		return err
	}
	startCmdLabel := fmt.Sprintf("%s '%s' service", task.Icon, task.GetName())
	r.registerProcessCmd(startCmdLabel, task.GetName(), startCmd, startStdinPipe)
	return err
}

func (r *Runner) runCheckServiceTask(task *config.Task, checkCmd *exec.Cmd, checkLogDone chan error) (err error) {
	logger.PrintfStarted("Check %s '%s' readiness on %s\n", task.Icon, task.GetName(), checkCmd.Dir)
	checkStdinPipe, err := checkCmd.StdinPipe()
	if err == nil {
		err = checkCmd.Start()
	}
	if err != nil {
		logger.PrintfError("Error checking service %s '%s' readiness:\n%s\n", task.Icon, task.GetName(), r.sprintfCmdArgs(checkCmd))
		fmt.Println(r.Spaces, err)
		return err
	}
	checkCmdLabel := fmt.Sprintf("%s '%s' readiness check", task.Icon, task.GetName())
	r.registerCommandCmd(checkCmdLabel, task.GetName(), checkCmd, checkStdinPipe)
	err = r.waitTaskCmd(task, checkCmd, checkCmdLabel, checkLogDone)
	r.unregisterCmd(checkCmdLabel)
	return err
}

func (r *Runner) waitTaskCmd(task *config.Task, cmd *exec.Cmd, cmdLabel string, logDone chan error) (err error) {
	executed := false
	ch := make(chan error)
	go func() {
		if waitErr := cmd.Wait(); waitErr != nil {
			logger.PrintfError("Error running %s:\n%s\n", cmdLabel, r.sprintfCmdArgs(cmd))
			fmt.Println(r.Spaces, waitErr)
			ch <- waitErr
			return
		}
		executed = true
		logger.PrintfSuccess("Successfully running %s\n", cmdLabel)
		if logErr := <-logDone; logErr != nil {
			logger.PrintfError("Error logging %s:\n", cmdLabel)
			fmt.Println(r.Spaces, logErr)
		} else {
			logger.PrintfSuccess("Successfully logging %s\n", cmdLabel)
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
		logger.PrintfError("%s\n", timeoutMessage)
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
	r.CmdInfoMutex.Lock()
	r.CmdInfo[label] = &CmdInfo{
		Cmd:       cmd,
		IsProcess: isProcess,
		StdInPipe: stdinPipe,
		TaskName:  taskName,
	}
	r.CmdInfoMutex.Unlock()
}

func (r *Runner) unregisterCmd(label string) {
	r.CmdInfoMutex.Lock()
	delete(r.CmdInfo, label)
	r.CmdInfoMutex.Unlock()
}

func (r *Runner) registerTask(taskName string) (success bool) {
	r.TaskStatusMutex.Lock()
	_, isStarted := r.TaskStatus[taskName]
	if isStarted {
		success = false
	} else {
		r.TaskStatus[taskName] = NewTaskStatus()
		success = true
	}
	r.TaskStatusMutex.Unlock()
	return success
}

func (r *Runner) finishTask(taskName string, err error) {
	r.TaskStatusMutex.Lock()
	r.TaskStatus[taskName].Finish(err)
	r.TaskStatusMutex.Unlock()
}

func (r *Runner) isTaskFinished(taskName string) (isFinished bool) {
	r.TaskStatusMutex.RLock()
	isFinished = r.TaskStatus[taskName].Finished
	r.TaskStatusMutex.RUnlock()
	return isFinished
}

func (r *Runner) isTaskError(taskName string) (err error) {
	r.TaskStatusMutex.RLock()
	err = r.TaskStatus[taskName].Error
	r.TaskStatusMutex.RUnlock()
	return err
}

func (r *Runner) waitTaskFinished(taskName string) (err error) {
	for {
		r.sleep(1 * time.Microsecond)
		if r.isTaskFinished(taskName) {
			return r.isTaskError(taskName)
		}
		if r.getKilledSignal() {
			return fmt.Errorf("Terminated")
		}
	}
}

func (r *Runner) sprintfCmdArgs(cmd *exec.Cmd) (output string) {
	d := logger.NewDecoration()
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
			row = fmt.Sprintf("%s   %s%s%s%s", r.Spaces, d.Faint, prefix, row, d.Normal)
			rows[index] = row
		}
		formattedArg := strings.Join(rows, "\n")
		formattedArgs = append(formattedArgs, formattedArg)
	}
	output = strings.Join(formattedArgs, "\n")
	return fmt.Sprintf(output)
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
	d := logger.NewDecoration()
	return fmt.Sprintf("%s* (PID=%d) %s%s", d.Faint, cmd.Process.Pid, label, d.Normal)
}

func (r *Runner) showStatus() {
	d := logger.NewDecoration()
	descriptionPrefix := r.Spaces + "    "
	processPrefix := r.Spaces + r.Spaces + " "
	processRows := []string{}
	r.CmdInfoMutex.Lock()
	for label, cmdInfo := range r.CmdInfo {
		cmd := cmdInfo.Cmd
		processRow := r.getProcessRow(label, cmd)
		processRows = append(processRows, processRow)
	}
	r.CmdInfoMutex.Unlock()
	statusCaption := r.getStatusCaption()
	r.StartTimeMutex.RLock()
	elapsedTime := time.Since(r.StartTime)
	elapsedTimeCaption := fmt.Sprintf("%s%sElapsed Time: %s%s\n", descriptionPrefix, d.Faint, elapsedTime, d.Normal)
	r.StartTimeMutex.RUnlock()
	currentTime := time.Now()
	currentTimeString := currentTime.Format("15:04:05")
	currentTimeCaption := fmt.Sprintf("%s%sCurrent Time: %s%s\n", descriptionPrefix, d.Faint, currentTimeString, d.Normal)
	activeProcessLabel := ""
	processCaption := ""
	if len(processRows) > 0 {
		activeProcessLabel = fmt.Sprintf("%s%sActive Process:%s\n", descriptionPrefix, d.Faint, d.Normal)
		processCaption = processPrefix + strings.Join(processRows, "\n"+processPrefix) + "\n"
	}
	logger.PrintfInspect("%s%s%s%s%s", statusCaption, elapsedTimeCaption, currentTimeCaption, activeProcessLabel, processCaption)
}

func (r *Runner) getStatusCaption() (statusCaption string) {
	d := logger.NewDecoration()
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
