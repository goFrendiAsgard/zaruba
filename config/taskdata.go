package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/google/uuid"
	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
	yaml "gopkg.in/yaml.v3"
)

type TaskData struct {
	task         *Task
	Zaruba       string
	Name         string
	ProjectName  string
	WorkPath     string
	DirPath      string
	FileLocation string
	Decoration   *output.Decoration
}

func NewTaskData(task *Task) (td *TaskData) {
	nextTask := *task
	nextTask.currentRecursiveLevel++
	return &TaskData{
		task:         &nextTask,
		Zaruba:       "\"${ZARUBA_HOME}/zaruba\"",
		Name:         task.GetName(),
		ProjectName:  task.Project.GetName(),
		WorkPath:     task.GetWorkPath(),
		DirPath:      filepath.Dir(task.GetFileLocation()),
		FileLocation: task.GetFileLocation(),
		Decoration:   task.Project.decoration,
	}
}

func (td *TaskData) GetWorkPath(path string) (absPath string) {
	return td.getAbsPath(td.WorkPath, path)
}

func (td *TaskData) GetRelativePath(path string) (absPath string) {
	return td.getAbsPath(td.DirPath, path)
}

func (td *TaskData) GetConfig(keys ...string) (val string, err error) {
	return td.task.GetConfig(keys...)
}

func (td *TaskData) GetSubConfigKeys(parentKeys ...string) (subKeys []string) {
	configKeys := td.task.GetConfigKeys()
	return str.GetSubKeys(configKeys, parentKeys)
}

func (td *TaskData) GetValue(keys ...string) (val string, err error) {
	return td.task.GetValue(keys...)
}

func (td *TaskData) GetSubValueKeys(parentKeys ...string) (subKeys []string) {
	valueKeys := td.task.GetValueKeys()
	return str.GetSubKeys(valueKeys, parentKeys)
}

func (td *TaskData) GetEnv(key string) (val string, err error) {
	return td.task.GetEnv(key)
}

func (td *TaskData) GetEnvs() (parsedEnv map[string]string, err error) {
	return td.task.GetEnvs()
}

func (td *TaskData) IsTrue(str string) (isTrue bool) {
	return boolean.IsTrue(str)
}

func (td *TaskData) IsFalse(str string) (isFalse bool) {
	return boolean.IsFalse(str)
}

func (td *TaskData) ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func (td *TaskData) EscapeShellArg(s string) (result string) {
	return str.EscapeShellArg(s)
}

func (td *TaskData) Indent(multiLineStr string, indentation string) (result string) {
	return str.Indent(multiLineStr, indentation)
}

func (td *TaskData) GetNewUUID() string {
	return uuid.NewString()
}

func (td *TaskData) Split(s, sep string) []string {
	return strings.Split(s, sep)
}

func (td *TaskData) Join(sep string, a []string) (string, error) {
	return strings.Join(a, sep), nil
}

func (td *TaskData) Trim(str, cutset string) (trimmedStr string) {
	return strings.Trim(str, cutset)
}

func (td *TaskData) ParseJSON(s string) (interface{}, error) {
	if s == "" {
		return make([]interface{}, 0), nil
	}
	var data interface{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (td *TaskData) ParseYAML(s string) (interface{}, error) {
	if s == "" {
		return make([]interface{}, 0), nil
	}
	var data interface{}
	if err := yaml.Unmarshal([]byte(s), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (td *TaskData) ReadFile(filePath string) (fileContent string, err error) {
	absFilePath := td.GetWorkPath(filePath)
	return file.ReadText(absFilePath)
}

func (td *TaskData) ListDir(dirPath string) (fileNames []string, err error) {
	absDirPath := td.GetWorkPath(dirPath)
	return file.ListDir(absDirPath)
}

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

func (td *TaskData) WriteFile(filePath string, content string) (err error) {
	return file.WriteText(filePath, content, 0755)
}

func (td *TaskData) getAbsPath(parentPath, path string) (absPath string) {
	if filepath.IsAbs(path) {
		return path
	}
	absParentPath, _ := filepath.Abs(parentPath)
	return filepath.Join(absParentPath, path)
}
