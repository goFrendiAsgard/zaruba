package config

import (
	"os"
	"path/filepath"
)

// GetTestDir retrieve test dir from environment variable
func GetTestDir() (absTestDir string) {
	testDir := os.Getenv("ZARUBA_TEST_DIR")
	if testDir == "" {
		return "/tmp"
	}
	absTestDir, err := filepath.Abs(testDir)
	if err != nil {
		return "/tmp"
	}
	return absTestDir
}
