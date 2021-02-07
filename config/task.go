package config

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"text/template"
	"time"

	"github.com/state-alchemists/zaruba/logger"
)

// Task is zaruba task
type Task struct {
	Start           []string              `yaml:"start,omitempty"`
	Check           []string              `yaml:"check,omitempty"`
	Timeout         string                `yaml:"timeout,omitempty"`
	Private         bool                  `yaml:"private,omitempty"`
	Extend          string                `yaml:"extend,omitempty"`
	Location        string                `yaml:"location,omitempty"`
	Config          map[string]string     `yaml:"config,omitempty"`
	LConfig         map[string][]string   `yaml:"lconfig,omitempty"`
	Env             map[string]*EnvConfig `yaml:"env,omitempty"`
	Dependencies    []string              `yaml:"dependencies,omitempty"`
	Description     string                `yaml:"description,omitempty"`
	Icon            string                `yaml:"icon,omitempty"`
	Logless         bool                  `yaml:"logless,omitempty"`
	BasePath        string                // Main yaml's location
	FileLocation    string                // File location where this task was declared
	Project         *ProjectConfig
	Name            string
	FunkyName       string
	TimeoutDuration time.Duration
}

// GetWorkPath get path of current task
func (task *Task) GetWorkPath() (path string) {
	path = task.getPath()
	if path != "" {
		return path
	}
	path, err := os.Getwd()
	if err != nil {
		return task.BasePath
	}
	return path
}

// GetKwarg getting config of a task
func (task *Task) GetKwarg(taskData *TaskData, keys ...string) (val string, err error) {
	key := strings.Join(keys, "::")
	pattern, exist := task.Project.Kwargs[key]
	if !exist {
		return "", nil
	}
	templateName := fmt.Sprintf("%s.kwargs.%s", task.Name, key)
	return task.getParsedPattern(taskData, templateName, pattern)
}

// GetAllKwarg getting all parsed env
func (task *Task) GetAllKwarg(taskData *TaskData) (parsedKwarg map[string]string, err error) {
	parsedKwarg = map[string]string{}
	for key := range taskData.task.Project.Kwargs {
		parsedKwarg[key], err = task.GetKwarg(taskData, key)
		if err != nil {
			return parsedKwarg, err
		}
	}
	return parsedKwarg, nil
}

// GetConfig getting config of a task
func (task *Task) GetConfig(taskData *TaskData, keys ...string) (val string, err error) {
	key := strings.Join(keys, "::")
	pattern, exist := task.Config[key]
	if !exist {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return "", nil
		}
		return parentTask.GetConfig(taskData, keys...)
	}
	templateName := fmt.Sprintf("%s.config.%s", task.Name, key)
	return task.getParsedPattern(taskData, templateName, pattern)
}

// GetAllConfig getting all parsed env
func (task *Task) GetAllConfig(taskData *TaskData) (parsedConfig map[string]string, err error) {
	parsedConfig = map[string]string{}
	for key := range taskData.task.Config {
		parsedConfig[key], err = task.GetConfig(taskData, key)
		if err != nil {
			return parsedConfig, err
		}
	}
	return parsedConfig, nil
}

// GetLConfig getting lconfig of a task
func (task *Task) GetLConfig(taskData *TaskData, keys ...string) (val []string, err error) {
	key := strings.Join(keys, "::")
	val = []string{}
	lConfig, exist := task.LConfig[key]
	if !exist {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return []string{}, nil
		}
		return parentTask.GetLConfig(taskData, keys...)
	}
	for index, pattern := range lConfig {
		templateName := fmt.Sprintf("%s.lconfig.%s[%d]", task.Name, key, index)
		element, err := task.getParsedPattern(taskData, templateName, pattern)
		if err != nil {
			return val, err
		}
		val = append(val, element)
	}
	return val, err
}

// GetAllLConfig getting all lConfig
func (task *Task) GetAllLConfig(taskData *TaskData) (parsedLConfig map[string][]string, err error) {
	parsedLConfig = map[string][]string{}
	for key := range taskData.task.LConfig {
		parsedLConfig[key], err = task.GetLConfig(taskData, key)
		if err != nil {
			return parsedLConfig, err
		}
	}
	return parsedLConfig, nil
}

// GetEnv getting env of a task
func (task *Task) GetEnv(taskData *TaskData, key string) (val string, err error) {
	envConfig, exist := task.Env[key]
	if !exist {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return os.Getenv(key), nil
		}
		return parentTask.GetEnv(taskData, key)
	}
	if envConfig.From != "" {
		if val = os.Getenv(envConfig.From); val != "" {
			return val, nil
		}
	}
	templateName := fmt.Sprintf("%s.env.%s", task.Name, key)
	return task.getParsedPattern(taskData, templateName, envConfig.Default)
}

// GetAllEnv getting all parsed env
func (task *Task) GetAllEnv(taskData *TaskData) (parsedEnv map[string]string, err error) {
	parsedEnv = map[string]string{}
	for key := range taskData.task.Env {
		parsedEnv[key], err = task.GetEnv(taskData, key)
		if err != nil {
			return parsedEnv, err
		}
	}
	return parsedEnv, nil
}

func (task *Task) getParsedPattern(taskData *TaskData, templateName, pattern string) (result string, err error) {
	lines := strings.Split(pattern, "\n")
	for index, line := range lines {
		lines[index] = fmt.Sprintf("%s | %s", fmt.Sprintf("%4d", index+1), line)
	}
	tmplName := fmt.Sprintf("\n%s:\n%s\n", templateName, strings.Join(lines, "\n"))
	tmpl, err := template.New(tmplName).Option("missingkey=zero").Parse(pattern)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err = tmpl.Execute(&b, taskData); err != nil {
		return "", err
	}
	result = b.String()
	return result, nil
}

func (task *Task) fillEnvTask() {
	for _, env := range task.Env {
		env.Task = task
	}
}

func (task *Task) init() (err error) {
	var timeErr error
	task.TimeoutDuration, timeErr = time.ParseDuration(task.Timeout)
	if timeErr != nil || task.TimeoutDuration <= 0 {
		task.TimeoutDuration = 5 * time.Minute
	}
	if task.Extend != "" {
		if _, exists := task.Project.Tasks[task.Extend]; !exists {
			return fmt.Errorf("Task %s is extended from %s but task %s doesn't exist", task.Name, task.Extend, task.Extend)
		}
	}
	task.generateIcon()
	task.generateFunkyName()
	return nil
}

func (task *Task) generateIcon() {
	if task.Icon == "" {
		icon := task.Project.IconGenerator.Create()
		task.Icon = icon
	}
}

func (task *Task) generateFunkyName() {
	repeat := 2 + task.Project.MaxPublishedTaskNameLength - len(task.Name) - len(task.Icon)
	if repeat < 0 {
		repeat = 0
	}
	paddedStr := strings.Repeat(" ", repeat)
	d := task.Project.Decoration
	paddedName := fmt.Sprintf("%s%s%s%s", d.GenerateColor(), task.Name, d.Normal, paddedStr)
	task.FunkyName = fmt.Sprintf("%s %s%s%s", paddedName, d.Faint, task.Icon, d.Normal)
}

func (task *Task) getPath() (path string) {
	if task.Location != "" {
		return filepath.Join(task.BasePath, task.Location)
	}
	if parentTask, exists := task.Project.Tasks[task.Extend]; exists {
		return parentTask.getPath()
	}
	return ""
}

// GetDependencies get unique dependencies of a task, recursively
func (task *Task) GetDependencies() (dependencies []string) {
	dependencies = task.getDependencies()
	sort.Strings(dependencies)
	return dependencies
}

func (task *Task) getDependencies() (dependencies []string) {
	dependencies = []string{}
	seen := map[string]bool{}
	for _, dependency := range task.Dependencies {
		if _, exist := seen[dependency]; exist {
			continue
		}
		seen[dependency] = true
		dependencies = append(dependencies, dependency)
	}
	if task.Extend == "" {
		return dependencies
	}
	// get parent's
	for _, dependency := range task.Project.Tasks[task.Extend].getDependencies() {
		if _, exist := seen[dependency]; exist {
			continue
		}
		seen[dependency] = true
		dependencies = append(dependencies, dependency)
	}
	return dependencies
}

// GetStartCmd get start command of a task
func (task *Task) GetStartCmd(logDone chan error) (cmd *exec.Cmd, exist bool, err error) {
	return task.getStartCmd(NewTaskData(task), logDone)
}

func (task *Task) getStartCmd(taskData *TaskData, logDone chan error) (cmd *exec.Cmd, exist bool, err error) {
	if len(task.Start) == 0 {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return cmd, false, fmt.Errorf("Cannot retrieve StartCmd from %s's parent", task.Name)
		}
		return parentTask.getStartCmd(taskData, logDone)
	}
	cmd, err = task.getCmd(taskData, "START", task.Start, logDone)
	return cmd, true, err
}

// GetCheckCmd get check command of a task
func (task *Task) GetCheckCmd(logDone chan error) (cmd *exec.Cmd, exist bool, err error) {
	return task.getCheckCmd(NewTaskData(task), logDone)
}

func (task *Task) getCheckCmd(taskData *TaskData, logDone chan error) (cmd *exec.Cmd, exist bool, err error) {
	if len(task.Check) == 0 {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return cmd, false, fmt.Errorf("Cannot retrieve CheckCmd from %s's parent", task.Name)
		}
		return parentTask.getCheckCmd(taskData, logDone)
	}
	cmd, err = task.getCmd(taskData, "CHECK", task.Check, logDone)
	return cmd, true, err
}

func (task *Task) getCmd(taskData *TaskData, cmdType string, commandPatternArgs []string, logDone chan error) (cmd *exec.Cmd, err error) {
	commandArgs := []string{}
	templateName := fmt.Sprintf("%s.%s", task.Name, strings.ToLower(cmdType))
	for _, pattern := range commandPatternArgs {
		arg, err := task.getParsedPattern(taskData, templateName, pattern)
		if err != nil {
			return cmd, err
		}
		commandArgs = append(commandArgs, arg)
	}
	name := commandArgs[0]
	args := commandArgs[1:]
	cmd = exec.Command(name, args...)
	cmd.Dir = taskData.task.GetWorkPath()
	cmd.Env = os.Environ()
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	for key := range taskData.task.Env {
		val, err := taskData.GetEnv(key)
		if err != nil {
			return cmd, err
		}
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, val))
	}
	// log stdout
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return cmd, err
	}
	outDone := make(chan error)
	go task.log(taskData, cmdType, "OUT", outPipe, outDone)
	// log stderr
	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return cmd, err
	}
	errDone := make(chan error)
	go task.log(taskData, cmdType, "ERR", errPipe, errDone)
	// combine stdout and stderr done
	go task.combineLogDone(outDone, errDone, logDone)
	return cmd, err
}

func (task *Task) combineLogDone(outDone, errDone, logDone chan error) {
	errErr := <-errDone
	outErr := <-outDone
	if outErr != nil {
		logDone <- outErr
		return
	}
	logDone <- errErr
}

func (task *Task) log(taskData *TaskData, cmdType, logType string, pipe io.ReadCloser, logDone chan error) {
	buf := bufio.NewScanner(pipe)
	d := task.Project.Decoration
	cmdIconType := task.getCmdIconType(cmdType)
	prefix := fmt.Sprintf("  %s%s%s %s", d.Faint, cmdIconType, d.Normal, taskData.task.FunkyName)
	logless := taskData.task.Logless
	print := logger.Printf
	if logType == "ERR" {
		print = logger.PrintfError
	}
	var err error = nil
	for buf.Scan() {
		content := buf.Text()
		print("%s %s\n", prefix, content)
		if !logless {
			if csvWriteErr := task.Project.CSVLogWriter.Log(logType, cmdType, taskData.Name, content, taskData.task.FunkyName); csvWriteErr != nil {
				err = csvWriteErr
			}
		}
	}
	logDone <- err
}

func (task *Task) getCmdIconType(cmdType string) string {
	if cmdType == "CHECK" {
		return "🔎"
	}
	return "🚀"
}
