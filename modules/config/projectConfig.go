package config

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// ProjectConfig configuration
type ProjectConfig struct {
	dirName                   string
	name                      string
	environments              *Environments
	components                map[string]*Component
	executions                []string
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

// GetEnvironments get environments of project
func (p *ProjectConfig) GetEnvironments() (environments *Environments) {
	return p.environments
}

// GetComponents get components of project
func (p *ProjectConfig) GetComponents() (components map[string]*Component) {
	return p.components
}

// GetComponentByName get component of project by name
func (p *ProjectConfig) GetComponentByName(name string) (component *Component) {
	return p.components[name]
}

// GetExecutions get executions order of projects
func (p *ProjectConfig) GetExecutions() (executions []string) {
	return p.executions
}

// GetLinks get links in the project
func (p *ProjectConfig) GetLinks() (links map[string][]string) {
	return p.links
}

// GetLinkDestinationList get link by source
func (p *ProjectConfig) GetLinkDestinationList(source string) (destinationList []string) {
	return p.links[source]
}

// ToYaml get yaml representation of projectConfig
func (p *ProjectConfig) ToYaml() (str string, err error) {
	pYaml := &ProjectConfigYaml{
		ProjectName: p.GetName(),
		Environments: EnvironmentsYaml{
			General:  p.GetEnvironments().general,
			Services: p.GetEnvironments().services,
		},
		Components: map[string]ComponentYaml{},
		Executions: p.GetExecutions(),
		Links:      p.GetLinks(),
	}
	for componentName, component := range p.GetComponents() {
		pYaml.Components[componentName] = ComponentYaml{
			Type:          component.GetType(),
			Origin:        component.GetOrigin(),
			Branch:        component.GetBranch(),
			Location:      component.GetLocation(),
			ImageName:     component.GetImageName(),
			Start:         component.GetStartCommand(),
			Run:           component.GetRunCommand(),
			ContainerName: component.GetContainerName(),
			Ports:         component.GetPorts(),
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
	if err != nil {
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
		origin := component.origin
		branch := component.branch
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

func (p *ProjectConfig) fromProjectConfigYaml(pYaml *ProjectConfigYaml, directory string) *ProjectConfig {
	// load pYaml into p
	p.dirName = directory
	p.name = pYaml.ProjectName
	p.environments = &Environments{
		general:  pYaml.Environments.General,
		services: pYaml.Environments.Services,
	}
	p.components = make(map[string]*Component)
	p.executions = pYaml.Executions
	p.links = pYaml.Links
	for componentName, component := range pYaml.Components {
		p.components[componentName] = &Component{
			componentType: component.Type,
			origin:        component.Origin,
			branch:        component.Branch,
			location:      component.Location,
			imageName:     component.ImageName,
			start:         component.Start,
			run:           component.Run,
			containerName: component.ContainerName,
			ports:         component.Ports,
			symbol:        component.Symbol,
			livenessCheck: component.LivenessCheck,
		}
	}
	return p
}

// Environments describe environment variables in general and for each services
type Environments struct {
	project  *ProjectConfig
	general  map[string]string
	services map[string]map[string]string
}

// GetGeneralVariables get general environment variables
func (e *Environments) GetGeneralVariables() (general map[string]string) {
	return e.general
}

// GetAllServicesVariables get all service variables (as map)
func (e *Environments) GetAllServicesVariables() (services map[string]map[string]string) {
	return e.services
}

// GetRuntimeVariables get variable of a service
func (e *Environments) GetRuntimeVariables(serviceName string) (variables map[string]string) {
	variables = map[string]string{}
	// get variables from general
	for generalVarName, generalVal := range e.general {
		actualVal := generalVal
		// prefer global env
		if os.Getenv(generalVarName) != "" {
			actualVal = os.Getenv(generalVarName)
		}
		variables[generalVarName] = actualVal
	}
	// get service variables
	if serviceEnv, exists := e.services[serviceName]; exists {
		for serviceVarName, serviceVal := range serviceEnv {
			actualVal := serviceVal
			// take actual service's variable value from previous variable
			for generalVarName, generalVal := range variables {
				if serviceVal == fmt.Sprintf("${%s}", generalVarName) || serviceVal == fmt.Sprintf("$%s", generalVarName) {
					actualVal = generalVal
					break
				}
			}
			variables[serviceVarName] = actualVal
		}
	}
	// inject container name
	if _, isExists := variables[serviceName]; !isExists {
		if component, isExists := e.project.GetComponents()[serviceName]; isExists {
			variables["CONTAINER_NAME"] = component.GetRuntimeContainerName()
		}
	}
	return variables
}

// GetQuotedRuntimeVariables get variable of a service
func (e *Environments) GetQuotedRuntimeVariables(serviceName string) (variables map[string]string) {
	variables = e.GetRuntimeVariables(serviceName)
	// quote variable value
	for key, val := range variables {
		unquotedVal, err := strconv.Unquote(val)
		if err != nil {
			unquotedVal = val
		}
		quotedVal := strconv.Quote(unquotedVal)
		variables[key] = quotedVal
	}
	return variables
}

// Component describe component specs
type Component struct {
	name          string
	project       *ProjectConfig
	componentType string
	origin        string
	branch        string
	location      string
	imageName     string
	start         string
	run           string
	livenessCheck string
	containerName string
	ports         map[int]int
	symbol        string
	runtimeSymbol string
	color         int
}

// GetType get component type
func (c *Component) GetType() (componentType string) {
	return c.componentType
}

// GetOrigin get component origin
func (c *Component) GetOrigin() (origin string) {
	return c.origin
}

// GetBranch get component branch
func (c *Component) GetBranch() (branch string) {
	return c.branch
}

// GetLocation get component location
func (c *Component) GetLocation() (location string) {
	return c.location
}

// GetName get component name
func (c *Component) GetName() (name string) {
	return c.name
}

// GetImageName get component imageName
func (c *Component) GetImageName() (imageName string) {
	return c.imageName
}

// GetStartCommand get component start command
func (c *Component) GetStartCommand() (start string) {
	return c.start
}

// GetRunCommand get component run command
func (c *Component) GetRunCommand() (run string) {
	return c.run
}

// GetLivenessCheckCommand get component run command
func (c *Component) GetLivenessCheckCommand() (livenessCheck string) {
	return c.livenessCheck
}

// GetSymbol get component run command
func (c *Component) GetSymbol() (symbol string) {
	return c.symbol
}

// GetContainerName get component container name
func (c *Component) GetContainerName() (containerName string) {
	return c.containerName
}

// GetPorts get component container name
func (c *Component) GetPorts() (ports map[int]int) {
	return c.ports
}

// GetRuntimeContainerName get container name for runtime
func (c *Component) GetRuntimeContainerName() (containerName string) {
	containerName = c.GetContainerName()
	if containerName == "" {
		projectName := c.project.GetName()
		componentName := c.name
		containerName = fmt.Sprintf("%s-%s", projectName, componentName)
	}
	return containerName
}

// GetRuntimeLocation get runtime location
func (c *Component) GetRuntimeLocation() (location string) {
	if c.GetType() == "service" {
		return c.GetLocation()
	}
	return c.project.GetDirName()
}

// GetRuntimeRunCommand get runtime run command
func (c *Component) GetRuntimeRunCommand() (command string) {
	command = c.GetRunCommand()
	if c.GetType() == "container" && command == "" {
		containerName := c.GetRuntimeContainerName()
		imageName := c.GetImageName()
		variables := c.project.GetEnvironments().GetQuotedRuntimeVariables(c.GetName())
		portMap := c.GetPorts()
		// parse ports
		portParams := "--publish-all"
		ports := []string{}
		for key, val := range portMap {
			ports = append(ports, fmt.Sprintf("-p %d:%d", key, val))
		}
		if len(ports) > 0 {
			portParams = strings.Join(ports, " ")
		}
		// parse environments
		environments := []string{}
		for key, val := range variables {
			environments = append(environments, fmt.Sprintf("-e %s=%s", key, val))
		}
		environmentParams := strings.Join(environments, " ")
		// create command
		command = fmt.Sprintf("docker run --name \"%s\" %s %s  -d \"%s\"", containerName, environmentParams, portParams, imageName)
	}
	return command
}

// GetRuntimeStartCommand get runtime start command
func (c *Component) GetRuntimeStartCommand() (command string) {
	command = c.GetStartCommand()
	if c.GetType() == "container" && command == "" {
		containerName := c.GetRuntimeContainerName()
		command = fmt.Sprintf("docker start \"%s\"", containerName)
	}
	return command
}

// GetRuntimeCommand get runtime command (run or start)
func (c *Component) GetRuntimeCommand() (command string) {
	if c.GetType() == "container" {
		startCommand := c.GetRuntimeStartCommand()
		runCommand := c.GetRuntimeRunCommand()
		containerName := c.GetRuntimeContainerName()
		return fmt.Sprintf("(%s || %s) && docker logs --follow %s", startCommand, runCommand, containerName)
	}
	return c.GetRuntimeStartCommand()
}

// GetRuntimeLivenessCheckCommand get runtime start command
func (c *Component) GetRuntimeLivenessCheckCommand() (command string) {
	command = c.GetLivenessCheckCommand()
	if command == "" {
		command = "echo ok"
	}
	return command
}

// GetRuntimeSymbol get component container name
func (c *Component) GetRuntimeSymbol() (runtimeSymbol string) {
	if c.runtimeSymbol == "" {
		if c.symbol != "" {
			c.runtimeSymbol = c.symbol
			return c.symbol
		}
		symbolList := []string{"🍏", "🍎", "🍌", "🍉", "🍇", "🍐", "🍊", "🍋", "🍓", "🍈", "🍒", "🍑", "🍍", "🥝", "🍅", "🍆", "🥑"}
		index := c.project.lastGeneratedSymbolIndex
		c.runtimeSymbol = symbolList[index]
		index++
		if index >= len(symbolList) {
			index = 0
		}
		c.project.lastGeneratedSymbolIndex = index
	}
	return c.runtimeSymbol
}

// GetRuntimeName get component name
func (c *Component) GetRuntimeName() (name string) {
	runtimeName := fmt.Sprintf("%s %s", c.GetRuntimeSymbol(), c.GetName())
	return fmt.Sprintf("%-12v", runtimeName)
}

// GetColor get component name
func (c *Component) GetColor() (color int) {
	if c.color == 0 {
		colorList := []int{92, 93, 94, 95, 96, 91}
		index := c.project.lastGeneratedColorIndex
		c.color = colorList[index]
		index++
		if index >= len(colorList) {
			index = 0
		}
		c.project.lastGeneratedColorIndex = index
	}
	return c.color
}
