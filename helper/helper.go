package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetWorkingProjectPath() (projectPath string, err error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return GetProjectPath(currentDir)
}

func GetProjectPath(currentDir string) (projectPath string, err error) {
	for {
		projectFilePath := filepath.Join(currentDir, "index.zaruba.yaml")
		if _, err := os.Stat(projectFilePath); err == nil {
			return currentDir, nil
		}
		projectFilePath = filepath.Join(currentDir, "index.zaruba.yml")
		if _, err := os.Stat(projectFilePath); err == nil {
			return currentDir, nil
		}
		if currentDir == "/" {
			break
		}
		currentDir = filepath.Dir(currentDir)
	}
	return "", fmt.Errorf("cannot determine project path")
}
