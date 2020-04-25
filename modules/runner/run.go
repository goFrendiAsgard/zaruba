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
	orderedExecutions, err := getOrderedExecutions(p, executions)
	if err != nil {
		errChan <- err
		executedChan <- true
		return
	}
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
func killCmdMap(projectDir string, p *config.ProjectConfig, cmdMap map[string]*exec.Cmd, orderedExecutions []string) {
	for index := len(orderedExecutions) - 1; index >= 0; index-- {
		serviceName := orderedExecutions[index]
		cmd := cmdMap[serviceName]
		if cmd.Process == nil {
			logger.Info("Process %s not found", serviceName)
			continue
		}
		component, err := p.GetComponentByName(serviceName)
		if err != nil {
			logger.Error("Failed to get %s component: %s", serviceName, err)
		}
		if component.GetType() == "container" {
			logger.Info("Stop %s container", serviceName)
			err = command.RunAndRedirect(projectDir, "docker", "stop", component.GetRuntimeContainerName())
			if err != nil {
				logger.Error("Failed to stop %s container: %s", serviceName, err)
			}
			continue
		}
		logger.Info("Kill %s process", serviceName)
		if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
			logger.Error("Failed to kill %s process: %s", serviceName, err)
		}
	}
}

func getCmdAndPipesMap(projectDir string, p *config.ProjectConfig, orderedExecutions []string) (cmdMap map[string]*exec.Cmd, err error) {
	cmdMap = map[string]*exec.Cmd{}
	for _, serviceName := range orderedExecutions {
		component, err := p.GetComponentByName(serviceName)
		if err != nil {
			killCmdMap(projectDir, p, cmdMap, orderedExecutions)
			return cmdMap, err
		}
		componentType := component.GetType()
		if componentType != "container" && componentType != "service" {
			continue
		}
		runtimeName, runtimeLocation, runtimeEnv, color := component.GetRuntimeName(), component.GetRuntimeLocation(), getServiceEnv(p, serviceName), component.GetColor()
		cmd, err := command.GetShellCmd(runtimeLocation, component.GetRuntimeCommand())
		cmd.Env = runtimeEnv
		outPipe, err := cmd.StdoutPipe()
		if err != nil {
			killCmdMap(projectDir, p, cmdMap, orderedExecutions)
			return cmdMap, err
		}
		errPipe, err := cmd.StderrPipe()
		if err != nil {
			killCmdMap(projectDir, p, cmdMap, orderedExecutions)
			return cmdMap, err
		}
		go logService(runtimeName, "OUT", color, outPipe)
		go logService(runtimeName, "ERR", color, errPipe)
		logger.Info("Start %s: %s", serviceName, strings.Join(cmd.Args, " "))
		cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
		err = cmd.Start()
		cmdMap[serviceName] = cmd
		if err != nil {
			killCmdMap(projectDir, p, cmdMap, orderedExecutions)
			return cmdMap, err
		}
		readyChan, errChan := make(chan bool), make(chan error)
		go checkReadiness(serviceName, runtimeLocation, component.GetRuntimeReadinessCheckCommand(), runtimeEnv, readyChan, errChan)
		<-readyChan
		err = <-errChan
		if err != nil {
			killCmdMap(projectDir, p, cmdMap, orderedExecutions)
			return cmdMap, err
		}
		logger.Info("%s ready", serviceName)
	}
	return cmdMap, err
}

func checkReadiness(serviceName, runtimeLocation, readinessCheckCommand string, runtimeEnv []string, readyChan chan bool, errChan chan error) {
	started := false
	counter := 0
	// set cmd
	for !started {
		cmd, err := command.GetShellCmd(runtimeLocation, readinessCheckCommand)
		if err != nil {
			readyChan <- false
			errChan <- err
		}
		cmd.Env = runtimeEnv
		time.Sleep(time.Second * 1)
		if counter == 0 {
			logger.Info("Checking readiness of %s", serviceName)
		}
		_, err = command.RunCmd(cmd)
		if err == nil {
			started = true
			break
		}
		if counter == 0 {
			logger.Info("Failed to confirm readiness of %s: %s", serviceName, err)
		}
		counter = counter + 1
		if counter == 10 {
			counter = 0
		}
	}
	readyChan <- started
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

func getOrderedExecutions(p *config.ProjectConfig, services []string) (dependencies []string, err error) {
	dependencies = []string{}
	for _, service := range services {
		component, err := p.GetComponentByName(service)
		if err != nil {
			return dependencies, err
		}
		currentDependencies := component.GetDependencies()
		subDependencies, err := getOrderedExecutions(p, currentDependencies)
		if err != nil {
			return dependencies, err
		}
		dependencies = append(subDependencies, dependencies...)
		dependencies = append(dependencies, service)
	}
	dependencies = getUniqueArr(dependencies)
	return dependencies, err
}

func getUniqueArr(arr []string) (uniqueArr []string) {
	uniqueArr = []string{}
	encountered := map[string]bool{}
	for _, element := range arr {
		if encountered[element] {
			continue
		}
		uniqueArr = append(uniqueArr, element)
		encountered[element] = true
	}
	return uniqueArr
}
