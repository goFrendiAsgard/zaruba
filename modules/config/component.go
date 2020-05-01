package config

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Component describe component specs
type Component struct {
	name           string         // name of the component (assigned when load config)
	project        *ProjectConfig // parent project
	color          int            // color (assigned automatically)
	runtimeSymbol  string         // emoji, will be `symbol` if assigned or random
	componentType  string         // service, library, container
	origin         string
	location       string
	image          string
	start          string
	containerName  string
	ports          map[int]int
	volumes        map[string]string
	symbol         string
	readinessCheck string
	dependencies   []string
	env            map[string]string
	runtimeName    string
}

// GetType get component type
func (c *Component) GetType() (componentType string) {
	return c.componentType
}

// GetOrigin get component origin
func (c *Component) GetOrigin() (origin string) {
	return c.origin
}

// GetLocation get component location
func (c *Component) GetLocation() (location string) {
	return c.location
}

// GetName get component name
func (c *Component) GetName() (name string) {
	return c.name
}

// GetImage get component image
func (c *Component) GetImage() (image string) {
	return c.image
}

// GetStartCommand get component start command
func (c *Component) GetStartCommand() (start string) {
	return c.start
}

// GetReadinessCheckCommand get component run command
func (c *Component) GetReadinessCheckCommand() (readinessCheck string) {
	return c.readinessCheck
}

// GetSymbol get component run command
func (c *Component) GetSymbol() (symbol string) {
	return c.symbol
}

// GetContainerName get component container name
func (c *Component) GetContainerName() (containerName string) {
	return c.containerName
}

// GetVolumes get component container name
func (c *Component) GetVolumes() (volumens map[string]string) {
	return c.volumes
}

// GetPorts get component container name
func (c *Component) GetPorts() (ports map[int]int) {
	return c.ports
}

// GetEnv get env of service
func (c *Component) GetEnv() (env map[string]string) {
	return c.env
}

// GetDependencies get service dependencies
func (c *Component) GetDependencies() (dependencies []string) {
	return c.dependencies
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

// GetRuntimeStartCommand get runtime run command
func (c *Component) GetRuntimeStartCommand() (command string) {
	command = c.GetStartCommand()
	if c.GetType() == "container" && command == "" {
		location := c.GetRuntimeLocation()
		projectName := c.project.name
		containerName := c.GetRuntimeContainerName()
		image := c.GetImage()
		env := c.GetQuotedRuntimeEnv()
		portMap := c.GetPorts()
		volumeMap := c.GetVolumes()
		// parse ports
		portParams := "--publish-all"
		ports := []string{}
		for key, val := range portMap {
			ports = append(ports, fmt.Sprintf("-p %d:%d", key, val))
		}
		if len(ports) > 0 {
			portParams = strings.Join(ports, " ")
		}
		// parse volume
		volumes := []string{}
		for key, val := range volumeMap {
			volumes = append(volumes, fmt.Sprintf("-v %s:%s", filepath.Join(location, key), val))
		}
		volumeParams := strings.Join(volumes, " ")
		// parse environments
		environments := []string{}
		for name, value := range env {
			environments = append(environments, fmt.Sprintf("-e %s=%s", name, value))
		}
		environmentParams := strings.Join(environments, " ")
		// create command
		command = fmt.Sprintf("docker run  --name \"%s\" --net=\"%s\" %s %s %s  -d \"%s\"", containerName, projectName, volumeParams, environmentParams, portParams, image)
	}
	return command
}

// GetRuntimeCommand get runtime command (run or start)
func (c *Component) GetRuntimeCommand() (command string) {
	if c.GetType() == "container" {
		runCommand := c.GetRuntimeStartCommand()
		containerName := c.GetRuntimeContainerName()
		startOrRunCommand := fmt.Sprintf("(docker start \"%s\" || %s)", containerName, runCommand)
		startOrRunAndFollowCommand := fmt.Sprintf("%s && docker logs --since 0m --follow %s", startOrRunCommand, containerName)
		return startOrRunAndFollowCommand
	}
	return c.GetStartCommand()
}

// GetRuntimeReadinessCheckCommand get runtime start command
func (c *Component) GetRuntimeReadinessCheckCommand() (command string) {
	command = c.GetReadinessCheckCommand()
	if command == "" {
		if c.componentType != "container" {
			command = "echo ok"
			return command
		}
		command = fmt.Sprintf("if [ $(docker inspect -f '{{.State.Running}}' \"%s\") = true ]; then echo ok; else echo notOk 1>&2; fi;", c.GetRuntimeContainerName())
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
	if c.runtimeName == "" {
		maxRuntimeNameLength := 12
		for _, otherC := range c.project.GetComponents() {
			runtimeName := fmt.Sprintf("%s %s", otherC.GetRuntimeSymbol(), otherC.GetName())
			if nameLength := utf8.RuneCountInString(runtimeName); nameLength > maxRuntimeNameLength {
				maxRuntimeNameLength = nameLength
			}
		}
		runtimeName := fmt.Sprintf("%s %s", c.GetRuntimeSymbol(), c.GetName())
		for utf8.RuneCountInString(runtimeName) < maxRuntimeNameLength {
			runtimeName += " "
		}
		c.runtimeName = runtimeName
	}
	return c.runtimeName
}

// GetColor get component name
func (c *Component) GetColor() (color int) {
	if c.color == 0 {
		colorList := []int{92, 93, 94, 95, 96}
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

// GetRuntimeEnv get runtime environment variables of a service
func (c *Component) GetRuntimeEnv() (env map[string]string) {
	venv := NewVirtualEnv()
	for otherServiceName, otherComponent := range c.project.components {
		if otherComponentType := otherComponent.GetType(); otherComponentType != "service" && otherComponentType != "container" {
			continue
		}
		if c.componentType == "container" && otherComponent.GetType() == "container" {
			venv.Add(otherServiceName, otherComponent.GetRuntimeContainerName())
		} else {
			venv.Add(otherServiceName, "0.0.0.0")
		}
	}
	// project env
	for name, value := range c.project.env {
		venv.Add(name, value)
	}
	// current service env
	for name, value := range c.env {
		venv.Add(name, value)
	}
	// current container name
	if componentType := c.GetType(); componentType == "container" {
		venv.Add("CONTAINER_NAME", c.GetRuntimeContainerName())
	}
	return venv.GetEnv()
}

// GetQuotedRuntimeEnv get runtime environment variables of a service
func (c *Component) GetQuotedRuntimeEnv() (env map[string]string) {
	env = c.GetRuntimeEnv()
	for name, value := range env {
		unquotedVal, err := strconv.Unquote(value)
		if err != nil {
			unquotedVal = value
		}
		env[name] = strconv.Quote(unquotedVal)
	}
	return env
}
