package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/env"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/str"
	yaml "gopkg.in/yaml.v2"
)

func SyncProjectEnvFiles(project *Project) (err error) {
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

func SyncProjectEnv(project *Project) (err error) {
	projectDir := filepath.Dir(project.GetFileLocation())
	updatedEnvRef := map[string]bool{}
	for taskName, task := range project.Tasks {
		if !strings.HasPrefix(taskName, "run") {
			continue
		}
		taskFileLocation := task.GetFileLocation()
		if !strings.HasPrefix(taskFileLocation, projectDir) {
			// we won't update any task declared outside our project
			continue
		}
		taskLocation := task.GetTaskLocation()
		if taskLocation == "" || taskLocation == projectDir {
			// we won't update any task declared outside our project
			continue
		}
		locationEnvMap, err := env.GetEnvByLocation(taskLocation)
		if err != nil {
			return err
		}
		envRefName := GetTaskEnvRefname(task)
		if envRefName == "" {
			// update taskEnv
			newEnvMap := getAdditionalEnvMap(task.Env, locationEnvMap)
			if err = setTaskEnv(task, newEnvMap); err != nil {
				return err
			}
			continue
		}
		if updatedEnvRef[envRefName] {
			continue
		}
		updatedEnvRef[envRefName] = true
		// update envRef
		newEnvMap := getAdditionalEnvMap(project.EnvRefMap[envRefName].Map, locationEnvMap)
		if err = setEnvRef(project.EnvRefMap[envRefName], newEnvMap); err != nil {
			return err
		}
	}
	return nil
}

func SetTaskEnv(task *Task, envMap map[string]string) (err error) {
	envRefName := GetTaskEnvRefname(task)
	if envRefName == "" {
		// update taskEnv
		return setTaskEnv(task, envMap)
	}
	// update envRef
	return setEnvRef(task.Project.EnvRefMap[envRefName], envMap)
}

func setTaskEnv(task *Task, envMap map[string]string) (err error) {
	if len(envMap) == 0 {
		return nil
	}
	taskName := task.GetName()
	envPrefix := strings.ToUpper(str.ToSnakeCase(GetTaskServiceName(taskName)))
	yamlLocation := task.GetFileLocation()
	p, err := loadRawProject(yamlLocation)
	if err != nil {
		return err
	}
	for envKey, envVal := range envMap {
		envFrom := getEnvFrom(envKey, envPrefix)
		if _, exist := p.Tasks[taskName].Env[envKey]; !exist {
			p.Tasks[taskName].Env[envKey] = &Env{}
		}
		p.Tasks[taskName].Env[envKey].From = envFrom
		p.Tasks[taskName].Env[envKey].Default = envVal
	}
	yamlContentB, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return file.WriteText(yamlLocation, string(yamlContentB), 0555)
}

func setEnvRef(envRef *EnvRef, envMap map[string]string) (err error) {
	if len(envMap) == 0 {
		return nil
	}
	envRefName := envRef.GetName()
	envPrefix := strings.ToUpper(str.ToSnakeCase(envRefName))
	yamlLocation := envRef.GetFileLocation()
	p, err := loadRawProject(yamlLocation)
	if err != nil {
		return err
	}
	for envKey, envVal := range envMap {
		envFrom := getEnvFrom(envKey, envPrefix)
		if _, exist := p.RawEnvRefMap[envRefName][envKey]; !exist {
			p.RawEnvRefMap[envRefName][envKey] = &Env{}
		}
		p.RawEnvRefMap[envRefName][envKey].From = envFrom
		p.RawEnvRefMap[envRefName][envKey].Default = envVal
	}
	yamlContentB, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return file.WriteText(yamlLocation, string(yamlContentB), 0555)
}

func getEnvFrom(envKey, envPrefix string) string {
	if !strings.HasPrefix(envKey, envPrefix) {
		return fmt.Sprintf("%s_%s", envPrefix, envKey)
	}
	return envKey
}

func getAdditionalEnvMap(existingEnvMap map[string]*Env, locationEnvMap map[string]string) (additionalEnvMap map[string]string) {
	additionalEnvMap = map[string]string{}
	for envKey, envVal := range locationEnvMap {
		if _, exist := existingEnvMap[envKey]; exist {
			continue
		}
		additionalEnvMap[envKey] = envVal
	}
	return additionalEnvMap
}
