package runner

import (
	"bufio"
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
	str, _ := p.ToColorizedYaml()
	logger.Info("Project Config Loaded: %s", str)
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
	// create network
	command.RunAndRedirect(projectDir, "docker", "network", "create", p.GetName())
	// run processes
	cmdMap, err := getCmdAndPipesMap(projectDir, p, orderedExecutions)
	if err != nil {
		killCmdMap(projectDir, p, cmdMap, orderedExecutions)
		errChan <- err
		executedChan <- true
		return
	}
	executedChan <- true
	// listen to stopChan
	<-stopChan
	killCmdMap(projectDir, p, cmdMap, orderedExecutions)
	errChan <- nil
}

func getServiceNames(p *config.ProjectConfig) (serviceNames []string) {
	serviceNames = []string{}
	for name, component := range p.GetComponents() {
		componentType := component.GetType()
		if componentType != "service" && componentType != "container" {
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
		errChans := make([]chan error, len(executionBatch))
		for index, serviceName := range executionBatch {
			errChans[index] = make(chan error)
			cmd := cmdMap[serviceName]
			if cmd.Process == nil {
				errChans[index] <- nil
				logger.Info("Process %s not found", serviceName)
				continue
			}
			go killCmd(projectDir, p, serviceName, cmd, errChans[index])
		}
		// wait all service run
		for _, errChan := range errChans {
			<-errChan
		}
	}
}

func killCmd(projectDir string, p *config.ProjectConfig, serviceName string, cmd *exec.Cmd, errChan chan error) {
	logger.Info("Kill %s process", serviceName)
	if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
		logger.Error("Failed to kill %s process: %s", serviceName, err)
		errChan <- err
	}
	errChan <- nil
}

func getCmdAndPipesMap(projectDir string, p *config.ProjectConfig, orderedExecutions [][]string) (cmdMap map[string]*exec.Cmd, err error) {
	cmdMap = map[string]*exec.Cmd{}
	for _, executionBatch := range orderedExecutions {
		errChans := make([]chan error, len(executionBatch))
		for index, serviceName := range executionBatch {
			component, err := p.GetComponentByName(serviceName)
			if err != nil {
				killCmdMap(projectDir, p, cmdMap, orderedExecutions)
				return cmdMap, err
			}
			// remove container
			// removeContainerIfNeeded(projectDir, component)
			// create cmd
			runtimeName, runtimeLocation, runtimeCommand, runtimeEnv, color := component.GetRuntimeName(), component.GetRuntimeLocation(), component.GetRuntimeCommand(), getServiceEnv(p, serviceName), component.GetColor()
			cmd, err := createServiceCmd(projectDir, serviceName, runtimeName, runtimeLocation, runtimeCommand, runtimeEnv, color)
			if err != nil {
				killCmdMap(projectDir, p, cmdMap, orderedExecutions)
				return cmdMap, err
			}
			// start cmd
			err = cmd.Start()
			cmdMap[serviceName] = cmd
			if err != nil {
				killCmdMap(projectDir, p, cmdMap, orderedExecutions)
				return cmdMap, err
			}
			errChans[index] = make(chan error)
			go checkServiceReadiness(serviceName, runtimeLocation, component.GetRuntimeReadinessCheckCommand(), runtimeEnv, errChans[index])
		}
		// wait all service run
		for _, errChan := range errChans {
			err = <-errChan
			if err != nil {
				killCmdMap(projectDir, p, cmdMap, orderedExecutions)
				return cmdMap, err
			}
		}
	}
	return cmdMap, err
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

func removeContainerIfNeeded(projectDir string, component *config.Component) {
	componentType := component.GetType()
	if componentType == "container" {
		containerName := component.GetRuntimeContainerName()
		command.RunAndRedirect(projectDir, "docker", "stop", containerName)
		command.RunAndRedirect(projectDir, "docker", "rm", containerName)
	}
}

func checkServiceReadiness(serviceName, runtimeLocation, readinessCheckCommand string, runtimeEnv []string, errChan chan error) {
	started := false
	failedCounter := 0
	logger.Info("Checking readiness of %s", serviceName)
	// set cmd
	for !started {
		cmd, err := command.GetShellCmd(runtimeLocation, readinessCheckCommand)
		if err != nil {
			errChan <- err
		}
		cmd.Env = runtimeEnv
		time.Sleep(time.Millisecond * 500)
		_, err = command.RunCmd(cmd)
		if err == nil {
			started = true
			break
		}
		// show failure and increase failedCounter
		if failedCounter == 0 {
			logger.Info("Failed to confirm readiness of %s: %s", serviceName, err)
		}
		failedCounter = failedCounter + 1
		if failedCounter == 100 {
			failedCounter = 0
		}
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
		log.Printf("\033[%dm[%s - %s]\033[0m %s", color, prefix, serviceName, buf.Text())
	}
	if err := buf.Err(); err != nil {
		logger.Error("%s: %s", serviceName, err)
	}
}

func getOrderedExecutions(p *config.ProjectConfig, executedServices []string) (executions [][]string, err error) {
	services := includeCandidate(p, executedServices)
	executions = [][]string{}
	leftServices := getLeftServices(services, executions)
	for len(leftServices) > 0 {
		currentBatch := []string{}
		for _, service := range leftServices {
			component, err := p.GetComponentByName(service)
			if err != nil {
				return executions, err
			}
			if isDependenciesFullfiled(component, executions) {
				currentBatch = append(currentBatch, service)
			}
		}
		executions = append(executions, currentBatch)
		leftServices = getLeftServices(services, executions)
	}
	return executions, err
}

func includeCandidate(p *config.ProjectConfig, services []string) []string {
	dependencies := []string{}
	for _, service := range services {
		component, err := p.GetComponentByName(service)
		if err != nil {
			logger.Fatal("Cannot get component ", service)
		}
		componentDependencies := component.GetDependencies()
		completeDependencies := includeCandidate(p, componentDependencies)
		for _, dependency := range completeDependencies {
			if !inArray(dependency, services) && !inArray(dependency, dependencies) {
				dependencies = append(dependencies, dependency)
			}
		}
	}
	return append(dependencies, services...)
}

func getLeftServices(services []string, executions [][]string) (left []string) {
	flattenExecutions := flattenArray(executions)
	left = []string{}
	for _, service := range services {
		if !inArray(service, flattenExecutions) {
			left = append(left, service)
		}
	}
	return left
}

func inArray(element string, arr []string) (found bool) {
	found = false
	for _, otherElement := range arr {
		if element == otherElement {
			found = true
			break
		}
	}
	return found
}

func isDependenciesFullfiled(component *config.Component, executions [][]string) (fullfilled bool) {
	flattenExecutions := flattenArray(executions)
	for _, dependency := range component.GetDependencies() {
		found := false
		for _, service := range flattenExecutions {
			if service == dependency {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func flattenArray(arr [][]string) (flatten []string) {
	flatten = []string{}
	for _, subArr := range arr {
		for _, element := range subArr {
			flatten = append(flatten, element)
		}
	}
	return flatten
}
