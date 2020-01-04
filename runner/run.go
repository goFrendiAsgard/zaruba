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
	log.Printf("[INFO] Run project `%s`", projectDir)
	p := config.LoadProjectConfig(projectDir)
	// get cmdDict and all it's output/error pipes
	cmdDict, outPipeDict, errPipeDict, err := getCmdDictAndPipes(projectDir, p)
	if err != nil {
		killCmdDict(p, cmdDict)
		errChan <- err
		executedChan <- true
		return
	}
	executedChan <- true
	// redirect error and output pipe
	for serviceName := range cmdDict {
		go logService(serviceName, "OUT", outPipeDict[serviceName])
		go logService(serviceName, "ERR", errPipeDict[serviceName])
	}
	// listen to stopChan
	<-stopChan
	killCmdDict(p, cmdDict)
	errChan <- nil
}

// kill based on p.Executions in reverse order
func killCmdDict(p *config.ProjectConfig, cmdDict map[string]*exec.Cmd) {
	for index := len(p.Executions) - 1; index >= 0; index-- {
		serviceName := p.Executions[index]
		cmd := cmdDict[serviceName]
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

func getCmdDictAndPipes(projectDir string, p *config.ProjectConfig) (cmdDict map[string]*exec.Cmd, outPipeDict, errPipeDict map[string]io.ReadCloser, err error) {
	cmdDict = map[string]*exec.Cmd{}
	outPipeDict = map[string]io.ReadCloser{}
	errPipeDict = map[string]io.ReadCloser{}
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
		outPipeDict[serviceName], err = cmd.StdoutPipe()
		if err != nil {
			return
		}
		errPipeDict[serviceName], err = cmd.StderrPipe()
		if err != nil {
			return
		}
		log.Printf("[INFO] Start %s: %s", serviceName, strings.Join(cmd.Args, " "))
		err = cmd.Start()
		cmdDict[serviceName] = cmd
		// if error, stop
		if err != nil {
			return
		}
	}
	return
}

func getServiceEnv(p *config.ProjectConfig, serviceName string) (environ []string) {
	environDict := map[string]string{}
	for key, val := range p.Environments.General {
		environDict[key] = val
	}
	if serviceEnv, exists := p.Environments.Services[serviceName]; exists {
		for key, val := range serviceEnv {
			environDict[key] = val
		}
	}
	// transform the map into array
	configEnv := []string{}
	for key, val := range environDict {
		configEnv = append(configEnv, fmt.Sprintf("%s=%s", key, val))
	}
	// merge the array with os.Environ
	environ = append(os.Environ(), configEnv...)
	return
}

func logService(serviceName, prefix string, readCloser io.ReadCloser) {
	buf := bufio.NewScanner(readCloser)
	for buf.Scan() {
		log.Printf("[%s | %s] %s", prefix, serviceName, buf.Text())
	}
	if err := buf.Err(); err != nil {
		log.Printf("[ERROR] %s > %s", serviceName, err)
	}
}
