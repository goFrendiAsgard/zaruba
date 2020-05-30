package runner

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

type processState struct {
	executed bool
	process  *exec.Cmd
}

func createProcessState(process *exec.Cmd) *processState {
	return &processState{
		executed: false,
		process:  process,
	}
}

// Runner stateful
type Runner struct {
	stoppedLock        *sync.RWMutex
	processesStateLock *sync.RWMutex
	stopped            bool
	processState       map[string]*processState
	p                  *config.ProjectConfig
	componentsToRun    map[string]*config.Component
	executionOrder     []string
}

// CreateRunner create runner
func CreateRunner(p *config.ProjectConfig, selectors []string) (r *Runner, err error) {
	components, err := p.GetComponentsBySelectors(selectors)
	if err != nil {
		return r, err
	}
	componentsToRun := map[string]*config.Component{}
	for name, component := range components {
		if component.GetType() != "container" && component.GetType() != "command" && component.GetType() != "service" {
			continue
		}
		componentsToRun[name] = component
	}
	r = &Runner{
		stoppedLock:        &sync.RWMutex{},
		processesStateLock: &sync.RWMutex{},
		stopped:            false,
		processState:       map[string]*processState{},
		p:                  p,
		componentsToRun:    componentsToRun,
		executionOrder:     []string{},
	}
	return r, err
}

// Run run components
func (r *Runner) Run(projectDir string, stopChan, executedChan chan bool, errChan chan error) {
	go r.stopByChan(stopChan, executedChan, errChan)
	runErrList := []chan error{}
	componentNameList := []string{}
	for componentName := range r.componentsToRun {
		componentNameList = append(componentNameList, componentName)
		runErr := make(chan error)
		go r.run(componentName, runErr)
		runErrList = append(runErrList, runErr)
	}
	var runAllErr error
	for _, runErr := range runErrList {
		err := <-runErr
		if runAllErr == nil && err != nil {
			runAllErr = err
			r.stop()
		}
	}
	executedChan <- true
	if runAllErr != nil {
		errChan <- runAllErr
		return
	}
	logger.Info("Successfully run: %s", strings.Join(componentNameList, ", "))
	logger.Info("Have fun !!!")
	// if components to run are all command, then kill everything
	if r.componentsToRunAreCommand() {
		r.stop()
		errChan <- nil
		return
	}
}

func (r *Runner) componentsToRunAreCommand() (isAllCommand bool) {
	isAllCommand = true
	for _, component := range r.componentsToRun {
		if component.GetType() != "command" {
			return false
		}
	}
	return isAllCommand
}

func (r *Runner) run(processName string, runErr chan error) {
	if r.isRegistered(processName) {
		go r.wait(processName, runErr)
		return
	}
	// get component
	component, err := r.p.GetComponentByName(processName)
	if err != nil {
		runErr <- err
		return
	}
	// wait/run dependencies
	if err := r.runOrWaitDependencies(component); err != nil {
		runErr <- err
		return
	}
	// create cmd
	cmd, err := r.createComponentCmd(component)
	if err != nil {
		runErr <- err
		return
	}
	// if the process is already registered, then just wait for it
	if r.isRegistered(processName) {
		go r.wait(processName, runErr)
		return
	}
	// if this is stopping don't run
	if r.isStopped() {
		runErr <- err
		return
	}
	// register and start
	r.register(processName, cmd)
	logger.Info("Starting %s", processName)
	if err := cmd.Start(); err != nil {
		runErr <- err
		return
	}
	// wait
	err = r.waitComponentReadiness(component, cmd)
	r.confirmExecution(processName)
	runErr <- err
}

func (r *Runner) waitComponentReadiness(component *config.Component, cmd *exec.Cmd) (err error) {
	componentName := component.GetName()
	switch component.GetType() {
	case "command":
		err = cmd.Wait()
		if err == nil {
			logger.Info("%s execution succeed", componentName)
		} else {
			logger.Error("%s execution failed: %s", componentName, err)
		}
	default:
		counter := 0
		for true && !r.isStopped() {
			if counter > 100000 {
				counter = 0
			}
			shouldLog := counter == 0
			if err := r.checkComponentReadiness(component, shouldLog); err == nil {
				logger.Info("%s is ready", componentName)
				break
			} else if shouldLog {
				logger.Error("%s is not ready: %s", componentName, err)
			}
			counter++
		}
	}
	return err
}

func (r *Runner) checkComponentReadiness(component *config.Component, shouldLog bool) (err error) {
	readinessURL := component.GetRuntimeReadinessURL()
	componentName := component.GetName()
	if readinessURL != "" {
		if shouldLog {
			logger.Info("Checking readines of %s: %s", componentName, readinessURL)
		}
		return r.checkReadinessByURL(readinessURL)
	}
	runtimeEnv := r.getComponentEnv(component)
	runtimeLocation := component.GetRuntimeLocation()
	runtimeReadinessCheckCommand := component.GetRuntimeReadinessCheckCommand()
	if shouldLog {
		logger.Info("Checking readines of %s: %s", componentName, runtimeReadinessCheckCommand)
	}
	cmd, err := command.GetShellCmd(runtimeLocation, runtimeReadinessCheckCommand)
	if err != nil {
		return err
	}
	cmd.Env = runtimeEnv
	if shouldLog {
		_, err = command.RunCmd(cmd)
		return err
	}
	_, err = command.RunCmdSilently(cmd)
	return err
}

func (r *Runner) checkReadinessByURL(readinessURL string) (err error) {
	resp, err := http.Get(readinessURL)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errorMessage := fmt.Sprintf("Response Code: %d", resp.StatusCode)
		return errors.New(errorMessage)
	}
	return nil
}

func (r *Runner) createComponentCmd(component *config.Component) (cmd *exec.Cmd, err error) {
	name := component.GetName()
	runtimeLocation := component.GetRuntimeLocation()
	runtimeCommand := component.GetRuntimeCommand()
	cmd, err = command.GetShellCmd(runtimeLocation, runtimeCommand)
	cmd.Env = r.getComponentEnv(component)
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return cmd, err
	}
	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return cmd, err
	}
	go r.logComponent(component, "OUT", outPipe)
	go r.logComponent(component, "ERR", errPipe)
	logger.Info("Creating command %s: %s", name, strings.Join(cmd.Args, " "))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	return cmd, err
}

func (r *Runner) getComponentEnv(component *config.Component) (environ []string) {
	environMap := component.GetRuntimeEnv()
	// transform the map into array
	configEnv := []string{}
	for key, val := range environMap {
		configEnv = append(configEnv, fmt.Sprintf("%s=%s", key, val))
	}
	// merge the array with os.Environ
	environ = append(os.Environ(), configEnv...)
	return environ
}

func (r *Runner) logComponent(component *config.Component, prefix string, readCloser io.ReadCloser) {
	runtimeName := component.GetRuntimeName()
	name := component.GetName()
	color := component.GetColor()
	buf := bufio.NewScanner(readCloser)
	for buf.Scan() {
		log.Printf("\033[%dm%s - %s\033[0m  %s", color, prefix, runtimeName, buf.Text())
	}
	if err := buf.Err(); err != nil {
		logger.Error("%s: %s", name, err)
	}
}

func (r *Runner) runOrWaitDependencies(component *config.Component) (err error) {
	if r.isStopped() {
		return err
	}
	dependencyErrChanList := []chan error{}
	for _, dependencyName := range component.GetDependencies() {
		dependencyErrChan := make(chan error)
		if r.isRegistered(dependencyName) {
			go r.wait(dependencyName, dependencyErrChan)
		} else {
			go r.run(dependencyName, dependencyErrChan)
		}
		dependencyErrChanList = append(dependencyErrChanList, dependencyErrChan)
	}
	for _, dependencyErr := range dependencyErrChanList {
		if err := <-dependencyErr; err != nil {
			return err
		}
	}
	return err
}

func (r *Runner) wait(processName string, errChan chan error) {
	for true {
		if r.isExecuted(processName) || r.isStopped() {
			break
		}
	}
	errChan <- nil
}

func (r *Runner) stopByChan(stopChan, executedChan chan bool, errChan chan error) {
	<-stopChan
	r.stop()
	errChan <- errors.New("Terminated")
	executedChan <- true
}

func (r *Runner) stop() {
	r.stoppedLock.Lock()
	defer r.stoppedLock.Unlock()
	r.stopped = true
	r.killall()
}

func (r *Runner) isStopped() bool {
	r.stoppedLock.RLock()
	defer r.stoppedLock.RUnlock()
	stopped := r.stopped
	return stopped
}

func (r *Runner) register(name string, process *exec.Cmd) {
	r.processesStateLock.Lock()
	defer r.processesStateLock.Unlock()
	if !r.isStopped() {
		r.processState[name] = createProcessState(process)
	}
}

func (r *Runner) confirmExecution(name string) {
	r.processesStateLock.Lock()
	defer r.processesStateLock.Unlock()
	r.executionOrder = append(r.executionOrder, name)
	r.processState[name].executed = true
}

func (r *Runner) isRegistered(name string) bool {
	r.processesStateLock.RLock()
	defer r.processesStateLock.RUnlock()
	_, exists := r.processState[name]
	return exists
}

func (r *Runner) isExecuted(name string) bool {
	if !r.isRegistered(name) {
		return false
	}
	r.processesStateLock.RLock()
	defer r.processesStateLock.RUnlock()
	return r.processState[name].executed
}

func (r *Runner) killall() {
	r.processesStateLock.Lock()
	defer r.processesStateLock.Unlock()
	for index := len(r.executionOrder) - 1; index >= 0; index-- {
		processName := r.executionOrder[index]
		process := r.processState[processName].process
		component, _ := r.p.GetComponentByName(processName)
		if component.GetType() == "command" {
			continue
		}
		processSuffix := "process"
		if component.GetType() == "container" {
			processSuffix = "container logger"
		}
		logger.Info("Killing %s %s", processName, processSuffix)
		if err := syscall.Kill(-process.Process.Pid, syscall.SIGTERM); err != nil {
			logger.Error("Failed to kill %s %s: %s", processName, processSuffix, err)
		}
	}
}
