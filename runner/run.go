package runner

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	p, err := config.LoadProjectConfig(projectDir)
	if err != nil {
		errChan <- err
		executedChan <- true
		return
	}
	// get cmdDict and all it's output/error pipes
	cmdDict, outPipeDict, errPipeDict, err := getCmdDictAndPipes(projectDir, p)
	if err != nil {
		killCmdDict(p, cmdDict)
		errChan <- err
		executedChan <- true
		return
	}
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
		var cmd *exec.Cmd
		if component.Type == "container" {
			cmd, err = command.GetShellCmd(projectDir, fmt.Sprintf("(docker start %s || %s) && docker logs --follow %s", component.ContainerName, component.Run, component.ContainerName))
		} else if component.Type == "service" {
			cmd, err = command.GetShellCmd(component.Location, component.Start)
		} else {
			continue
		}
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

func logService(serviceName, prefix string, readCloser io.ReadCloser) {
	buf := bufio.NewScanner(readCloser)
	for buf.Scan() {
		log.Printf("[%s] %s > %s", prefix, serviceName, buf.Text())
	}
	if err := buf.Err(); err != nil {
		log.Printf("[ERROR] %s > %s", serviceName, err)
	}
}
