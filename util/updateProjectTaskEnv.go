package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/str"
)

func UpdateProjectTaskEnv(project *config.Project) (err error) {
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

		locationEnvMap, err := GetEnvByLocation(taskLocation)
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

func updateTaskEnv(task *config.Task, locationEnvMap map[string]string) (err error) {
	additionalEnvMap := getAdditionalEnvMap(task.Env, locationEnvMap)
	if len(additionalEnvMap) == 0 {
		return nil
	}
	taskName := task.GetName()
	serviceName := GetTaskServiceName(taskName)
	envPrefix := strings.ToUpper(str.ToSnakeCase(serviceName))
	yamlLocation := task.GetFileLocation()
	b, err := ioutil.ReadFile(yamlLocation)
	if err != nil {
		return err
	}
	fileContentStr := string(b)
	lines := strings.Split(fileContentStr, "\n")
	// look for task's env
	taskNamePattern := fmt.Sprintf("^ +%s:.*$", taskName)
	envLineIndex, submatch, err := str.GetLastSubmatch(lines, "tasks:.*$", taskNamePattern, "^ +env:.*$")
	if err != nil {
		return err
	}
	if envLineIndex == -1 {
		return fmt.Errorf("env of task %s found on %s", taskName, yamlLocation)
	}
	// get single indentation
	singleIndentation, err := str.GetSingleIndentation(submatch[0], 2)
	if err != nil {
		return err
	}
	lines[envLineIndex] = fmt.Sprintf("%senv:", str.Repeat(singleIndentation, 2))
	// get additional env yaml
	additionalEnvLines := getAdditionalYamlEnvLines(additionalEnvMap, envPrefix, singleIndentation, 2)
	// construct new lines
	newLines := []string{}
	newLines = append(newLines, lines[:envLineIndex+1]...)
	newLines = append(newLines, additionalEnvLines...)
	if envLineIndex+1 < len(lines) {
		newLines = append(newLines, lines[envLineIndex+1:]...)
	}
	// save
	newFileContentStr := strings.Join(newLines, "\n")
	return ioutil.WriteFile(yamlLocation, []byte(newFileContentStr), 0755)
}

func updateEnvRef(envRef *config.EnvRef, locationEnvMap map[string]string) (err error) {
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
	lines := strings.Split(fileContentStr, "\n")
	// look for envRefName
	envRefNamePattern := fmt.Sprintf("^ +%s:.*$", envRefName)
	envRefNameLineIndex, submatch, err := str.GetLastSubmatch(lines, "^envs:.*$", envRefNamePattern)
	if err != nil {
		return err
	}
	if envRefNameLineIndex == -1 {
		return fmt.Errorf("env %s not found on %s", envRefName, yamlLocation)
	}
	// get single indentation
	singleIndentation, err := str.GetSingleIndentation(submatch[0], 1)
	if err != nil {
		return err
	}
	// clean envRefName line
	lines[envRefNameLineIndex] = fmt.Sprintf("%s%s:", singleIndentation, envRefName)
	// get additional env yaml
	additionalEnvLines := getAdditionalYamlEnvLines(additionalEnvMap, envPrefix, singleIndentation, 1)
	// construct new lines
	newLines := []string{}
	newLines = append(newLines, lines[:envRefNameLineIndex+1]...)
	newLines = append(newLines, additionalEnvLines...)
	if envRefNameLineIndex+1 < len(lines) {
		newLines = append(newLines, lines[envRefNameLineIndex+1:]...)
	}
	// save
	newFileContentStr := strings.Join(newLines, "\n")
	return ioutil.WriteFile(yamlLocation, []byte(newFileContentStr), 0755)
}

func getAdditionalYamlEnvLines(additionalEnvMap map[string]string, envPrefix string, singleIndentation string, indentationLevel int) (envLines []string) {
	blockIndentationStr := str.Repeat(singleIndentation, indentationLevel)
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
		envLines = append(envLines, fmt.Sprintf("%s%s%s:", blockIndentationStr, singleIndentation, envKey))
		envLines = append(envLines, fmt.Sprintf("%s%s%sfrom: %s", blockIndentationStr, singleIndentation, singleIndentation, envFrom))
		envLines = append(envLines, fmt.Sprintf("%s%s%sdefault: %s", blockIndentationStr, singleIndentation, singleIndentation, envVal))
	}
	return envLines
}

func getAdditionalEnvMap(existingEnvMap map[string]*config.Env, locationEnvMap map[string]string) (additionalEnvMap map[string]string) {
	additionalEnvMap = map[string]string{}
	for envKey, envVal := range locationEnvMap {
		if _, exist := existingEnvMap[envKey]; exist {
			continue
		}
		additionalEnvMap[envKey] = envVal
	}
	return additionalEnvMap
}
