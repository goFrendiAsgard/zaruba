package config

// Input configuration
type Input struct {
	DefaultValue string `yaml:"default,omitempty"`
	Description  string `yaml:"description,omitempty"`
	FileLocation string
}
