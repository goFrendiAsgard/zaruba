package dsl

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
	"strconv"
	"strings"
	"syscall"
	"text/template"
	"time"
)

// Task is zaruba task
type Task struct {
	Start                   []string          `yaml:"start,omitempty"`
	MaxStartRetryStr        string            `yaml:"maxStartRetry,omitempty"`
	StartRetryDelayStr      string            `yaml:"startRetryDelay,omitempty"`
	Check                   []string          `yaml:"check,omitempty"`
	MaxCheckRetryStr        string            `yaml:"maxCheckRetry,omitempty"`
	CheckRetryDelayStr      string            `yaml:"checkRetryDelay,omitempty"`
	TimeoutStr              string            `yaml:"timeout,omitempty"`
	Private                 bool              `yaml:"private,omitempty"`
	AutoTerminateStr        string            `yaml:"autoTerminate,omitempty"`
	Extend                  string            `yaml:"extend,omitempty"`
	Extends                 []string          `yaml:"extends,omitempty"`
	Location                string            `yaml:"location,omitempty"`
	ConfigRef               string            `yaml:"configRef,omitempty"`
	ConfigRefs              []string          `yaml:"configRefs,omitempty"`
	Configs                 map[string]string `yaml:"configs,omitempty"`
	EnvRef                  string            `yaml:"envRef,omitempty"`
	EnvRefs                 []string          `yaml:"envRefs,omitempty"`
	Envs                    map[string]*Env   `yaml:"envs,omitempty"`
	Dependencies            []string          `yaml:"dependencies,omitempty"`
	Inputs                  []string          `yaml:"inputs,omitempty"`
	Description             string            `yaml:"description,omitempty"`
	Icon                    string            `yaml:"icon,omitempty"`
	SaveLog                 string            `yaml:"saveLog,omitempty"`
	SyncEnv                 string            `yaml:"syncEnv,omitempty"`
	SyncEnvLocation         string            `yaml:"syncEnvLocation,omitempty"`
	Project                 *Project          `yaml:"_project,omitempty"`
	fileLocation            string            // File location where this task was declared
	uuid                    string            // Unique identifier of current task
	name                    string            // Current task name
	generatedRandomName     string            // Random name
	logPrefix               string            // Task prefix for logging
	timeoutDuration         time.Duration
	startRetry              int
	startRetryDelayDuration time.Duration
	checkRetry              int
	checkRetryDelayDuration time.Duration
	tpl                     *Tpl
	maxRecursiveLevel       int
	currentRecursiveLevel   int
	color                   string
	icon                    string
	isIconGenerated         bool
}

func (task *Task) init() {
	task.isIconGenerated = false
	task.maxRecursiveLevel = 100
	task.currentRecursiveLevel = 0
	task.generateIcon()
	task.generateColor()
	task.generateLogPrefix()
	task.generateUUID()
	task.generateGeneratedRandomName()
}

func (task *Task) GetUUID() (uuid string) {
	return task.uuid
}

func (task *Task) GetGeneratedRandomName() (name string) {
	return task.generatedRandomName
}

func (task *Task) GetName() (name string) {
	return task.name
}

func (task *Task) GetColor() (color string) {
	return task.color
}

func (task *Task) GetIcon() (icon string) {
	return task.icon
}

func (task *Task) GetDecoratedIcon() (decoratedIcon string) {
	return task.Project.Decoration.Icon(task.icon)
}

func (task *Task) getDefaultMaxStartRetry() int {
	return 3
}

func (task *Task) getMaxStartRetry() int {
	startRetry, err := strconv.Atoi(task.MaxStartRetryStr)
	if err == nil {
		return startRetry
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		parentStartRetry := parentTask.getMaxStartRetry()
		if parentStartRetry > 0 {
			return parentStartRetry
		}
	}
	return task.getDefaultMaxStartRetry()
}

func (task *Task) GetMaxStartRetry() int {
	if task.startRetry > 0 {
		return task.startRetry
	}
	startRetry := task.getMaxStartRetry()
	if startRetry <= 0 {
		startRetry = task.getDefaultMaxStartRetry()
	}
	task.startRetry = startRetry
	return startRetry
}

func (task *Task) getDefaultStartRetryDelayDuration() time.Duration {
	return 1 * time.Second
}

func (task *Task) getStartRetryDelayDuration() time.Duration {
	startRetryDelayDuration, err := time.ParseDuration(task.StartRetryDelayStr)
	if err == nil {
		return startRetryDelayDuration
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		parentStartRetryDelayDuration := parentTask.getStartRetryDelayDuration()
		if parentStartRetryDelayDuration > 0 {
			return parentStartRetryDelayDuration
		}
	}
	return task.getDefaultStartRetryDelayDuration()
}

func (task *Task) GetStartRetryDelayDuration() time.Duration {
	if task.startRetryDelayDuration > 0 {
		return task.startRetryDelayDuration
	}
	startRetryDelayDuration := task.getStartRetryDelayDuration()
	if startRetryDelayDuration < 0 {
		startRetryDelayDuration = task.getDefaultStartRetryDelayDuration()
	}
	task.startRetryDelayDuration = startRetryDelayDuration
	return startRetryDelayDuration
}

func (task *Task) getDefaultMaxCheckRetry() int {
	// 0 means infinite
	return 0
}

func (task *Task) getMaxCheckRetry() int {
	checkRetry, err := strconv.Atoi(task.MaxCheckRetryStr)
	if err == nil {
		return checkRetry
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		parentCheckRetry := parentTask.getMaxCheckRetry()
		if parentCheckRetry > 0 {
			return parentCheckRetry
		}
	}
	return task.getDefaultMaxCheckRetry()
}

func (task *Task) GetMaxCheckRetry() int {
	if task.checkRetry > 0 {
		return task.checkRetry
	}
	checkRetry := task.getMaxCheckRetry()
	if checkRetry < 0 {
		checkRetry = task.getDefaultMaxCheckRetry()
	}
	task.checkRetry = checkRetry
	return checkRetry
}

func (task *Task) getDefaultCheckRetryDelayDuration() time.Duration {
	return 1 * time.Second
}

func (task *Task) getCheckRetryDelayDuration() time.Duration {
	checkRetryDelayDuration, err := time.ParseDuration(task.CheckRetryDelayStr)
	if err == nil {
		return checkRetryDelayDuration
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		parentCheckRetryDelayDuration := parentTask.getCheckRetryDelayDuration()
		if parentCheckRetryDelayDuration > 0 {
			return parentCheckRetryDelayDuration
		}
	}
	return task.getDefaultCheckRetryDelayDuration()
}

func (task *Task) GetCheckRetryDelayDuration() time.Duration {
	if task.checkRetryDelayDuration > 0 {
		return task.checkRetryDelayDuration
	}
	checkRetryDelayDuration := task.getCheckRetryDelayDuration()
	if checkRetryDelayDuration < 0 {
		checkRetryDelayDuration = task.getDefaultCheckRetryDelayDuration()
	}
	task.checkRetryDelayDuration = checkRetryDelayDuration
	return checkRetryDelayDuration
}

func (task *Task) getTimeoutDuration() time.Duration {
	timeoutDuration, err := time.ParseDuration(task.TimeoutStr)
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

func (task *Task) GetIsSaveLog() bool {
	if task.Project.Util.Bool.IsFalse(task.SaveLog) {
		return false
	}
	parentTaskNames := task.GetParentTaskNames()
	if len(parentTaskNames) > 0 {
		parentTaskName := parentTaskNames[0]
		parentTask := task.Project.Tasks[parentTaskName]
		if parentTask.GetIsSaveLog() {
			return true
		}
	}
	return true
}

func (task *Task) GetShouldSyncEnv() bool {
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
		if parentTask.GetShouldSyncEnv() {
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

func (task *Task) GetWorkPath() (workPath string) {
	if taskLocation := task.GetLocation(); taskLocation != "" {
		return taskLocation
	}
	workPath, _ = os.Getwd()
	return workPath
}

func (task *Task) GetAutoTerminate() (autoTerminate bool) {
	if task.Project.Util.Bool.IsTrue(task.AutoTerminateStr) {
		return true
	}
	if task.Project.Util.Bool.IsFalse(task.AutoTerminateStr) {
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

func (task *Task) IsHavingStartCmd() bool {
	if len(task.Start) > 0 {
		return true
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		if parentTask.IsHavingStartCmd() {
			return true
		}
	}
	return false
}

func (task *Task) IsHavingCheckCmd() bool {
	if len(task.Check) > 0 {
		return true
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		if parentTask.IsHavingCheckCmd() {
			return true
		}
	}
	return false
}

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
	configKeys := task.GetConfigKeys()
	for _, configKey := range configKeys {
		upperSnakeConfigKey := task.Project.Util.Str.ToUpperSnake(configKey)
		configTemplatePattern := fmt.Sprintf("{{ .GetConfig \"%s\" }}", configKey)
		// normalize ${ZARUBA_CONFIG_<CONFIG_NAME>}
		pattern = strings.ReplaceAll(pattern, fmt.Sprintf("${ZARUBA_CONFIG_%s}", upperSnakeConfigKey), configTemplatePattern)
		// normalize $ZARUBA_CONFIG_<CONFIG_NAME>
		pattern = strings.ReplaceAll(pattern, fmt.Sprintf("$ZARUBA_CONFIG_%s", upperSnakeConfigKey), configTemplatePattern)
	}
	inputMap, _, _ := task.Project.GetInputs([]string{task.GetName()})
	for inputKey := range inputMap {
		upperSnakeInputKey := task.Project.Util.Str.ToUpperSnake(inputKey)
		inputTemplatePattern := fmt.Sprintf("{{ .GetValue \"%s\" }}", inputKey)
		// normalize ${ZARUBA_INPUT_<INPUT_NAME>}
		pattern = strings.ReplaceAll(pattern, fmt.Sprintf("${ZARUBA_INPUT_%s}", upperSnakeInputKey), inputTemplatePattern)
		// normalize $ZARUBA_INPUT_<INPUT_NAME>
		pattern = strings.ReplaceAll(pattern, fmt.Sprintf("$ZARUBA_INPUT_%s", upperSnakeInputKey), inputTemplatePattern)
	}
	if task.tpl == nil {
		task.tpl = NewTpl(task)
	}
	templateName := task.getTemplateName(templateNamePrefix, pattern)
	tmpl, err := template.New(templateName).Option("missingkey=zero").Parse(pattern)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err = tmpl.Execute(&b, task.tpl); err != nil {
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
	if !task.isIconGenerated {
		task.icon = task.Icon
		if task.icon == "" {
			task.icon = task.Project.Decoration.GenerateIcon()
		}
		task.isIconGenerated = true
	}
}

func (task *Task) generateUUID() {
	if task.uuid == "" {
		task.uuid = task.Project.Util.Str.NewUUID()
	}
}

func (task *Task) generateGeneratedRandomName() {
	if task.generatedRandomName == "" {
		task.generatedRandomName = task.Project.Util.Str.NewName()
	}
}

func (task *Task) generateColor() {
	if task.color == "" {
		d := task.Project.Decoration
		color := d.Faint
		if !task.Private {
			color = d.GenerateColor()
		}
		task.color = color
	}
}

func (task *Task) generateLogPrefix() {
	logTaskName := task.GetName()
	if len(logTaskName) > task.Project.maxPublishedTaskNameLength {
		strLen := task.Project.maxPublishedTaskNameLength - 3
		logTaskName = logTaskName[:strLen] + "..."
	} else {
		repeat := task.Project.maxPublishedTaskNameLength - len(logTaskName)
		logTaskName = logTaskName + strings.Repeat(" ", repeat)
	}
	d := task.Project.Decoration
	task.logPrefix = fmt.Sprintf("%s %s%s%s", task.GetDecoratedIcon(), task.color, logTaskName, d.Normal)
}

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

func (task *Task) GetStartCmd() (cmd *exec.Cmd, err error) {
	cmdPatterns, err := task.GetStartCmdPatterns()
	if err != nil {
		return cmd, err
	}
	cmd, err = task.getCmd("START", cmdPatterns)
	return cmd, err
}

func (task *Task) GetStartCmdPatterns() (cmdPatterns []string, err error) {
	if !task.IsHavingStartCmd() {
		return cmdPatterns, fmt.Errorf("cannot retrieve start cmd from any parent task of %s", task.GetName())
	}
	if len(task.Start) > 0 {
		return task.Start, nil
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		cmdPatterns, err = parentTask.GetStartCmdPatterns()
		if err != nil {
			return cmdPatterns, err
		}
	}
	return cmdPatterns, err
}

func (task *Task) GetCheckCmd() (cmd *exec.Cmd, err error) {
	cmdPatterns, err := task.GetCheckCmdPatterns()
	if err != nil {
		return cmd, err
	}
	cmd, err = task.getCmd("CHECK", cmdPatterns)
	return cmd, err
}

func (task *Task) GetCheckCmdPatterns() (cmdPatterns []string, err error) {
	if !task.IsHavingCheckCmd() {
		return cmdPatterns, fmt.Errorf("cannot retrieve check cmd from any parent task of %s", task.GetName())
	}
	if len(task.Check) > 0 {
		return task.Check, nil
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		parentTask := task.Project.Tasks[parentTaskName]
		cmdPatterns, err = parentTask.GetCheckCmdPatterns()
		if err != nil {
			return cmdPatterns, err
		}
	}
	return cmdPatterns, nil
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
	// task env
	if err = task.setCmdEnv(cmd); err != nil {
		return cmd, err
	}
	sessionId := task.Project.GetSessionId()
	// log stdout
	outPipe, _ := cmd.StdoutPipe()
	go task.readLogFromBuffer(sessionId, cmdType, "OUT", outPipe, task.Project.StdoutChan, task.Project.StdoutRecordChan)
	// log stderr
	errPipe, _ := cmd.StderrPipe()
	go task.readLogFromBuffer(sessionId, cmdType, "ERR", errPipe, task.Project.StderrChan, task.Project.StderrRecordChan)
	// combine stdout and stderr done
	return cmd, err
}

func (task *Task) setCmdEnv(cmd *exec.Cmd) error {
	// env
	envMap, err := task.GetEnvs()
	if err != nil {
		return err
	}
	for key, val := range envMap {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, val))
	}
	// config
	configKeys := task.GetConfigKeys()
	for _, configKey := range configKeys {
		configEnvKey := fmt.Sprintf("ZARUBA_CONFIG_%s", task.Project.Util.Str.ToUpperSnake(configKey))
		val, err := task.GetConfig(configKey)
		if err != nil {
			return err
		}
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", configEnvKey, val))
	}
	return nil
}

func (task *Task) readLogFromBuffer(sessionId, cmdType, logType string, pipe io.ReadCloser, logChan chan string, logRecordChan chan []string) {
	buf := bufio.NewScanner(pipe)
	isSaveLog := task.GetIsSaveLog()
	outputWgAdditionPerRow := 1
	if isSaveLog {
		outputWgAdditionPerRow = 2
	}
	cmdIconType := task.getCmdIconType(cmdType)
	logPrefix := fmt.Sprintf("%s %s", cmdIconType, task.logPrefix)
	taskName := task.GetName()
	isFirstTime := true
	previousChan := make(chan bool)
	nextChan := make(chan bool)
	for buf.Scan() {
		task.Project.OutputWgMutex.Lock()
		task.Project.OutputWg.Add(outputWgAdditionPerRow)
		task.Project.OutputWgMutex.Unlock()
		content := buf.Text()
		// previous and next chan is necessary to make sure that logChan and logRecordChan get message in order
		if isFirstTime {
			isFirstTime = false
			go task.sendLog(cmdType, logType, logPrefix, sessionId, taskName, content, isSaveLog, previousChan, nextChan, logChan, logRecordChan)
			previousChan <- true
			continue
		}
		previousChan = nextChan
		nextChan = make(chan bool)
		go task.sendLog(cmdType, logType, logPrefix, sessionId, taskName, content, isSaveLog, previousChan, nextChan, logChan, logRecordChan)
	}
}

func (task *Task) sendLog(cmdType, logType, logPrefix, sessionId, taskName, content string, saveLog bool, previousChan, nextChan chan bool, logChan chan string, logRecordChan chan []string) {
	d := task.Project.Decoration
	now := time.Now()
	decoratedContent := ""
	if task.Project.showLogTime {
		nowRoundStr := fmt.Sprintf("%-12s", now.Format("15:04:05.999"))
		decoratedContent = fmt.Sprintf("%s %s%s%s %s\n", logPrefix, d.Faint, nowRoundStr, d.Normal, content)
	} else {
		decoratedContent = fmt.Sprintf("%s %s\n", logPrefix, content)
	}
	<-previousChan
	logChan <- decoratedContent
	if saveLog {
		nowStr := now.String()
		rowContent := []string{nowStr, logType, cmdType, taskName, content, sessionId}
		logRecordChan <- rowContent
	}
	nextChan <- true
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
		return task.Project.Decoration.InspectIcon
	}
	return task.Project.Decoration.RunIcon
}
