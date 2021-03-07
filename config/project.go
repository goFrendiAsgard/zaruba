package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	FileLocation               string
	BasePath                   string
	Values                     map[string]string
	SortedTaskNames            []string
	SortedInputNames           []string
	MaxPublishedTaskNameLength int
	IconGenerator              *iconer.Generator
	Decoration                 *logger.Decoration
	CSVLogWriter               *logger.CSVLogWriter
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
	project.CSVLogWriter = logger.NewCSVLogWriter(logFile)
	project.IconGenerator = iconer.NewGenerator()
	project.Decoration = logger.NewDecoration()
	project.setSortedTaskNames()
	project.setSortedInputNames()
	project.linkToTasks()
	project.setDefaultValues()
	err = project.checkInputs()
	return project, err
}

func loadProject(configFile string) (project *Project, err error) {
	d := logger.NewDecoration()
	parsedConfigFile := os.ExpandEnv(configFile)
	logger.PrintfStarted("%sLoading %s%s\n", d.Faint, parsedConfigFile, d.Normal)
	project = &Project{
		Includes: []string{},
		Tasks:    map[string]*Task{},
		Inputs:   map[string]*Input{},
		Values:   map[string]string{},
	}
	b, err := ioutil.ReadFile(parsedConfigFile)
	if err != nil {
		return project, err
	}
	if err = yaml.Unmarshal(b, project); err != nil {
		return project, err
	}
	project.reverseInclusion() // we need to reverse inclusion, so that the first include file will always be overridden by the later
	project.FileLocation = parsedConfigFile
	if !filepath.IsAbs(project.FileLocation) {
		project.FileLocation, err = filepath.Abs(project.FileLocation)
		if err != nil {
			return project, err
		}
	}
	project.BasePath = filepath.Dir(project.FileLocation)
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
	return filepath.Base(filepath.Dir(project.FileLocation))
}

// AddGlobalEnv add global environment for a projectConfig
func (project *Project) AddGlobalEnv(pairOrFile string) {
	pairParts := strings.SplitN(pairOrFile, "=", 2)
	if len(pairParts) == 2 {
		key := pairParts[0]
		val := pairParts[1]
		os.Setenv(key, val)
		return
	}
	godotenv.Load(pairOrFile)
}

// AddValues add value for a project
func (project *Project) AddValues(pairOrFile string) (err error) {
	pairParts := strings.SplitN(pairOrFile, "=", 2)
	if len(pairParts) == 2 {
		key := pairParts[0]
		val := pairParts[1]
		project.Values[key] = val
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
		project.Values[key] = val
	}
	return nil
}

// SetValue set value for a project
func (project *Project) SetValue(key, value string) {
	project.Values[key] = value
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
		for inputName, input := range subInputs {
			inputs[inputName] = input
		}
		inputOrder = append(inputOrder, subInputOrder...)
		// include task's inputs
		for _, inputName := range task.Inputs {
			input, inputExist := project.Inputs[inputName]
			if !inputExist {
				return inputs, inputOrder, fmt.Errorf("Input %s is not defined", inputName)
			}
			inputs[inputName] = input
			inputOrder = append(inputOrder, inputName)
		}
	}
	return inputs, inputOrder, err
}

// Init all tasks
func (project *Project) Init() (err error) {
	for _, taskName := range project.SortedTaskNames {
		task := project.Tasks[taskName]
		err := task.init()
		if err != nil {
			return err
		}
	}
	return err
}

func (project *Project) setTaskFileLocation() {
	for _, task := range project.Tasks {
		task.FileLocation = project.FileLocation
		task.BasePath = project.BasePath
	}
}

func (project *Project) setInputFileLocation() {
	for _, input := range project.Inputs {
		input.FileLocation = project.FileLocation
	}
}

func (project *Project) setDefaultValues() {
	for inputName, input := range project.Inputs {
		project.SetValue(inputName, input.DefaultValue)
	}
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
			parsedIncludeLocation = filepath.Join(project.BasePath, parsedIncludeLocation)
		}
		includedProject, err := loadProject(parsedIncludeLocation)
		if err != nil {
			return err
		}
		// cascade inputs
		for inputName, input := range includedProject.Inputs {
			_, inputAlreadyDeclared := project.Inputs[inputName]
			if inputAlreadyDeclared {
				return fmt.Errorf("Cannot declare input `%s` on `%s` because the input was already declared on `%s`", inputName, parsedIncludeLocation, input.FileLocation)
			}
			project.Inputs[inputName] = input
		}
		// cascade tasks
		for taskName, task := range includedProject.Tasks {
			existingTask, taskAlreadyDeclared := project.Tasks[taskName]
			if taskAlreadyDeclared {
				return fmt.Errorf("Cannot declare task `%s` on `%s` because the task was already declared on `%s`", taskName, parsedIncludeLocation, existingTask.FileLocation)
			}
			project.Tasks[taskName] = task
		}
	}
	return err
}

func (project *Project) setSortedTaskNames() {
	project.SortedTaskNames = []string{}
	project.MaxPublishedTaskNameLength = 0
	for taskName, task := range project.Tasks {
		if !task.Private && len(taskName) > project.MaxPublishedTaskNameLength {
			project.MaxPublishedTaskNameLength = len(taskName)
		}
		project.SortedTaskNames = append(project.SortedTaskNames, taskName)
	}
	if project.MaxPublishedTaskNameLength > 15 {
		project.MaxPublishedTaskNameLength = 15
	}
	sort.Strings(project.SortedTaskNames)
}

func (project *Project) setSortedInputNames() {
	project.SortedInputNames = []string{}
	for inputName := range project.Inputs {
		project.SortedInputNames = append(project.SortedInputNames, inputName)
	}
	sort.Strings(project.SortedInputNames)
}

func (project *Project) linkToTasks() {
	for taskName, task := range project.Tasks {
		task.Project = project
		task.Name = taskName
		task.linkToEnvs()
	}
}
