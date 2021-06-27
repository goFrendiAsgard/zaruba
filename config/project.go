package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/output"
	yaml "gopkg.in/yaml.v2"
)

// Project is zaruba configuration
type Project struct {
	Includes                   []string                     `yaml:"includes,omitempty"`
	Tasks                      map[string]*Task             `yaml:"tasks,omitempty"`
	Name                       string                       `yaml:"name,omitempty"`
	Inputs                     map[string]*Variable         `yaml:"inputs,omitempty"`
	RawEnvRefMap               map[string]map[string]Env    `yaml:"envs,omitempty"`
	RawConfigRefMap            map[string]map[string]string `yaml:"configs,omitempty"`
	EnvRefMap                  map[string]EnvRef
	ConfigRefMap               map[string]ConfigRef
	fileLocation               string
	values                     map[string]string
	sortedTaskNames            []string
	sortedInputNames           []string
	maxPublishedTaskNameLength int
	decoration                 *output.Decoration
	logger                     output.Logger
	dataLogger                 output.RecordLogger
	IsInitialized              bool
}

// NewProject create new Config from Yaml File
func NewProject(logger output.Logger, dataLogger output.RecordLogger, decoration *output.Decoration, projectFile string, defaultIncludes []string) (p *Project, err error) {
	if os.Getenv("ZARUBA_HOME") == "" {
		executable, _ := os.Executable()
		os.Setenv("ZARUBA_HOME", filepath.Dir(executable))
	}
	p, err = loadProject(logger, decoration, projectFile, defaultIncludes)
	if err != nil {
		return p, err
	}
	p.logger = logger
	p.dataLogger = dataLogger
	p.decoration = decoration
	p.setSortedTaskNames()
	p.setSortedInputNames()
	p.linkToTasks()
	p.linkToInputs()
	p.setDefaultValues()
	if err = p.validateTask(); err != nil {
		return p, err
	}
	for _, taskName := range p.sortedTaskNames {
		task := p.Tasks[taskName]
		task.init()
	}
	return p, err
}

func loadProject(logger output.Logger, d *output.Decoration, projectFile string, defaultIncludes []string) (p *Project, err error) {
	parsedProjectFile, _ := filepath.Abs(os.ExpandEnv(projectFile))
	p = &Project{
		Includes:                   []string{},
		RawEnvRefMap:               map[string]map[string]Env{},
		RawConfigRefMap:            map[string]map[string]string{},
		Tasks:                      map[string]*Task{},
		Inputs:                     map[string]*Variable{},
		values:                     map[string]string{},
		EnvRefMap:                  map[string]EnvRef{},
		ConfigRefMap:               map[string]ConfigRef{},
		IsInitialized:              false,
		maxPublishedTaskNameLength: 17,
	}
	keyValidator := NewKeyValidator(parsedProjectFile)
	b, err := keyValidator.Validate()
	if err != nil {
		return p, err
	}
	if err = yaml.Unmarshal(b, p); err != nil {
		return p, fmt.Errorf("error parsing YAML '%s': %s", parsedProjectFile, err)
	}
	p.include(parsedProjectFile, defaultIncludes)
	p.fileLocation = parsedProjectFile
	p.setTaskFileLocation()
	p.setInputFileLocation()
	p.setProjectEnvRefMap()
	p.setProjectConfigRefMap()
	// cascade project, add inclusion's property to this project
	if err = p.cascadeIncludes(logger, d); err != nil {
		return p, err
	}
	return p, err
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

// GetName get projectName
func (p *Project) GetName() (name string) {
	if p.Name != "" {
		return p.Name
	}
	return filepath.Base(filepath.Dir(p.fileLocation))
}

// GetFileLocation
func (p *Project) GetFileLocation() (fileLocation string) {
	return p.fileLocation
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

// AddGlobalEnv add global environment for a projectConfig
func (p *Project) AddGlobalEnv(pairOrFile string) (err error) {
	if p.IsInitialized {
		return fmt.Errorf("cannot AddGlobalEnv, project has been initialized")
	}
	pairParts := strings.SplitN(pairOrFile, "=", 2)
	if len(pairParts) == 2 {
		key := pairParts[0]
		val := pairParts[1]
		os.Setenv(key, val)
		return nil
	}
	return godotenv.Load(pairOrFile)
}

// AddValue add value for a project
func (p *Project) AddValue(pairOrFile string) (err error) {
	if p.IsInitialized {
		return fmt.Errorf("cannot AddValue, project has been initialized")
	}
	pairParts := strings.SplitN(pairOrFile, "=", 2)
	if len(pairParts) == 2 {
		key := pairParts[0]
		val := pairParts[1]
		p.values[key] = val
		return nil
	}
	b, err := ioutil.ReadFile(pairOrFile)
	if err != nil {
		return err
	}
	values := map[string]string{}
	if err = yaml.Unmarshal(b, values); err != nil {
		return err
	}
	for key, val := range values {
		p.values[key] = val
	}
	return nil
}

// SetValue set value for a project
func (p *Project) SetValue(key, value string) (err error) {
	if p.IsInitialized {
		return fmt.Errorf("cannot SetValue, project has been initialized")
	}
	p.values[key] = value
	return nil
}

// GetInputs given task names
func (p *Project) GetInputs(taskNames []string) (inputs map[string]*Variable, inputOrder []string, err error) {
	inputs = map[string]*Variable{}
	inputOrder = []string{}
	for _, taskName := range taskNames {
		task, taskExist := p.Tasks[taskName]
		if !taskExist {
			return inputs, inputOrder, fmt.Errorf("task '%s' is not exist", taskName)
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
	r, _ := regexp.Compile("[^a-zA-Z0-9]+")
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
		inputEnvKey := "ZARUBA_INPUT_" + string(r.ReplaceAll([]byte(strings.ToUpper(key)), []byte("_")))
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
			return fmt.Errorf("task '%s' is not exist", taskName)
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

func (p *Project) setProjectEnvRefMap() {
	for baseEnvName, baseEnvMap := range p.RawEnvRefMap {
		p.EnvRefMap[baseEnvName] = EnvRef{
			fileLocation: p.fileLocation,
			name:         baseEnvName,
			Map:          baseEnvMap,
		}
	}
}

func (p *Project) setProjectConfigRefMap() {
	for baseConfigName, baseConfigMap := range p.RawConfigRefMap {
		p.ConfigRefMap[baseConfigName] = ConfigRef{
			fileLocation: p.fileLocation,
			name:         baseConfigName,
			Map:          baseConfigMap,
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
		p.SetValue(inputName, input.DefaultValue)
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
		haveCheckCmd := task.HaveCheckCmd()
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
	for _, parentTaskName := range task.getParentTaskNames() {
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
			if _, baseEnvExist := p.EnvRefMap[task.EnvRef]; !baseEnvExist {
				return fmt.Errorf("undeclared envRef on '%s': Env '%s' is required at tasks[%s][envRef]", task.GetFileLocation(), task.EnvRef, taskName)
			}
			return nil
		}
		for index, baseEnvKey := range task.EnvRefs {
			if _, baseEnvExist := p.EnvRefMap[baseEnvKey]; !baseEnvExist {
				return fmt.Errorf("undeclared envRefs on '%s': Env '%s' is required at tasks[%s][envRefs][%d]", task.GetFileLocation(), baseEnvKey, taskName, index)
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
			if _, baseConfigExist := p.ConfigRefMap[task.ConfigRef]; !baseConfigExist {
				return fmt.Errorf("undeclared configRef on '%s': Config '%s' is required at tasks[%s][configRef]", task.GetFileLocation(), task.ConfigRef, taskName)
			}
			return nil
		}
		for index, baseConfigKey := range task.ConfigRefs {
			if _, baseConfigExist := p.ConfigRefMap[baseConfigKey]; !baseConfigExist {
				return fmt.Errorf("undeclared configRefs on '%s': Config '%s' is required at tasks[%s][configRefs][%d]", task.GetFileLocation(), baseConfigKey, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) cascadeIncludes(logger output.Logger, d *output.Decoration) (err error) {
	for _, includeLocation := range p.Includes {
		parsedIncludeLocation := os.ExpandEnv(includeLocation)
		if !filepath.IsAbs(parsedIncludeLocation) {
			parsedIncludeLocation = filepath.Join(filepath.Dir(p.fileLocation), parsedIncludeLocation)
		}
		includedProject, err := loadProject(logger, d, parsedIncludeLocation, []string{})
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
		existingBaseEnv, baseEnvAlreadyDeclared := p.EnvRefMap[envRefName]
		if baseEnvAlreadyDeclared {
			if envRef.fileLocation == existingBaseEnv.fileLocation {
				continue
			}
			return fmt.Errorf("redundant envs declaration on '%s': Env ref '%s' was already declared on '%s'", envRef.fileLocation, envRefName, existingBaseEnv.fileLocation)
		}
		p.EnvRefMap[envRefName] = envRef
	}
	return nil
}

func (p *Project) cascadeConfigRef(parsedIncludeLocation string, includedProject *Project) (err error) {
	for configRefName, configRef := range includedProject.ConfigRefMap {
		existingBaseConfig, baseConfigAlreadyDeclared := p.ConfigRefMap[configRefName]
		if baseConfigAlreadyDeclared {
			if configRef.fileLocation == existingBaseConfig.fileLocation {
				continue
			}
			return fmt.Errorf("redundant configs declaration on '%s': Config ref '%s' was already declared '%s'", configRef.fileLocation, configRefName, existingBaseConfig.fileLocation)
		}
		p.ConfigRefMap[configRefName] = configRef
	}
	return nil
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
