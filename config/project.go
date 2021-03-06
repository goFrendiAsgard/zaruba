package config

import (
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
	Includes                   []string         `yaml:"includes,omitempty"`
	Tasks                      map[string]*Task `yaml:"tasks,omitempty"`
	Name                       string           `yaml:"name,omitempty"`
	FileLocation               string
	Values                     map[string]string
	SortedTaskNames            []string
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
	project.Values = map[string]string{}
	project.IconGenerator = iconer.NewGenerator()
	project.Decoration = logger.NewDecoration()
	if err != nil {
		return project, err
	}
	project.setSortedTaskNames()
	project.linkTasks()
	return project, err
}

func loadProject(configFile string) (project *Project, err error) {
	d := logger.NewDecoration()
	parsedConfigFile := os.ExpandEnv(configFile)
	logger.PrintfStarted("%sLoading %s%s\n", d.Faint, parsedConfigFile, d.Normal)
	project = &Project{
		Includes: []string{},
		Tasks:    map[string]*Task{},
	}
	b, err := ioutil.ReadFile(parsedConfigFile)
	if err != nil {
		return project, err
	}
	if err = yaml.Unmarshal(b, project); err != nil {
		return project, err
	}
	project.reverseInclusion() // we need to reverse inclusion, so that the first include file will always be overridden by the later
	absConfigFile := parsedConfigFile
	if !filepath.IsAbs(absConfigFile) {
		absConfigFile, err = filepath.Abs(absConfigFile)
		if err != nil {
			return project, err
		}
	}
	project.FileLocation = absConfigFile
	project.fillTaskFileLocationAndDirPath(absConfigFile)
	inclusionParentDir := filepath.Dir(absConfigFile)
	if err = project.loadInclusion(inclusionParentDir); err != nil {
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

// AddValues add global environment for a projectConfig
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

// Init parse all tasks
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

func (project *Project) loadInclusion(parentDir string) (err error) {
	for _, includeLocation := range project.Includes {
		parsedIncludeLocation := os.ExpandEnv(includeLocation)
		if !filepath.IsAbs(parsedIncludeLocation) {
			parsedIncludeLocation = filepath.Join(parentDir, parsedIncludeLocation)
		}
		includeConf, err := loadProject(parsedIncludeLocation)
		if err != nil {
			return err
		}
		for taskName, task := range includeConf.Tasks {
			if _, exists := project.Tasks[taskName]; !exists {
				project.Tasks[taskName] = task
			}
		}
	}
	return err
}

func (project *Project) fillTaskFileLocationAndDirPath(absConfigFile string) {
	absConfigDir := filepath.Dir(absConfigFile)
	for _, task := range project.Tasks {
		task.FileLocation = absConfigFile
		task.BasePath = absConfigDir
	}
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

func (project *Project) linkTasks() {
	for taskName, task := range project.Tasks {
		task.Project = project
		task.Name = taskName
		task.linkEnvs()
	}
}
