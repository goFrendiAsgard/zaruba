package watcher

import (
	"io/ioutil"
	"path"

	"github.com/state-alchemists/zaruba/config"
	"gopkg.in/yaml.v2"
)

// HookMap is project hook configuration map
type HookMap map[string]HookConfig

// HookConfig is single configuration for each file hook
type HookConfig struct {
	PostTriggers []string `yaml:"post-triggers"`
	PreTriggers  []string `yaml:"pre-triggers"`
	Links        []string `yaml:"links"`
}

/*
// NewHookMap load new TemplateConfig from a template
func NewHookMap(templatePath string) (HookConfig, error) {
	tc := make(HookMap)
	configFile := path.Join(templatePath, config.HookFile)
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return tc, err
	}
	err = yaml.Unmarshal([]byte(data), &tc)
	return tc, err
}
*/
