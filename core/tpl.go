package core

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/state-alchemists/zaruba/output"
)

type Tpl struct {
	task                *Task
	ZarubaHome          string
	ZarubaBin           string
	Name                string
	ProjectName         string
	UUID                string
	GeneratedRandomName string
	WorkDirPath         string
	TaskDirPath         string
	ProjectDirPath      string
	FileLocation        string
	Decoration          *output.Decoration
	Util                *CoreUtil
}

func NewTpl(task *Task) (td *Tpl) {
	zarubaHome := os.Getenv("ZARUBA_HOME")
	if zarubaHome == "" {
		executable, _ := os.Executable()
		zarubaHome = filepath.Dir(executable)
	}
	nextTask := *task
	nextTask.currentRecursiveLevel++
	return &Tpl{
		task:                &nextTask,
		ZarubaHome:          zarubaHome,
		ZarubaBin:           filepath.Join(zarubaHome, "zaruba"),
		Name:                task.GetName(),
		ProjectName:         task.Project.GetName(),
		UUID:                task.GetUUID(),
		GeneratedRandomName: task.GetGeneratedRandomName(),
		WorkDirPath:         task.GetWorkPath(),
		TaskDirPath:         filepath.Dir(task.GetFileLocation()),
		ProjectDirPath:      filepath.Dir(task.Project.GetFileLocation()),
		FileLocation:        task.GetFileLocation(),
		Decoration:          task.Project.Decoration,
		Util:                task.Project.Util,
	}
}

func (tpl *Tpl) GetWorkPath(path string) (absPath string) {
	return tpl.getAbsPath(tpl.WorkDirPath, path)
}

func (tpl *Tpl) GetTaskPath(path string) (absPath string) {
	return tpl.getAbsPath(tpl.TaskDirPath, path)
}

func (tpl *Tpl) GetProjectPath(path string) (absPath string) {
	return tpl.getAbsPath(tpl.ProjectDirPath, path)
}

func (tpl *Tpl) GetConfig(key string) (val string, err error) {
	return tpl.task.GetConfig(key)
}

func (tpl *Tpl) GetConfigs(keyPattern string) (parsedConfig map[string]string, err error) {
	return tpl.task.GetConfigs(keyPattern)
}

func (tpl *Tpl) GetPorts() []int {
	ports := []int{}
	portStr, _ := tpl.GetConfig("ports")
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

func (tpl *Tpl) GetSubValueKeys(parentKeys ...string) (subKeys []string) {
	keys := tpl.task.GetValueKeys()
	seen := map[string]bool{}
	parentKey := strings.Join(parentKeys, "::")
	prefixLength := len(parentKey) + len("::")
	subKeys = []string{}
	for _, key := range keys {
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

func (tpl *Tpl) GetValue(keys ...string) (val string, err error) {
	return tpl.task.GetValue(keys...)
}

func (tpl *Tpl) GetEnv(key string) (val string, err error) {
	return tpl.task.GetEnv(key)
}

func (tpl *Tpl) GetEnvs() (parsedEnv map[string]string, err error) {
	return tpl.task.GetEnvs()
}

func (tpl *Tpl) ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func (tpl *Tpl) GetDockerImageName() string {
	dockerImagePrefix := ""
	useImagePrefix, _ := tpl.GetConfig("useImagePrefix")
	if tpl.Util.Bool.IsTrue(useImagePrefix) {
		dockerImagePrefix, _ = tpl.GetConfig("imagePrefix")
	}
	dockerImageName, _ := tpl.GetConfig("imageName")
	if dockerImageName == "" {
		defaultServiceName, _ := tpl.Util.Path.GetDefaultAppName(tpl.TaskDirPath)
		dockerImageName = tpl.task.Project.Util.Str.ToKebab(defaultServiceName)
	}
	if dockerImagePrefix == "" {
		return dockerImageName
	}
	return fmt.Sprintf("%s/%s", dockerImagePrefix, dockerImageName)
}

func (tpl *Tpl) ParseFile(filePath string) (parsedStr string, err error) {
	absFilePath := tpl.GetWorkPath(filePath)
	pattern, err := tpl.Util.File.ReadText(absFilePath)
	if err != nil {
		return "", err
	}
	templateName := fmt.Sprintf("File: %s", absFilePath)
	tmpl, err := template.New(templateName).Option("missingkey=zero").Parse(pattern)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err = tmpl.Execute(&b, tpl); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (tpl *Tpl) Template(content string) (escapedStr string) {
	return fmt.Sprintf("{{ %s }}", content)
}

func (tpl *Tpl) getAbsPath(parentPath, path string) (absPath string) {
	if filepath.IsAbs(path) {
		return path
	}
	absParentPath, _ := filepath.Abs(parentPath)
	return filepath.Join(absParentPath, path)
}
