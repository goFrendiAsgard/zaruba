package command

import (
	"os"
	"os/exec"
)

// RunSingle run a single command
func RunSingle(shell []string, dir string, environ []string, command string) error {
	commandList := append(shell, command)
	cmd := exec.Command(commandList[0], commandList[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = environ
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return err
	}
	return err
}

// RunMultiple run multiple commands
func RunMultiple(shell []string, dir string, environ []string, commands []string) error {
	for _, command := range commands {
		err := RunSingle(shell, dir, environ, command)
		if err != nil {
			return err
		}
	}
	return nil
}
