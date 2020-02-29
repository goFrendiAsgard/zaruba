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
	return cmd, err
}

// GetShellCmd get Cmd object for running shell command
func GetShellCmd(dir, script string) (cmd *exec.Cmd, err error) {
	shell, shellArg := config.GetShellAndShellArg()
	args := []string{shellArg, script}
	return GetCmd(dir, shell, args...)
}

// RunCmd a single command
func RunCmd(cmd *exec.Cmd) (err error) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("[INFO] Run `%s` on `%s`", strings.Join(cmd.Args, " "), cmd.Dir)
	err = cmd.Run()
	return err
}

// Run run command
func Run(dir, command string, args ...string) (err error) {
	cmd, err := GetCmd(dir, command, args...)
	if err != nil {
		return err
	}
	return RunCmd(cmd)
}

// RunScript run script
func RunScript(dir, script string) (err error) {
	cmd, err := GetShellCmd(dir, script)
	if err != nil {
		return err
	}
	return RunCmd(cmd)
}
