package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/logger"
)

func getSubKeys(dictionary map[string]string, parentKeys []string) (subKeys []string) {
	seen := map[string]bool{}
	parentKey := strings.Join(parentKeys, "::")
	prefixLength := len(parentKey) + len("::")
	subKeys = []string{}
	for key := range dictionary {
		if !strings.HasPrefix(key, parentKey+"::") {
			continue
		}
		childKey := key[prefixLength:]
		if childKey == "" {
			continue
		}
		childKeyParts := strings.SplitN(childKey, "::", 2)
		subkey := childKeyParts[0]
		seen[subkey] = true
	}
	for key := range seen {
		subKeys = append(subKeys, key)
	}
	return subKeys
}

// TaskData is struct sent to template
type TaskData struct {
	task         *Task
	Name         string
	ProjectName  string
	BasePath     string
	WorkPath     string
	DirPath      string
	FileLocation string
	Decoration   logger.Decoration
}

// NewTaskData create new task data
func NewTaskData(task *Task) (td *TaskData) {
	return &TaskData{
		task:         task,
		Name:         task.GetName(),
		ProjectName:  task.Project.GetName(),
		BasePath:     task.GetBasePath(),
		WorkPath:     task.GetWorkPath(),
		DirPath:      filepath.Dir(task.GetFileLocation()),
		FileLocation: task.GetFileLocation(),
		Decoration:   *logger.NewDecoration(),
	}
}

// GetConfig get config of task data
func (td *TaskData) GetConfig(keys ...string) (val string, err error) {
	return td.task.GetConfig(td, keys...)
}

// GetConfigs get all environment
func (td *TaskData) GetConfigs() (parsedEnv map[string]string, err error) {
	return td.task.GetConfigs(td)
}

// GetSubConfigKeys get config subkeys
func (td *TaskData) GetSubConfigKeys(keys ...string) (subKeys []string) {
	return getSubKeys(td.task.Config, keys)
}

// GetLConfig get config of task data
func (td *TaskData) GetLConfig(keys ...string) (val []string, err error) {
	return td.task.GetLConfig(td, keys...)
}

// GetLConfigs get all environment
func (td *TaskData) GetLConfigs() (parsedEnv map[string][]string, err error) {
	return td.task.GetLConfigs(td)
}

// GetValue get keyword argument
func (td *TaskData) GetValue(keys ...string) (val string, err error) {
	return td.task.GetValue(td, keys...)
}

// GetSubValueKeys get keyword argument subkeys
func (td *TaskData) GetSubValueKeys(keys ...string) (subKeys []string) {
	return getSubKeys(td.task.Project.values, keys)
}

// GetValues get all keyword arguments
func (td *TaskData) GetValues() (parsedEnv map[string]string, err error) {
	return td.task.GetValues(td)
}

// GetEnv get environment
func (td *TaskData) GetEnv(key string) (val string, err error) {
	return td.task.GetEnv(td, key)
}

// GetEnvs get all environment
func (td *TaskData) GetEnvs() (parsedEnv map[string]string, err error) {
	return td.task.GetEnvs(td)
}

// getAbsPath of any string
func (td *TaskData) getAbsPath(parentPath, path string) (absPath string) {
	if filepath.IsAbs(path) {
		return path
	}
	absParentPath, err := filepath.Abs(parentPath)
	if err != nil {
		absParentPath = parentPath
	}
	return filepath.Join(absParentPath, path)
}

// GetWorkPath get workPath (path relative to task.location)
func (td *TaskData) GetWorkPath(path string) (absPath string) {
	return td.getAbsPath(td.WorkPath, path)
}

// GetBasePath get basePath (path relative to main yaml's directory)
func (td *TaskData) GetBasePath(path string) (absPath string) {
	return td.getAbsPath(td.BasePath, path)
}

// GetRelativePath get basePath (path relateive to task's definition directory)
func (td *TaskData) GetRelativePath(path string) (absPath string) {
	return td.getAbsPath(td.DirPath, path)
}

// GetTask get other task
func (td *TaskData) GetTask(taskName string) (otherTd *TaskData, err error) {
	task, taskFound := td.task.Project.Tasks[taskName]
	if !taskFound {
		return nil, fmt.Errorf("Task %s is not exist", taskName)
	}
	return NewTaskData(task), nil
}

// GetDefaultShell get default shell
func (td *TaskData) GetDefaultShell() (shell string) {
	if _, err := os.Stat("/usr/bin/bash"); !os.IsNotExist(err) {
		return "bash"
	}
	if _, err := os.Stat("/bin/bash"); !os.IsNotExist(err) {
		return "bash"
	}
	return "sh"
}

// IsTrue check if string represent "true"
func (td *TaskData) IsTrue(str string) (isTrue bool) {
	return boolean.IsTrue(str)
}

// IsFalse check if string represent "false"
func (td *TaskData) IsFalse(str string) (isFalse bool) {
	return boolean.IsFalse(str)
}

// Trim trim string
func (td *TaskData) Trim(str, cutset string) (trimmedStr string) {
	return strings.Trim(str, cutset)
}

// GetDockerImagePrefix get docker image prefix
func (td *TaskData) GetDockerImagePrefix() (dockerImagePrefix string) {
	// if useImagePrefix is false
	useImagePrefix, _ := td.GetConfig("useImagePrefix")
	if useImagePrefix != "" && td.IsFalse(useImagePrefix) {
		return ""
	}
	// Try to get prefix from dockerEnv config, docker.env value, or "default"
	dockerEnvConfig, _ := td.GetConfig("dockerEnv")
	dockerEnvValue, _ := td.GetValue("docker.env")
	for _, dockerEnv := range []string{dockerEnvConfig, dockerEnvValue, "default"} {
		if dockerEnv == "" {
			continue
		}
		if dockerImagePrefix, _ := td.GetValue("dockerImagePrefix", dockerEnv); dockerImagePrefix != "" {
			return fmt.Sprintf("%s/", dockerImagePrefix)
		}
		return "local/"
	}
	return ""
}
