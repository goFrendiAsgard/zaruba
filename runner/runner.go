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
	defer r.cmdInfoMutex.Unlock()
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
		r.logger.DPrintfSuccess("Reach %s %s%s wrapper%s\n", task.GetDecoratedIcon(), task.GetColor(), task.GetName(), r.decoration.Normal)
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
	startCmdLabel := fmt.Sprintf("%s %s%s runner%s", task.GetDecoratedIcon(), task.GetColor(), task.GetName(), r.decoration.Normal)
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
	maxRetry := task.GetMaxStartRetry()
	retryDelayDuration := task.GetStartRetryDelayDuration()
	cmdLabel := fmt.Sprintf("%s %s%s starter%s", task.GetDecoratedIcon(), task.GetColor(), task.GetName(), r.decoration.Normal)
	cmdMaker := task.GetStartCmd
	cmd, startStdinPipe, err := r.createCmdAndStdinPipe(cmdMaker)
	if err != nil {
		return err
	}
	r.logger.DPrintfStarted("Running %s %s on %s\n", cmdLabel, r.getRetryAttemptCaption(1, maxRetry), cmd.Dir)
	cmd.Start()
	r.registerCmd(cmdLabel, task, cmdMaker, cmd, startStdinPipe, maxRetry, retryDelayDuration, true)
	return nil
}

func (r *Runner) checkLongRunningTask(task *dsl.Task) (err error) {
	maxCheckRetry := task.GetMaxCheckRetry()
	checkRetryDelayDuration := task.GetCheckRetryDelayDuration()
	checkCmdLabel := fmt.Sprintf("%s %s%s readiness checker%s", task.GetDecoratedIcon(), task.GetColor(), task.GetName(), r.decoration.Normal)
	err = r.waitSimpleCmd(checkCmdLabel, task, task.GetCheckCmd, maxCheckRetry, checkRetryDelayDuration)
	r.unregisterCmd(checkCmdLabel)
	return err
}

func (r *Runner) waitSimpleCmd(cmdLabel string, task *dsl.Task, cmdMaker func() (*exec.Cmd, error), maxRetry int, retryDelayDuration time.Duration) (err error) {
	timeoutDuration := task.GetTimeoutDuration()
	executed := false
	ch := make(chan error)
	// run task
	go func() {
		var cmdErr error
		var cmd *exec.Cmd
		var stdinPipe io.WriteCloser
		for attempt := 1; r.shouldRetry(attempt, maxRetry); attempt++ {
			cmd, stdinPipe, cmdErr = r.createCmdAndStdinPipe(cmdMaker)
			if cmdErr != nil {
				ch <- cmdErr
				return
			}
			r.logger.DPrintfStarted("Running %s %s on %s\n", cmdLabel, r.getRetryAttemptCaption(attempt, maxRetry), cmd.Dir)
			cmdErr = cmd.Start()
			r.registerCmd(cmdLabel, task, cmdMaker, cmd, stdinPipe, maxRetry, retryDelayDuration, false)
			if cmdErr != nil {
				r.handleCmdStartFailure(cmdLabel, cmdErr, cmd, attempt, maxRetry, retryDelayDuration)
				continue
			}
			cmdErr = cmd.Wait()
			// no error, quit this function
			if cmdErr == nil {
				r.handleCmdWaitSuccess(cmdLabel, attempt, maxRetry)
				executed = true
				ch <- nil
				return
			}
			// any error
			r.handleCmdWaitFailure(cmdLabel, cmdErr, cmd, attempt, maxRetry, retryDelayDuration)
		}
		ch <- cmdErr
	}()
	// checking timeout
	go func() {
		r.sleep(timeoutDuration)
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

func (r *Runner) waitLongRunningCmd(ch chan error) {
	seen := map[string]bool{}
	for {
		r.sleep(50 * time.Millisecond)
		if r.getKilledSignal() {
			ch <- fmt.Errorf("Terminated")
			return
		}
		r.cmdInfoMutex.Lock()
		for cmdLabel, cmdInfo := range r.cmdInfo {
			if _, exist := seen[cmdLabel]; exist || !cmdInfo.IsProcess {
				continue
			}
			seen[cmdLabel] = true
			cmd := cmdInfo.Cmd
			go r.handleLongRunningCmd(cmdLabel, cmdInfo, cmd, ch)
		}
		r.cmdInfoMutex.Unlock()
	}
}

func (r *Runner) handleLongRunningCmd(cmdLabel string, cmdInfo *CmdInfo, cmd *exec.Cmd, ch chan error) {
	r.cmdInfoMutex.Lock()
	cmdTask := cmdInfo.Task
	cmdTaskName := cmdTask.GetName()
	cmdMaker := cmdInfo.CmdMaker
	maxRetry := cmdInfo.MaxRetry
	retryDelayDuration := cmdInfo.RetryDelayDuration
	r.cmdInfoMutex.Unlock()
	var stdinPipe io.WriteCloser
	var cmdErr error
	for attempt := 1; r.shouldRetry(attempt, maxRetry); attempt++ {
		// for attempt == 1, cmd has already been started. So we don't have to re-create and start cmd
		if attempt > 1 {
			cmd, stdinPipe, cmdErr = r.createCmdAndStdinPipe(cmdMaker)
			if cmdErr != nil {
				ch <- cmdErr
				return
			}
			r.logger.DPrintfStarted("Running %s %s on %s\n", cmdLabel, r.getRetryAttemptCaption(attempt, maxRetry), cmd.Dir)
			cmdErr = cmd.Start()
			r.registerCmd(cmdLabel, cmdTask, cmdMaker, cmd, stdinPipe, maxRetry, retryDelayDuration, true)
			if cmdErr != nil {
				r.handleCmdStartFailure(cmdLabel, cmdErr, cmd, attempt, maxRetry, retryDelayDuration)
				continue
			}
		}
		cmdErr = cmd.Wait()
		// no error, quit loop
		if cmdErr == nil {
			r.handleCmdWaitSuccess(cmdLabel, attempt, maxRetry)
			break
		}
		// any error
		r.handleCmdWaitFailure(cmdLabel, cmdErr, cmd, attempt, maxRetry, retryDelayDuration)
	}
	if cmdErr != nil {
		r.handleCommonLongRunningCmdFailure("exited", cmdLabel, cmdErr, cmd)
		ch <- cmdErr
		return
	}
	if !r.isTaskFinished(cmdTaskName) {
		cmdErr = fmt.Errorf("%s stopped before ready", cmdLabel)
		r.handleCommonLongRunningCmdFailure("stopped", cmdLabel, cmdErr, cmd)
		ch <- cmdErr
		return
	}
	r.unregisterCmd(cmdLabel)
	r.logger.DPrintfError("%s exited without any error message\n", cmdLabel)
	ch <- fmt.Errorf("%s exited without any error message", cmdLabel)
}

func (r *Runner) handleCommonLongRunningCmdFailure(reason string, cmdLabel string, err error, cmd *exec.Cmd) {
	errMessage := r.getCmdErrorMessage(err, cmd)
	r.logger.DPrintfError("%s %s:%s\n", cmdLabel, reason, errMessage)
	r.unregisterCmd(cmdLabel)
	r.setSurpressWaitErrorSignal()

}

func (r *Runner) handleCmdWaitFailure(cmdLabel string, err error, cmd *exec.Cmd, attempt, maxRetry int, retryDelayDuration time.Duration) {
	r.handleCmdCommonFailure("Exit", cmdLabel, err, cmd, attempt, maxRetry)
	r.unregisterCmd(cmdLabel)
	if attempt != maxRetry && r.shouldRetry(attempt, maxRetry) {
		r.sleep(retryDelayDuration)
	}
}

func (r *Runner) handleCmdWaitSuccess(cmdLabel string, attempt, maxRetry int) {
	r.logger.DPrintfSuccess("Successfully running %s %s\n", cmdLabel, r.getRetryAttemptCaption(attempt, maxRetry))
}

func (r *Runner) handleCmdStartFailure(cmdLabel string, err error, cmd *exec.Cmd, attempt, maxRetry int, retryDelayDuration time.Duration) {
	r.handleCmdCommonFailure("Cannot start", cmdLabel, err, cmd, attempt, maxRetry)
	r.unregisterCmd(cmdLabel)
	if attempt != maxRetry && r.shouldRetry(attempt, maxRetry) {
		r.sleep(retryDelayDuration)
	}
}

func (r *Runner) handleCmdCommonFailure(logPrefix string, cmdLabel string, err error, cmd *exec.Cmd, attempt, maxRetry int) {
	errMessage := r.getCmdErrorMessage(err, cmd)
	r.logger.DPrintfError("%s %s %s:%s\n", logPrefix, cmdLabel, r.getRetryAttemptCaption(attempt, maxRetry), errMessage)
}

func (r *Runner) getCmdErrorMessage(err error, cmd *exec.Cmd) string {
	if !r.getKilledSignal() && !r.getSurpressWaitErrorSignal() {
		return fmt.Sprintf("\n%s\n%s", r.sprintfCmdArgs(cmd), err)
	}
	return fmt.Sprintf(" %s", err)
}

func (r *Runner) shouldRetry(attempt, maxRetry int) bool {
	if r.getKilledSignal() {
		return false
	}
	isInfiniteRetry := maxRetry < 1 // 0 or less indicate infinity
	if isInfiniteRetry {
		return true
	}
	return attempt <= maxRetry
}

func (r *Runner) getRetryAttemptCaption(attempt, maxRetry int) string {
	if maxRetry == 0 {
		return fmt.Sprintf("%s(Attempt %d of infinite)%s", r.decoration.Faint, attempt, r.decoration.Normal)
	}
	return fmt.Sprintf("%s(Attempt %d of %d)%s", r.decoration.Faint, attempt, maxRetry, r.decoration.Normal)
}

func (r *Runner) createCmdAndStdinPipe(cmdMaker func() (*exec.Cmd, error)) (cmd *exec.Cmd, stdinPipe io.WriteCloser, cmdErr error) {
	cmd, cmdErr = cmdMaker()
	if cmdErr != nil {
		return cmd, stdinPipe, cmdErr
	}
	stdinPipe, cmdErr = cmd.StdinPipe()
	return cmd, stdinPipe, cmdErr
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
