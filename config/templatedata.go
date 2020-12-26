package config

import (
	"fmt"
	"path/filepath"

	"github.com/state-alchemists/zaruba/logger"
)

// TaskData is struct sent to template
type TaskData struct {
	task         *Task
	Name         string
	ProjectName  string
	BasePath     string
	WorkPath     string
	DirPath      string
	FileLocation string
	Kwargs       Dictionary
	Env          Dictionary
	Config       Dictionary
	LConfig      map[string][]string
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
		Kwargs:       task.Project.Kwargs,
		Env:          task.ParsedEnv,
		Config:       task.ParsedConfig,
		LConfig:      task.ParsedLConfig,
		Decoration:   *logger.NewDecoration(),
	}
}

// GetEnv of TaskData
func (td *TaskData) GetEnv(key string) (val string) {
	return td.task.GetEnv(key)
}

// GetAbsPath of any string
func (td *TaskData) GetAbsPath(parentPath, path string) (absPath string) {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(parentPath, path)
}

// GetTaskConfig get config from other task
func (td *TaskData) GetTaskConfig(taskName, config string) (value string, err error) {
	task, taskFound := td.task.Project.Tasks[taskName]
	if !taskFound {
		return "", fmt.Errorf("Task %s is not exist", taskName)
	}
	parsedConfig, err := task.getParsedConfig()
	if err != nil {
		return "", err
	}
	value, configFound := parsedConfig[config]
	if !configFound {
		return "", fmt.Errorf("Config %s on task %s is not exist", config, taskName)
	}
	return value, nil
}
