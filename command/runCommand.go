package command

import (
	"os"
	"os/exec"
)

// Run a single command
func Run(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)
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
