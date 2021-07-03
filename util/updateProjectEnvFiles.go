package util

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/config"
)

func UpdateProjectEnvFiles(project *config.Project) (err error) {
	projectDir := filepath.Dir(project.GetFileLocation())
	files, err := ioutil.ReadDir(projectDir)
	if err != nil {
		return err
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
		fileEnvMap, err := godotenv.Read(filepath.Join(projectDir, fileName))
		if err != nil {
			return err
		}
		for _, task := range project.Tasks {
			for _, envKey := range task.GetEnvKeys() {
				envObj, declared := task.GetEnvObject(envKey)
				if !declared {
					continue
				}
				envFrom := envObj.From
				if envFrom == "" {
					continue
				}
				if _, keyExist := fileEnvMap[envFrom]; keyExist {
					continue
				}
				envDefault := envObj.Default
				fileEnvMap[envFrom] = envDefault
			}
		}
		godotenv.Write(fileEnvMap, filepath.Join(projectDir, fileName))
	}
	return nil
}
