package runner

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
)

// Run a project config
func Run(projectDir string, stopChan chan bool, errChan chan error) {
	projectDir, err := filepath.Abs(projectDir)
	if err != nil {
		errChan <- err
		return
	}
	log.Printf("[INFO] Run project `%s`", projectDir)
	p, err := config.LoadProjectConfig(projectDir)
	if err != nil {
		errChan <- err
		return
	}
	cmdDict, err := getCmdDict(p)
	if err != nil {
		errChan <- err
		return
	}
	/*
		// listen to sigterm
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			stopChan <- true
		}()
	*/
	// listen to stopChan
	<-stopChan
	killCmdDict(cmdDict)
	errChan <- nil
}

func killCmdDict(cmdDict map[string]*exec.Cmd) {
	for serviceName, cmd := range cmdDict {
		log.Printf("[INFO] Kill %s process", serviceName)
		if cmd.Process == nil {
			continue
		}
		err := cmd.Process.Kill()
		if err != nil {
			log.Printf("[ERROR] Failed to kill %s process: %s", serviceName, err)
		}
	}
}

func getCmdDict(p *config.ProjectConfig) (cmdDict map[string]*exec.Cmd, err error) {
	cmdDict = map[string]*exec.Cmd{}
	for _, serviceName := range p.Executions {
		component := p.Components[serviceName]
		var cmd *exec.Cmd
		if component.Type == "container" {
			cmd, err = command.GetShellCmd(p.ProjectDir, fmt.Sprintf("(docker start %s || %s) && docker logs --follow %s", component.ContainerName, component.Run, component.ContainerName))
		} else if component.Type == "service" {
			cmd, err = command.GetShellCmd(component.Location, component.Start)
		} else {
			continue
		}
		log.Printf("[INFO] Start %s", serviceName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		/*
			go logServiceOutput(serviceName, cmd)
			go logServiceError(serviceName, cmd)
		*/
		err = cmd.Start()
		cmdDict[serviceName] = cmd
		// if error, stop
		if err != nil {
			return
		}
	}
	return
}

func logServiceOutput(serviceName string, cmd *exec.Cmd) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("[ERROR] Cannot get stdoutPipe: %s", err)
		return
	}
	logService(serviceName, "OUT", stdout)
}

func logServiceError(serviceName string, cmd *exec.Cmd) {
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Printf("[ERROR] Cannot get stderrPipe: %s", err)
		return
	}
	logService(serviceName, "ERR", stderr)
}

func logService(serviceName, prefix string, readCloser io.ReadCloser) {
	buf := bufio.NewScanner(readCloser)
	go func() {
		for buf.Scan() {
			log.Printf("[%s] %s > %s", prefix, serviceName, buf.Text())
		}
	}()
	go func() {
		if err := buf.Err(); err != nil {
			log.Printf("[ERROR] %s > %s", serviceName, err)
		}
	}()
	forever := make(chan bool)
	<-forever
}
