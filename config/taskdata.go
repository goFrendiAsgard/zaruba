package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
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
	Decoration   *output.Decoration
}

// NewTaskData create new task data
func NewTaskData(task *Task) (td *TaskData) {
	nextTask := *task
	nextTask.currentRecursiveLevel++
	return &TaskData{
		task:         &nextTask,
		Name:         task.GetName(),
		ProjectName:  task.Project.GetName(),
		BasePath:     task.GetBasePath(),
		WorkPath:     task.GetWorkPath(),
		DirPath:      filepath.Dir(task.GetFileLocation()),
		FileLocation: task.GetFileLocation(),
		Decoration:   task.Project.decoration,
	}
}

// GetConfig get config of task data
func (td *TaskData) GetConfig(keys ...string) (val string, err error) {
	return td.task.GetConfig(keys...)
}

// GetSubConfigKeys get config subkeys
func (td *TaskData) GetSubConfigKeys(parentKeys ...string) (subKeys []string) {
	configKeys := td.task.GetConfigKeys()
	return str.GetSubKeys(configKeys, parentKeys)
}

// GetLConfig get config of task data
func (td *TaskData) GetLConfig(keys ...string) (val []string, err error) {
	return td.task.GetLConfig(keys...)
}

// GetSubLConfigKeys get config subkeys
func (td *TaskData) GetSubLConfigKeys(parentKeys ...string) (subKeys []string) {
	lConfigKeys := td.task.GetLConfigKeys()
	return str.GetSubKeys(lConfigKeys, parentKeys)
}

// GetValue get keyword argument
func (td *TaskData) GetValue(keys ...string) (val string, err error) {
	return td.task.GetValue(keys...)
}

// GetSubValueKeys get keyword argument subkeys
func (td *TaskData) GetSubValueKeys(parentKeys ...string) (subKeys []string) {
	valueKeys := td.task.GetValueKeys()
	return str.GetSubKeys(valueKeys, parentKeys)
}

// GetEnv get environment
func (td *TaskData) GetEnv(key string) (val string, err error) {
	return td.task.GetEnv(key)
}

// GetEnvs get all environment
func (td *TaskData) GetEnvs() (parsedEnv map[string]string, err error) {
	return td.task.GetEnvs()
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
		return nil, fmt.Errorf("task '%s' is not exist", taskName)
	}
	return NewTaskData(task), nil
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

// ReadFile file
func (td *TaskData) ReadFile(filePath string) (fileContent string, err error) {
	absFilePath := td.GetWorkPath(filePath)
	fileContentB, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		return "", err
	}
	return string(fileContentB), err
}

// ListDir directory
func (td *TaskData) ListDir(dirPath string) (fileNames []string, err error) {
	absDirPath := td.GetWorkPath(dirPath)
	fileNames = []string{}
	files, err := ioutil.ReadDir(absDirPath)
	if err != nil {
		return fileNames, err
	}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames, nil
}

// ParseFile parse file
func (td *TaskData) ParseFile(filePath string) (parsedStr string, err error) {
	absFilePath := td.GetWorkPath(filePath)
	pattern, err := td.ReadFile(absFilePath)
	if err != nil {
		return "", err
	}
	templateName := fmt.Sprintf("File: %s", absFilePath)
	tmpl, err := template.New(templateName).Option("missingkey=zero").Parse(pattern)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err = tmpl.Execute(&b, td); err != nil {
		return "", err
	}
	return b.String(), nil
}

// ReplaceAllWith
func (td *TaskData) ReplaceAllWith(s string, replacements ...string) (result string) {
	return str.ReplaceAllWith(s, replacements...)
}

// EscapeShellValue
func (td *TaskData) EscapeShellValue(s string, quoteList ...string) (result string) {
	quote := "\""
	if len(quoteList) > 0 {
		quote = quoteList[0]
	}
	return str.EscapeShellValue(s, quote)
}

// DoubleQuoteShellValue
func (td *TaskData) DoubleQuoteShellValue(s string) (result string) {
	return str.DoubleQuoteShellValue(s)
}

// SingleQuoteShellValue
func (td *TaskData) SingleQuoteShellValue(s string) (result string) {
	return str.SingleQuoteShellValue(s)
}
