package config

// Variable configuration
type Variable struct {
	defaultValue string `yaml:"default,omitempty"`
	description  string `yaml:"description,omitempty"`
}
