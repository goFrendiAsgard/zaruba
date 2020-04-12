package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Component describe component specs
type Component struct {
	name           string         // name of the component (assigned when load config)
	project        *ProjectConfig // parent project
	color          int            // color (assigned automatically)
	runtimeSymbol  string         // emoji, will be `symbol` if assigned or random
	componentType  string         // service, library, container
	origin         string
	branch         string
	location       string
	image          string
	start          string
	run            string
	containerName  string
	ports          map[int]int
	symbol         string
	readinessCheck string
	dependencies   []string
	env            map[string]string
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

// GetImage get component image
func (c *Component) GetImage() (image string) {
	return c.image
}

// GetStartCommand get component start command
func (c *Component) GetStartCommand() (start string) {
	return c.start
}

// GetRunCommand get component run command
func (c *Component) GetRunCommand() (run string) {
	return c.run
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
		image := c.GetImage()
		env := c.GetQuotedRuntimeEnv()
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
		for name, value := range env {
			environments = append(environments, fmt.Sprintf("-e %s=%s", name, value))
		}
		environmentParams := strings.Join(environments, " ")
		// create command
		command = fmt.Sprintf("docker run --name \"%s\" %s %s  -d \"%s\"", containerName, environmentParams, portParams, image)
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

// GetRuntimeReadinessCheckCommand get runtime start command
func (c *Component) GetRuntimeReadinessCheckCommand() (command string) {
	command = c.GetReadinessCheckCommand()
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

// GetEnv get env of service
func (c *Component) GetEnv() (env map[string]string) {
	return c.env
}

// GetDependencies get service dependencies
func (c *Component) GetDependencies() (dependencies []string) {
	return c.dependencies
}

// GetRuntimeEnv get runtime environment variables of a service
func (c *Component) GetRuntimeEnv() (env map[string]string) {
	env = map[string]string{}
	// other service/container name
	for otherServiceName, otherComponent := range c.project.components {
		if otherComponentType := otherComponent.GetType(); otherComponentType != "service" && otherComponentType != "container" {
			continue
		}
		env[otherServiceName] = "0.0.0.0"
	}
	// project env
	for name, value := range c.project.env {
		if osValue := os.Getenv(name); osValue != "" {
			value = osValue
		}
		env[name] = value
	}
	// current service env
	for name, value := range c.env {
		env[name] = value
	}
	// current container name
	if componentType := c.GetType(); componentType == "container" {
		env["CONTAINER_NAME"] = c.GetRuntimeContainerName()
	}
	env = c.parseEnv(env)
	return env
}

func (c *Component) parseEnv(env map[string]string) map[string]string {
	names := []string{}
	for name := range env {
		names = append(names, name)
	}
	for index, name := range names {
		value := env[name]
		for _, prevName := range names[:index] {
			prevValue := env[prevName]
			value = strings.ReplaceAll(value, fmt.Sprintf("${%s}", prevName), prevValue)
			value = strings.ReplaceAll(value, fmt.Sprintf("$%s", prevName), prevValue)
		}
		env[name] = value
	}
	return env
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
