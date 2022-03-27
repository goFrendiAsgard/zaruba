package core

import (
	"os"
	"path/filepath"
)

func SetDefaultEnv() {
	executable, _ := os.Executable()
	if os.Getenv("ZARUBA_HOME") == "" {
		os.Setenv("ZARUBA_HOME", filepath.Dir(executable))
	}
	if os.Getenv("ZARUBA_BIN") == "" {
		os.Setenv("ZARUBA_BIN", executable)
	}
	if os.Getenv("ZARUBA_SHELL") == "" {
		os.Setenv("ZARUBA_SHELL", "bash")
	}
	if os.Getenv("ZARUBA_SCRIPTS") == "" {
		os.Setenv("ZARUBA_SCRIPTS", "")
	}
	if os.Getenv("ZARUBA_DECORATION") == "" {
		os.Setenv("ZARUBA_DECORATION", "default")
	}
}
