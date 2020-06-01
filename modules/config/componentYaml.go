package config

// ComponentYaml describe component specs
type ComponentYaml struct {
	Type           string            `yaml:"type"`           // Component's type: "service", "library", "container", "command"
	Labels         map[string]string `yaml:"labels"`         // Component's label
	Origin         string            `yaml:"origin"`         // Component's git origin
	Location       string            `yaml:"location"`       // location of the component
	Image          string            `yaml:"image"`          // container's image
	Start          string            `yaml:"start"`          // start service command
	ContainerName  string            `yaml:"containerName"`  // container name
	Ports          map[string]string `yaml:"ports"`          // exposed ports (host:container)
	Volumes        map[string]string `yaml:"volumes"`        // volumes (host:container)
	Symbol         string            `yaml:"symbol"`         // emoji representation
	ReadinessCheck string            `yaml:"readinessCheck"` // check readiness command
	ReadinessURL   string            `yaml:"readinessUrl"`   // check readiness command
	Dependencies   []string          `yaml:"dependencies"`   // dependencies
	Env            map[string]string `yaml:"env"`            // environment
}
