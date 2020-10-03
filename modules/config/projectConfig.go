package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

// ProjectConfig configuration
type ProjectConfig struct {
	dirName                      string // directory name (assigned automatically)
	ignores                      []string
	name                         string
	defaultBranch                string
	env                          map[string]string
	components                   map[string]*Component
	links                        map[string][]string
	sortedLinkSources            []string
	isSortedLinkSourcesCached    bool
	lastGeneratedSymbolIndexLock *sync.RWMutex
	lastGeneratedSymbolIndex     int
	lastGeneratedColorIndexLock  *sync.RWMutex
	lastGeneratedColorIndex      int
}

// GetIgnores get ignored of project
func (p *ProjectConfig) GetIgnores() (ignores []string) {
	return p.ignores
}

// GetName get name of project
func (p *ProjectConfig) GetName() (projectName string) {
	return p.name
}

// GetDirName get name of project
func (p *ProjectConfig) GetDirName() (projectDirName string) {
	return p.dirName
}

// GetDefaultBranch get branch of porject
func (p *ProjectConfig) GetDefaultBranch() (projectDefaultBranch string) {
	return p.defaultBranch
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

// GetComponentsByLabels get component by labels
func (p *ProjectConfig) GetComponentsByLabels(labelStrings []string) (components map[string]*Component) {
	components = map[string]*Component{}
	for _, labelString := range labelStrings {
		pair := strings.SplitN(labelString, ":", 2)
		queryKey, queryVal := pair[0], pair[1]
		for componentName, component := range p.GetComponents() {
			labels := component.GetLabels()
			for key, val := range labels {
				if queryKey == key && isLabelValueMatch(queryVal, val) {
					components[componentName] = component
				}
			}
		}
	}
	return components
}

func isLabelValueMatch(queryVal, val string) (match bool) {
	// val only contains single value and match
	if queryVal == val {
		return true
	}
	// val contains multiple value (space sparated) and the first value is match
	if strings.HasPrefix(val, fmt.Sprintf("%s ", queryVal)) {
		return true
	}
	// val contains multiple value (space sparated) and the last value is match
	if strings.HasSuffix(val, fmt.Sprintf(" %s", queryVal)) {
		return true
	}
	// val contains multiple value (space sparated) and any 1..N-1 value is match
	if strings.Contains(val, fmt.Sprintf(" %s ", queryVal)) {
		return true
	}
	return false
}

// GetComponentsByNamesOrLabels get component by names or labels
func (p *ProjectConfig) GetComponentsByNamesOrLabels(namesOrLabels []string) (components map[string]*Component, err error) {
	components = map[string]*Component{}
	for _, nameOrLabel := range namesOrLabels {
		// by label
		if strings.Contains(nameOrLabel, ":") {
			byLabelComponents := p.GetComponentsByLabels([]string{nameOrLabel})
			for name, component := range byLabelComponents {
				components[name] = component
			}
			continue
		}
		// by name
		byNameComponent, err := p.GetComponentByName(nameOrLabel)
		if err != nil {
			return components, err
		}
		components[nameOrLabel] = byNameComponent
	}
	return components, err
}

// GetComponentsBySelectors get component by selector
func (p *ProjectConfig) GetComponentsBySelectors(selectors []string) (components map[string]*Component, err error) {
	if len(selectors) == 0 {
		selectors = []string{"scenario:default"}
	}
	components, err = p.GetComponentsByNamesOrLabels(selectors)
	exists := false
	for range components {
		exists = true
		break
	}
	if !exists {
		components = p.GetComponents()
	}
	return components, err
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

func (p *ProjectConfig) getRelativeLocation(basePath, location string) (relativeLocation string) {
	relativeLocation, err := filepath.Rel(basePath, location)
	if err != nil {
		relativeLocation = location
	}
	return relativeLocation
}

func (p *ProjectConfig) getRelativeLinks(basePath string) (relativeLinks map[string][]string) {
	relativeLinks = map[string][]string{}
	for src, dstList := range p.GetLinks() {
		relativeDstList := []string{}
		for _, dst := range dstList {
			relativeDst := p.getRelativeLocation(basePath, dst)
			relativeDstList = append(relativeDstList, relativeDst)
		}
		relativeLinks[p.getRelativeLocation(basePath, src)] = relativeDstList
	}
	return relativeLinks
}

// ToYaml get yaml representation of projectConfig
func (p *ProjectConfig) ToYaml() (str string, err error) {
	basePath := p.GetDirName()
	pYaml := &ProjectConfigYaml{
		Ignores:       p.GetIgnores(),
		Name:          p.GetName(),
		Env:           p.GetEnv(),
		DefaultBranch: p.GetDefaultBranch(),
		Components:    map[string]ComponentYaml{},
		Links:         p.getRelativeLinks(basePath),
	}
	for componentName, component := range p.GetComponents() {
		componentLocation := component.GetLocation()
		pYaml.Components[componentName] = ComponentYaml{
			Type:           component.GetType(),
			Labels:         component.GetLabels(),
			Origin:         component.GetOrigin(),
			Location:       p.getRelativeLocation(basePath, componentLocation),
			Image:          component.GetImage(),
			Start:          component.GetStartCommand(),
			ContainerName:  component.GetContainerName(),
			Ports:          component.GetPorts(),
			Volumes:        component.GetVolumes(),
			ReadinessCheck: component.GetReadinessCheckCommand(),
			ReadinessURL:   component.GetReadinessURL(),
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
		if location == "" {
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
	branch := pYaml.DefaultBranch
	if branch == "" {
		branch = "main"
	}
	// load pYaml into p
	p.ignores = pYaml.Ignores
	p.dirName = directory
	p.defaultBranch = branch
	p.env = pYaml.Env
	p.name = pYaml.Name
	p.components = make(map[string]*Component)
	p.links = pYaml.Links
	for componentName, cYaml := range pYaml.Components {
		p.components[componentName] = &Component{
			labels:              cYaml.Labels,
			componentType:       cYaml.Type,
			origin:              cYaml.Origin,
			location:            cYaml.Location,
			image:               cYaml.Image,
			start:               cYaml.Start,
			containerName:       cYaml.ContainerName,
			ports:               cYaml.Ports,
			volumes:             cYaml.Volumes,
			symbol:              cYaml.Symbol,
			readinessCheck:      cYaml.ReadinessCheck,
			readinessURL:        cYaml.ReadinessURL,
			dependencies:        cYaml.Dependencies,
			venvLock:            &sync.RWMutex{},
			env:                 cYaml.Env,
			runtimeSymbolLock:   &sync.RWMutex{},
			runtimeLocationLock: &sync.RWMutex{},
			runtimeNameLock:     &sync.RWMutex{},
			colorLock:           &sync.RWMutex{},
		}
	}
	return p
}
