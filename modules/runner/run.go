package runner

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Run a project config.
func Run(projectDir string, p *config.ProjectConfig, executions []string, stopChan, executedChan chan bool, errChan chan error) {
	if len(executions) == 0 {
		executions = getServiceNames(p)
	}
	// get ordered execution
	orderedExecutions, err := getOrderedExecutions(p, executions)
	if err != nil {
		errChan <- err
		executedChan <- true
		return
	}
	// create docker network
	command.RunAndRedirect(projectDir, "docker", "network", "create", p.GetName())
	// prepare channels
	cmdMap := map[string]*exec.Cmd{}
	executed := false
	resultOfGetCmdMapChan := make(chan resultOfGetCmdMap)
	stopOnProgressChan := make(chan bool)
	// listen for stop channel
	go func() {
		<-stopChan
		if executed { // if execution is complete, kill cmdMap and send error to errChan. This will end everything
			killCmdMap(projectDir, p, cmdMap, orderedExecutions)
			errChan <- err // when everything is finished, we send something to errChan
		} else { // if execution is not complete, tell "getCmdMap" to stop immedieately
			stopOnProgressChan <- true
		}
	}()
	// run container, services, and commands
	go getCmdMap(projectDir, p, orderedExecutions, stopOnProgressChan, resultOfGetCmdMapChan)
	resultOfGetCmdMap := <-resultOfGetCmdMapChan
	cmdMap, err = resultOfGetCmdMap.cmdMap, resultOfGetCmdMap.err
	executed = true
	if err != nil { // if there is error (possibly because of "stopChan"), kill all process, send data to executedChan and errChan
		killCmdMap(projectDir, p, cmdMap, orderedExecutions)
		executedChan <- executed
		errChan <- err
		return
	}
	executedChan <- executed             // still waiting, because we don't send anything to errChan
	if !isServiceExists(p, executions) { // unless we only have "command" and "container" in this session. In that case, kill process
		killCmdMap(projectDir, p, cmdMap, orderedExecutions)
		errChan <- nil
	}
}

func isServiceExists(p *config.ProjectConfig, executions []string) bool {
	for _, execution := range executions {
		component, _ := p.GetComponentByName(execution)
		if component.GetType() == "service" {
			return true
		}
	}
	return false
}

func getServiceNames(p *config.ProjectConfig) (serviceNames []string) {
	serviceNames = []string{}
	for name, component := range p.GetComponents() {
		componentType := component.GetType()
		if componentType != "service" && componentType != "container" && componentType != "command" {
			continue
		}
		serviceNames = append(serviceNames, name)
	}
	return serviceNames
}

// kill based on p.Executions in reverse order
func killCmdMap(projectDir string, p *config.ProjectConfig, cmdMap map[string]*exec.Cmd, orderedExecutions [][]string) {
	for batchIndex := len(orderedExecutions) - 1; batchIndex >= 0; batchIndex-- {
		executionBatch := orderedExecutions[batchIndex]
		errChans := []chan error{}
		for _, serviceName := range executionBatch {
			errChan := make(chan error)
			errChans = append(errChans, errChan)
			cmd := cmdMap[serviceName]
			if cmd == nil || cmd.Process == nil {
				errChan <- nil
				logger.Info("Process %s not found", serviceName)
				continue
			}
			go killCmd(projectDir, p, serviceName, cmd, errChan)
		}
		// wait all service killed
		for _, errChan := range errChans {
			<-errChan
		}
	}
}

func killCmd(projectDir string, p *config.ProjectConfig, serviceName string, cmd *exec.Cmd, errChan chan error) {
	component, _ := p.GetComponentByName(serviceName)
	componentType := component.GetType()
	if componentType == "command" {
		errChan <- nil
		return
	}
	processType := "service"
	if componentType == "docker" {
		processType = "container logger"
	}
	logger.Info("Kill %s %s", processType, serviceName)
	if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
		logger.Error("Failed to kill %s %s: %s", processType, serviceName, err)
		errChan <- err
		return
	}
	logger.Info("%s %s killed", processType, serviceName)
	errChan <- nil
}

type resultOfGetCmdMap struct {
	cmdMap map[string]*exec.Cmd
	err    error
}

func createResultOfGetCmdMap(cmdMap map[string]*exec.Cmd, err error) resultOfGetCmdMap {
	return resultOfGetCmdMap{
		cmdMap: cmdMap,
		err:    err,
	}
}

func getCmdMap(projectDir string, p *config.ProjectConfig, orderedExecutions [][]string, stopProgressChan chan bool, resultOfGetCmdMapChan chan resultOfGetCmdMap) {
	cmdMap := map[string]*exec.Cmd{}
	// if there is stop signal, stop working
	go func() {
		stopProgress, notClosed := <-stopProgressChan
		if stopProgress || notClosed {
			resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, errors.New("Receiving stop signal"))
		}
	}()
	for _, executionBatch := range orderedExecutions {
		serviceErrChans := []chan error{}
		for _, serviceName := range executionBatch {
			component, err := p.GetComponentByName(serviceName)
			if err != nil {
				resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, err)
			}
			// create cmd
			runtimeName, runtimeLocation, runtimeCommand, runtimeEnv, color := component.GetRuntimeName(), component.GetRuntimeLocation(), component.GetRuntimeCommand(), getServiceEnv(p, serviceName), component.GetColor()
			cmd, err := createServiceCmd(projectDir, serviceName, runtimeName, runtimeLocation, runtimeCommand, runtimeEnv, color)
			if err != nil {
				resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, err)
			}
			// start cmd
			err = cmd.Start()
			cmdMap[serviceName] = cmd
			if err != nil {
				resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, err)
			}
			serviceErrChan := make(chan error)
			serviceErrChans = append(serviceErrChans, serviceErrChan)
			if component.GetType() == "command" {
				go checkCommandFinished(cmd, serviceErrChan)
			} else {
				go checkServiceReadiness(serviceName, runtimeLocation, component.GetRuntimeReadinessCheckCommand(), runtimeEnv, serviceErrChan)
			}
		}
		// wait all service run
		for _, serviceErrChan := range serviceErrChans {
			err := <-serviceErrChan
			if err != nil {
				resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, err)
				return
			}
		}
	}
	resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, nil)
}

func createServiceCmd(projectDir, serviceName, runtimeName, runtimeLocation string, runtimeCommand string, runtimeEnv []string, color int) (cmd *exec.Cmd, err error) {
	cmd, err = command.GetShellCmd(runtimeLocation, runtimeCommand)
	cmd.Env = runtimeEnv
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return cmd, err
	}
	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return cmd, err
	}
	go logService(runtimeName, "OUT", color, outPipe)
	go logService(runtimeName, "ERR", color, errPipe)
	logger.Info("Start %s: %s", serviceName, strings.Join(cmd.Args, " "))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	return cmd, err
}

func checkCommandFinished(cmd *exec.Cmd, errChan chan error) {
	err := cmd.Wait()
	errChan <- err
}

func checkServiceReadiness(serviceName, runtimeLocation, readinessCheckCommand string, runtimeEnv []string, errChan chan error) {
	started := false
	failedCounter := 0
	outputInterval := 20
	// set cmd
	for !started {
		cmd, err := command.GetShellCmd(runtimeLocation, readinessCheckCommand)
		if err != nil {
			errChan <- err
		}
		cmd.Env = runtimeEnv
		if failedCounter == 0 {
			logger.Info("Checking readiness of %s", serviceName)
			_, err = command.RunCmd(cmd)
		} else {
			_, err = command.RunCmdSilently(cmd)
		}
		if err == nil {
			started = true
			break
		}
		// show failure and increase failedCounter
		if failedCounter == 0 {
			logger.Info("Failed to confirm readiness of %s: %s", serviceName, err)
		}
		failedCounter = failedCounter + 1
		if failedCounter == outputInterval {
			failedCounter = 0
		}
		time.Sleep(time.Millisecond * 500)
	}
	logger.Info("%s ready", serviceName)
	errChan <- nil
}

func getServiceEnv(p *config.ProjectConfig, serviceName string) (environ []string) {
	c, err := p.GetComponentByName(serviceName)
	if err != nil {
		return environ
	}
	environMap := c.GetRuntimeEnv()
	// transform the map into array
	configEnv := []string{}
	for key, val := range environMap {
		configEnv = append(configEnv, fmt.Sprintf("%s=%s", key, val))
	}
	// merge the array with os.Environ
	environ = append(os.Environ(), configEnv...)
	return environ
}

func logService(serviceName, prefix string, color int, readCloser io.ReadCloser) {
	buf := bufio.NewScanner(readCloser)
	for buf.Scan() {
		log.Printf("\033[%dm%s - %s\033[0m  %s", color, prefix, serviceName, buf.Text())
	}
	if err := buf.Err(); err != nil {
		logger.Error("%s: %s", serviceName, err)
	}
}
