package config

import (
	"os"
	"runtime"
	"strings"
)

// GetShell retrieve template dir from environment variable
func GetShell() []string {
	shell := os.Getenv("ZARUBA_SHELL")
	if shell == "" {
		if runtime.GOOS == "windows" {
			shell = "c:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
		} else {
			shell = "/bin/sh"
		}
	}
	shellArgs := strings.Split(os.Getenv("ZARUBA_SHELL_ARGS"), " ")
	if len(shellArgs) == 1 && shellArgs[0] == "" {
		if runtime.GOOS != "windows" {
			shellArgs = []string{"-c"}
		}
	}
	return append([]string{shell}, shellArgs...)
}
