package config

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// ProjectConfig configuration
type ProjectConfig struct {
	dirName                   string // directory name (assigned automatically)
	name                      string
	env                       map[string]string
	components                map[string]*Component
	links                     map[string][]string
	sortedLinkSources         []string
	isSortedLinkSourcesCached bool
	lastGeneratedSymbolIndex  int
	lastGeneratedColorIndex   int
}

// GetName get name of project
func (p *ProjectConfig) GetName() (projectName string) {
	return p.name
}

// GetDirName get name of project
func (p *ProjectConfig) GetDirName() (projectDirName string) {
	return p.dirName
}

// GetComponents get components of project
func (p *ProjectConfig) GetComponents() (components map[string]*Component) {
	return p.components
}

// GetComponentByName get component of project by name
func (p *ProjectConfig) GetComponentByName(name string) (component *Component, err error) {
	component, exists := p.components[name]
	if !exists {
		errorMessage := fmt.Sprintf("Cannot find component `%s`", name)
		err = errors.New(errorMessage)
	}
	return component, err
}

// GetLinks get links in the project
func (p *ProjectConfig) GetLinks() (links map[string][]string) {
	return p.links
}

// GetLinkDestinationList get link by source
func (p *ProjectConfig) GetLinkDestinationList(source string) (destinationList []string) {
	return p.links[source]
}

// GetEnv get env of project
func (p *ProjectConfig) GetEnv() (env map[string]string) {
	return p.env
}

// ToYaml get yaml representation of projectConfig
func (p *ProjectConfig) ToYaml() (str string, err error) {
	pYaml := &ProjectConfigYaml{
		Name:       p.GetName(),
		Env:        p.GetEnv(),
		Components: map[string]ComponentYaml{},
		Links:      p.GetLinks(),
	}
	for componentName, component := range p.GetComponents() {
		pYaml.Components[componentName] = ComponentYaml{
			Type:           component.GetType(),
			Origin:         component.GetOrigin(),
			Branch:         component.GetBranch(),
			Location:       component.GetLocation(),
			Image:          component.GetImage(),
			Start:          component.GetStartCommand(),
			Run:            component.GetRunCommand(),
			ContainerName:  component.GetContainerName(),
			Ports:          component.GetPorts(),
			ReadinessCheck: component.GetReadinessCheckCommand(),
			Dependencies:   component.GetDependencies(),
			Env:            component.GetEnv(),
		}
	}

	d, err := yaml.Marshal(*pYaml)
	if err != nil {
		return str, err
	}
	str = string(d)
	return str, err
}

// ToColorizedYaml get yaml representation (colorized)
func (p *ProjectConfig) ToColorizedYaml() (str string, err error) {
	str, err = p.ToYaml()
	if err == nil {
		str = fmt.Sprintf("\n\033[36m%s\033[0m", str)
	}
	return str, err
}

// GetSortedLinkSources get sorted link sources
func (p *ProjectConfig) GetSortedLinkSources() (sortedSources []string) {
	if p.isSortedLinkSourcesCached {
		return p.sortedLinkSources
	}
	sortedSources = []string{}
	for source := range p.links {
		sortedSources = append(sortedSources, source)
	}
	// sort keys
	sort.SliceStable(sortedSources, func(i int, j int) bool {
		firstSource, secondSource := sortedSources[i], sortedSources[j]
		// get destination
		firstDestinations := p.links[firstSource]
		// compare
		for _, destination := range firstDestinations {
			if strings.HasPrefix(destination, secondSource) {
				return true
			}
		}
		return false
	})
	p.sortedLinkSources = sortedSources
	p.isSortedLinkSourcesCached = true
	return sortedSources
}

// GetSubrepoPrefixMap get map of all eligible component's subrepoPrefix (for git subtree)
func (p *ProjectConfig) GetSubrepoPrefixMap(projectDir string) (subRepoPrefixMap map[string]string) {
	subRepoPrefixMap = map[string]string{}
	for componentName, component := range p.components {
		location := component.location
		branch := component.branch
		if location == "" || branch == "" {
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

func (p *ProjectConfig) fromProjectConfigYaml(pYaml *ProjectConfigYaml, directory string) *ProjectConfig {
	// load pYaml into p
	p.dirName = directory
	p.env = pYaml.Env
	p.name = pYaml.Name
	p.components = make(map[string]*Component)
	p.links = pYaml.Links
	for componentName, component := range pYaml.Components {
		p.components[componentName] = &Component{
			componentType:  component.Type,
			origin:         component.Origin,
			branch:         component.Branch,
			location:       component.Location,
			image:          component.Image,
			start:          component.Start,
			run:            component.Run,
			containerName:  component.ContainerName,
			ports:          component.Ports,
			symbol:         component.Symbol,
			readinessCheck: component.ReadinessCheck,
			dependencies:   component.Dependencies,
			env:            component.Env,
		}
	}
	return p
}
