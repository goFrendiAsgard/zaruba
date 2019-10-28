package creator

import (
	"github.com/state-alchemists/zaruba/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

// TemplateConfig is template configuration
type TemplateConfig map[string]ModeConfig

// ModeConfig is mode configuration
type ModeConfig struct {
	Copy   map[string]string `yaml:"copy"`
	Modify map[string]string `yaml:"modify"`
	Hook   []string          `yaml:"hook"`
}

// Load from yaml
func (tc *TemplateConfig) Load(project string) error {
	configFile := path.Join(project, config.TemplateConfigFile)
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), &tc)
	return err
}
