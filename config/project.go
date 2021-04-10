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
	"github.com/state-alchemists/zaruba/monitor"
	yaml "gopkg.in/yaml.v2"
)

// Project is zaruba configuration
type Project struct {
	Includes                   []string                       `yaml:"includes,omitempty"`
	Tasks                      map[string]*Task               `yaml:"tasks,omitempty"`
	Name                       string                         `yaml:"name,omitempty"`
	Inputs                     map[string]*Input              `yaml:"inputs,omitempty"`
	RawBaseEnv                 map[string]map[string]BaseEnv  `yaml:"envs,omitempty"`
	RawBaseConfig              map[string]map[string]string   `yaml:"configs,omitempty"`
	RawBaseLConfig             map[string]map[string][]string `yaml:"lconfigs,omitempty"`
	baseEnv                    map[string]ProjectBaseEnv
	baseConfig                 map[string]ProjectBaseConfig
	baseLConfig                map[string]ProjectBaseLConfig
	fileLocation               string
	basePath                   string
	values                     map[string]string
	sortedTaskNames            []string
	sortedInputNames           []string
	maxPublishedTaskNameLength int
	decoration                 *monitor.Decoration
	logger                     monitor.Logger
	dataLogger                 monitor.RecordLogger
	IsInitialized              bool
}

// NewProject create new Config from Yaml File
func NewProject(logger monitor.Logger, dataLogger monitor.RecordLogger, decoration *monitor.Decoration, configFile string) (p *Project, err error) {
	if os.Getenv("ZARUBA_HOME") == "" {
		executable, err := os.Executable()
		if err != nil {
			return p, err
		}
		os.Setenv("ZARUBA_HOME", filepath.Dir(executable))
	}
	p, err = loadProject(logger, decoration, configFile)
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
	if err = p.setDefaultValues(); err != nil {
		return p, err
	}
	err = p.validate()
	return p, err
}

func loadProject(logger monitor.Logger, d *monitor.Decoration, projectFile string) (p *Project, err error) {
	parsedProjectFile := os.ExpandEnv(projectFile)
	logger.DPrintfStarted("%sLoading %s%s\n", d.Faint, parsedProjectFile, d.Normal)
	p = &Project{
		Includes:       []string{},
		RawBaseEnv:     map[string]map[string]BaseEnv{},
		RawBaseConfig:  map[string]map[string]string{},
		RawBaseLConfig: map[string]map[string][]string{},
		Tasks:          map[string]*Task{},
		Inputs:         map[string]*Input{},
		values:         map[string]string{},
		baseEnv:        map[string]ProjectBaseEnv{},
		baseConfig:     map[string]ProjectBaseConfig{},
		baseLConfig:    map[string]ProjectBaseLConfig{},
		IsInitialized:  false,
	}
	b, err := ioutil.ReadFile(parsedProjectFile)
	if err != nil {
		return p, err
	}
	if err = yaml.Unmarshal(b, p); err != nil {
		return p, err
	}
	p.reverseInclusion() // we need to reverse inclusion, so that the first include file will always be overridden by the later
	p.fileLocation = parsedProjectFile
	if !filepath.IsAbs(p.fileLocation) {
		p.fileLocation, err = filepath.Abs(p.fileLocation)
		if err != nil {
			return p, err
		}
	}
	p.basePath = filepath.Dir(p.fileLocation)
	p.setTaskFileLocation()
	p.setInputFileLocation()
	p.setProjectBaseEnv()
	p.setProjectBaseConfig()
	p.setProjectBaseLConfig()
	// cascade project, add inclusion's property to this project
	if err = p.cascadeIncludes(logger, d); err != nil {
		return p, err
	}
	return p, err
}

func (p *Project) reverseInclusion() {
	headIndex := 0
	tailIndex := len(p.Includes) - 1
	for headIndex < tailIndex {
		p.Includes[headIndex], p.Includes[tailIndex] = p.Includes[tailIndex], p.Includes[headIndex]
		headIndex++
		tailIndex--
	}
}

// GetName get projectName
func (p *Project) GetName() (name string) {
	if p.Name != "" {
		return p.Name
	}
	return filepath.Base(filepath.Dir(p.fileLocation))
}

// GetBasePath get basePath
func (p *Project) GetBasePath() (basePath string) {
	return p.basePath
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
		return fmt.Errorf("Cannot AddGlobalEnv, project has been initialized")
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
		return fmt.Errorf("Cannot AddValue, project has been initialized")
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
		return fmt.Errorf("Cannot SetValue, project has been initialized")
	}
	p.values[key] = value
	return nil
}

// GetInputs given task names
func (p *Project) GetInputs(taskNames []string) (inputs map[string]*Input, inputOrder []string, err error) {
	inputs = map[string]*Input{}
	inputOrder = []string{}
	for _, taskName := range taskNames {
		task, taskExist := p.Tasks[taskName]
		if !taskExist {
			return inputs, inputOrder, fmt.Errorf("Task %s is not exist", taskName)
		}
		// include task's dependencies and parent's inputs first
		subTaskNames := []string{}
		if task.Extend != "" {
			subTaskNames = append(subTaskNames, task.Extend)
		}
		subTaskNames = append(subTaskNames, task.Dependencies...)
		subInputs, subInputOrder, err := p.GetInputs(subTaskNames)
		if err != nil {
			return inputs, inputOrder, err
		}
		for _, inputName := range subInputOrder {
			subInput := subInputs[inputName]
			if _, inputRegistered := inputs[inputName]; !inputRegistered {
				inputOrder = append(inputOrder, inputName)
				inputs[inputName] = subInput
			}
		}
		// include task's inputs
		for _, inputName := range task.Inputs {
			input := p.Inputs[inputName]
			if _, inputRegistered := inputs[inputName]; !inputRegistered {
				inputs[inputName] = input
				inputOrder = append(inputOrder, inputName)
			}
		}
	}
	return inputs, inputOrder, err
}

// Init all tasks
func (p *Project) Init() (err error) {
	for _, taskName := range p.sortedTaskNames {
		task := p.Tasks[taskName]
		err := task.init()
		if err != nil {
			return err
		}
	}
	r, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return err
	}
	for key, value := range p.values {
		parsedValue := os.ExpandEnv(value)
		p.values[key] = parsedValue
		inputEnvKey := "ZARUBA_INPUT_" + string(r.ReplaceAll([]byte(strings.ToUpper(key)), []byte("_")))
		os.Setenv(inputEnvKey, parsedValue)
	}
	p.IsInitialized = true
	return err
}

func (p *Project) setProjectBaseEnv() {
	for baseEnvName, baseEnvMap := range p.RawBaseEnv {
		p.baseEnv[baseEnvName] = ProjectBaseEnv{
			fileLocation: p.fileLocation,
			name:         baseEnvName,
			BaseEnvMap:   baseEnvMap,
		}
	}
}

func (p *Project) setProjectBaseConfig() {
	for baseConfigName, baseConfigMap := range p.RawBaseConfig {
		p.baseConfig[baseConfigName] = ProjectBaseConfig{
			fileLocation:  p.fileLocation,
			name:          baseConfigName,
			BaseConfigMap: baseConfigMap,
		}
	}
}

func (p *Project) setProjectBaseLConfig() {
	for baseLConfigName, baseLConfigMap := range p.RawBaseLConfig {
		p.baseLConfig[baseLConfigName] = ProjectBaseLConfig{
			fileLocation:   p.fileLocation,
			name:           baseLConfigName,
			BaseLConfigMap: baseLConfigMap,
		}
	}
}

func (p *Project) setTaskFileLocation() {
	for _, task := range p.Tasks {
		task.fileLocation = p.fileLocation
		task.basePath = p.basePath
	}
}

func (p *Project) setInputFileLocation() {
	for _, input := range p.Inputs {
		input.fileLocation = p.fileLocation
	}
}

func (p *Project) setDefaultValues() (err error) {
	for inputName, input := range p.Inputs {
		if err = p.SetValue(inputName, input.DefaultValue); err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) validate() (err error) {
	if err = p.validateTaskExtend(); err != nil {
		return err
	}
	if err = p.validateTaskInputs(); err != nil {
		return err
	}
	if err = p.validateTaskBaseEnv(); err != nil {
		return err
	}
	if err = p.validateTaskBaseConfig(); err != nil {
		return err
	}
	if err = p.validateTaskBaseLConfig(); err != nil {
		return err
	}
	return nil
}

func (p *Project) validateTaskExtend() (err error) {
	for taskName, task := range p.Tasks {
		if len(task.Extends) > 0 && task.Extend != "" {
			return fmt.Errorf("Redundant declaration. %s has both `extend` and `extends`", taskName)
		}
		if task.Extend != "" {
			if _, parentTaskExist := p.Tasks[task.Extend]; !parentTaskExist {
				return fmt.Errorf("Task %s is required at %s.extend but it was not declared", task.Extend, taskName)
			}
		}
		for index, parentTaskName := range task.Extends {
			if _, parentTaskExist := p.Tasks[parentTaskName]; !parentTaskExist {
				return fmt.Errorf("Task %s is required at %s.extends[%d] but it was not declared", parentTaskName, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) validateTaskInputs() (err error) {
	for taskName, task := range p.Tasks {
		for index, inputName := range task.Inputs {
			if _, inputExist := p.Inputs[inputName]; !inputExist {
				return fmt.Errorf("Input %s is required for %s.inputs[%d] but it was not declared", inputName, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) validateTaskBaseEnv() (err error) {
	for taskName, task := range p.Tasks {
		if len(task.EnvRefs) > 0 && task.EnvRef != "" {
			return fmt.Errorf("Redundant declaration. %s has both `envRef` and `envRefs`", taskName)
		}
		if task.EnvRef != "" {
			if _, baseEnvExist := p.baseEnv[task.EnvRef]; !baseEnvExist {
				return fmt.Errorf("Env %s is required at %s.envRef but it was not declared", task.EnvRef, taskName)
			}
			return nil
		}
		for index, baseEnvKey := range task.EnvRefs {
			if _, baseEnvExist := p.baseEnv[baseEnvKey]; !baseEnvExist {
				return fmt.Errorf("Env %s is required at %s.envRefs[%d] but it was not declared", baseEnvKey, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) validateTaskBaseConfig() (err error) {
	for taskName, task := range p.Tasks {
		if len(task.ConfigRefs) > 0 && task.ConfigRef != "" {
			return fmt.Errorf("Redundant declaration. %s has both `configRef` and `configRefs`", taskName)
		}
		if task.ConfigRef != "" {
			if _, baseConfigExist := p.baseConfig[task.ConfigRef]; !baseConfigExist {
				return fmt.Errorf("Config %s is required at %s.configRef but it was not declared", task.ConfigRef, taskName)
			}
			return nil
		}
		for index, baseConfigKey := range task.ConfigRefs {
			if _, baseConfigExist := p.baseConfig[baseConfigKey]; !baseConfigExist {
				return fmt.Errorf("Config %s is required at %s.configRefs[%d] but it was not declared", baseConfigKey, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) validateTaskBaseLConfig() (err error) {
	for taskName, task := range p.Tasks {
		if len(task.LConfigRefs) > 0 && task.LConfigRef != "" {
			return fmt.Errorf("Redundant declaration. %s has both `lconfigRef` and `lconfigRefs`", taskName)
		}
		if task.LConfigRef != "" {
			if _, baseLConfigExist := p.baseLConfig[task.LConfigRef]; !baseLConfigExist {
				return fmt.Errorf("LConfig %s is required at %s.lconfigRef but it was not declared", task.LConfigRef, taskName)
			}
			return nil
		}
		for index, baseLConfigKey := range task.LConfigRefs {
			if _, baseLConfigExist := p.baseLConfig[baseLConfigKey]; !baseLConfigExist {
				return fmt.Errorf("LConfig %s is required at %s.lconfigRefs[%d] but it was not declared", baseLConfigKey, taskName, index)
			}
		}
	}
	return nil
}

func (p *Project) cascadeIncludes(logger monitor.Logger, d *monitor.Decoration) (err error) {
	for _, includeLocation := range p.Includes {
		parsedIncludeLocation := os.ExpandEnv(includeLocation)
		if !filepath.IsAbs(parsedIncludeLocation) {
			parsedIncludeLocation = filepath.Join(p.basePath, parsedIncludeLocation)
		}
		includedProject, err := loadProject(logger, d, parsedIncludeLocation)
		if err != nil {
			return err
		}
		if err = p.cascadeInputs(parsedIncludeLocation, includedProject); err != nil {
			return err
		}
		if err = p.cascadeTasks(parsedIncludeLocation, includedProject); err != nil {
			return err
		}
		if err = p.cascadeBaseEnv(parsedIncludeLocation, includedProject); err != nil {
			return err
		}
		if err = p.cascadeBaseConfig(parsedIncludeLocation, includedProject); err != nil {
			return err
		}
		if err = p.cascadeBaseLConfig(parsedIncludeLocation, includedProject); err != nil {
			return err
		}
	}
	return err
}

func (p *Project) cascadeInputs(parsedIncludeLocation string, includedProject *Project) (err error) {
	for inputName, input := range includedProject.Inputs {
		_, inputAlreadyDeclared := p.Inputs[inputName]
		if inputAlreadyDeclared {
			return fmt.Errorf("Cannot declare input `%s` on `%s` because the input was already declared on `%s`", inputName, parsedIncludeLocation, input.fileLocation)
		}
		p.Inputs[inputName] = input
	}
	return nil
}

func (p *Project) cascadeTasks(parsedIncludeLocation string, includedProject *Project) (err error) {
	for taskName, task := range includedProject.Tasks {
		existingTask, taskAlreadyDeclared := p.Tasks[taskName]
		if taskAlreadyDeclared {
			return fmt.Errorf("Cannot declare task `%s` on `%s` because it was already declared on `%s`", taskName, parsedIncludeLocation, existingTask.GetFileLocation())
		}
		p.Tasks[taskName] = task
	}
	return nil
}

func (p *Project) cascadeBaseEnv(parsedIncludeLocation string, includedProject *Project) (err error) {
	for baseEnvName, baseEnv := range includedProject.baseEnv {
		existingBaseEnv, baseEnvAlreadyDeclared := p.baseEnv[baseEnvName]
		if baseEnvAlreadyDeclared {
			return fmt.Errorf("Cannot declare project env `%s` on `%s` because it was already declared on `%s`", baseEnvName, parsedIncludeLocation, existingBaseEnv.fileLocation)
		}
		p.baseEnv[baseEnvName] = baseEnv
	}
	return nil
}

func (p *Project) cascadeBaseConfig(parsedIncludeLocation string, includedProject *Project) (err error) {
	for baseConfigName, baseConfig := range includedProject.baseConfig {
		existingBaseConfig, baseConfigAlreadyDeclared := p.baseConfig[baseConfigName]
		if baseConfigAlreadyDeclared {
			return fmt.Errorf("Cannot declare project config `%s` on `%s` because it was already declared on `%s`", baseConfigName, parsedIncludeLocation, existingBaseConfig.fileLocation)
		}
		p.baseConfig[baseConfigName] = baseConfig
	}
	return nil
}

func (p *Project) cascadeBaseLConfig(parsedIncludeLocation string, includedProject *Project) (err error) {
	for baseLConfigName, baseLConfig := range includedProject.baseLConfig {
		existingBaseLConfig, baseLConfigAlreadyDeclared := p.baseLConfig[baseLConfigName]
		if baseLConfigAlreadyDeclared {
			return fmt.Errorf("Cannot declare project lconfig `%s` on `%s` because it was already declared on `%s`", baseLConfigName, parsedIncludeLocation, existingBaseLConfig.fileLocation)
		}
		p.baseLConfig[baseLConfigName] = baseLConfig
	}
	return nil
}

func (p *Project) setSortedTaskNames() {
	p.sortedTaskNames = []string{}
	p.maxPublishedTaskNameLength = 0
	for taskName, task := range p.Tasks {
		if !task.Private && len(taskName) > p.maxPublishedTaskNameLength {
			p.maxPublishedTaskNameLength = len(taskName)
		}
		p.sortedTaskNames = append(p.sortedTaskNames, taskName)
	}
	if p.maxPublishedTaskNameLength > 15 {
		p.maxPublishedTaskNameLength = 15
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
