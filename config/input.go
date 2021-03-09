package config

// Input configuration
type Input struct {
	DefaultValue string `yaml:"default,omitempty"`
	Description  string `yaml:"description,omitempty"`
	Project      *Project
	fileLocation string
	name         string
}

// GetName get name of input
func (input *Input) GetName() (name string) {
	return input.name
}
