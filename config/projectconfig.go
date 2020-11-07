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

// ProjectConfig is zaruba configuration
type ProjectConfig struct {
	Includes          []string         `yaml:"includes"`
	Tasks             map[string]*Task `yaml:"tasks"`
	Name              string           `yaml:"name"`
	FileLocation      string
	Kwargs            map[string]string
	Generator         *iconer.Generator
	SortedTaskNames   []string
	MaxTaskNameLength int
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
	conf.Kwargs = map[string]string{}
	conf.Generator = iconer.NewGenerator()
	if err != nil {
		return conf, err
	}
	conf.generateProperties()
	conf.fillTaskAndEnv()
	return conf, err
}

func parseStr(rawStr string) (parsedStr string) {
	parsedStr = rawStr
	for _, environ := range os.Environ() {
		pair := strings.SplitN(environ, "=", 2)
		key := pair[0]
		val := os.Getenv(pair[0])
		parsedStr = strings.ReplaceAll(parsedStr, fmt.Sprintf("${%s}", key), val)
		parsedStr = strings.ReplaceAll(parsedStr, fmt.Sprintf("$%s", key), val)
	}
	return parsedStr
}

func loadConfigRecursively(configFile string) (conf *ProjectConfig, err error) {
	d := logger.NewDecoration()
	configFile = parseStr(configFile)
	logger.PrintfStarted("%sLoading %s%s\n", d.Dim, configFile, d.Normal)
	conf = &ProjectConfig{
		Includes: []string{},
		Tasks:    map[string]*Task{},
	}
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return conf, err
	}
	if err = yaml.Unmarshal(b, conf); err != nil {
		return conf, err
	}
	absConfigFile := configFile
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

// GetName get projectName
func (conf *ProjectConfig) GetName() (name string) {
	if conf.Name != "" {
		return conf.Name
	}
	return filepath.Base(filepath.Dir(conf.FileLocation))
}

// GetPublishedTask get all published task
func (conf *ProjectConfig) GetPublishedTask() (publishedTasks map[string]*Task) {
	publishedTasks = map[string]*Task{}
	for taskName, task := range conf.Tasks {
		if !task.Private {
			publishedTasks[taskName] = task
		}
	}
	return publishedTasks
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
		if err = task.init(); err != nil {
			return err
		}
	}
	return err
}

func (conf *ProjectConfig) loadInclusion(parentDir string) (err error) {
	for _, include := range conf.Includes {
		includeLocation := parseStr(include)
		if !filepath.IsAbs(includeLocation) {
			includeLocation = filepath.Join(parentDir, includeLocation)
		}
		includeConf, err := loadConfigRecursively(includeLocation)
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
	conf.MaxTaskNameLength = 0
	for taskName, task := range conf.Tasks {
		if !task.Private && len(taskName) > conf.MaxTaskNameLength {
			conf.MaxTaskNameLength = len(taskName)
		}
		conf.SortedTaskNames = append(conf.SortedTaskNames, taskName)
	}
	if conf.MaxTaskNameLength > 20 {
		conf.MaxTaskNameLength = 20
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
