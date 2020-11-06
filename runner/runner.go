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
	TaskNames       []string
	Conf            *config.ProjectConfig
	TaskStatus      map[string]*TaskStatus
	TaskStatusMutex *sync.RWMutex
	CommandCmds     map[string]*exec.Cmd
	CommandCmdMutex *sync.RWMutex
	ProcessCmds     map[string]*exec.Cmd
	ProcessCmdMutex *sync.RWMutex
	Killed          bool
	KilledMutex     *sync.RWMutex
	Spaces          string
}

// NewRunner create new runner
func NewRunner(conf *config.ProjectConfig, taskNames []string) (runner *Runner) {
	return &Runner{
		TaskNames:       taskNames,
		Conf:            conf,
		TaskStatus:      map[string]*TaskStatus{},
		TaskStatusMutex: &sync.RWMutex{},
		CommandCmds:     map[string]*exec.Cmd{},
		CommandCmdMutex: &sync.RWMutex{},
		ProcessCmds:     map[string]*exec.Cmd{},
		ProcessCmdMutex: &sync.RWMutex{},
		Killed:          false,
		KilledMutex:     &sync.RWMutex{},
		Spaces:          "     ",
	}
}

// Run Tasks
func (r *Runner) Run() (err error) {
	ch := make(chan error)
	go r.handleTerminationSignal(ch)
	go r.run(ch)
	go r.waitAnyProcessError(ch)
	err = <-ch
	if err != nil {
		if !r.getKilledSignal() {
			r.Terminate()
		}
		return err
	}
	return nil
}

// Terminate all processes
func (r *Runner) Terminate() {
	logger.PrintfError("Terminating\n")
	r.propagateKilledSignal()
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
				r.unregisterProcessCmd(currentLabel)
				if err != nil {
					ch <- err
				} else {
					logger.PrintfError("%s exited without any error message\n", currentLabel)
				}
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

func (r *Runner) propagateKilledSignal() {
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
	logger.PrintfSuccess("Job Complete !!!\n")
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
	if err := r.runTaskNames(task.Dependencies); err != nil {
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
	logger.PrintfStarted("Check '%s' readiness on %s\n", task.Name, checkCmd.Dir)
	startCmdLabel := fmt.Sprintf("'%s' service", task.Name)
	r.registerProcessCmd(startCmdLabel, startCmd)
	// checker
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
		time.Sleep(task.TimeoutDuration)
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

func (r *Runner) registerProcessCmd(label string, cmd *exec.Cmd) {
	r.ProcessCmdMutex.Lock()
	r.ProcessCmds[label] = cmd
	r.ProcessCmdMutex.Unlock()
}

func (r *Runner) unregisterProcessCmd(label string) {
	r.ProcessCmdMutex.Lock()
	delete(r.ProcessCmds, label)
	r.ProcessCmdMutex.Unlock()
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

func (r *Runner) sprintfCmdArgs(cmd *exec.Cmd) string {
	d := logger.NewDecoration()
	formattedArgs := []string{}
	for _, arg := range cmd.Args {
		rows := strings.Split(arg, "\n")
		formattedArg := strings.Join(rows, fmt.Sprintf("\n%s     ", r.Spaces))
		formattedArgs = append(formattedArgs, fmt.Sprintf("%s   * %s", r.Spaces, formattedArg))
	}
	return fmt.Sprintf("%s%s%s", d.Dim, strings.Join(formattedArgs, "\n"), d.Normal)
}
