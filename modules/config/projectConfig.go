package config

import (
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// ProjectConfig configuration
type ProjectConfig struct {
	ProjectName  string               `yaml:"name"`
	Environments Environments         `yaml:"environments"`
	Components   map[string]Component `yaml:"components"`
	Executions   []string             `yaml:"executions"`
	Links        map[string][]string  `yaml:"links"`
}

// Environments describe environment variables in general and for each services
type Environments struct {
	General  map[string]string            `yaml:"general"`
	Services map[string]map[string]string `yaml:"services"`
}

// Component describe component specs
type Component struct {
	Type          string `yaml:"type"`          // Component type: "service", "library", "other", or "container"
	Origin        string `yaml:"origin"`        // Component's git origin
	Branch        string `yaml:"branch"`        // Component's branch
	Location      string `yaml:"location"`      // location of the component
	Start         string `yaml:"start"`         // command to start service component
	Run           string `yaml:"run"`           // command to run container component
	ContainerName string `yaml:"containerName"` // container name of container component
}

// ToYaml get yaml representation of projectConfig
func (p *ProjectConfig) ToYaml() (str string, err error) {
	d, err := yaml.Marshal(*p)
	if err != nil {
		return str, err
	}
	str = string(d)
	return str, err
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
	return sortedSources
}

// GetSubrepoPrefixMap get map of all eligible component's subrepoPrefix (for git subtree)
func (p *ProjectConfig) GetSubrepoPrefixMap(projectDir string) (subRepoPrefixMap map[string]string) {
	subRepoPrefixMap = map[string]string{}
	for componentName, component := range p.Components {
		location := component.Location
		origin := component.Origin
		branch := component.Branch
		if location == "" || origin == "" || branch == "" {
			continue
		}
		subRepoPrefix := getSubrepoPrefix(projectDir, location)
		subRepoPrefixMap[componentName] = subRepoPrefix
	}
	return subRepoPrefixMap
}

func getSubrepoPrefix(projectDir, location string) string {
	if !strings.HasPrefix(location, projectDir) {
		return location
	}
	return strings.Trim(strings.TrimPrefix(location, projectDir), string(os.PathSeparator))
}
