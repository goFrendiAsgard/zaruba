package runner

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Run a project config.
func Run(projectDir string, p *config.ProjectConfig, stopChan, executedChan chan bool, errChan chan error) {
	str, _ := p.ToColorizedYaml()
	logger.Info("Project Config Loaded: %s", str)
	cmdMap, err := getCmdAndPipesMap(projectDir, p)
	if err != nil {
		killCmdMap(p, cmdMap)
		errChan <- err
		executedChan <- true
		return
	}
	executedChan <- true
	// listen to stopChan
	<-stopChan
	killCmdMap(p, cmdMap)
	errChan <- nil
}

// kill based on p.Executions in reverse order
func killCmdMap(p *config.ProjectConfig, cmdMap map[string]*exec.Cmd) {
	executions := p.GetExecutions()
	for index := len(executions) - 1; index >= 0; index-- {
		serviceName := executions[index]
		cmd := cmdMap[serviceName]
		if cmd.Process == nil {
			logger.Info("Process %s not found", serviceName)
			continue
		}
		logger.Info("Kill %s process", serviceName)
		err := cmd.Process.Kill()
		if err != nil {
			logger.Error("Failed to kill %s process: %s", serviceName, err)
		}
	}
}

func getCmdAndPipesMap(projectDir string, p *config.ProjectConfig) (cmdMap map[string]*exec.Cmd, err error) {
	cmdMap = map[string]*exec.Cmd{}
	for _, serviceName := range p.GetExecutions() {
		component, err := p.GetComponentByName(serviceName)
		if err != nil {
			killCmdMap(p, cmdMap)
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
			killCmdMap(p, cmdMap)
			return cmdMap, err
		}
		errPipe, err := cmd.StderrPipe()
		if err != nil {
			killCmdMap(p, cmdMap)
			return cmdMap, err
		}
		go logService(runtimeName, "OUT", color, outPipe)
		go logService(runtimeName, "ERR", color, errPipe)
		logger.Info("Start %s: %s", serviceName, strings.Join(cmd.Args, " "))
		err = cmd.Start()
		cmdMap[serviceName] = cmd
		if err != nil {
			killCmdMap(p, cmdMap)
			return cmdMap, err
		}
		startedChan, errChan := make(chan bool), make(chan error)
		go checkReadiness(serviceName, runtimeLocation, component.GetRuntimeReadinessCheckCommand(), runtimeEnv, startedChan, errChan)
		<-startedChan
		err = <-errChan
		if err != nil {
			killCmdMap(p, cmdMap)
			return cmdMap, err
		}
		logger.Info("%s started", serviceName)
	}
	return cmdMap, err
}

func checkReadiness(serviceName, runtimeLocation, readinessCheckCommand string, runtimeEnv []string, startedChan chan bool, errChan chan error) {
	started := false
	// set cmd
	for !started {
		cmd, err := command.GetShellCmd(runtimeLocation, readinessCheckCommand)
		if err != nil {
			startedChan <- false
			errChan <- err
		}
		cmd.Env = runtimeEnv
		time.Sleep(time.Second * 1)
		logger.Info("Checking readiness of %s", serviceName)
		_, err = command.RunCmd(cmd)
		if err == nil {
			started = true
			break
		}
		logger.Info("Failed to confirm readiness of %s: %s", serviceName, err)
	}
	startedChan <- started
	errChan <- nil
}

func getServiceEnv(p *config.ProjectConfig, serviceName string) (environ []string) {
	environMap := p.GetEnvironments().GetRuntimeVariables(serviceName)
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
