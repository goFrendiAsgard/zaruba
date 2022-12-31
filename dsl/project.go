package dsl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/output"
	yaml "gopkg.in/yaml.v3"
)

// Project is zaruba configuration
type Project struct {
	Name                       string                       `yaml:"name,omitempty"`
	Includes                   []string                     `yaml:"includes,omitempty"`
	Inputs                     map[string]*Variable         `yaml:"inputs,omitempty"`
	Tasks                      map[string]*Task             `yaml:"tasks,omitempty"`
	RawEnvRefMap               map[string]map[string]*Env   `yaml:"envs,omitempty"`
	RawConfigRefMap            map[string]map[string]string `yaml:"configs,omitempty"`
	EnvRefMap                  map[string]*EnvRef           `yaml:"_envrefmap,omitempty"`
	ConfigRefMap               map[string]*ConfigRef        `yaml:"_configrefmap,omitempty"`
	IsInitialized              bool                         `yaml:"_isInitialized,omitempty"`
	StdoutChan                 chan string
	StdoutRecordChan           chan []string
	StderrChan                 chan string
	StderrRecordChan           chan []string
	OutputWg                   *sync.WaitGroup
	OutputWgMutex              *sync.Mutex
	fileLocation               string
	values                     map[string]string
	sortedTaskNames            []string
	sortedInputNames           []string
	maxPublishedTaskNameLength int
	Decoration                 *output.Decoration
	Util                       *DSLUtil
	additionalEnvNames         []string
	additionalValueNames       []string
	showLogTime                bool
	sessionId                  string
}

func NewDefaultProject(projectFile string) (p *Project, err error) {
	return NewProject(projectFile, output.NewDefaultDecoration(), true)
}

func NewProject(projectFile string, decoration *output.Decoration, showLogTime bool) (p *Project, err error) {
	defaultIncludes := []string{"${ZARUBA_HOME}/core.zaruba.yaml"}
	for _, script := range strings.Split(os.Getenv("ZARUBA_SCRIPTS"), ":") {
		if script == "" {
			continue
		}
		defaultIncludes = append(defaultIncludes, script)
	}
	return NewCustomProject(projectFile, decoration, showLogTime, defaultIncludes)
}

// NewCustomProject create new Config from Yaml File
func NewCustomProject(projectFile string, decoration *output.Decoration, showLogTime bool, defaultIncludes []string) (p *Project, err error) {
	p, err = loadProject(decoration, projectFile, defaultIncludes)
	if err != nil {
		return p, err
	}
	p.StdoutChan = make(chan string)
	p.StdoutRecordChan = make(chan []string)
	p.StderrChan = make(chan string)
	p.StderrRecordChan = make(chan []string)
	p.OutputWg = new(sync.WaitGroup)
	p.OutputWgMutex = new(sync.Mutex)
	p.Decoration = decoration
	p.showLogTime = showLogTime
	p.setSortedTaskNames()
	p.setSortedInputNames()
	p.linkToTasks()
	p.linkToInputs()
	p.setDefaultValues()
	p.generateSessionId()
	if err = p.validateTask(); err != nil {
		return p, err
	}
	for _, taskName := range p.sortedTaskNames {
		task := p.Tasks[taskName]
		task.init()
	}
	return p, err
}

func loadProject(d *output.Decoration, projectFile string, defaultIncludes []string) (p *Project, err error) {
	parsedProjectFile, err := parseProjectFile(projectFile)
	if err != nil {
		return p, err
	}
	p, err = newRawProject(parsedProjectFile)
	if err != nil {
		return p, err
	}
	p.include(parsedProjectFile, defaultIncludes)
	p.fileLocation = parsedProjectFile
	p.setTaskFileLocation()
	p.setInputFileLocation()
	p.setProjectEnvRefMap()
	p.setProjectConfigRefMap()
	// cascade project, add inclusion's property to this project
	if err = p.cascadeIncludes(d); err != nil {
		return p, err
	}
	return p, err
}

func parseProjectFile(projectFile string) (parsedProjectFile string, err error) {
	if projectFile == "" {
		return "", nil
	}
	return filepath.Abs(os.ExpandEnv(projectFile))
}

func newRawProject(parsedProjectFile string) (p *Project, err error) {
	p = &Project{
		Includes:                   []string{},
		RawEnvRefMap:               map[string]map[string]*Env{},
		RawConfigRefMap:            map[string]map[string]string{},
		Tasks:                      map[string]*Task{},
		Inputs:                     map[string]*Variable{},
		values:                     map[string]string{},
		EnvRefMap:                  map[string]*EnvRef{},
		ConfigRefMap:               map[string]*ConfigRef{},
		IsInitialized:              false,
		maxPublishedTaskNameLength: 20,
		Util:                       NewDSLUtil(),
	}
	if parsedProjectFile == "" {
		return p, nil
	}
	keyValidator := NewKeyValidator(parsedProjectFile)
	b, err := keyValidator.Validate()
	if err != nil {
		return p, err
	}
	if err = yaml.Unmarshal(b, p); err != nil {
		return p, fmt.Errorf("error parsing YAML '%s': %s", parsedProjectFile, err)
	}
	return p, nil
}

func (p *Project) include(parsedProjectFile string, defaultIncludes []string) {
	for _, script := range defaultIncludes {
		if strings.Trim(script, " ") == "" {
			continue
		}
		parsedScript, _ := filepath.Abs(os.ExpandEnv(script))
		if parsedScript == parsedProjectFile {
			continue
		}
		shouldIncludeParsedScript := true
		for _, currentInclude := range p.Includes {
			parsedCurrentInclude, _ := filepath.Abs(os.ExpandEnv(currentInclude))
			if parsedScript == parsedCurrentInclude {
				shouldIncludeParsedScript = false
				break
			}
		}
		if shouldIncludeParsedScript {
			p.Includes = append(p.Includes, parsedScript)
		}
	}
}

func (p *Project) GetSessionId() (sessionId string) {
	return p.sessionId
}

// GetName get projectName
func (p *Project) GetName() (name string) {
	if p.Name != "" {
		return p.Name
	}
	return filepath.Base(p.GetDirPath())
}

// GetFileLocation
func (p *Project) GetFileLocation() (fileLocation string) {
	return p.fileLocation
}

func (p *Project) GetDirPath() (dirPath string) {
	if p.fileLocation == "" {
		dirPath, _ = os.Getwd()
		return dirPath
	}
	return filepath.Dir(p.fileLocation)
}

// GetSortedInputNames get sorted input names
func (p *Project) GetSortedInputNames() (sortedInputNames []string) {
	return p.sortedInputNames
}

// GetSortedTaskNames get sorted task names
func (p *Project) GetSortedTaskNames() (sortedTaskNames []string) {
	return p.sortedTaskNames
}

// GetValues get value
func (p *Project) GetValues() (values map[string]string) {
	return p.values
}

// GetValue get value
func (p *Project) GetValue(key string) (value string) {
	return p.values[key]
}

// IsValueExist is value exist
func (p *Project) IsValueExist(key string) (exist bool) {
	_, exist = p.values[key]
	return exist
}

// AddEnv add global environment for a projectConfig
func (p *Project) AddEnv(newEnv string) (err error) {
	if p.IsInitialized {
		return fmt.Errorf("cannot AddEnv, project has been initialized")
	}
	p.additionalEnvNames = append(p.additionalEnvNames, newEnv)
	// load env from json string
	envMap := map[string]string{}
	if err := json.Unmarshal([]byte(newEnv), &envMap); err == nil {
		for key := range envMap {
			val := envMap[key]
			os.Setenv(key, val)
		}
		return nil
	}
	// load env from file
	if _, err := os.Stat(newEnv); !os.IsNotExist(err) {
		return godotenv.Load(newEnv)
	}
	// load env from string
	pairParts := strings.SplitN(newEnv, "=", 2)
	if len(pairParts) == 2 {
		key := pairParts[0]
		val := pairParts[1]
		os.Setenv(key, val)
		return nil
	}
	return fmt.Errorf("invalid env: %s", newEnv)
}

func (p *Project) GetAdditionalEnvNames() []string {
	return p.additionalEnvNames
}

// AddValue add value for a project
func (p *Project) AddValue(newValue string) (err error) {
	if p.IsInitialized {
		return fmt.Errorf("cannot AddValue, project has been initialized")
	}
	p.additionalValueNames = append(p.additionalValueNames, newValue)
	// load values from json string
	valueMap := map[string]string{}
	if err := json.Unmarshal([]byte(newValue), &valueMap); err == nil {
		for key := range valueMap {
			val := valueMap[key]
			p.values[key] = val
		}
		return nil
	}
	// load values from file
	if _, err := os.Stat(newValue); !os.IsNotExist(err) {
		b, err := ioutil.ReadFile(newValue)
		if err != nil {
			return err
		}
		keyValues := map[string]string{}
		if err = yaml.Unmarshal(b, keyValues); err != nil {
			return err
		}
		for key, val := range keyValues {
			p.values[key] = val
		}
		return nil
	}
	// load values from string
	pairParts := strings.SplitN(newValue, "=", 2)
	if len(pairParts) == 2 {
		key := pairParts[0]
		val := pairParts[1]
		p.values[key] = val
		return nil
	}
	return fmt.Errorf("invalid value: %s", newValue)
}

// SetValue set value for a project
func (p *Project) SetValue(key, value string) (err error) {
	if p.IsInitialized {
		return fmt.Errorf("cannot SetValue, project has been initialized")
	}
	p.additionalValueNames = append(p.additionalValueNames, fmt.Sprintf("%s=%s", key, value))
	p.setValue(key, value)
	return nil
}

func (p *Project) GetAdditionalValueNames() []string {
	return p.additionalValueNames
}

func (p *Project) setValue(key, value string) {
	p.values[key] = value
}

// GetInputs given task names
func (p *Project) GetInputs(taskNames []string) (inputs map[string]*Variable, inputOrder []string, err error) {
	inputs = map[string]*Variable{}
	inputOrder = []string{}
	for _, taskName := range taskNames {
		task, taskExist := p.Tasks[taskName]
		if !taskExist {
			return inputs, inputOrder, fmt.Errorf("task '%s' does not exist", taskName)
		}
		// include task's dependencies
		dependencyTaskNames := []string{}
		dependencyTaskNames = append(dependencyTaskNames, task.Dependencies...)
		dependencyInputs, dependencyInputOrder, _ := p.GetInputs(dependencyTaskNames)
		for _, inputName := range dependencyInputOrder {
			subInput := dependencyInputs[inputName]
			if _, inputRegistered := inputs[inputName]; !inputRegistered {
				inputOrder = append(inputOrder, inputName)
				inputs[inputName] = subInput
			}
		}
		// include task's inputs
		for _, inputName := range task.Inputs {
			input := p.Inputs[inputName]
			if _, inputRegistered := inputs[inputName]; !inputRegistered {
				inputOrder = append(inputOrder, inputName)
				inputs[inputName] = input
			}
		}
	}
	return inputs, inputOrder, err
}

// Init all tasks
func (p *Project) Init() (err error) {
	for key, value := range p.values {
		parsedValue := os.ExpandEnv(value)
		// validate (allow empty value, but throw error if value is set and invalid)
		if parsedValue != "" {
			if input, inputExist := p.Inputs[key]; inputExist {
				if err = input.Validate(parsedValue); err != nil {
					return err
				}
			}
		}
		// inject envvars (useful for secret inputs)
		p.values[key] = parsedValue
		inputEnvKey := "ZARUBA_INPUT_" + p.Util.Str.ToUpperSnake(key)
		os.Setenv(inputEnvKey, parsedValue)
	}
	p.IsInitialized = true
	return err
}

func (p *Project) GetAutoTerminate(taskNames []string) (autoTerminate bool) {
	for _, taskName := range taskNames {
		if !p.Tasks[taskName].GetAutoTerminate() {
			return false
		}
	}
	return true
}

// ValidateByTaskNames validate by task names and throw error if invalid
func (p *Project) ValidateByTaskNames(taskNames []string) (err error) {
	for _, taskName := range taskNames {
		task, taskExist := p.Tasks[taskName]
		if !taskExist {
			return fmt.Errorf("task '%s' does not exist", taskName)
		}
		for _, inputName := range task.Inputs {
			value := p.values[inputName]
			if input, inputExist := p.Inputs[inputName]; inputExist {
				if err = input.Validate(value); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *Project) GetValueFileNames() (valueFileNames []string, err error) {
	projectDir := filepath.Dir(p.GetFileLocation())
	files, err := ioutil.ReadDir(projectDir)
	if err != nil {
		return valueFileNames, err
	}
	valueFileNames = []string{}
	for _, file := range files {
		isDir := file.IsDir()
		if isDir {
			continue
		}
		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".values.yaml") || fileName == ".previous.values.yaml" {
			continue
		}
		valueFileNames = append(valueFileNames, filepath.Join(projectDir, fileName))
	}
	return valueFileNames, nil
}

func (p *Project) GetEnvFileNames() (envFileNames []string, err error) {
	projectDir := filepath.Dir(p.GetFileLocation())
	files, err := ioutil.ReadDir(projectDir)
	if err != nil {
		return envFileNames, err
	}
	envFileNames = []string{}
	for _, file := range files {
		isDir := file.IsDir()
		if isDir {
			continue
		}
		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".env") && !strings.HasSuffix(fileName, ".env.template") {
			continue
		}
		envFileNames = append(envFileNames, filepath.Join(projectDir, fileName))
	}
	return envFileNames, nil
}

func (p *Project) setProjectEnvRefMap() {
	for envRefName, envRefMap := range p.RawEnvRefMap {
		p.EnvRefMap[envRefName] = &EnvRef{
			fileLocation: p.fileLocation,
			name:         envRefName,
			Map:          envRefMap,
		}
	}
}

func (p *Project) setProjectConfigRefMap() {
	for configRefName, configRefMap := range p.RawConfigRefMap {
		p.ConfigRefMap[configRefName] = &ConfigRef{
			fileLocation: p.fileLocation,
			name:         configRefName,
			Map:          configRefMap,
		}
	}
}

func (p *Project) setTaskFileLocation() {
	for _, task := range p.Tasks {
		task.fileLocation = p.fileLocation
	}
}

func (p *Project) setInputFileLocation() {
	for _, input := range p.Inputs {
		input.fileLocation = p.fileLocation
	}
}

func (p *Project) setDefaultValues() {
	for inputName, input := range p.Inputs {
		p.setValue(inputName, input.DefaultValue)
	}
}

func (p *Project) validateTask() (err error) {
	if err = p.validateTaskExtend(); err != nil {
		return err
	}
	if err = p.validateTaskDependencies(); err != nil {
		return err
	}
	if err = p.validateTaskAcyclic(); err != nil {
		return err
	}
	if err = p.validateTaskInputs(); err != nil {
		return err
	}
	if err = p.validateTaskEnvRef(); err != nil {
		return err
	}
	if err = p.validateTaskConfigRef(); err != nil {
		return err
	}
	if err = p.validateTaskAutoTerminte(); err != nil {
		return err
	}
	return nil
}

func (p *Project) validateTaskAutoTerminte() (err error) {
	for _, task := range p.Tasks {
		autoTerminate := task.GetAutoTerminate()
		haveCheckCmd := task.IsHavingCheckCmd()
		if haveCheckCmd && autoTerminate {
			return fmt.Errorf("restricted autoTerminate value on '%s': Task '%s' is autoTerminate, but tasks[%s][check] is not empty (also check task's parents)", task.GetFileLocation(), task.GetName(), task.GetName())
		}
	}
	return nil
}

func (p *Project) validateTaskAcyclic() (err error) {
	for _, task := range p.Tasks {
		isRecursive, recursiveTaskname := p.isTaskRecursive(task, []string{})
		if isRecursive {
			recursiveTask := p.Tasks[recursiveTaskname]
			return fmt.Errorf("recursive task on '%s': Task '%s' is recursively need itself", recursiveTask.GetFileLocation(), recursiveTask.GetName())
		}
	}
	return nil
}

func (p *Project) isTaskRecursive(task *Task, previousTaskNames []string) (isRecursive bool, recursiveTaskName string) {
	taskName := task.GetName()
	for _, previousTaskName := range previousTaskNames {
		if previousTaskName == taskName {
			return true, taskName
		}
	}
	previousTaskNames = append(previousTaskNames, task.GetName())
	for _, dependencyTaskName := range task.Dependencies {
		subTask := p.Tasks[dependencyTaskName]
		if isRecursive, recursiveTaskName := p.isTaskRecursive(subTask, previousTaskNames); isRecursive {
			return true, recursiveTaskName
		}
	}
	for _, parentTaskName := range task.GetParentTaskNames() {
		subTask := p.Tasks[parentTaskName]
		if isRecursive, recursiveTaskName := p.isTaskRecursive(subTask, previousTaskNames); isRecursive {
			return true, recursiveTaskName
		}
	}
	return false, ""
}

func (p *Project) validateTaskDependencies() (err error) {
	for taskName, task := range p.Tasks {
		for index, dependencyTaskName := range task.Dependencies {
			if _, dependencyTaskExist := p.Tasks[dependencyTaskName]; !dependencyTaskExist {
				return fmt.Errorf("undeclared task dependency on '%s': Task '%s' is required at tasks[%s][dependencies][%d]", task.GetFileLocation(), dependencyTaskName, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) validateTaskExtend() (err error) {
	for taskName, task := range p.Tasks {
		if len(task.Extends) > 0 && task.Extend != "" {
			return fmt.Errorf("redundant key declaration on '%s': Task '%s' has both `extend` and `extends`", task.GetFileLocation(), taskName)
		}
		if task.Extend != "" {
			if _, parentTaskExist := p.Tasks[task.Extend]; !parentTaskExist {
				return fmt.Errorf("undeclared parent task on '%s': Task '%s' is required at tasks[%s][extend]", task.GetFileLocation(), task.Extend, taskName)
			}
		}
		for index, parentTaskName := range task.Extends {
			if _, parentTaskExist := p.Tasks[parentTaskName]; !parentTaskExist {
				return fmt.Errorf("undeclared parent task on '%s': Task '%s' is required at tasks[%s][extends][%d]", task.GetFileLocation(), parentTaskName, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) validateTaskInputs() (err error) {
	for taskName, task := range p.Tasks {
		for index, inputName := range task.Inputs {
			if _, inputExist := p.Inputs[inputName]; !inputExist {
				return fmt.Errorf("undeclared input task on '%s': Input '%s' is required at %s[inputs][%d]", task.GetFileLocation(), inputName, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) validateTaskEnvRef() (err error) {
	for taskName, task := range p.Tasks {
		if len(task.EnvRefs) > 0 && task.EnvRef != "" {
			return fmt.Errorf("redundant key declaration on '%s': Task '%s' has both `envRef` and `envRefs`", task.GetFileLocation(), taskName)
		}
		if task.EnvRef != "" {
			if _, envRefExist := p.EnvRefMap[task.EnvRef]; !envRefExist {
				return fmt.Errorf("undeclared envRef on '%s': Env '%s' is required at tasks[%s][envRef]", task.GetFileLocation(), task.EnvRef, taskName)
			}
		}
		for index, envRefName := range task.EnvRefs {
			if _, envRefExist := p.EnvRefMap[envRefName]; !envRefExist {
				return fmt.Errorf("undeclared envRefs on '%s': Env '%s' is required at tasks[%s][envRefs][%d]", task.GetFileLocation(), envRefName, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) validateTaskConfigRef() (err error) {
	for taskName, task := range p.Tasks {
		if len(task.ConfigRefs) > 0 && task.ConfigRef != "" {
			return fmt.Errorf("redundant key declaration on '%s': Task '%s' has both `config` and `configRefs`", task.GetFileLocation(), taskName)
		}
		if task.ConfigRef != "" {
			if _, configRefExist := p.ConfigRefMap[task.ConfigRef]; !configRefExist {
				return fmt.Errorf("undeclared configRef on '%s': Config '%s' is required at tasks[%s][configRef]", task.GetFileLocation(), task.ConfigRef, taskName)
			}
		}
		for index, configRefName := range task.ConfigRefs {
			if _, configRefExist := p.ConfigRefMap[configRefName]; !configRefExist {
				return fmt.Errorf("undeclared configRefs on '%s': Config '%s' is required at tasks[%s][configRefs][%d]", task.GetFileLocation(), configRefName, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) cascadeIncludes(d *output.Decoration) (err error) {
	for _, includeLocation := range p.Includes {
		parsedIncludeLocation := os.ExpandEnv(includeLocation)
		if !filepath.IsAbs(parsedIncludeLocation) {
			parsedIncludeLocation = filepath.Join(filepath.Dir(p.fileLocation), parsedIncludeLocation)
		}
		includedProject, err := loadProject(d, parsedIncludeLocation, []string{})
		if err != nil {
			return err
		}
		if err = p.cascadeInputs(parsedIncludeLocation, includedProject); err != nil {
			return err
		}
		if err = p.cascadeTasks(parsedIncludeLocation, includedProject); err != nil {
			return err
		}
		if err = p.cascadeEnvRef(parsedIncludeLocation, includedProject); err != nil {
			return err
		}
		if err = p.cascadeConfigRef(parsedIncludeLocation, includedProject); err != nil {
			return err
		}
	}
	return err
}

func (p *Project) cascadeInputs(parsedIncludeLocation string, includedProject *Project) (err error) {
	for inputName, input := range includedProject.Inputs {
		existingInput, inputAlreadyDeclared := p.Inputs[inputName]
		if inputAlreadyDeclared {
			if input.fileLocation == existingInput.fileLocation {
				continue
			}
			return fmt.Errorf("redundant input declaration on '%s': Input '%s' was already declared on '%s'", input.fileLocation, inputName, existingInput.fileLocation)
		}
		p.Inputs[inputName] = input
	}
	return nil
}

func (p *Project) cascadeTasks(parsedIncludeLocation string, includedProject *Project) (err error) {
	for taskName, task := range includedProject.Tasks {
		existingTask, taskAlreadyDeclared := p.Tasks[taskName]
		if taskAlreadyDeclared {
			if task.fileLocation == existingTask.fileLocation {
				continue
			}
			return fmt.Errorf("redundant task declaration on '%s': Task '%s' was already declared on '%s'", task.fileLocation, taskName, existingTask.fileLocation)
		}
		p.Tasks[taskName] = task
	}
	return nil
}

func (p *Project) cascadeEnvRef(parsedIncludeLocation string, includedProject *Project) (err error) {
	for envRefName, envRef := range includedProject.EnvRefMap {
		existingEnvRef, envRefAlreadyDeclared := p.EnvRefMap[envRefName]
		if envRefAlreadyDeclared {
			if envRef.fileLocation == existingEnvRef.fileLocation {
				continue
			}
			return fmt.Errorf("redundant envs declaration on '%s': Env '%s' was already declared on '%s'", envRef.fileLocation, envRefName, existingEnvRef.fileLocation)
		}
		p.EnvRefMap[envRefName] = envRef
	}
	return nil
}

func (p *Project) cascadeConfigRef(parsedIncludeLocation string, includedProject *Project) (err error) {
	for configRefName, configRef := range includedProject.ConfigRefMap {
		existingConfigRef, configRefAlreadyDeclared := p.ConfigRefMap[configRefName]
		if configRefAlreadyDeclared {
			if configRef.fileLocation == existingConfigRef.fileLocation {
				continue
			}
			return fmt.Errorf("redundant configs declaration on '%s': Config '%s' was already declared '%s'", configRef.fileLocation, configRefName, existingConfigRef.fileLocation)
		}
		p.ConfigRefMap[configRefName] = configRef
	}
	return nil
}

func (p *Project) generateSessionId() {
	p.sessionId = fmt.Sprintf("%s-%s", p.Util.Str.NewName(), p.Util.Str.NewUUID())
}

func (p *Project) setSortedTaskNames() {
	p.sortedTaskNames = []string{}
	for taskName := range p.Tasks {
		p.sortedTaskNames = append(p.sortedTaskNames, taskName)
	}
	sort.Strings(p.sortedTaskNames)
}

func (p *Project) setSortedInputNames() {
	p.sortedInputNames = []string{}
	for inputName := range p.Inputs {
		p.sortedInputNames = append(p.sortedInputNames, inputName)
	}
	sort.Strings(p.sortedInputNames)
}

func (p *Project) linkToTasks() {
	for taskName, task := range p.Tasks {
		task.Project = p
		task.name = taskName
		task.linkToEnvs()
	}
}

func (p *Project) linkToInputs() {
	for inputName, input := range p.Inputs {
		input.Project = p
		input.name = inputName
	}
}

func (p *Project) ProcessOutputWg() {
	p.OutputWgMutex.Lock()
	p.OutputWg.Done()
	p.OutputWgMutex.Unlock()
}
