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
	Start           []string              `yaml:"start"`
	Check           []string              `yaml:"check"`
	Timeout         string                `yaml:"timeout"`
	Private         bool                  `yaml:"private"`
	Extend          string                `yaml:"extend"`
	Location        string                `yaml:"location"`
	Config          map[string]string     `yaml:"config"`
	LConfig         map[string][]string   `yaml:"lconfig"`
	Env             map[string]*EnvConfig `yaml:"env"`
	Dependencies    []string              `yaml:"dependencies"`
	Description     string                `yaml:"description"`
	Icon            string                `yaml:"icon"`
	FileLocation    string
	Project         *ProjectConfig
	BasePath        string
	ParsedEnv       map[string]string
	ParsedConfig    map[string]string
	ParsedLConfig   map[string][]string
	Name            string
	FunkyName       string
	TimeoutDuration time.Duration
}

func (task *Task) fillEnvTask() {
	for _, env := range task.Env {
		env.Task = task
	}
}

// GetEnv get env for a task, parsed and os's
func (task *Task) GetEnv(key string) (val string) {
	if val, exists := task.ParsedEnv[key]; exists {
		return val
	}
	return os.Getenv(key)
}

func (task *Task) init() (err error) {
	task.TimeoutDuration, err = time.ParseDuration(task.Timeout)
	if err != nil || task.TimeoutDuration <= 0 {
		task.TimeoutDuration = 5 * time.Minute
	}
	if task.Extend != "" {
		if _, exists := task.Project.Tasks[task.Extend]; !exists {
			return fmt.Errorf("Task %s is extending %s but task %s doesn't exist", task.Name, task.Extend, task.Extend)
		}
	}
	task.ParsedEnv = task.getParsedEnv()
	if task.ParsedConfig, err = task.getParsedConfig(); err != nil {
		return err
	}
	if task.ParsedLConfig, err = task.getParsedLConfig(); err != nil {
		return err
	}
	task.generateIcon()
	task.generateFunkyName()
	return err
}

func (task *Task) generateIcon() {
	if task.Icon == "" {
		icon := task.Project.IconGenerator.Create()
		task.Icon = icon
	}
}

func (task *Task) generateFunkyName() {
	repeat := task.Project.MaxTaskNameLength - len(task.Name)
	if repeat < 0 {
		repeat = 0
	}
	paddedStr := strings.Repeat(" ", repeat)
	d := task.Project.Decoration
	paddedName := fmt.Sprintf("%s%s%s%s", d.GenerateColor(), task.Name, d.Normal, paddedStr)
	task.FunkyName = fmt.Sprintf("%s %s%s%s", paddedName, d.Faint, task.Icon, d.Normal)
}

func (task *Task) getParsedEnv() (parsedEnv map[string]string) {
	parsedEnv = map[string]string{}
	for envName, envConfig := range task.Env {
		val := os.Getenv(envConfig.From)
		if val == "" {
			val = envConfig.Default
		}
		parsedEnv[envName] = val
	}
	parentTask, exists := task.Project.Tasks[task.Extend]
	if !exists {
		return parsedEnv
	}
	parentParsedEnv := parentTask.getParsedEnv()
	for key, val := range parentParsedEnv {
		if _, exists := parsedEnv[key]; !exists {
			parsedEnv[key] = val
		}
	}
	return parsedEnv
}

func (task *Task) getParsedConfig() (parsedConfig map[string]string, err error) {
	parsedConfig = map[string]string{}
	for configName, configPattern := range task.Config {
		if parsedConfig[configName], err = task.parseCurentTaskTemplatePattern(configPattern); err != nil {
			return parsedConfig, err
		}
	}
	parentTask, exists := task.Project.Tasks[task.Extend]
	if !exists {
		return parsedConfig, nil
	}
	parentParsedConfig, err := parentTask.getParsedConfig()
	if err != nil {
		return parsedConfig, err
	}
	for configName, val := range parentParsedConfig {
		if _, exists := parsedConfig[configName]; !exists {
			parsedConfig[configName] = val
		}
	}
	return parsedConfig, err
}

func (task *Task) getParsedLConfig() (parsedLConfig map[string][]string, err error) {
	parsedLConfig = map[string][]string{}
	for configName, configPatterns := range task.LConfig {
		parsedLConfig[configName] = []string{}
		for _, configPattern := range configPatterns {
			val, err := task.parseCurentTaskTemplatePattern(configPattern)
			if err != nil {
				return parsedLConfig, err
			}
			parsedLConfig[configName] = append(parsedLConfig[configName], val)
		}
	}
	parentTask, exists := task.Project.Tasks[task.Extend]
	if !exists {
		return parsedLConfig, nil
	}
	parentParsedLConfig, err := parentTask.getParsedLConfig()
	if err != nil {
		return parsedLConfig, err
	}
	for configName, val := range parentParsedLConfig {
		if _, exists := parsedLConfig[configName]; !exists {
			parsedLConfig[configName] = val
		}
	}
	return parsedLConfig, err
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
func (task *Task) GetStartCmd() (cmd *exec.Cmd, exist bool, err error) {
	return task.getStartCmd(NewTaskData(task))
}

func (task *Task) getStartCmd(taskData *TaskData) (cmd *exec.Cmd, exist bool, err error) {
	if len(task.Start) == 0 {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return cmd, false, fmt.Errorf("Cannot retrieve StartCmd from %s's parent", task.Name)
		}
		return parentTask.getStartCmd(taskData)
	}
	cmd, err = task.getCmd("⏳", task.Start, taskData)
	return cmd, true, err
}

// GetCheckCmd get check command of a task
func (task *Task) GetCheckCmd() (cmd *exec.Cmd, exist bool, err error) {
	return task.getCheckCmd(NewTaskData(task))
}

func (task *Task) getCheckCmd(taskData *TaskData) (cmd *exec.Cmd, exist bool, err error) {
	if len(task.Check) == 0 {
		parentTask, exists := task.Project.Tasks[task.Extend]
		if !exists {
			return cmd, false, fmt.Errorf("Cannot retrieve CheckCmd from %s's parent", task.Name)
		}
		return parentTask.getCheckCmd(taskData)
	}
	cmd, err = task.getCmd("🔍", task.Check, taskData)
	return cmd, true, err
}

func (task *Task) getCmd(cmdIconType string, commandPatternArgs []string, taskData *TaskData) (cmd *exec.Cmd, err error) {
	commandArgs := []string{}
	for _, pattern := range commandPatternArgs {
		arg, err := task.parseTemplatePattern(pattern, taskData)
		if err != nil {
			return cmd, err
		}
		commandArgs = append(commandArgs, arg)
	}
	name := commandArgs[0]
	args := commandArgs[1:]
	cmd = exec.Command(name, args...)
	cmd.Dir = taskData.task.GetWorkPath()
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	for key, val := range task.ParsedEnv {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, val))
	}
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return cmd, err
	}
	go task.log(cmdIconType, "OUT", outPipe, taskData)
	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return cmd, err
	}
	go task.log(cmdIconType, "ERR", errPipe, taskData)
	return cmd, err
}

func (task *Task) log(cmdIconType, logType string, pipe io.ReadCloser, taskData *TaskData) {
	buf := bufio.NewScanner(pipe)
	d := task.Project.Decoration
	prefix := fmt.Sprintf("  %s%s %s", d.Normal, cmdIconType, taskData.task.FunkyName)
	for buf.Scan() {
		content := buf.Text()
		if logType == "ERR" {
			logger.PrintfError("%s %s\n", prefix, content)
		} else {
			logger.Printf("%s %s\n", prefix, content)
		}
	}
}

func (task *Task) parseCurentTaskTemplatePattern(pattern string) (val string, err error) {
	return task.parseTemplatePattern(pattern, NewTaskData(task))
}

func (task *Task) parseTemplatePattern(pattern string, taskData *TaskData) (val string, err error) {
	tmpl, err := template.New(pattern).Parse(pattern)
	if err != nil {
		return val, err
	}
	var b bytes.Buffer
	if err = tmpl.Execute(&b, taskData); err != nil {
		return val, err
	}
	val = b.String()
	return val, err
}
