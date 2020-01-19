package command

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/modules/config"
)

// GetCmd get Cmd object
func GetCmd(dir, command string, args ...string) (cmd *exec.Cmd, err error) {
	cmd = exec.Command(command, args...)
	cmd.Env = os.Environ()
	cmd.Dir, err = filepath.Abs(dir)
	return
}

// GetShellCmd get Cmd object for running shell command
func GetShellCmd(dir, script string) (cmd *exec.Cmd, err error) {
	shell, shellArg := config.GetShell()
	args := []string{shellArg, script}
	return GetCmd(dir, shell, args...)
}

// Run a single command
func Run(cmd *exec.Cmd) (err error) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("[INFO] Run `%s` on `%s`", strings.Join(cmd.Args, " "), cmd.Dir)
	err = cmd.Run()
	return err
}
