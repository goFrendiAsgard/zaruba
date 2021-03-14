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
	"github.com/state-alchemists/zaruba/iconer"
	"github.com/state-alchemists/zaruba/logger"
	yaml "gopkg.in/yaml.v2"
)

// Project is zaruba configuration
type Project struct {
	Includes                   []string          `yaml:"includes,omitempty"`
	Tasks                      map[string]*Task  `yaml:"tasks,omitempty"`
	Name                       string            `yaml:"name,omitempty"`
	Inputs                     map[string]*Input `yaml:"inputs,omitempty"`
	fileLocation               string
	basePath                   string
	values                     map[string]string
	sortedTaskNames            []string
	sortedInputNames           []string
	maxPublishedTaskNameLength int
	iconGenerator              *iconer.Generator
	decoration                 *logger.Decoration
	csvLogWriter               *logger.CSVLogWriter
	IsInitialized              bool
}

// NewProject create new Config from Yaml File
func NewProject(configFile string) (project *Project, err error) {
	if os.Getenv("ZARUBA_HOME") == "" {
		executable, err := os.Executable()
		if err != nil {
			return project, err
		}
		os.Setenv("ZARUBA_HOME", filepath.Dir(executable))
	}
	project, err = loadProject(configFile)
	if err != nil {
		return project, err
	}
	dir := os.ExpandEnv(filepath.Dir(configFile))
	logFile := filepath.Join(dir, "log.zaruba.csv")
	project.csvLogWriter = logger.NewCSVLogWriter(logFile)
	project.iconGenerator = iconer.NewGenerator()
	project.decoration = logger.NewDecoration()
	project.setSortedTaskNames()
	project.setSortedInputNames()
	project.linkToTasks()
	project.linkToInputs()
	if err = project.setDefaultValues(); err != nil {
		return project, err
	}
	err = project.checkInputs()
	return project, err
}

func loadProject(configFile string) (project *Project, err error) {
	d := logger.NewDecoration()
	parsedConfigFile := os.ExpandEnv(configFile)
	logger.PrintfStarted("%sLoading %s%s\n", d.Faint, parsedConfigFile, d.Normal)
	project = &Project{
		Includes:      []string{},
		Tasks:         map[string]*Task{},
		Inputs:        map[string]*Input{},
		values:        map[string]string{},
		IsInitialized: false,
	}
	b, err := ioutil.ReadFile(parsedConfigFile)
	if err != nil {
		return project, err
	}
	if err = yaml.Unmarshal(b, project); err != nil {
		return project, err
	}
	project.reverseInclusion() // we need to reverse inclusion, so that the first include file will always be overridden by the later
	project.fileLocation = parsedConfigFile
	if !filepath.IsAbs(project.fileLocation) {
		project.fileLocation, err = filepath.Abs(project.fileLocation)
		if err != nil {
			return project, err
		}
	}
	project.basePath = filepath.Dir(project.fileLocation)
	project.setTaskFileLocation()
	project.setInputFileLocation()
	// cascade project, add inclusion's property to this project
	if err = project.cascadeIncludes(); err != nil {
		return project, err
	}
	return project, err
}

func (project *Project) reverseInclusion() {
	i := 0
	j := len(project.Includes) - 1
	for i < j {
		project.Includes[i], project.Includes[j] = project.Includes[j], project.Includes[i]
		i++
		j--
	}
}

// GetName get projectName
func (project *Project) GetName() (name string) {
	if project.Name != "" {
		return project.Name
	}
	return filepath.Base(filepath.Dir(project.fileLocation))
}

// GetBasePath get basePath
func (project *Project) GetBasePath() (basePath string) {
	return project.basePath
}

// GetSortedInputNames get sorted input names
func (project *Project) GetSortedInputNames() (sortedInputNames []string) {
	return project.sortedInputNames
}

// GetSortedTaskNames get sorted task names
func (project *Project) GetSortedTaskNames() (sortedTaskNames []string) {
	return project.sortedTaskNames
}

// GetValues get value
func (project *Project) GetValues() (values map[string]string) {
	return project.values
}

// GetValue get value
func (project *Project) GetValue(key string) (value string) {
	return project.values[key]
}

// IsValueExist is value exist
func (project *Project) IsValueExist(key string) (exist bool) {
	_, exist = project.values[key]
	return exist
}

// AddGlobalEnv add global environment for a projectConfig
func (project *Project) AddGlobalEnv(pairOrFile string) (err error) {
	if project.IsInitialized {
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
func (project *Project) AddValue(pairOrFile string) (err error) {
	if project.IsInitialized {
		return fmt.Errorf("Cannot AddValue, project has been initialized")
	}
	pairParts := strings.SplitN(pairOrFile, "=", 2)
	if len(pairParts) == 2 {
		key := pairParts[0]
		val := pairParts[1]
		project.values[key] = val
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
		project.values[key] = val
	}
	return nil
}

// SetValue set value for a project
func (project *Project) SetValue(key, value string) (err error) {
	if project.IsInitialized {
		return fmt.Errorf("Cannot AddValue, project has been initialized")
	}
	project.values[key] = value
	return nil
}

// GetInputs given task names
func (project *Project) GetInputs(taskNames []string) (inputs map[string]*Input, inputOrder []string, err error) {
	inputs = map[string]*Input{}
	inputOrder = []string{}
	for _, taskName := range taskNames {
		task, taskExist := project.Tasks[taskName]
		if !taskExist {
			return inputs, inputOrder, fmt.Errorf("Task %s is not exist", taskName)
		}
		// include task's dependencies and parent's inputs first
		subTaskNames := []string{}
		if task.Extend != "" {
			subTaskNames = append(subTaskNames, task.Extend)
		}
		subTaskNames = append(subTaskNames, task.Dependencies...)
		subInputs, subInputOrder, err := project.GetInputs(subTaskNames)
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
			input, inputDeclared := project.Inputs[inputName]
			if !inputDeclared {
				return inputs, inputOrder, fmt.Errorf("Input %s is not defined", inputName)
			}
			if _, inputRegistered := inputs[inputName]; !inputRegistered {
				inputs[inputName] = input
				inputOrder = append(inputOrder, inputName)
			}
		}
	}
	return inputs, inputOrder, err
}

// Init all tasks
func (project *Project) Init() (err error) {
	for _, taskName := range project.sortedTaskNames {
		task := project.Tasks[taskName]
		err := task.init()
		if err != nil {
			return err
		}
	}
	r, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return err
	}
	for key, value := range project.values {
		parsedValue := os.ExpandEnv(value)
		project.values[key] = parsedValue
		inputEnvKey := "ZARUBA_INPUT_" + string(r.ReplaceAll([]byte(strings.ToUpper(key)), []byte("_")))
		os.Setenv(inputEnvKey, parsedValue)
	}
	project.IsInitialized = true
	return err
}

func (project *Project) setTaskFileLocation() {
	for _, task := range project.Tasks {
		task.fileLocation = project.fileLocation
		task.basePath = project.basePath
	}
}

func (project *Project) setInputFileLocation() {
	for _, input := range project.Inputs {
		input.fileLocation = project.fileLocation
	}
}

func (project *Project) setDefaultValues() (err error) {
	for inputName, input := range project.Inputs {
		if err = project.SetValue(inputName, input.DefaultValue); err != nil {
			return err
		}
	}
	return nil
}

func (project *Project) checkInputs() (err error) {
	for taskName, task := range project.Tasks {
		for _, inputName := range task.Inputs {
			_, inputExist := project.Inputs[inputName]
			if !inputExist {
				return fmt.Errorf("Input %s is required by %s but it was not declared", inputName, taskName)
			}
		}
	}
	return nil
}

func (project *Project) cascadeIncludes() (err error) {
	for _, includeLocation := range project.Includes {
		parsedIncludeLocation := os.ExpandEnv(includeLocation)
		if !filepath.IsAbs(parsedIncludeLocation) {
			parsedIncludeLocation = filepath.Join(project.basePath, parsedIncludeLocation)
		}
		includedProject, err := loadProject(parsedIncludeLocation)
		if err != nil {
			return err
		}
		// cascade inputs
		for inputName, input := range includedProject.Inputs {
			_, inputAlreadyDeclared := project.Inputs[inputName]
			if inputAlreadyDeclared {
				return fmt.Errorf("Cannot declare input `%s` on `%s` because the input was already declared on `%s`", inputName, parsedIncludeLocation, input.fileLocation)
			}
			project.Inputs[inputName] = input
		}
		// cascade tasks
		for taskName, task := range includedProject.Tasks {
			existingTask, taskAlreadyDeclared := project.Tasks[taskName]
			if taskAlreadyDeclared {
				return fmt.Errorf("Cannot declare task `%s` on `%s` because the task was already declared on `%s`", taskName, parsedIncludeLocation, existingTask.GetFileLocation())
			}
			project.Tasks[taskName] = task
		}
	}
	return err
}

func (project *Project) setSortedTaskNames() {
	project.sortedTaskNames = []string{}
	project.maxPublishedTaskNameLength = 0
	for taskName, task := range project.Tasks {
		if !task.Private && len(taskName) > project.maxPublishedTaskNameLength {
			project.maxPublishedTaskNameLength = len(taskName)
		}
		project.sortedTaskNames = append(project.sortedTaskNames, taskName)
	}
	if project.maxPublishedTaskNameLength > 15 {
		project.maxPublishedTaskNameLength = 15
	}
	sort.Strings(project.sortedTaskNames)
}

func (project *Project) setSortedInputNames() {
	project.sortedInputNames = []string{}
	for inputName := range project.Inputs {
		project.sortedInputNames = append(project.sortedInputNames, inputName)
	}
	sort.Strings(project.sortedInputNames)
}

func (project *Project) linkToTasks() {
	for taskName, task := range project.Tasks {
		task.Project = project
		task.name = taskName
		task.linkToEnvs()
	}
}

func (project *Project) linkToInputs() {
	for inputName, input := range project.Inputs {
		input.Project = project
		input.name = inputName
	}
}
