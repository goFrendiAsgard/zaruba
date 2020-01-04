package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// Environments describe environment variables in general and for each services
type Environments struct {
	General  map[string]string            `yaml:"general"`
	Services map[string]map[string]string `yaml:"services"`
}

// Component describe component specs
type Component struct {
	Type          string `yaml:"type"`          // Component type: "service", "library", "other", or "container"
	Origin        string `yaml:"string"`        // Component's git origin
	Branch        string `yaml:"branch"`        // Component's branch
	Location      string `yaml:"location"`      // location of the component
	Start         string `yaml:"start"`         // command to start service component
	Run           string `yaml:"run"`           // command to run container component
	ContainerName string `yaml:"containerName"` // container name of container component
}

// ProjectConfig configuration
type ProjectConfig struct {
	Environments Environments         `yaml:"environments"`
	Components   map[string]Component `yaml:"components"`
	Executions   []string             `yaml:"executions"`
	Links        map[string][]string  `yaml:"links"`
}

// ToYaml get yaml representation of projectConfig
func (p *ProjectConfig) ToYaml() (str string, err error) {
	d, err := yaml.Marshal(p)
	if err == nil {
		return
	}
	str = string(d)
	return
}

func (p *ProjectConfig) adjustLocation(absDirPath string) {
	// adjust component's location
	for componentName, component := range p.Components {
		component.Location = getAbsLocation(absDirPath, component.Location)
		p.Components[componentName] = component
	}
	// adjust component's link
	newLinks := make(map[string][]string)
	for source, destinations := range p.Links {
		newSource := getAbsLocation(absDirPath, source)
		newLinks[newSource] = []string{}
		for _, destination := range destinations {
			newLinks[newSource] = append(newLinks[newSource], getAbsLocation(absDirPath, destination))
		}
	}
	p.Links = newLinks
}

// GetSortedLinkSources get sorted link sources
func (p *ProjectConfig) GetSortedLinkSources() (sortedSources []string) {
	sortedSources = []string{}
	for source := range p.Links {
		sortedSources = append(sortedSources, source)
	}
	// sort keys
	sort.SliceStable(sortedSources, func(i int, j int) bool {
		firstSource, secondSource := sortedSources[i], sortedSources[j]
		// get destination
		firstDestinations := p.Links[firstSource]
		// compare
		for _, destination := range firstDestinations {
			if strings.HasPrefix(destination, secondSource) {
				return true
			}
		}
		return false
	})
	return
}

// NewProjectConfig create new ProjectConfig
func NewProjectConfig() (p *ProjectConfig) {
	p = &ProjectConfig{
		Environments: Environments{
			General:  make(map[string]string),
			Services: make(map[string]map[string]string),
		},
		Components: make(map[string]Component),
		Executions: []string{},
		Links:      make(map[string][]string),
	}
	return
}

// LoadProjectConfig load project configuration from project directory
func LoadProjectConfig(projectDir string) (p *ProjectConfig) {
	p = NewProjectConfig()
	projectDir, err := filepath.Abs(projectDir)
	if err != nil {
		log.Printf("[ERROR] Invalid project directory `%s`: %s", projectDir, err)
		return
	}
	// read file's content
	b, err := ioutil.ReadFile(filepath.Join(projectDir, "zaruba.config.yaml"))
	if err != nil {
		log.Printf("[ERROR] Cannot read `zaruba.config.yaml`: %s", err)
		return
	}
	str := string(b)
	// create new ProjectConfig and unmarshal
	err = yaml.Unmarshal([]byte(str), p)
	if err != nil {
		log.Printf("[ERROR] Cannot load `zaruba.config.yaml`: %s", err)
		return
	}
	p.adjustLocation(projectDir)
	return
}

func getAbsLocation(absDirPath, location string) string {
	if filepath.IsAbs(location) {
		return location
	}
	return filepath.Join(absDirPath, location)
}
