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
	"syscall"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Run a project config.
func Run(projectDir string, p *config.ProjectConfig, selectors []string, stopChan, executedChan chan bool, errChan chan error) {
	// get ordered executions
	executableComponentNames, err := getExecutableComponentNames(p, selectors)
	orderedExecutions := [][]string{}
	if err == nil {
		orderedExecutions, err = getOrderedExecutions(p, executableComponentNames)
	}
	if err != nil {
		executedChan <- true
		errChan <- err
		return
	}
	// create docker network
	if isComponentTypeExists(p, "container", orderedExecutions) {
		command.RunAndRedirect(projectDir, "docker", "network", "create", p.GetName())
	}
	// prepare channels
	cmdMap := map[string]*exec.Cmd{}
	resultOfGetCmdMapChan := make(chan resultOfGetCmdMap)
	stopOnProgressChan := make(chan bool)
	executed := false
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
	resultOfGetCmdMap, executed := <-resultOfGetCmdMapChan, true
	cmdMap, err = resultOfGetCmdMap.cmdMap, resultOfGetCmdMap.err
	if err != nil { // if there is error (possibly because of "stopChan"), kill all process, send data to executedChan and errChan
		executedChan <- executed
		killCmdMap(projectDir, p, cmdMap, orderedExecutions)
		errChan <- err
		return
	}
	executedChan <- executed                                     // still waiting, because we don't send anything to errChan
	if !isComponentTypeExists(p, "service", orderedExecutions) { // unless we only have "command" and "container" in this session. In that case, kill process
		killCmdMap(projectDir, p, cmdMap, orderedExecutions)
		errChan <- err
	}
}

func isComponentTypeExists(p *config.ProjectConfig, componentType string, orderedExecutions [][]string) bool {
	for _, executions := range orderedExecutions {
		for _, execution := range executions {
			component, _ := p.GetComponentByName(execution)
			if component.GetType() == componentType {
				return true
			}
		}
	}
	return false
}

func getExecutableComponentNames(p *config.ProjectConfig, executions []string) (componentNames []string, err error) {
	componentNames = []string{}
	var components map[string]*config.Component
	if len(executions) == 0 {
		components = p.GetComponents()
	} else {
		components, err = p.GetComponentsByNamesOrLabels(executions)
		if err != nil {
			return componentNames, err
		}
	}
	for name, component := range components {
		componentType := component.GetType()
		if componentType != "service" && componentType != "container" && componentType != "command" {
			continue
		}
		componentNames = append(componentNames, name)
	}
	return componentNames, err
}

// kill based on p.Executions in reverse order
func killCmdMap(projectDir string, p *config.ProjectConfig, cmdMap map[string]*exec.Cmd, orderedExecutions [][]string) {
	for batchIndex := len(orderedExecutions) - 1; batchIndex >= 0; batchIndex-- {
		executionBatch := orderedExecutions[batchIndex]
		errChans := []chan error{}
		for _, serviceName := range executionBatch {
			errChan := make(chan error)
			errChans = append(errChans, errChan)
			cmd, exists := cmdMap[serviceName]
			if !exists || cmd == nil || cmd.Process == nil {
				logger.Info("Process %s not found", serviceName)
				go func() {
					errChan <- nil
				}()
				continue
			}
			go killCmd(p, serviceName, cmd, errChan)
		}
		// wait all service killed
		for _, errChan := range errChans {
			<-errChan
		}
	}
}

func killCmd(p *config.ProjectConfig, serviceName string, cmd *exec.Cmd, errChan chan error) {
	component, _ := p.GetComponentByName(serviceName)
	componentType := component.GetType()
	if componentType == "command" { // ignore command
		errChan <- nil
		return
	}
	processLabel := getProcessLabel(componentType, serviceName)
	logger.Info("Kill %s", processLabel)
	if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
		logger.Error("Failed to kill %s: %s", processLabel, err)
		errChan <- err
		return
	}
	logger.Info("Succesfully kill %s", processLabel)
	errChan <- nil
}

func getProcessLabel(componentType, serviceName string) (processLabel string) {
	if componentType == "docker" {
		return fmt.Sprintf("container logger %s", serviceName)
	}
	return fmt.Sprintf("service %s", serviceName)
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

func getCmdMap(projectDir string, p *config.ProjectConfig, orderedExecutions [][]string, stopOnProgressChan chan bool, resultOfGetCmdMapChan chan resultOfGetCmdMap) {
	cmdMap, forcedStop := map[string]*exec.Cmd{}, false
	// if there is stop signal, stop working
	go func() {
		<-stopOnProgressChan
		resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, errors.New("Receiving stop signal"))
		forcedStop = true
	}()
	for _, executionBatch := range orderedExecutions {
		if forcedStop {
			return
		}
		serviceErrChans := []chan error{}
		for _, serviceName := range executionBatch {
			component, err := p.GetComponentByName(serviceName)
			if err != nil {
				resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, err)
				return
			}
			// create cmd
			runtimeEnv := getServiceEnv(p, serviceName)
			cmd, err := createServiceCmd(serviceName, component, runtimeEnv)
			if err != nil {
				resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, err)
				return
			}
			if forcedStop {
				return
			}
			// start cmd
			err = cmd.Start()
			cmdMap[serviceName] = cmd
			if err != nil {
				resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, err)
				return
			}
			// wait for component to be executed correctly
			serviceErrChan := make(chan error)
			serviceErrChans = append(serviceErrChans, serviceErrChan)
			go checkComponent(component, cmd, serviceName, runtimeEnv, serviceErrChan)
		}
		// wait all service run
		if serviceErr := waitServiceErrChans(serviceErrChans); serviceErr != nil {
			resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, serviceErr)
			return
		}
	}
	resultOfGetCmdMapChan <- createResultOfGetCmdMap(cmdMap, nil)
}

func waitServiceErrChans(serviceErrChans []chan error) (err error) {
	for _, serviceErrChan := range serviceErrChans {
		err = <-serviceErrChan
	}
	return err
}

func createServiceCmd(serviceName string, component *config.Component, runtimeEnv []string) (cmd *exec.Cmd, err error) {
	runtimeName := component.GetRuntimeName()
	runtimeLocation := component.GetRuntimeLocation()
	runtimeCommand := component.GetRuntimeCommand()
	color := component.GetColor()
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

func checkComponent(component *config.Component, cmd *exec.Cmd, serviceName string, runtimeEnv []string, serviceErrChan chan error) {
	if component.GetType() == "command" {
		checkCommandFinished(cmd, serviceErrChan)
	} else {
		checkServiceReadiness(component, serviceName, runtimeEnv, serviceErrChan)
	}
}

func checkCommandFinished(cmd *exec.Cmd, errChan chan error) {
	err := cmd.Wait()
	errChan <- err
}

func checkServiceReadiness(component *config.Component, serviceName string, runtimeEnv []string, errChan chan error) {
	var err error
	started := false
	trialCounter := 0
	resetTrialCounterInterval := 20
	runtimeReadinessURL := component.GetRuntimeReadinessURL()
	// set cmd
	for !started {
		verboseOutput := trialCounter == 0
		if runtimeReadinessURL == "" {
			err = runReadinessCmdCheck(verboseOutput, serviceName, component, runtimeEnv)
		} else {
			err = runReadinessURLCheck(verboseOutput, serviceName, runtimeReadinessURL)
		}
		if err == nil {
			logger.Info("%s ready", serviceName)
			errChan <- nil
			return
		}
		// show failure and increase failedCounter
		if trialCounter == 0 {
			logger.Info("Failed to confirm readiness of %s: %s", serviceName, err)
		}
		trialCounter = trialCounter + 1
		if trialCounter == resetTrialCounterInterval {
			trialCounter = 0
		}
		time.Sleep(time.Millisecond * 500)
	}
	logger.Info("%s is not ready", serviceName)
	errChan <- err
}

func runReadinessURLCheck(verboseOutput bool, serviceName string, runtimeReadinessURL string) (err error) {
	if verboseOutput {
		logger.Info("Checking readiness URL of %s", serviceName)
	}
	_, err = http.Get(runtimeReadinessURL)
	return err
}

func runReadinessCmdCheck(verboseOutput bool, serviceName string, component *config.Component, runtimeEnv []string) (err error) {
	cmd, err := command.GetShellCmd(component.GetRuntimeLocation(), component.GetRuntimeReadinessCheckCommand())
	if err != nil {
		return err
	}
	cmd.Env = runtimeEnv
	if verboseOutput {
		logger.Info("Checking readiness of %s", serviceName)
		_, err = command.RunCmd(cmd)
		return err
	}
	_, err = command.RunCmdSilently(cmd)
	return err
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
