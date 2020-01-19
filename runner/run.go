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

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
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
	p, err := config.LoadProjectConfig(projectDir)
	if err != nil {
		errChan <- err
		executedChan <- true
		return
	}
	log.Printf("[INFO] Run project `%s`", projectDir)
	// get cmdMap, run them, and get their output/error pipes
	cmdMap, outPipeMap, errPipeMap, err := getCmdAndPipesMap(projectDir, p)
	if err != nil {
		killCmdMap(p, cmdMap)
		errChan <- err
		executedChan <- true
		return
	}
	executedChan <- true
	// redirect error and output pipe
	for serviceName := range cmdMap {
		go logService(serviceName, "OUT", outPipeMap[serviceName])
		go logService(serviceName, "ERR", errPipeMap[serviceName])
	}
	// listen to stopChan
	<-stopChan
	killCmdMap(p, cmdMap)
	errChan <- nil
}

// kill based on p.Executions in reverse order
func killCmdMap(p *config.ProjectConfig, cmdMap map[string]*exec.Cmd) {
	for index := len(p.Executions) - 1; index >= 0; index-- {
		serviceName := p.Executions[index]
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

func getCmdAndPipesMap(projectDir string, p *config.ProjectConfig) (cmdMap map[string]*exec.Cmd, outPipeMap, errPipeMap map[string]io.ReadCloser, err error) {
	cmdMap = map[string]*exec.Cmd{}
	outPipeMap = map[string]io.ReadCloser{}
	errPipeMap = map[string]io.ReadCloser{}
	for _, serviceName := range p.Executions {
		component := p.Components[serviceName]
		serviceEnv := getServiceEnv(p, serviceName)
		// get cmd
		var cmd *exec.Cmd
		switch component.Type {
		case "container":
			cmd, err = command.GetShellCmd(projectDir, fmt.Sprintf("(docker start %s || %s) && docker logs --follow %s", component.ContainerName, component.Run, component.ContainerName))
		case "service":
			cmd, err = command.GetShellCmd(component.Location, component.Start)
		default:
			continue
		}
		// set cmd.Env
		cmd.Env = serviceEnv
		// get pipes
		outPipeMap[serviceName], err = cmd.StdoutPipe()
		if err != nil {
			return
		}
		errPipeMap[serviceName], err = cmd.StderrPipe()
		if err != nil {
			return
		}
		log.Printf("[INFO] Start %s: %s", serviceName, strings.Join(cmd.Args, " "))
		// run
		err = cmd.Start()
		cmdMap[serviceName] = cmd
		// if error, stop
		if err != nil {
			return
		}
	}
	return
}

func getServiceEnv(p *config.ProjectConfig, serviceName string) (environ []string) {
	environMap := map[string]string{}
	for key, val := range p.Environments.General {
		environMap[key] = val
	}
	if serviceEnv, exists := p.Environments.Services[serviceName]; exists {
		for key, val := range serviceEnv {
			environMap[key] = val
		}
	}
	// transform the map into array
	configEnv := []string{}
	for key, val := range environMap {
		configEnv = append(configEnv, fmt.Sprintf("%s=%s", key, val))
	}
	// merge the array with os.Environ
	environ = append(os.Environ(), configEnv...)
	return
}

func logService(serviceName, prefix string, readCloser io.ReadCloser) {
	buf := bufio.NewScanner(readCloser)
	for buf.Scan() {
		log.Printf("[%s - %s] %s", prefix, serviceName, buf.Text())
	}
	if err := buf.Err(); err != nil {
		log.Printf("[ERROR] %s > %s", serviceName, err)
	}
}
