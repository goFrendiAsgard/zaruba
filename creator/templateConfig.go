package creator

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/state-alchemists/zaruba/config"
	"gopkg.in/yaml.v2"
)

// TemplateConfig is template configuration
type TemplateConfig map[string]ModeConfig

func (tc TemplateConfig) getBaseMode() ModeConfig {
	if baseMode, ok := tc["base"]; ok {
		return baseMode
	}
	return ModeConfig{}
}

// GetMode get modeConfig from TemplateConfig
func (tc TemplateConfig) GetMode(mode string) (ModeConfig, error) {
	baseConfig := tc.getBaseMode()
	if mode == "" || mode == "base" {
		return baseConfig, nil
	}
	if modeConfig, ok := tc[mode]; ok {
		modeConfig.Copy = mergeMaps(baseConfig.Copy, modeConfig.Copy)
		modeConfig.CopyAndSubstitute = mergeMaps(baseConfig.CopyAndSubstitute, modeConfig.CopyAndSubstitute)
		modeConfig.Substitutions = mergeMaps(baseConfig.Substitutions, modeConfig.Substitutions)
		modeConfig.PostTriggers = append(baseConfig.PostTriggers, modeConfig.PostTriggers...)
		modeConfig.PreTriggers = append(baseConfig.PreTriggers, modeConfig.PreTriggers...)
		return modeConfig, nil
	}
	return baseConfig, fmt.Errorf("Undefined mode: `%s`", mode)
}

// ModeConfig is mode configuration
type ModeConfig struct {
	Copy              map[string]string `yaml:"copy"`
	CopyAndSubstitute map[string]string `yaml:"copy-and-substitute"`
	Substitutions     map[string]string `yaml:"substitutions"`
	PostTriggers      []string          `yaml:"post-triggers"`
	PreTriggers       []string          `yaml:"pre-triggers"`
}

// NewTemplateConfig load new TemplateConfig from a template
func NewTemplateConfig(templatePath string) (TemplateConfig, error) {
	tc := make(TemplateConfig)
	configFile := path.Join(templatePath, config.TemplateConfigFile)
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return tc, err
	}
	err = yaml.Unmarshal([]byte(data), &tc)
	return tc, err
}

func mergeMaps(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
