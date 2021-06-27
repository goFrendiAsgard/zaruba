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
	"github.com/state-alchemists/zaruba/str"
)

// Task is zaruba task
type Task struct {
	Start                 []string          `yaml:"start,omitempty"`
	Check                 []string          `yaml:"check,omitempty"`
	Timeout               string            `yaml:"timeout,omitempty"`
	Private               bool              `yaml:"private,omitempty"`
	AutoTerminate         string            `yaml:"autoTerminate,omitempty"`
	Extend                string            `yaml:"extend,omitempty"`
	Extends               []string          `yaml:"extends,omitempty"`
	Location              string            `yaml:"location,omitempty"`
	ConfigRef             string            `yaml:"configRef,omitempty"`
	ConfigRefs            []string          `yaml:"configRefs,omitempty"`
	Config                map[string]string `yaml:"config,omitempty"`
	EnvRef                string            `yaml:"envRef,omitempty"`
	EnvRefs               []string          `yaml:"envRefs,omitempty"`
	Env                   map[string]*Env   `yaml:"env,omitempty"`
	Dependencies          []string          `yaml:"dependencies,omitempty"`
	Inputs                []string          `yaml:"inputs,omitempty"`
	Description           string            `yaml:"description,omitempty"`
	Icon                  string            `yaml:"icon,omitempty"`
	SaveLog               string            `yaml:"saveLog,omitempty"`
	fileLocation          string            // File location where this task was declared
	Project               *Project
	name                  string
	logPrefix             string
	timeoutDuration       time.Duration
	td                    *TaskData
	maxRecursiveLevel     int
	currentRecursiveLevel int
}

func (task *Task) init() {
	task.maxRecursiveLevel = 100
	task.currentRecursiveLevel = 0
	var timeErr error
	task.timeoutDuration, timeErr = time.ParseDuration(task.Timeout)
	if timeErr != nil || task.timeoutDuration <= 0 {
		task.timeoutDuration = 5 * time.Minute
	}
	task.generateIcon()
	task.generateLogPrefix()
}

// GetName get task name
func (task *Task) GetName() (name string) {
	return task.name
}

// GetTimeoutDuration get timeout duration of a task
func (task *Task) GetTimeoutDuration() time.Duration {
	return task.timeoutDuration
}

// GetFileLocation get file location of a task
func (task *Task) GetFileLocation() (fileLocation string) {
	return task.fileLocation
}

// GetWorkPath get path of current task
func (task *Task) GetWorkPath() (workPath string) {
	workPath = task.getPath()
	if workPath != "" {
		return workPath
	}
	workPath, _ = os.Getwd()
	return workPath
}

// GetAutoTerminate
func (task *Task) GetAutoTerminate() (autoTerminate bool) {
	if boolean.IsTrue(task.AutoTerminate) {
		return true
	}
	if boolean.IsFalse(task.AutoTerminate) {
		return false
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		if parentTask.GetAutoTerminate() {
			return true
		}
	}
	return false
}

// HaveStartCmd return whether task has start command or not
func (task *Task) HaveStartCmd() bool {
	if len(task.Start) > 0 {
		return true
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		if parentTask.HaveStartCmd() {
			return true
		}
	}
	return false
}

// HaveCheckCmd return whether task has check command or not
func (task *Task) HaveCheckCmd() bool {
	if len(task.Check) > 0 {
		return true
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		if parentTask.HaveCheckCmd() {
			return true
		}
	}
	return false
}

// GetValue getting config of a task
func (task *Task) GetValue(keys ...string) (val string, err error) {
	key := strings.Join(keys, "::")
	pattern, exist := task.Project.GetValue(key), task.Project.IsValueExist(key)
	if !exist {
		return "", nil
	}
	templateName := fmt.Sprintf("%s[values][%s]", task.GetName(), key)
	return task.getParsedPattern(templateName, pattern)
}

func (task *Task) GetValueKeys() (keys []string) {
	keys = []string{}
	for key := range task.Project.values {
		keys = append(keys, key)
	}
	return keys
}

// GetConfig getting config of a task
func (task *Task) GetConfig(keys ...string) (val string, err error) {
	key := strings.Join(keys, "::")
	if pattern, declared := task.GetConfigPattern(key); declared {
		templateName := fmt.Sprintf("%s[config][%s]", task.GetName(), key)
		return task.getParsedPattern(templateName, pattern)
	}
	return "", nil
}

func (task *Task) GetConfigKeys() (keys []string) {
	keys = []string{}
	for key := range task.Config {
		keys = append(keys, key)
	}
	for _, baseConfigKey := range task.getConfigRefKeys() {
		for key := range task.Project.ConfigRefMap[baseConfigKey].Map {
			keys = append(keys, key)
		}
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		parentKeys := parentTask.GetConfigKeys()
		keys = append(keys, parentKeys...)
	}
	return str.GetUniqueElements(keys)
}

func (task *Task) GetConfigPattern(key string) (pattern string, declared bool) {
	if pattern, declared = task.Config[key]; declared {
		return pattern, true
	}
	for _, baseConfigKey := range task.getConfigRefKeys() {
		projectBaseConfig := task.Project.ConfigRefMap[baseConfigKey]
		if pattern, declared = projectBaseConfig.Map[key]; declared {
			return pattern, true
		}
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		if pattern, declared = parentTask.GetConfigPattern(key); declared {
			return pattern, true
		}
	}
	return "", false
}

// GetEnvs getting all parsed env
func (task *Task) GetEnvs() (parsedEnv map[string]string, err error) {
	parsedEnv = map[string]string{}
	for _, key := range task.GetEnvKeys() {
		parsedEnv[key], err = task.GetEnv(key)
		if err != nil {
			return parsedEnv, err
		}
	}
	return parsedEnv, nil
}

// GetEnv getting env of a task
func (task *Task) GetEnv(key string) (val string, err error) {
	if env, declared := task.GetEnvObject(key); declared {
		if env.From != "" {
			if val = os.Getenv(env.From); val != "" {
				return val, nil
			}
		}
		templateNamePrefix := fmt.Sprintf("%s[env][%s]", task.GetName(), key)
		return task.getParsedPattern(templateNamePrefix, env.Default)
	}
	return os.Getenv(key), nil
}

func (task *Task) GetEnvKeys() (keys []string) {
	keys = []string{}
	for key := range task.Env {
		keys = append(keys, key)
	}
	for _, baseEnvKey := range task.getEnvRefKeys() {
		for key := range task.Project.EnvRefMap[baseEnvKey].Map {
			keys = append(keys, key)
		}
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		parentKeys := parentTask.GetEnvKeys()
		keys = append(keys, parentKeys...)
	}
	return str.GetUniqueElements(keys)
}

func (task *Task) GetEnvObject(key string) (env *Env, declared bool) {
	if env, declared = task.Env[key]; declared {
		return env, declared
	}
	for _, baseEnvKey := range task.getEnvRefKeys() {
		projectBaseEnv := task.Project.EnvRefMap[baseEnvKey]
		if baseEnv, declared := projectBaseEnv.Map[key]; declared {
			return &Env{From: baseEnv.From, Default: baseEnv.Default}, true
		}
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		if env, declared = parentTask.GetEnvObject(key); declared {
			return env, true
		}
	}
	return nil, false
}

func (task *Task) getParentTaskNames() (parentTaskNames []string) {
	if task.Extend != "" {
		return []string{task.Extend}
	}
	return task.Extends
}

func (task *Task) getConfigRefKeys() (parentTaskNames []string) {
	if task.ConfigRef != "" {
		return []string{task.ConfigRef}
	}
	return task.ConfigRefs
}

func (task *Task) getEnvRefKeys() (parentTaskNames []string) {
	if task.EnvRef != "" {
		return []string{task.EnvRef}
	}
	return task.EnvRefs
}

func (task *Task) getParsedPattern(templateNamePrefix, pattern string) (result string, err error) {
	if task.currentRecursiveLevel >= task.maxRecursiveLevel {
		return "", fmt.Errorf("max recursive parsing on %s: %s", templateNamePrefix, pattern)
	}
	if task.td == nil {
		task.td = NewTaskData(task)
	}
	templateName := task.getTemplateName(templateNamePrefix, pattern)
	tmpl, err := template.New(templateName).Option("missingkey=zero").Parse(pattern)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err = tmpl.Execute(&b, task.td); err != nil {
		return "", err
	}
	result = b.String()
	return result, nil
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

func (task *Task) generateIcon() {
	if task.Icon == "" {
		task.Icon = task.Project.decoration.GenerateIcon()
	}
}

func (task *Task) generateLogPrefix() {
	taskName := task.GetName()
	if len(taskName) > task.Project.maxPublishedTaskNameLength {
		strLen := task.Project.maxPublishedTaskNameLength - 3
		taskName = taskName[:strLen] + "..."
	} else {
		repeat := task.Project.maxPublishedTaskNameLength - len(taskName)
		taskName = taskName + strings.Repeat(" ", repeat)
	}
	d := task.Project.decoration
	color := d.Faint
	if !task.Private {
		color = d.GenerateColor()
	}
	task.logPrefix = fmt.Sprintf("%s%s%s %s", color, taskName, d.Normal, d.Icon(task.Icon))
}

func (task *Task) getPath() (path string) {
	if task.Location != "" {
		return filepath.Join(filepath.Dir(task.fileLocation), task.Location)
	}
	parentTaskNames := task.getParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
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
	seen := map[string]bool{}
	for _, dependency := range task.Dependencies {
		seen[dependency] = true
	}
	dependencies = append([]string{}, task.Dependencies...)
	for _, dependencyTaskName := range dependencies {
		subDependencies := task.Project.Tasks[dependencyTaskName].getDependencies()
		for _, subDependency := range subDependencies {
			if !seen[subDependency] {
				dependencies = append(dependencies, subDependency)
				seen[subDependency] = true
			}
		}
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		subDependencies := task.Project.Tasks[parentTaskName].getDependencies()
		for _, subDependency := range subDependencies {
			if !seen[subDependency] {
				dependencies = append(dependencies, subDependency)
				seen[subDependency] = true
			}
		}
	}
	return dependencies
}

// GetStartCmd get start command of a task
func (task *Task) GetStartCmd(logDone chan error) (cmd *exec.Cmd, exist bool, err error) {
	cmdPatterns, exist, err := task.GetStartCmdPatterns()
	if err != nil || !exist {
		return cmd, exist, err
	}
	cmd, err = task.getCmd("START", cmdPatterns, logDone)
	return cmd, exist, err
}

func (task *Task) GetStartCmdPatterns() (cmdPatterns []string, exist bool, err error) {
	if len(task.Start) > 0 {
		return task.Start, true, nil
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		cmdPatterns, exist, err = parentTask.GetStartCmdPatterns()
		if err != nil || exist {
			return cmdPatterns, exist, err
		}
	}
	return cmdPatterns, false, fmt.Errorf("cannot retrieve start cmd from any parent task of %s", task.GetName())
}

// GetCheckCmd get check command of a task
func (task *Task) GetCheckCmd(logDone chan error) (cmd *exec.Cmd, exist bool, err error) {
	cmdPatterns, exist, err := task.GetCheckCmdPatterns()
	if err != nil || !exist {
		return cmd, exist, err
	}
	cmd, err = task.getCmd("CHECK", cmdPatterns, logDone)
	return cmd, exist, err
}

func (task *Task) GetCheckCmdPatterns() (cmdPatterns []string, exist bool, err error) {
	if len(task.Check) > 0 {
		return task.Check, true, nil
	}
	for _, parentTaskName := range task.getParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		cmdPatterns, exist, err = parentTask.GetCheckCmdPatterns()
		if err != nil || exist {
			return cmdPatterns, exist, err
		}
	}
	return cmdPatterns, false, fmt.Errorf("cannot retrieve check cmd from any parent task of %s", task.GetName())
}

func (task *Task) getCmd(cmdType string, commandPatternArgs []string, logDone chan error) (cmd *exec.Cmd, err error) {
	commandArgs := []string{}
	templateNamePrefix := fmt.Sprintf("%s[%s]", task.GetName(), strings.ToLower(cmdType))
	for _, pattern := range commandPatternArgs {
		arg, err := task.getParsedPattern(templateNamePrefix, pattern)
		if err != nil {
			return cmd, err
		}
		commandArgs = append(commandArgs, arg)
	}
	name, args := commandArgs[0], commandArgs[1:]
	cmd = exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Dir = task.GetWorkPath()
	cmd.Env = os.Environ()
	envs, err := task.GetEnvs()
	if err != nil {
		return cmd, err
	}
	for key, val := range envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, val))
	}
	// log stdout
	outPipe, _ := cmd.StdoutPipe()
	outDone := make(chan error)
	go task.log(cmdType, "OUT", outPipe, outDone)
	// log stderr
	errPipe, _ := cmd.StderrPipe()
	errDone := make(chan error)
	go task.log(cmdType, "ERR", errPipe, errDone)
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

func (task *Task) log(cmdType, logType string, pipe io.ReadCloser, logDone chan error) {
	buf := bufio.NewScanner(pipe)
	d := task.Project.decoration
	cmdIconType := task.getCmdIconType(cmdType)
	prefix := fmt.Sprintf("%s %s", cmdIconType, task.logPrefix)
	saveLog := task.SaveLog == "" || boolean.IsTrue(task.SaveLog)
	print := task.Project.logger.DPrintf
	if logType == "ERR" {
		print = task.Project.logger.DPrintfError
	}
	taskName := task.GetName()
	var err error = nil
	for buf.Scan() {
		content := buf.Text()
		now := time.Now()
		nowRoundStr := fmt.Sprintf("%-12s", now.Format("15:04:05.999"))
		print("%s %s%s%s %s\n", prefix, d.Faint, nowRoundStr, d.Normal, content)
		if saveLog {
			nowStr := now.String()
			if csvWriteErr := task.Project.dataLogger.Log(nowStr, logType, cmdType, taskName, content); csvWriteErr != nil {
				err = csvWriteErr
			}
		}
	}
	logDone <- err
}

func (task *Task) getCmdIconType(cmdType string) string {
	if cmdType == "CHECK" {
		return task.Project.decoration.Inspect
	}
	return task.Project.decoration.Run
}
