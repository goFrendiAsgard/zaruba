package config

import (
	"os"
	"runtime"
)

// GetShell retrieve template dir from environment variable
func GetShell() string {
	shell := os.Getenv("ZARUBA_SHELL")
	if shell == "" {
		if runtime.GOOS == "windows" {
			shell = "c:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
		}
		shell = "/bin/bash"
	}
	return shell
}
