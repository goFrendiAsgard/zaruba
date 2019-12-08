package command

import (
	"os"
	"os/exec"
)

// Run a single command
func Run(shell []string, dir string, command string) error {
	commandList := append(shell, command)
	cmd := exec.Command(commandList[0], commandList[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return err
	}
	return err
}
