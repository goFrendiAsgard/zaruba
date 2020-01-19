package config

import (
	"os"
)

// GetShell get shell and args
func GetShell() (shell, shellArg string) {
	shell = os.Getenv("ZARUBA_SHELL")
	shellArg = os.Getenv("ZARUBA_SHELL_ARG")
	if shell == "" {
		shell = "/bin/bash"
		if shellArg == "" {
			shellArg = "-c"
		}
	}
	return
}
