package core

// Env is task environment
type Env struct {
	From    string `yaml:"from"`
	Default string `yaml:"default"`
	Task    *Task  `yaml:"_task,omitempty"`
}
