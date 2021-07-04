package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
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
	// look for "tasks"
	taskLineIndex := -1
	for lineIndex, line := range lines {
		if strings.HasPrefix(line, "tasks:") {
			taskLineIndex = lineIndex
			break
		}
	}
	if taskLineIndex == -1 {
		return fmt.Errorf("no tasks found on %s", yamlLocation)
	}
	// Look for taskName
	indentationStr, taskNameLineIndex := "", -1
	rex := regexp.MustCompile(fmt.Sprintf("^( +)%s:(.*)$", taskName))
	for lineIndex := taskLineIndex + 1; lineIndex < len(lines); lineIndex++ {
		line := lines[lineIndex]
		matches := rex.FindAllStringSubmatch(line, -1)
		if len(matches) == 1 {
			taskNameLineIndex = lineIndex
			indentationStr = matches[0][1]
			suffix := strings.TrimLeft(matches[0][2], " ")
			if strings.HasPrefix(suffix, "{}") {
				lines[lineIndex] = fmt.Sprintf("%s%s:", indentationStr, taskName)
			}
			break
		}
	}
	if taskNameLineIndex == -1 {
		return fmt.Errorf("task %s not found on %s", taskName, yamlLocation)
	}
	// Look for env
	envLineIndex := -1
	rex = regexp.MustCompile("^ +env:(.*)$")
	for lineIndex := taskLineIndex + 1; lineIndex < len(lines); lineIndex++ {
		line := lines[lineIndex]
		matches := rex.FindAllStringSubmatch(line, -1)
		if len(matches) == 1 {
			envLineIndex = lineIndex
			suffix := strings.TrimLeft(matches[0][1], " ")
			if strings.HasPrefix(suffix, "{}") {
				lines[lineIndex] = fmt.Sprintf("%s%senv:", indentationStr, indentationStr)
			}
			break
		}
	}
	if envLineIndex == -1 {
		return fmt.Errorf("env of task %s found on %s", taskName, yamlLocation)
	}
	// get additional env yaml
	additionalEnvLines := getAdditionalEnvLines(additionalEnvMap, envPrefix, indentationStr, 1)
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
	// look for "envs:"
	envLineIndex := -1
	for lineIndex, line := range lines {
		if strings.HasPrefix(line, "envs:") {
			envLineIndex = lineIndex
			break
		}
	}
	if envLineIndex == -1 {
		return fmt.Errorf("no envs found on %s", yamlLocation)
	}
	// look for envRefKey
	indentationStr, envRefNameLineIndex := "", -1
	rex := regexp.MustCompile(fmt.Sprintf("^( +)%s:(.*)$", envRefName))
	for lineIndex := envLineIndex + 1; lineIndex < len(lines); lineIndex++ {
		line := lines[lineIndex]
		matches := rex.FindAllStringSubmatch(line, -1)
		if len(matches) == 1 {
			envRefNameLineIndex = lineIndex
			indentationStr = matches[0][1]
			suffix := strings.TrimLeft(matches[0][2], " ")
			if strings.HasPrefix(suffix, "{}") {
				lines[lineIndex] = fmt.Sprintf("%s%s:", indentationStr, envRefName)
			}
			break
		}
	}
	if envLineIndex == -1 {
		return fmt.Errorf("env %s not found on %s", envRefName, yamlLocation)
	}
	// get additional env yaml
	additionalEnvLines := getAdditionalEnvLines(additionalEnvMap, envPrefix, indentationStr, 1)
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

func getAdditionalEnvLines(additionalEnvMap map[string]string, envPrefix string, indentationStr string, blockIndentationRepeat int) (envLines []string) {
	blockIndentationStr := ""
	for i := 0; i < blockIndentationRepeat; i++ {
		blockIndentationStr += indentationStr
	}
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
		envLines = append(envLines, fmt.Sprintf("%s%s%s:", blockIndentationStr, indentationStr, envKey))
		envLines = append(envLines, fmt.Sprintf("%s%s%sfrom: %s", blockIndentationStr, indentationStr, indentationStr, envFrom))
		envLines = append(envLines, fmt.Sprintf("%s%s%sdefault: %s", blockIndentationStr, indentationStr, indentationStr, envVal))
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
