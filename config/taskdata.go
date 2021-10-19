package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/utility"
	yaml "gopkg.in/yaml.v3"
)

type TaskData struct {
	task           *Task
	ZarubaHome     string
	ZarubaBin      string
	Name           string
	ProjectName    string
	UUID           string
	WorkDirPath    string
	TaskDirPath    string
	ProjectDirPath string
	FileLocation   string
	Decoration     *output.Decoration
	Util           *utility.Util
}

func NewTaskData(task *Task) (td *TaskData) {
	zarubaHome := os.Getenv("ZARUBA_HOME")
	if zarubaHome == "" {
		executable, _ := os.Executable()
		zarubaHome = filepath.Dir(executable)
	}
	nextTask := *task
	nextTask.currentRecursiveLevel++
	return &TaskData{
		task:           &nextTask,
		ZarubaHome:     zarubaHome,
		ZarubaBin:      filepath.Join(zarubaHome, "zaruba"),
		Name:           task.GetName(),
		ProjectName:    task.Project.GetName(),
		UUID:           task.GetUUID(),
		WorkDirPath:    task.GetWorkPath(),
		TaskDirPath:    filepath.Dir(task.GetFileLocation()),
		ProjectDirPath: filepath.Dir(task.Project.GetFileLocation()),
		FileLocation:   task.GetFileLocation(),
		Decoration:     task.Project.Decoration,
		Util:           task.Project.Util,
	}
}

func (td *TaskData) GetWorkPath(path string) (absPath string) {
	return td.getAbsPath(td.WorkDirPath, path)
}

func (td *TaskData) GetTaskPath(path string) (absPath string) {
	return td.getAbsPath(td.TaskDirPath, path)
}

func (td *TaskData) GetProjectPath(path string) (absPath string) {
	return td.getAbsPath(td.ProjectDirPath, path)
}

func (td *TaskData) GetConfig(keys ...string) (val string, err error) {
	return td.task.GetConfig(keys...)
}

func (td *TaskData) GetPorts() []int {
	ports := []int{}
	portStr, _ := td.GetConfig("ports")
	lines := strings.Split(portStr, "\n")
	for _, line := range lines {
		line = strings.Trim(line, " \"'")
		if line == "" {
			continue
		}
		portParts := strings.Split(line, ":")
		if len(portParts) > 1 {
			port, _ := strconv.Atoi(portParts[1])
			ports = append(ports, port)
			continue
		}
		port, _ := strconv.Atoi(portParts[0])
		ports = append(ports, port)
	}
	return ports
}

func (td *TaskData) GetSubConfigKeys(parentKeys ...string) (subKeys []string) {
	configKeys := td.task.GetConfigKeys()
	return td.task.Project.Util.Str.GetSubKeys(configKeys, parentKeys)
}

func (td *TaskData) GetValue(keys ...string) (val string, err error) {
	return td.task.GetValue(keys...)
}

func (td *TaskData) GetSubValueKeys(parentKeys ...string) (subKeys []string) {
	valueKeys := td.task.GetValueKeys()
	return td.task.Project.Util.Str.GetSubKeys(valueKeys, parentKeys)
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

func (td *TaskData) GetDockerImageName() string {
	dockerImagePrefix := ""
	useImagePrefix, _ := td.GetConfig("useImagePrefix")
	if boolean.IsTrue(useImagePrefix) {
		dockerImagePrefix, _ = td.GetConfig("imagePrefix")
	}
	dockerImageName, _ := td.GetConfig("imageName")
	if dockerImageName == "" {
		defaultServiceName, _ := GetDefaultServiceName(td.TaskDirPath)
		dockerImageName = td.task.Project.Util.Str.ToKebab(defaultServiceName)
	}
	if dockerImagePrefix == "" {
		return dockerImageName
	}
	return fmt.Sprintf("%s/%s", dockerImagePrefix, dockerImageName)
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

func (t *TaskData) IsFileExist(filePath string) (exist bool, err error) {
	if _, statErr := os.Stat(filePath); statErr == nil {
		return true, nil
	} else if os.IsNotExist(statErr) {
		return false, nil
	} else {
		return false, statErr
	}
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
	absFilePath := td.GetWorkPath(filePath)
	return file.WriteText(absFilePath, content, 0755)
}

func (td *TaskData) Template(content string) (escapedStr string) {
	return fmt.Sprintf("{{ %s }}", content)
}

func (td *TaskData) getAbsPath(parentPath, path string) (absPath string) {
	if filepath.IsAbs(path) {
		return path
	}
	absParentPath, _ := filepath.Abs(parentPath)
	return filepath.Join(absParentPath, path)
}
