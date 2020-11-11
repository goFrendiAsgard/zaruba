package runner

import (
	"fmt"
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

// Runner is used to run tasks
type Runner struct {
	TaskNames           []string
	Conf                *config.ProjectConfig
	TaskStatus          map[string]*TaskStatus
	TaskStatusMutex     *sync.RWMutex
	CommandCmds         map[string]*exec.Cmd
	CommandCmdMutex     *sync.RWMutex
	ProcessCmds         map[string]*exec.Cmd
	ProcessCmdTaskNames map[string]string
	ProcessCmdMutex     *sync.RWMutex
	Killed              bool
	KilledMutex         *sync.RWMutex
	Done                bool
	DoneMutex           *sync.RWMutex
	StatusInterval      time.Duration
	StartTimeMutex      *sync.RWMutex
	StartTime           time.Time
	Spaces              string
}

// NewRunner create new runner
func NewRunner(conf *config.ProjectConfig, taskNames []string, statusInterval time.Duration) (runner *Runner) {
	return &Runner{
		TaskNames:           taskNames,
		Conf:                conf,
		TaskStatus:          map[string]*TaskStatus{},
		TaskStatusMutex:     &sync.RWMutex{},
		CommandCmds:         map[string]*exec.Cmd{},
		CommandCmdMutex:     &sync.RWMutex{},
		ProcessCmds:         map[string]*exec.Cmd{},
		ProcessCmdTaskNames: map[string]string{},
		ProcessCmdMutex:     &sync.RWMutex{},
		Killed:              false,
		KilledMutex:         &sync.RWMutex{},
		Done:                false,
		DoneMutex:           &sync.RWMutex{},
		StatusInterval:      statusInterval,
		Spaces:              "     ",
		StartTimeMutex:      &sync.RWMutex{},
	}
}

// Run Tasks
func (r *Runner) Run() (err error) {
	r.StartTime = time.Now()
	ch := make(chan error)
	go r.handleTerminationSignal(ch)
	go r.run(ch)
	go r.waitAnyProcessError(ch)
	go r.showStatusByInterval()
	err = <-ch
	if err == nil && r.getKilledSignal() {
		return fmt.Errorf("Terminated")
	}
	if !r.getKilledSignal() {
		r.Terminate()
	}
	return err
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

func (r *Runner) showStatusByInterval() {
	for true {
		r.sleep(r.StatusInterval)
		if r.getKilledSignal() {
			return
		}
		r.showStatus()
	}
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
	r.CommandCmdMutex.Lock()
	for label, cmd := range r.CommandCmds {
		processRow := r.getProcessRow(label, cmd)
		processRows = append(processRows, processRow)
	}
	r.CommandCmdMutex.Unlock()
	r.ProcessCmdMutex.Lock()
	for label, cmd := range r.ProcessCmds {
		processRow := r.getProcessRow(label, cmd)
		processRows = append(processRows, processRow)
	}
	r.ProcessCmdMutex.Unlock()
	done := r.getDoneSignal()
	statusCaption := fmt.Sprintf("%sJob Starting...%s\n", d.Bold, d.Normal)
	if done {
		statusCaption = fmt.Sprintf("%s%sJob Running...%s\n", d.Bold, d.Green, d.Normal)
	}
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

// Terminate all processes
func (r *Runner) Terminate() {
	logger.PrintfError("Terminating\n")
	r.setKilledSignal()
	// kill unfinished commands
	r.CommandCmdMutex.Lock()
	for label, cmd := range r.CommandCmds {
		logger.PrintfKill("Kill %s\n", label)
		if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGINT); err != nil {
			fmt.Println(r.Spaces, err)
		}
	}
	r.CommandCmdMutex.Unlock()
	// kill running processes
	r.ProcessCmdMutex.Lock()
	for label, cmd := range r.ProcessCmds {
		logger.PrintfKill("Kill %s\n", label)
		if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGINT); err != nil {
			fmt.Println(r.Spaces, err)
		}
	}
	r.ProcessCmdMutex.Unlock()
}

func (r *Runner) waitAnyProcessError(ch chan error) {
	seen := map[string]bool{}
	for true {
		r.ProcessCmdMutex.Lock()
		for label, cmd := range r.ProcessCmds {
			if _, exist := seen[label]; exist {
				continue
			}
			seen[label] = true
			currentLabel, currentCmd := label, cmd
			go func() {
				err := currentCmd.Wait()
				currentTaskName := r.getProcessCmdTaskName(currentLabel)
				r.unregisterProcessCmd(currentLabel)
				if err != nil {
					logger.PrintfError("%s exited with error:\n%s\n", currentLabel, r.sprintfCmdArgs(currentCmd))
					fmt.Println(err)
					ch <- err
					return
				}
				if !r.isTaskFinished(currentTaskName) {
					logger.PrintfError("%s stopped before ready:\n%s\n", currentLabel, r.sprintfCmdArgs(currentCmd))
					ch <- fmt.Errorf("%s stopped before ready", currentLabel)
					return
				}
				logger.PrintfError("%s exited without any error message\n", currentLabel)
			}()
		}
		r.ProcessCmdMutex.Unlock()
	}
}

func (r *Runner) handleTerminationSignal(ch chan error) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	sig := <-signalChannel
	switch sig {
	case os.Interrupt:
		ch <- fmt.Errorf("Receiving SIGINT")
		break
	case syscall.SIGTERM:
		ch <- fmt.Errorf("Receiving SIGTERM")
		break
	default:
		ch <- fmt.Errorf("Receiving termination signal")
	}
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
	if err := r.runTaskNames(r.TaskNames); err != nil {
		ch <- err
		return
	}
	r.setDoneSignal()
	r.showStatus()
	d := logger.NewDecoration()
	logger.PrintfSuccess("%s%sJob Complete !!! 🎉🎉🎉%s\n", d.Bold, d.Green, d.Normal)
	// wait until no process left
	for true {
		processExist := false
		r.ProcessCmdMutex.RLock()
		for range r.ProcessCmds {
			processExist = true
			break
		}
		r.ProcessCmdMutex.RUnlock()
		if !processExist {
			ch <- nil
		}
	}
}

func (r *Runner) runTaskNames(taskNames []string) (err error) {
	tasks := []*config.Task{}
	for _, taskName := range taskNames {
		task, exists := r.Conf.Tasks[taskName]
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
	if !r.registerTask(task.Name) {
		ch <- r.waitTaskFinished(task.Name)
		return
	}
	if err := r.runTaskNames(task.GetDependencies()); err != nil {
		ch <- err
		return
	}
	startCmd, startExist, startErr := task.GetStartCmd()
	if !startExist {
		logger.PrintfSuccess("Reach '%s' wrapper\n", task.Name)
		r.finishTask(task.Name, nil)
		ch <- nil
		return
	}
	if startErr != nil {
		ch <- startErr
		return
	}
	checkCmd, checkExist, checkErr := task.GetCheckCmd()
	if !checkExist {
		err := r.runCommandTask(task, startCmd)
		r.finishTask(task.Name, err)
		ch <- err
		return
	}
	if checkErr != nil {
		ch <- checkErr
		return
	}
	err := r.runServiceTask(task, startCmd, checkCmd)
	r.finishTask(task.Name, err)
	ch <- err
}

func (r *Runner) runCommandTask(task *config.Task, startCmd *exec.Cmd) (err error) {
	logger.PrintfStarted("Run '%s' command on %s\n", task.Name, startCmd.Dir)
	if err = startCmd.Start(); err != nil {
		logger.PrintfError("Error running command '%s':\n%s\n", task.Name, r.sprintfCmdArgs(startCmd))
		fmt.Println(r.Spaces, err)
		return err
	}
	startCmdLabel := fmt.Sprintf("'%s' command", task.Name)
	r.registerCommandCmd(startCmdLabel, startCmd)
	err = r.runTaskCmdWithTimeout(task, startCmd, startCmdLabel)
	r.unregisterCommandCmd(startCmdLabel)
	return err
}

func (r *Runner) runServiceTask(task *config.Task, startCmd *exec.Cmd, checkCmd *exec.Cmd) (err error) {
	logger.PrintfStarted("Run '%s' service on %s\n", task.Name, startCmd.Dir)
	if err = startCmd.Start(); err != nil {
		logger.PrintfError("Error running service '%s':\n%s\n", task.Name, r.sprintfCmdArgs(startCmd))
		fmt.Println(r.Spaces, err)
		return err
	}
	startCmdLabel := fmt.Sprintf("'%s' service", task.Name)
	r.registerProcessCmd(startCmdLabel, startCmd, task)
	// checker
	logger.PrintfStarted("Check '%s' readiness on %s\n", task.Name, checkCmd.Dir)
	if err = checkCmd.Start(); err != nil {
		logger.PrintfError("Error checking service '%s' readiness:\n%s\n", task.Name, r.sprintfCmdArgs(checkCmd))
		fmt.Println(r.Spaces, err)
		return err
	}
	checkCmdLabel := fmt.Sprintf("'%s' readiness check", task.Name)
	r.registerCommandCmd(checkCmdLabel, checkCmd)
	err = r.runTaskCmdWithTimeout(task, checkCmd, checkCmdLabel)
	r.unregisterCommandCmd(checkCmdLabel)
	return err
}

func (r *Runner) runTaskCmdWithTimeout(task *config.Task, cmd *exec.Cmd, cmdLabel string) (err error) {
	executed := false
	ch := make(chan error)
	go func() {
		if waitErr := cmd.Wait(); waitErr != nil {
			logger.PrintfError("Error running %s:\n%s\n", cmdLabel, r.sprintfCmdArgs(cmd))
			fmt.Println(r.Spaces, waitErr)
			ch <- waitErr
			return
		}
		logger.PrintfSuccess("Successfully running %s\n", cmdLabel)
		executed = true
		ch <- nil
		return
	}()
	go func() {
		r.sleep(task.TimeoutDuration)
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

func (r *Runner) registerCommandCmd(label string, cmd *exec.Cmd) {
	r.CommandCmdMutex.Lock()
	r.CommandCmds[label] = cmd
	r.CommandCmdMutex.Unlock()
}

func (r *Runner) unregisterCommandCmd(label string) {
	r.CommandCmdMutex.Lock()
	delete(r.CommandCmds, label)
	r.CommandCmdMutex.Unlock()
}

func (r *Runner) registerProcessCmd(label string, cmd *exec.Cmd, task *config.Task) {
	r.ProcessCmdMutex.Lock()
	r.ProcessCmds[label] = cmd
	r.ProcessCmdTaskNames[label] = task.Name
	r.ProcessCmdMutex.Unlock()
}

func (r *Runner) unregisterProcessCmd(label string) {
	r.ProcessCmdMutex.Lock()
	delete(r.ProcessCmds, label)
	delete(r.ProcessCmdTaskNames, label)
	r.ProcessCmdMutex.Unlock()
}

func (r *Runner) getProcessCmdTaskName(label string) (taskName string) {
	r.ProcessCmdMutex.Lock()
	taskName = r.ProcessCmdTaskNames[label]
	r.ProcessCmdMutex.Unlock()
	return taskName
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
	for true {
		if r.isTaskFinished(taskName) {
			return r.isTaskError(taskName)
		}
		if r.getKilledSignal() {
			return fmt.Errorf("Terminated")
		}
	}
	return nil
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
