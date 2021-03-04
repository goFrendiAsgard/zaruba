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

// ProjectConfig is zaruba configuration
type ProjectConfig struct {
	Includes                   []string         `yaml:"includes,omitempty"`
	Tasks                      map[string]*Task `yaml:"tasks,omitempty"`
	Name                       string           `yaml:"name,omitempty"`
	FileLocation               string
	Kwargs                     map[string]string
	IconGenerator              *iconer.Generator
	SortedTaskNames            []string
	MaxPublishedTaskNameLength int
	Decoration                 *logger.Decoration
	CSVLogWriter               *logger.CSVLogWriter
}

// NewConfig create new Config from Yaml File
func NewConfig(configFile string) (conf *ProjectConfig, err error) {
	if os.Getenv("ZARUBA_HOME") == "" {
		executable, err := os.Executable()
		if err != nil {
			return conf, err
		}
		os.Setenv("ZARUBA_HOME", filepath.Dir(executable))
	}
	conf, err = loadConfigRecursively(configFile)
	if err != nil {
		return conf, err
	}
	dir := os.ExpandEnv(filepath.Dir(configFile))
	logFile := filepath.Join(dir, "log.zaruba.csv")
	conf.CSVLogWriter = logger.NewCSVLogWriter(logFile)
	conf.Kwargs = map[string]string{}
	conf.IconGenerator = iconer.NewGenerator()
	conf.Decoration = logger.NewDecoration()
	if err != nil {
		return conf, err
	}
	conf.generateProperties()
	conf.fillTaskAndEnv()
	return conf, err
}

func loadConfigRecursively(configFile string) (conf *ProjectConfig, err error) {
	d := logger.NewDecoration()
	parsedConfigFile := os.ExpandEnv(configFile)
	logger.PrintfStarted("%sLoading %s%s\n", d.Faint, parsedConfigFile, d.Normal)
	conf = &ProjectConfig{
		Includes: []string{},
		Tasks:    map[string]*Task{},
	}
	b, err := ioutil.ReadFile(parsedConfigFile)
	if err != nil {
		return conf, err
	}
	if err = yaml.Unmarshal(b, conf); err != nil {
		return conf, err
	}
	conf.reverseInclusion() // we need to reverse inclusion, so that the first include file will always be overridden by the later
	absConfigFile := parsedConfigFile
	if !filepath.IsAbs(absConfigFile) {
		absConfigFile, err = filepath.Abs(absConfigFile)
		if err != nil {
			return conf, err
		}
	}
	conf.FileLocation = absConfigFile
	conf.fillTaskFileLocationAndDirPath(absConfigFile)
	inclusionParentDir := filepath.Dir(absConfigFile)
	if err = conf.loadInclusion(inclusionParentDir); err != nil {
		return conf, err
	}
	return conf, err
}

func (conf *ProjectConfig) reverseInclusion() {
	i := 0
	j := len(conf.Includes) - 1
	for i < j {
		conf.Includes[i], conf.Includes[j] = conf.Includes[j], conf.Includes[i]
		i++
		j--
	}
}

// GetName get projectName
func (conf *ProjectConfig) GetName() (name string) {
	if conf.Name != "" {
		return conf.Name
	}
	return filepath.Base(filepath.Dir(conf.FileLocation))
}

// AddGlobalEnv add global environment for a projectConfig
func (conf *ProjectConfig) AddGlobalEnv(pairOrFile string) {
	pairParts := strings.SplitN(pairOrFile, "=", 2)
	if len(pairParts) == 2 {
		key := pairParts[0]
		val := pairParts[1]
		os.Setenv(key, val)
		return
	}
	godotenv.Load(pairOrFile)
}

// AddKwargs add global environment for a projectConfig
func (conf *ProjectConfig) AddKwargs(pairOrFile string) (err error) {
	pairParts := strings.SplitN(pairOrFile, "=", 2)
	if len(pairParts) == 2 {
		key := pairParts[0]
		val := pairParts[1]
		conf.Kwargs[key] = val
		return nil
	}
	b, err := ioutil.ReadFile(pairOrFile)
	if err != nil {
		return err
	}
	kwargs := map[string]string{}
	if err = yaml.Unmarshal(b, kwargs); err != nil {
		return err
	}
	for key, val := range kwargs {
		conf.Kwargs[key] = val
	}
	return nil
}

// Init parse all tasks
func (conf *ProjectConfig) Init() (err error) {
	for _, taskName := range conf.SortedTaskNames {
		task := conf.Tasks[taskName]
		err := task.init()
		if err != nil {
			return err
		}
	}
	return err
}

func (conf *ProjectConfig) loadInclusion(parentDir string) (err error) {
	for _, includeLocation := range conf.Includes {
		parsedIncludeLocation := os.ExpandEnv(includeLocation)
		if !filepath.IsAbs(parsedIncludeLocation) {
			parsedIncludeLocation = filepath.Join(parentDir, parsedIncludeLocation)
		}
		includeConf, err := loadConfigRecursively(parsedIncludeLocation)
		if err != nil {
			return err
		}
		for taskName, task := range includeConf.Tasks {
			if _, exists := conf.Tasks[taskName]; !exists {
				conf.Tasks[taskName] = task
			}
		}
	}
	return err
}

func (conf *ProjectConfig) fillTaskFileLocationAndDirPath(absConfigFile string) {
	absConfigDir := filepath.Dir(absConfigFile)
	for _, task := range conf.Tasks {
		task.FileLocation = absConfigFile
		task.BasePath = absConfigDir
	}
}

func (conf *ProjectConfig) generateProperties() {
	conf.SortedTaskNames = []string{}
	conf.MaxPublishedTaskNameLength = 0
	for taskName, task := range conf.Tasks {
		if !task.Private && len(taskName) > conf.MaxPublishedTaskNameLength {
			conf.MaxPublishedTaskNameLength = len(taskName)
		}
		conf.SortedTaskNames = append(conf.SortedTaskNames, taskName)
	}
	if conf.MaxPublishedTaskNameLength > 15 {
		conf.MaxPublishedTaskNameLength = 15
	}
	sort.Strings(conf.SortedTaskNames)
}

func (conf *ProjectConfig) fillTaskAndEnv() {
	for taskName, task := range conf.Tasks {
		task.Project = conf
		task.Name = taskName
		task.fillEnvTask()
	}
}
