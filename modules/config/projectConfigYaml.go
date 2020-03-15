package config

// ProjectConfigYaml configuration
type ProjectConfigYaml struct {
	ProjectName  string                   `yaml:"name"`
	Environments EnvironmentsYaml         `yaml:"environments"`
	Components   map[string]ComponentYaml `yaml:"components"`
	Executions   []string                 `yaml:"executions"`
	Links        map[string][]string      `yaml:"links"`
}

// EnvironmentsYaml describe environment variables in general and for each services
type EnvironmentsYaml struct {
	General  map[string]string            `yaml:"general"`
	Services map[string]map[string]string `yaml:"services"`
}

// ComponentYaml describe component specs
type ComponentYaml struct {
	Type          string      `yaml:"type"`          // Component type: "service", "library", "other", or "container"
	Origin        string      `yaml:"origin"`        // Component's git origin
	Branch        string      `yaml:"branch"`        // Component's branch
	Location      string      `yaml:"location"`      // location of the component
	ImageName     string      `yaml:"imageName"`     // image name of container component
	Start         string      `yaml:"start"`         // command to start service component
	Run           string      `yaml:"run"`           // command to run container component
	ContainerName string      `yaml:"containerName"` // container name of container component
	Ports         map[int]int `yaml:"ports"`         // container exposed ports (host:container)
	Symbol        string      `yaml:"symbol"`        // emoji (when running the component)
	LivenessCheck string      `yaml:"livenessCheck"` // command to check service liveness
}
