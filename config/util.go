package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func GetTaskServiceName(taskName string) (serviceName string) {
	if strings.HasPrefix(taskName, "run") && taskName != "run" {
		upperServiceName := strings.TrimPrefix(taskName, "run")
		return strings.ToLower(string(upperServiceName[0])) + upperServiceName[1:]
	}
	return taskName
}

func GetTaskEnvRefname(task *Task) (envRefName string) {
	if task.EnvRef != "" {
		return task.EnvRef
	}
	if len(task.EnvRefs) > 0 {
		return task.EnvRefs[0]
	}
	return ""
}

func GetDefaultServiceName(location string) (serviceName string, err error) {
	absPath, err := filepath.Abs(location)
	if err != nil {
		return "", err
	}
	baseName := filepath.Base(absPath)
	pattern := regexp.MustCompile(`[^A-Za-z0-9]`)
	spacedBaseName := (pattern.ReplaceAllString(baseName, " "))
	titledBaseName := strings.Title(spacedBaseName)
	serviceName = strings.ReplaceAll(titledBaseName, " ", "")
	if len(serviceName) > 0 {
		bts := []byte(serviceName)
		lc := bytes.ToLower([]byte{bts[0]})
		rest := bts[1:]
		serviceName = string(bytes.Join([][]byte{lc, rest}, nil))
	}
	return serviceName, err
}

func SetProjectValue(fileName, key, value string) (err error) {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if value == "" {
		return fmt.Errorf("value cannot be empty")
	}
	fileContentB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	configMap := map[string]string{}
	if err := yaml.Unmarshal(fileContentB, &configMap); err != nil {
		return err
	}
	configMap[key] = value
	newFileContentB, err := yaml.Marshal(configMap)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, newFileContentB, 0755)
}
