package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/conversion"
	"github.com/state-alchemists/zaruba/monitor"
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
	Decoration   *monitor.Decoration
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
		Decoration:   task.Project.decoration,
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
func (td *TaskData) GetSubConfigKeys(keys ...string) (subKeys []string, err error) {
	configs, err := td.GetConfigs()
	if err != nil {
		return subKeys, err
	}
	return str.GetSubKeys(conversion.NormalizeMapStringValue(configs), keys), nil
}

// GetLConfig get config of task data
func (td *TaskData) GetLConfig(keys ...string) (val []string, err error) {
	return td.task.GetLConfig(td, keys...)
}

// GetLConfigs get all environment
func (td *TaskData) GetLConfigs() (parsedEnv map[string][]string, err error) {
	return td.task.GetLConfigs(td)
}

// GetSubLConfigKeys get config subkeys
func (td *TaskData) GetSubLConfigKeys(keys ...string) (subKeys []string, err error) {
	lConfigs, err := td.GetLConfigs()
	if err != nil {
		return subKeys, err
	}
	return str.GetSubKeys(conversion.NormalizeMapListStringValue(lConfigs), keys), nil
}

// GetValue get keyword argument
func (td *TaskData) GetValue(keys ...string) (val string, err error) {
	return td.task.GetValue(td, keys...)
}

// GetValues get all keyword arguments
func (td *TaskData) GetValues() (parsedEnv map[string]string, err error) {
	return td.task.GetValues(td)
}

// GetSubValueKeys get keyword argument subkeys
func (td *TaskData) GetSubValueKeys(keys ...string) (subKeys []string, err error) {
	values, err := td.GetValues()
	if err != nil {
		return subKeys, err
	}
	return str.GetSubKeys(conversion.NormalizeMapStringValue(values), keys), err
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
