package config

// ProjectConfigYaml configuration
type ProjectConfigYaml struct {
	Name       string                   `yaml:"name"`       // project name
	Env        map[string]string        `yaml:"env"`        // general environment
	Components map[string]ComponentYaml `yaml:"components"` // components
	Links      map[string][]string      `yaml:"links"`      // links
}
