package config

// ProjectConfigYaml configuration
type ProjectConfigYaml struct {
	Ignores       []string                 `yaml:"ignores"`       // Ignore list
	Name          string                   `yaml:"name"`          // project name
	DefaultBranch string                   `yaml:"defaultBranch"` // default branch name
	Env           map[string]string        `yaml:"env"`           // general environment
	Components    map[string]ComponentYaml `yaml:"components"`    // components
	Links         map[string][]string      `yaml:"links"`         // links
}
