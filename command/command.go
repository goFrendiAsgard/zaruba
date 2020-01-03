package command

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/state-alchemists/zaruba/stringformat"
)

// GetCmd get Cmd object
func GetCmd(dir, command string, args ...string) (cmd *exec.Cmd, err error) {
	cmd = exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Dir, err = filepath.Abs(dir)
	return
}

// Run a single command
func Run(dir, command string, args ...string) (err error) {
	cmd, err := GetCmd(dir, command, args...)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Run `%s` on `%s` %s", command, cmd.Dir, stringformat.SprintArgs(args))
	err = cmd.Run()
	return err
}
