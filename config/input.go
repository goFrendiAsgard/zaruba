package config

// Input configuration
type Input struct {
	DefaultValue string   `yaml:"default,omitempty"`
	Description  string   `yaml:"description,omitempty"`
	Secret       bool     `yaml:"secret,omitempty"`
	Validation   string   `yaml:"validation,omitempty"`
	Options      []string `yaml:"options,omitempty"`
	Prompt       string   `yaml:"prompt,omitempty"`
	AllowCustom  string   `yaml:"allowCustom,omitempty"`
	Project      *Project
	fileLocation string
	name         string
}

// GetName get name of input
func (input *Input) GetName() (name string) {
	return input.name
}
