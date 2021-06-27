package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/config"
)

func UpdateProjectEnvFiles(project *config.Project, serviceName string, location string) (err error) {
	absDirPath, err := filepath.Abs(location)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(absDirPath)
	if err != nil {
		return err
	}
	envRef, envRefExist := project.EnvRefMap[serviceName]
	if !envRefExist {
		return fmt.Errorf("envRef %s doesn't exist", serviceName)
	}
	for _, file := range files {
		isDir := file.IsDir()
		if isDir {
			continue
		}
		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".env") && !strings.HasSuffix(fileName, ".env.template") {
			continue
		}
		fileEnvMap, err := godotenv.Read(filepath.Join(absDirPath, fileName))
		if err != nil {
			return err
		}
		for key, env := range envRef.Map {
			if _, keyExist := fileEnvMap[key]; keyExist {
				continue
			}
			envFrom, envDefault := env.From, env.Default
			fileEnvMap[envFrom] = envDefault
		}
		godotenv.Write(fileEnvMap, filepath.Join(absDirPath, fileName))
	}
	return nil
}
