package config

// ProjectConfigYaml configuration
type ProjectConfigYaml struct {
	Ignores    []string                 `yaml:"ignores"`    // Ignore list
	Name       string                   `yaml:"name"`       // project name
	Env        map[string]string        `yaml:"env"`        // general environment
	Components map[string]ComponentYaml `yaml:"components"` // components
	Links      map[string][]string      `yaml:"links"`      // links
}
