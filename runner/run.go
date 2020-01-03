package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
)

// Run a project config
func Run(p *config.ProjectConfig, stopChan chan bool, errChan chan error) {
	cmdList, err := getCmdList(p)
	if err != nil {
		errChan <- err
		return
	}
	// run all cmds
	for _, cmd := range cmdList {
		err := cmd.Start()
		if err != nil {
			killCmdList(cmdList)
			errChan <- err
			return
		}
	}
	// listen to sigterm
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		killCmdList(cmdList)
		errChan <- nil
	}()
	// listen to stopChan
	stopped := <-stopChan
	if stopped {
		killCmdList(cmdList)
	}
	errChan <- nil
}

func killCmdList(cmdList []*exec.Cmd) {
	for _, cmd := range cmdList {
		if cmd.Process == nil {
			continue
		}
		err := cmd.Process.Kill()
		if err != nil {
			log.Printf("[ERROR] Failed to kill process: %s", err)
		}
	}
}

func getCmdList(p *config.ProjectConfig) (cmdList []*exec.Cmd, err error) {
	shell, shellArg := config.GetShell()
	// get cmdList
	cmdList = []*exec.Cmd{}
	for _, serviceName := range p.Executions {
		component := p.Components[serviceName]
		var cmd *exec.Cmd
		if component.Type == "container" {
			args := append([]string{shellArg}, fmt.Sprintf("docker start %s || %s", component.ContainerName, component.Run))
			cmd, err = command.GetCmd(p.ProjectDir, shell, args...)
		} else if component.Type == "service" {
			args := append([]string{shellArg}, component.Start)
			cmd, err = command.GetCmd(p.ProjectDir, shell, args...)
		} else {
			continue
		}
		// if error, stop
		if err != nil {
			return
		}
		cmdList = append(cmdList, cmd)
	}
	return
}
