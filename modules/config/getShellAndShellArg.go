package config

import (
	"os"
)

// GetShellAndShellArg get shell and args
func GetShellAndShellArg() (shell, shellArg string) {
	shell = os.Getenv("ZARUBA_SHELL")
	shellArg = os.Getenv("ZARUBA_SHELL_ARG")
	if shell == "" {
		shell = "/bin/bash"
		if shellArg == "" {
			shellArg = "-c"
		}
	}
	return shell, shellArg
}
