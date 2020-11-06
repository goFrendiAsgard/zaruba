package config

import "path/filepath"

// TaskData is struct sent to template
type TaskData struct {
	task        *Task
	Name        string
	ProjectName string
	BasePath    string
	WorkPath    string
	Kwargs      Dictionary
	Env         Dictionary
	Config      Dictionary
	LConfig     map[string][]string
}

// NewTaskData create new task data
func NewTaskData(task *Task) (td *TaskData) {
	return &TaskData{
		task:        task,
		Name:        task.Name,
		ProjectName: task.Project.GetName(),
		BasePath:    task.BasePath,
		WorkPath:    task.GetWorkPath(),
		Kwargs:      task.Project.Kwargs,
		Env:         task.ParsedEnv,
		Config:      task.ParsedConfig,
		LConfig:     task.ParsedLConfig,
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
