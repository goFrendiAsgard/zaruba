package runner

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
)

// Run a project config.
func Run(projectDir string, stopChan, executedChan chan bool, errChan chan error) {
	projectDir, err := filepath.Abs(projectDir)
	if err != nil {
		errChan <- err
		executedChan <- true
		return
	}
	log.Printf("[INFO] Load config of project `%s`", projectDir)
	p, err := config.NewProjectConfig(projectDir)
	if err != nil {
		errChan <- err
		executedChan <- true
		return
	}
	str, _ := p.ToYaml()
	log.Printf("[INFO] Project Config Loaded:\n\033[33m%s\033[0m", str)
	// get cmdMap, run them, and get their output/error pipes
	log.Printf("[INFO] Run project `%s`", projectDir)
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
			log.Printf("[INFO] Process %s not found", serviceName)
			continue
		}
		log.Printf("[INFO] Kill %s process", serviceName)
		err := cmd.Process.Kill()
		if err != nil {
			log.Printf("[ERROR] Failed to kill %s process: %s", serviceName, err)
		}
	}
}

func getCmdAndPipesMap(projectDir string, p *config.ProjectConfig) (cmdMap map[string]*exec.Cmd, err error) {
	cmdMap = map[string]*exec.Cmd{}
	for _, serviceName := range p.GetExecutions() {
		component := p.GetComponentByName(serviceName)
		componentType := component.GetType()
		// get cmd
		if componentType != "container" && componentType != "service" {
			continue
		}
		runtimeName := component.GetRuntimeName()
		runtimeLocation := component.GetRuntimeLocation()
		runtimeEnv := getServiceEnv(p, serviceName)
		color := component.GetColor()
		cmd, err := command.GetShellCmd(runtimeLocation, component.GetRuntimeCommand())
		// set cmd.Env
		cmd.Env = runtimeEnv
		// get pipes
		outPipe, err := cmd.StdoutPipe()
		if err != nil {
			return cmdMap, err
		}
		errPipe, err := cmd.StderrPipe()
		if err != nil {
			return cmdMap, err
		}
		go logService(runtimeName, "OUT", color, outPipe)
		go logService(runtimeName, "ERR", color, errPipe)
		// run
		log.Printf("[INFO] Start %s: %s", serviceName, strings.Join(cmd.Args, " "))
		err = cmd.Start()
		cmdMap[serviceName] = cmd
		// if error, stop
		if err != nil {
			return cmdMap, err
		}
		// check whether the service is running or not
		startedChan := make(chan bool)
		errChan := make(chan error)
		go checkLiveness(serviceName, runtimeLocation, component.GetRuntimeLivenessCheckCommand(), runtimeEnv, startedChan, errChan)
		<-startedChan
		err = <-errChan
		if err != nil {
			return cmdMap, err
		}
		log.Printf("[INFO] %s started", serviceName)
	}
	return cmdMap, err
}

func checkLiveness(serviceName, runtimeLocation, livenessCheckCommand string, runtimeEnv []string, startedChan chan bool, errChan chan error) {
	started := false
	// set cmd
	for !started {
		cmd, err := command.GetShellCmd(runtimeLocation, livenessCheckCommand)
		if err != nil {
			startedChan <- false
			errChan <- err
		}
		cmd.Env = runtimeEnv
		time.Sleep(time.Second * 1)
		log.Printf("[INFO] Checking liveness of %s", serviceName)
		_, err = command.RunCmd(cmd)
		if err == nil {
			started = true
			break
		}
		log.Printf("[INFO] Failed to confirm liveness of %s: %s", serviceName, err)
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
		log.Printf("[ERROR] %s: %s", serviceName, err)
	}
}
