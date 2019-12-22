package command

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Run a single command
func Run(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Dir, _ = filepath.Abs(dir)
	err := cmd.Run()
	return err
}
