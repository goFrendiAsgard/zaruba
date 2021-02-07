package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/logger"
)

func getSubKeys(dictionary map[string]string, parentKeys []string) (subKeys []string) {
	seen := map[string]bool{}
	parentKey := strings.Join(parentKeys, "::")
	prefixLength := len(parentKey) + len("::")
	subKeys = []string{}
	for key := range dictionary {
		if !strings.HasPrefix(key, parentKey) {
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
		Name:         task.Name,
		ProjectName:  task.Project.GetName(),
		BasePath:     task.BasePath,
		WorkPath:     task.GetWorkPath(),
		DirPath:      filepath.Dir(task.FileLocation),
		FileLocation: task.FileLocation,
		Decoration:   *logger.NewDecoration(),
	}
}

// GetConfig get config of task data
func (td *TaskData) GetConfig(keys ...string) (val string, err error) {
	return td.task.GetConfig(td, keys...)
}

// GetAllConfig get all environment
func (td *TaskData) GetAllConfig() (parsedEnv map[string]string, err error) {
	return td.task.GetAllConfig(td)
}

// GetConfigSubKeys get config subkeys
func (td *TaskData) GetConfigSubKeys(keys ...string) (subKeys []string) {
	return getSubKeys(td.task.Config, keys)
}

// GetLConfig get config of task data
func (td *TaskData) GetLConfig(keys ...string) (val []string, err error) {
	return td.task.GetLConfig(td, keys...)
}

// GetAllLConfig get all environment
func (td *TaskData) GetAllLConfig() (parsedEnv map[string][]string, err error) {
	return td.task.GetAllLConfig(td)
}

// GetKwarg get keyword argument
func (td *TaskData) GetKwarg(keys ...string) (val string, err error) {
	return td.task.GetKwarg(td, keys...)
}

// GetKwargSubKeys get keyword argument subkeys
func (td *TaskData) GetKwargSubKeys(keys ...string) (subKeys []string) {
	return getSubKeys(td.task.Project.Kwargs, keys)
}

// GetAllKwarg get all keyword arguments
func (td *TaskData) GetAllKwarg() (parsedEnv map[string]string, err error) {
	return td.task.GetAllKwarg(td)
}

// GetEnv get environment
func (td *TaskData) GetEnv(key string) (val string, err error) {
	return td.task.GetEnv(td, key)
}

// GetAllEnv get all environment
func (td *TaskData) GetAllEnv() (parsedEnv map[string]string, err error) {
	return td.task.GetAllEnv(td)
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

// GetWorkPath get workPath
func (td *TaskData) GetWorkPath(path string) (absPath string) {
	return td.getAbsPath(td.WorkPath, path)
}

// GetBasePath get basePath
func (td *TaskData) GetBasePath(path string) (absPath string) {
	return td.getAbsPath(td.BasePath, path)
}

// GetRelativePath get basePath
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
