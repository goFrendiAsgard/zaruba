package command

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// GetCmd get cmd object
func GetCmd(dir, command string, args ...string) (cmd *exec.Cmd, err error) {
	cmd = exec.Command(command, args...)
	cmd.Env = os.Environ()
	cmd.Dir, err = filepath.Abs(dir)
	return cmd, err
}

// GetShellCmd get cmd object for running shell command
func GetShellCmd(dir, script string) (cmd *exec.Cmd, err error) {
	shell, shellArg := config.GetShellAndShellArg()
	args := []string{shellArg, script}
	return GetCmd(dir, shell, args...)
}

// RunCmd run cmd object
func RunCmd(cmd *exec.Cmd) (output string, err error) {
	logger.Info("Run `%s` on `%s`", strings.Join(cmd.Args, " "), cmd.Dir)
	outputB, err := cmd.Output()
	if err != nil {
		return output, err
	}
	output = string(outputB)
	return output, err
}

// Run run command
func Run(dir, command string, args ...string) (output string, err error) {
	cmd, err := GetCmd(dir, command, args...)
	cmd.Stderr = os.Stderr
	if err != nil {
		return output, err
	}
	return RunCmd(cmd)
}

// RunScript run script
func RunScript(dir, script string) (output string, err error) {
	cmd, err := GetShellCmd(dir, script)
	cmd.Stderr = os.Stderr
	if err != nil {
		return output, err
	}
	return RunCmd(cmd)
}

// RunCmdAndRedirect run cmd object
func RunCmdAndRedirect(cmd *exec.Cmd) (err error) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	logger.Info("Run `%s` on `%s`", strings.Join(cmd.Args, " "), cmd.Dir)
	err = cmd.Run()
	return err
}

// RunAndRedirect run command
func RunAndRedirect(dir, command string, args ...string) (err error) {
	cmd, err := GetCmd(dir, command, args...)
	if err != nil {
		return err
	}
	return RunCmdAndRedirect(cmd)
}

// RunScriptAndRedirect run script
func RunScriptAndRedirect(dir, script string) (err error) {
	cmd, err := GetShellCmd(dir, script)
	if err != nil {
		return err
	}
	return RunCmdAndRedirect(cmd)
}
