package core

import (
	"fmt"
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
	if os.Getenv("ZARUBA_LOG_TIME") == "" {
		os.Setenv("ZARUBA_LOG_TIME", "true")
	}
	if os.Getenv("ZARUBA_LOG_STATUS_TIME_INTERVAL") == "" {
		os.Setenv("ZARUBA_LOG_STATUS_TIME_INTERVAL", "5m")
	}
	if os.Getenv("ZARUBA_LOG_STATUS_LINE_INTERVAL") == "" {
		os.Setenv("ZARUBA_LOG_STATUS_LINE_INTERVAL", "40")
	}
	if os.Getenv("ZARUBA_ENV") == "" {
		os.Setenv("ZARUBA_ENV", "")
	}
	if os.Getenv("ZARUBA_MAX_LOG_FILE_SIZE") == "" {
		os.Setenv("ZARUBA_MAX_LOG_FILE_SIZE", fmt.Sprintf("%d", 5*1024*1024))
	}
}
