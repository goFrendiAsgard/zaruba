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

	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/logger"
)

// Task is zaruba task
type Task struct {
	Start           []string            `yaml:"start,omitempty"`
	Check           []string            `yaml:"check,omitempty"`
	Timeout         string              `yaml:"timeout,omitempty"`
	Private         bool                `yaml:"private,omitempty"`
	Extend          string              `yaml:"extend,omitempty"`
	Location        string              `yaml:"location,omitempty"`
	Config          map[string]string   `yaml:"config,omitempty"`
	LConfig         map[string][]string `yaml:"lconfig,omitempty"`
	Env             map[string]*Env     `yaml:"env,omitempty"`
	Dependencies    []string            `yaml:"dependencies,omitempty"`
	Inputs          []string            `yaml:"inputs,omitempty"`
	Description     string              `yaml:"description,omitempty"`
	Icon            string              `yaml:"icon,omitempty"`
	SaveLog         string              `yaml:"saveLog,omitempty"`
	basePath        string              // Main yaml's location
	fileLocation    string              // File location where this task was declared
	Project         *Project
	name            string
	logPrefix       string
	timeoutDuration time.Duration
}

// GetName get task name
func (task *Task) GetName() (name string) {
	return task.name
}

// GetTimeoutDuration get timeout duration of a task
func (task *Task) GetTimeoutDuration() time.Duration {
	return task.timeoutDuration
}

// GetBasePath get file location of a task
func (task *Task) GetBasePath() (basePath string) {
	return task.basePath
}

// GetFileLocation get file location of a task
func (task *Task) GetFileLocation() (fileLocation string) {
	return task.fileLocation
}

// GetWorkPath get path of current task
func (task *Task) GetWorkPath() (path string) {
	path = task.getPath()
	if path != "" {
		return path
	}
	path, err := os.Getwd()
	if err != nil {
		return task.basePath
	}
	return path
}

// GetValue getting config of a task
func (task *Task) GetValue(td *TaskData, keys ...string) (val string, err error) {
	key := strings.Join(keys, "::")
	pattern, exist := task.Project.GetValue(key), task.Project.IsValueExist(key)
	if !exist {
		return "", nil
	}
	templateName := fmt.Sprintf("%s.values.%s", task.GetName(), key)
	return task.getParsedPattern(td, templateName, pattern)
}

// GetValues getting all parsed env
func (task *Task) GetValues(td *TaskData) (parsedValues map[string]string, err error) {
	parsedValues = map[string]string{}
	for key := range td.task.Project.values {
		parsedValues[key], err = task.GetValue(td, key)
		if err != nil {
			return parsedValues, err
		}
	}
	return parsedValues, nil
}

// GetConfig getting config of a task
func (task *Task) GetConfig(td *TaskData, keys ...string) (val string, err error) {
	key := strings.Join(keys, "::")
	pattern, exist := task.Config[key]
	if !exist {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return "", nil
		}
		return parentTask.GetConfig(td, keys...)
	}
	templateName := fmt.Sprintf("%s.config.%s", task.GetName(), key)
	return task.getParsedPattern(td, templateName, pattern)
}

// GetConfigs getting all parsed env
func (task *Task) GetConfigs(td *TaskData) (parsedConfig map[string]string, err error) {
	parsedConfig = map[string]string{}
	for key := range td.task.Config {
		parsedConfig[key], err = task.GetConfig(td, key)
		if err != nil {
			return parsedConfig, err
		}
	}
	return parsedConfig, nil
}

// GetLConfig getting lconfig of a task
func (task *Task) GetLConfig(td *TaskData, keys ...string) (val []string, err error) {
	key := strings.Join(keys, "::")
	val = []string{}
	lConfig, exist := task.LConfig[key]
	if !exist {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return []string{}, nil
		}
		return parentTask.GetLConfig(td, keys...)
	}
	for index, pattern := range lConfig {
		templateName := fmt.Sprintf("%s.lconfig.%s[%d]", task.GetName(), key, index)
		element, err := task.getParsedPattern(td, templateName, pattern)
		if err != nil {
			return val, err
		}
		val = append(val, element)
	}
	return val, err
}

// GetLConfigs getting all lConfig
func (task *Task) GetLConfigs(td *TaskData) (parsedLConfig map[string][]string, err error) {
	parsedLConfig = map[string][]string{}
	for key := range td.task.LConfig {
		parsedLConfig[key], err = task.GetLConfig(td, key)
		if err != nil {
			return parsedLConfig, err
		}
	}
	return parsedLConfig, nil
}

// GetEnv getting env of a task
func (task *Task) GetEnv(td *TaskData, key string) (val string, err error) {
	envConfig, exist := task.Env[key]
	if !exist {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return os.Getenv(key), nil
		}
		return parentTask.GetEnv(td, key)
	}
	if envConfig.From != "" {
		if val = os.Getenv(envConfig.From); val != "" {
			return val, nil
		}
	}
	templateNamePrefix := fmt.Sprintf("%s.env.%s", task.GetName(), key)
	return task.getParsedPattern(td, templateNamePrefix, envConfig.Default)
}

// GetEnvs getting all parsed env
func (task *Task) GetEnvs(td *TaskData) (parsedEnv map[string]string, err error) {
	parsedEnv = map[string]string{}
	for key := range td.task.Env {
		parsedEnv[key], err = task.GetEnv(td, key)
		if err != nil {
			return parsedEnv, err
		}
	}
	return parsedEnv, nil
}

func (task *Task) getParsedPattern(td *TaskData, templateNamePrefix, pattern string) (result string, err error) {
	templateName := task.getTemplateName(templateNamePrefix, pattern)
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

func (task *Task) getTemplateName(templateNamePrefix, pattern string) (templateName string) {
	lines := strings.Split(pattern, "\n")
	if len(lines) == 2 && lines[1] == "" {
		lines = []string{lines[0]}
	}
	if len(lines) > 1 {
		for index, line := range lines {
			lines[index] = fmt.Sprintf("%s | %s", fmt.Sprintf("%4d", index+1), line)
		}
	}
	return fmt.Sprintf("\n%s:\n%s\n", templateNamePrefix, strings.Join(lines, "\n"))
}

func (task *Task) linkToEnvs() {
	for _, env := range task.Env {
		env.Task = task
	}
}

func (task *Task) init() (err error) {
	var timeErr error
	task.timeoutDuration, timeErr = time.ParseDuration(task.Timeout)
	if timeErr != nil || task.timeoutDuration <= 0 {
		task.timeoutDuration = 5 * time.Minute
	}
	if task.Extend != "" {
		if _, exists := task.Project.Tasks[task.Extend]; !exists {
			return fmt.Errorf("Task %s is extended from %s but task %s doesn't exist", task.GetName(), task.Extend, task.Extend)
		}
	}
	task.generateIcon()
	task.generateLogPrefix()
	return nil
}

func (task *Task) generateIcon() {
	if task.Icon == "" {
		icon := task.Project.iconGenerator.Create()
		task.Icon = icon
	}
}

func (task *Task) generateLogPrefix() {
	repeat := 2 + task.Project.maxPublishedTaskNameLength - len(task.GetName()) - len(task.Icon)
	if repeat < 0 {
		repeat = 0
	}
	paddedStr := strings.Repeat(" ", repeat)
	d := task.Project.decoration
	paddedName := fmt.Sprintf("%s%s%s%s", d.GenerateColor(), task.GetName(), d.Normal, paddedStr)
	task.logPrefix = fmt.Sprintf("%s %s%s%s", paddedName, d.Faint, task.Icon, d.Normal)
}

func (task *Task) getPath() (path string) {
	if task.Location != "" {
		return filepath.Join(task.basePath, task.Location)
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

func (task *Task) getStartCmd(td *TaskData, logDone chan error) (cmd *exec.Cmd, exist bool, err error) {
	if len(task.Start) == 0 {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return cmd, false, fmt.Errorf("Cannot retrieve StartCmd from %s's parent", task.GetName())
		}
		return parentTask.getStartCmd(td, logDone)
	}
	cmd, err = task.getCmd(td, "START", task.Start, logDone)
	return cmd, true, err
}

// GetCheckCmd get check command of a task
func (task *Task) GetCheckCmd(logDone chan error) (cmd *exec.Cmd, exist bool, err error) {
	return task.getCheckCmd(NewTaskData(task), logDone)
}

func (task *Task) getCheckCmd(td *TaskData, logDone chan error) (cmd *exec.Cmd, exist bool, err error) {
	if len(task.Check) == 0 {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return cmd, false, fmt.Errorf("Cannot retrieve CheckCmd from %s's parent", task.GetName())
		}
		return parentTask.getCheckCmd(td, logDone)
	}
	cmd, err = task.getCmd(td, "CHECK", task.Check, logDone)
	return cmd, true, err
}

func (task *Task) getCmd(td *TaskData, cmdType string, commandPatternArgs []string, logDone chan error) (cmd *exec.Cmd, err error) {
	commandArgs := []string{}
	templateName := fmt.Sprintf("%s.%s", task.GetName(), strings.ToLower(cmdType))
	for _, pattern := range commandPatternArgs {
		arg, err := task.getParsedPattern(td, templateName, pattern)
		if err != nil {
			return cmd, err
		}
		commandArgs = append(commandArgs, arg)
	}
	name := commandArgs[0]
	args := commandArgs[1:]
	cmd = exec.Command(name, args...)
	cmd.Dir = td.task.GetWorkPath()
	cmd.Env = os.Environ()
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	for key := range td.task.Env {
		val, err := td.GetEnv(key)
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
	go task.log(td, cmdType, "OUT", outPipe, outDone)
	// log stderr
	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return cmd, err
	}
	errDone := make(chan error)
	go task.log(td, cmdType, "ERR", errPipe, errDone)
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

func (task *Task) log(td *TaskData, cmdType, logType string, pipe io.ReadCloser, logDone chan error) {
	buf := bufio.NewScanner(pipe)
	d := task.Project.decoration
	cmdIconType := task.getCmdIconType(cmdType)
	prefix := fmt.Sprintf("  %s%s%s %s", d.Faint, cmdIconType, d.Normal, td.task.logPrefix)
	saveLog := td.task.SaveLog == "" || boolean.IsTrue(td.task.SaveLog)
	print := logger.Printf
	if logType == "ERR" {
		print = logger.PrintfError
	}
	var err error = nil
	for buf.Scan() {
		content := buf.Text()
		print("%s %s\n", prefix, content)
		if saveLog {
			if csvWriteErr := task.Project.csvLogWriter.Log(logType, cmdType, td.Name, content, td.task.logPrefix); csvWriteErr != nil {
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
