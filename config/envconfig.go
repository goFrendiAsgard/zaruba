package config

// EnvConfig is task environment
type EnvConfig struct {
	From    string `yaml:"from"`
	Default string `yaml:"default"`
	Parent  *Task
}
