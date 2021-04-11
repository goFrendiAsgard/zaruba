package config

import (
	"fmt"
	"regexp"
)

// Variable configuration
type Variable struct {
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
func (variable *Variable) GetName() (name string) {
	return variable.name
}

// Validate validate a value
func (input *Variable) Validate(value string) (err error) {
	if input.Validation == "" {
		return nil
	}
	valid, err := regexp.Match(input.Validation, []byte(value))
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("Value %s does not match %s", value, input.Validation)
	}
	return nil
}
