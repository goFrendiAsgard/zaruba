package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/state-alchemists/zaruba/modules/file"
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
	Origin        string `yaml:"origin"`        // Component's git origin
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
	d, err := yaml.Marshal(*p)
	if err != nil {
		return str, err
	}
	str = string(d)
	return str, err
}

func (p *ProjectConfig) adjustLocation(absDirPath string) {
	// adjust component's location
	for componentName, component := range p.Components {
		component.Location = file.GetAbsoluteLocation(absDirPath, component.Location)
		p.Components[componentName] = component
	}
	// adjust component's link
	newLinks := make(map[string][]string)
	for source, destinations := range p.Links {
		newSource := file.GetAbsoluteLocation(absDirPath, source)
		newLinks[newSource] = []string{}
		for _, destination := range destinations {
			newLinks[newSource] = append(newLinks[newSource], file.GetAbsoluteLocation(absDirPath, destination))
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

// NewProjectConfig create new ProjectConfig
func NewProjectConfig() (p *ProjectConfig) {
	return &ProjectConfig{
		Environments: Environments{
			General:  make(map[string]string),
			Services: make(map[string]map[string]string),
		},
		Components: make(map[string]Component),
		Executions: []string{},
		Links:      make(map[string][]string),
	}
}

// LoadProjectConfig load project configuration from project directory
func LoadProjectConfig(projectDir string) (p *ProjectConfig, err error) {
	allDirs, err := file.GetAllFiles(projectDir, file.NewOption().SetIsOnlyDir(true))
	if err != nil {
		return p, err
	}
	p = NewProjectConfig()
	for _, directory := range allDirs {
		subP, loadSubErr := LoadSingleProjectConfig(directory)
		if loadSubErr != nil {
			if os.IsNotExist(loadSubErr) {
				continue
			}
			err = loadSubErr
			break
		}
		// merge general environment
		for generalSubEnvName, generalSubEnv := range subP.Environments.General {
			if _, exists := p.Environments.General[generalSubEnvName]; !exists {
				p.Environments.General[generalSubEnvName] = generalSubEnv
			}
		}
		// merge service environment
		for serviceName, serviceEnvMap := range subP.Environments.Services {
			// if p doesn't have any environment for the service, add it
			if _, exists := p.Environments.Services[serviceName]; !exists {
				p.Environments.Services[serviceName] = serviceEnvMap
				continue
			}
			// p already has environment for the service, cascade it
			for serviceSubEnvName, serviceSubEnv := range serviceEnvMap {
				if _, exists := p.Environments.Services[serviceName][serviceSubEnvName]; !exists {
					p.Environments.Services[serviceName][serviceSubEnvName] = serviceSubEnv
				}
			}
		}
		// merge component
		for componentName, component := range subP.Components {
			if _, exists := p.Components[componentName]; !exists {
				p.Components[componentName] = component
			}
		}
		// merge executions
		for _, subExecution := range subP.Executions {
			exists := false
			for _, execution := range p.Executions {
				if execution == subExecution {
					exists = true
					break
				}
			}
			if !exists {
				p.Executions = append(p.Executions, subExecution)
			}
		}
		// merge links
		for libPath, subLinks := range subP.Links {
			// if p doesn't have any link for libPath, add it
			if _, exists := p.Links[libPath]; !exists {
				p.Links[libPath] = subLinks
				continue
			}
			for _, subLink := range subLinks {
				exists := false
				for _, link := range p.Links[libPath] {
					if subLink == link {
						exists = true
						break
					}
				}
				if !exists {
					p.Links[libPath] = append(p.Links[libPath], subLink)
				}
			}
		}
	}
	return p, err
}

func getSubrepoPrefix(projectDir, location string) string {
	if !strings.HasPrefix(location, projectDir) {
		return location
	}
	return strings.Trim(strings.TrimPrefix(location, projectDir), string(os.PathSeparator))
}

// LoadSingleProjectConfig load project configuration from a directory
func LoadSingleProjectConfig(directory string) (p *ProjectConfig, err error) {
	p = NewProjectConfig()
	directory, err = filepath.Abs(directory)
	if err != nil {
		return p, err
	}
	// read file's content
	b, err := ioutil.ReadFile(filepath.Join(directory, "zaruba.config.yaml"))
	if err != nil {
		return p, err
	}
	str := string(b)
	// create new ProjectConfig and unmarshal
	err = yaml.Unmarshal([]byte(str), p)
	if err != nil {
		return p, err
	}
	p.adjustLocation(directory)
	return p, err
}
