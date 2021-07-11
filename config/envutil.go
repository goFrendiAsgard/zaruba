package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/env"
	"github.com/state-alchemists/zaruba/str"
)

func UpdateProjectEnvFiles(project *Project) (err error) {
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

func UpdateProjectTaskEnv(project *Project) (err error) {
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
			if err = updateTaskEnv(task, locationEnvMap); err != nil {
				return err
			}
			continue
		}
		if updatedEnvRef[envRefName] {
			continue
		}
		updatedEnvRef[envRefName] = true
		// update envRef
		if err = updateEnvRef(project.EnvRefMap[envRefName], locationEnvMap); err != nil {
			return err
		}
	}
	return nil
}

func updateTaskEnv(task *Task, locationEnvMap map[string]string) (err error) {
	additionalEnvMap := getAdditionalEnvMap(task.Env, locationEnvMap)
	if len(additionalEnvMap) == 0 {
		return nil
	}
	taskName := task.GetName()
	envPrefix := strings.ToUpper(str.ToSnakeCase(GetTaskServiceName(taskName)))
	yamlLocation := task.GetFileLocation()
	b, err := ioutil.ReadFile(yamlLocation)
	if err != nil {
		return err
	}
	fileContentStr := string(b)
	patterns := []string{
		"tasks:.*$",
		fmt.Sprintf("^ +%s:.*$", taskName),
		"^ +env:.*$",
	}
	suplements := []string{
		"tasks:",
		fmt.Sprintf("  %s:", taskName),
		"    env: {}",
	}
	lines, _ := str.InsertIfNotFound(strings.Split(fileContentStr, "\n"), patterns, suplements)
	// look for task's env
	envLineIndex, submatch, err := str.GetFirstMatch(lines, suplements)
	if err != nil {
		return err
	}
	if envLineIndex == -1 {
		return fmt.Errorf("env of task %s not found on %s", taskName, yamlLocation)
	}
	// get single indentation
	singleIndentation, err := str.GetSingleIndentation(submatch[0], 2)
	if err != nil {
		return err
	}
	// prepare replacement
	envLine := fmt.Sprintf("%senv:", str.Repeat(singleIndentation, 2))
	additionalEnvLines := getAdditionalYamlEnvLines(additionalEnvMap, envPrefix, singleIndentation, 2)
	replacement := append([]string{envLine}, additionalEnvLines...)
	newLines, err := str.ReplaceLine(lines, envLineIndex, replacement)
	if err != nil {
		return err
	}
	// save
	newFileContentStr := strings.Join(newLines, "\n")
	return ioutil.WriteFile(yamlLocation, []byte(newFileContentStr), 0755)
}

func updateEnvRef(envRef *EnvRef, locationEnvMap map[string]string) (err error) {
	additionalEnvMap := getAdditionalEnvMap(envRef.Map, locationEnvMap)
	if len(additionalEnvMap) == 0 {
		return nil
	}
	envRefName := envRef.GetName()
	envPrefix := strings.ToUpper(str.ToSnakeCase(envRefName))
	yamlLocation := envRef.GetFileLocation()
	b, err := ioutil.ReadFile(yamlLocation)
	if err != nil {
		return err
	}
	fileContentStr := string(b)
	patterns := []string{
		"^envs:.*$",
		fmt.Sprintf("^ +%s:.*$", envRefName),
	}
	suplements := []string{
		"envs:",
		fmt.Sprintf("  %s: {}", envRefName),
	}
	lines, _ := str.InsertIfNotFound(strings.Split(fileContentStr, "\n"), patterns, suplements)
	// look for envRefName
	envRefNameLineIndex, submatch, _ := str.GetFirstMatch(lines, patterns)
	if envRefNameLineIndex == -1 {
		return fmt.Errorf("env %s not found on %s", envRefName, yamlLocation)
	}
	// get single indentation
	indentation, err := str.GetSingleIndentation(submatch[0], 1)
	if err != nil {
		return err
	}
	// prepare replacement
	envRefNameLine := fmt.Sprintf("%s%s:", indentation, envRefName)
	additionalEnvLines := getAdditionalYamlEnvLines(additionalEnvMap, envPrefix, indentation, 1)
	replacement := append([]string{envRefNameLine}, additionalEnvLines...)
	newLines, err := str.ReplaceLine(lines, envRefNameLineIndex, replacement)
	if err != nil {
		return err
	}
	// save
	newFileContentStr := strings.Join(newLines, "\n")
	return ioutil.WriteFile(yamlLocation, []byte(newFileContentStr), 0755)
}

func getAdditionalYamlEnvLines(additionalEnvMap map[string]string, envPrefix string, indentation string, indentationLevel int) (envLines []string) {
	blockIndentationStr := str.Repeat(indentation, indentationLevel)
	// sort keyList
	envKeyList := []string{}
	for envKey := range additionalEnvMap {
		envKeyList = append(envKeyList, envKey)
	}
	sort.Strings(envKeyList)
	envLines = []string{}
	// loop
	for _, envKey := range envKeyList {
		envVal := additionalEnvMap[envKey]
		envFrom := envKey
		if !strings.HasPrefix(envKey, envPrefix) {
			envFrom = fmt.Sprintf("%s_%s", envPrefix, envKey)
		}
		envLines = append(envLines, fmt.Sprintf("%s%s%s:", blockIndentationStr, indentation, envKey))
		envLines = append(envLines, fmt.Sprintf("%s%s%sfrom: %s", blockIndentationStr, indentation, indentation, envFrom))
		envLines = append(envLines, fmt.Sprintf("%s%s%sdefault: %s", blockIndentationStr, indentation, indentation, envVal))
	}
	return envLines
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
