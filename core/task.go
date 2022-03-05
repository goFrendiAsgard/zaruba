package core

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"syscall"
	"text/template"
	"time"
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
	Configs               map[string]string `yaml:"configs,omitempty"`
	EnvRef                string            `yaml:"envRef,omitempty"`
	EnvRefs               []string          `yaml:"envRefs,omitempty"`
	Envs                  map[string]*Env   `yaml:"envs,omitempty"`
	Dependencies          []string          `yaml:"dependencies,omitempty"`
	Inputs                []string          `yaml:"inputs,omitempty"`
	Description           string            `yaml:"description,omitempty"`
	Icon                  string            `yaml:"icon,omitempty"`
	SaveLog               string            `yaml:"saveLog,omitempty"`
	SyncEnv               string            `yaml:"syncEnv,omitempty"`
	SyncEnvLocation       string            `yaml:"syncEnvLocation,omitempty"`
	Project               *Project          `yaml:"_project,omitempty"`
	fileLocation          string            // File location where this task was declared
	uuid                  string            // Unique identifier of current task
	name                  string            // Current task name
	generatedRandomName   string            // Random name
	logPrefix             string            // Task prefix for logging
	timeoutDuration       time.Duration
	td                    *Tpl
	maxRecursiveLevel     int
	currentRecursiveLevel int
}

func (task *Task) init() {
	task.maxRecursiveLevel = 100
	task.currentRecursiveLevel = 0
	task.generateIcon()
	task.generateLogPrefix()
	task.generateUUID()
	task.generateGeneratedRandomName()
}

// GetUUID get task uid
func (task *Task) GetUUID() (uuid string) {
	return task.uuid
}

// GetGeneratedRandomName get generated random name
func (task *Task) GetGeneratedRandomName() (name string) {
	return task.generatedRandomName
}

// GetName get task name
func (task *Task) GetName() (name string) {
	return task.name
}

func (task *Task) getTimeoutDuration() time.Duration {
	timeoutDuration, err := time.ParseDuration(task.Timeout)
	if err == nil {
		return timeoutDuration
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		parentTimeoutDuration := parentTask.getTimeoutDuration()
		if parentTimeoutDuration > 0 {
			return parentTimeoutDuration
		}
	}
	return 0
}

// GetTimeoutDuration get timeout duration of a task
func (task *Task) GetTimeoutDuration() time.Duration {
	if task.timeoutDuration > 0 {
		return task.timeoutDuration
	}
	timeoutDuration := task.getTimeoutDuration()
	if timeoutDuration <= 0 {
		timeoutDuration = 10 * time.Minute
	}
	task.timeoutDuration = timeoutDuration
	return timeoutDuration
}

// GetFileLocation get file location of a task
func (task *Task) GetFileLocation() (fileLocation string) {
	return task.fileLocation
}

func (task *Task) GetLocation() (path string) {
	if task.Location != "" {
		return filepath.Join(filepath.Dir(task.fileLocation), task.Location)
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		parentTaskLocation := parentTask.GetLocation()
		if parentTaskLocation != "" {
			return parentTaskLocation
		}
	}
	return ""
}

func (task *Task) GetSaveLog() bool {
	if task.Project.Util.Bool.IsFalse(task.SaveLog) {
		return false
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		if parentTask.GetSaveLog() {
			return true
		}
	}
	return true
}

func (task *Task) ShouldSyncEnv() bool {
	if task.Project.Util.Bool.IsFalse(task.SyncEnv) {
		return false
	}
	if task.Project.Util.Bool.IsTrue(task.SyncEnv) {
		return true
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		if parentTask.ShouldSyncEnv() {
			return true
		}
	}
	return false
}

func (task *Task) GetSyncEnvLocation() (path string) {
	if task.SyncEnvLocation != "" {
		return filepath.Join(filepath.Dir(task.fileLocation), task.SyncEnvLocation)
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		parentTaskSyncEnvLocation := parentTask.GetSyncEnvLocation()
		if parentTaskSyncEnvLocation != "" {
			return parentTaskSyncEnvLocation
		}
	}
	return task.GetLocation()
}

// GetWorkPath get path of current task
func (task *Task) GetWorkPath() (workPath string) {
	if taskLocation := task.GetLocation(); taskLocation != "" {
		return taskLocation
	}
	workPath, _ = os.Getwd()
	return workPath
}

// GetAutoTerminate
func (task *Task) GetAutoTerminate() (autoTerminate bool) {
	if task.Project.Util.Bool.IsTrue(task.AutoTerminate) {
		return true
	}
	if task.Project.Util.Bool.IsFalse(task.AutoTerminate) {
		return false
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
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
	for _, parentTaskName := range task.GetParentTaskNames() {
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
	for _, parentTaskName := range task.GetParentTaskNames() {
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

// GetConfigs getting all parsed config which key matching a pattern
func (task *Task) GetConfigs(keyPattern string) (parsedConfig map[string]string, err error) {
	parsedConfig = map[string]string{}
	for _, key := range task.GetConfigKeys() {
		match, err := regexp.MatchString(keyPattern, key)
		if err != nil {
			return parsedConfig, err
		}
		if !match {
			continue
		}
		parsedConfig[key], err = task.GetConfig(key)
		if err != nil {
			return parsedConfig, err
		}
	}
	return parsedConfig, nil
}

// GetConfig getting config of a task
func (task *Task) GetConfig(key string) (val string, err error) {
	if pattern, declared := task.GetConfigPattern(key); declared {
		templateName := fmt.Sprintf("%s[config][%s]", task.GetName(), key)
		return task.getParsedPattern(templateName, pattern)
	}
	return "", nil
}

func (task *Task) GetConfigKeys() (keys []string) {
	keys = []string{}
	for key := range task.Configs {
		keys = append(keys, key)
	}
	for _, envRefName := range task.getConfigRefKeys() {
		for key := range task.Project.ConfigRefMap[envRefName].Map {
			keys = append(keys, key)
		}
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		parentKeys := parentTask.GetConfigKeys()
		keys = append(keys, parentKeys...)
	}
	return task.getUniqueElements(keys)
}

func (task *Task) GetConfigPattern(key string) (pattern string, declared bool) {
	if pattern, declared = task.Configs[key]; declared {
		return pattern, true
	}
	for _, configRefName := range task.getConfigRefKeys() {
		projectBaseConfig := task.Project.ConfigRefMap[configRefName]
		if pattern, declared = projectBaseConfig.Map[key]; declared {
			return pattern, true
		}
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
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
	for key := range task.Envs {
		keys = append(keys, key)
	}
	for _, envRefName := range task.getEnvRefKeys() {
		for key := range task.Project.EnvRefMap[envRefName].Map {
			keys = append(keys, key)
		}
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		parentKeys := parentTask.GetEnvKeys()
		keys = append(keys, parentKeys...)
	}
	return task.getUniqueElements(keys)
}

func (task *Task) GetEnvObject(key string) (env *Env, declared bool) {
	if env, declared = task.Envs[key]; declared {
		return env, declared
	}
	for _, envRefName := range task.getEnvRefKeys() {
		projectBaseEnv := task.Project.EnvRefMap[envRefName]
		if envObject, declared := projectBaseEnv.Map[key]; declared {
			return &Env{From: envObject.From, Default: envObject.Default}, true
		}
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		if env, declared = parentTask.GetEnvObject(key); declared {
			return env, true
		}
	}
	return nil, false
}

func (task *Task) GetFirstEnvRefName() (envRefName string) {
	if task.EnvRef != "" {
		return task.EnvRef
	}
	if len(task.EnvRefs) > 0 {
		return task.EnvRefs[0]
	}
	return ""
}

func (task *Task) GetFirstConfigRefName() (configRefName string) {
	if task.ConfigRef != "" {
		return task.ConfigRef
	}
	if len(task.ConfigRefs) > 0 {
		return task.ConfigRefs[0]
	}
	return ""
}

func (task *Task) GetParentTaskNames() (parentTaskNames []string) {
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
		task.td = NewTpl(task)
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
	for _, env := range task.Envs {
		env.Task = task
	}
}

func (task *Task) generateIcon() {
	if task.Icon == "" {
		task.Icon = task.Project.Decoration.GenerateIcon()
	}
}

func (task *Task) generateUUID() {
	if task.uuid == "" {
		task.uuid = task.Project.Util.Str.NewUUID()
	}
}

func (task *Task) generateGeneratedRandomName() {
	if task.uuid == "" {
		task.generatedRandomName = task.Project.Util.Str.NewName()
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
	d := task.Project.Decoration
	color := d.Faint
	if !task.Private {
		color = d.GenerateColor()
	}
	task.logPrefix = fmt.Sprintf("%s%s%s %s", color, taskName, d.Normal, d.Icon(task.Icon))
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
	for _, parentTaskName := range task.GetParentTaskNames() {
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
func (task *Task) GetStartCmd() (cmd *exec.Cmd, exist bool, err error) {
	cmdPatterns, exist, err := task.GetStartCmdPatterns()
	if err != nil || !exist {
		return cmd, exist, err
	}
	cmd, err = task.getCmd("START", cmdPatterns)
	return cmd, exist, err
}

func (task *Task) GetStartCmdPatterns() (cmdPatterns []string, exist bool, err error) {
	if len(task.Start) > 0 {
		return task.Start, true, nil
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		cmdPatterns, exist, err = parentTask.GetStartCmdPatterns()
		if err != nil || exist {
			return cmdPatterns, exist, err
		}
	}
	return cmdPatterns, false, fmt.Errorf("cannot retrieve start cmd from any parent task of %s", task.GetName())
}

// GetCheckCmd get check command of a task
func (task *Task) GetCheckCmd() (cmd *exec.Cmd, exist bool, err error) {
	cmdPatterns, exist, err := task.GetCheckCmdPatterns()
	if err != nil || !exist {
		return cmd, exist, err
	}
	cmd, err = task.getCmd("CHECK", cmdPatterns)
	return cmd, exist, err
}

func (task *Task) GetCheckCmdPatterns() (cmdPatterns []string, exist bool, err error) {
	if len(task.Check) > 0 {
		return task.Check, true, nil
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		cmdPatterns, exist, err = parentTask.GetCheckCmdPatterns()
		if err != nil || exist {
			return cmdPatterns, exist, err
		}
	}
	return cmdPatterns, false, fmt.Errorf("cannot retrieve check cmd from any parent task of %s", task.GetName())
}

func (task *Task) getCmd(cmdType string, commandPatternArgs []string) (cmd *exec.Cmd, err error) {
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
	go task.log(cmdType, "OUT", outPipe, task.Project.StdoutChan, task.Project.StdoutRecordChan)
	// log stderr
	errPipe, _ := cmd.StderrPipe()
	go task.log(cmdType, "ERR", errPipe, task.Project.StderrChan, task.Project.StderrRecordChan)
	// combine stdout and stderr done
	return cmd, err
}

func (task *Task) log(cmdType, logType string, pipe io.ReadCloser, logChan chan string, logRecordChan chan []string) {
	buf := bufio.NewScanner(pipe)
	d := task.Project.Decoration
	cmdIconType := task.getCmdIconType(cmdType)
	prefix := fmt.Sprintf("%s %s", cmdIconType, task.logPrefix)
	saveLog := task.GetSaveLog()
	taskName := task.GetName()
	for buf.Scan() {
		content := buf.Text()
		now := time.Now()
		nowRoundStr := fmt.Sprintf("%-12s", now.Format("15:04:05.999"))
		decoratedContent := fmt.Sprintf("%s %s%s%s %s\n", prefix, d.Faint, nowRoundStr, d.Normal, content)
		logChan <- decoratedContent
		if saveLog {
			nowStr := now.String()
			rowContent := []string{nowStr, logType, cmdType, taskName, content}
			logRecordChan <- rowContent
		}
	}
}

func (task *Task) getUniqueElements(arr []string) (result []string) {
	result = []string{}
	seen := map[string]bool{}
	for _, element := range arr {
		if _, exist := seen[element]; exist {
			continue
		}
		result = append(result, element)
		seen[element] = true
	}
	return result
}

func (task *Task) getCmdIconType(cmdType string) string {
	if cmdType == "CHECK" {
		return task.Project.Decoration.Inspect
	}
	return task.Project.Decoration.Run
}
