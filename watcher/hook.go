package watcher

import (
	"io/ioutil"
	"log"
	"path"

	"github.com/state-alchemists/zaruba/config"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

// HookConfig is project hook configuration map
type HookConfig map[string]SingleHookConfig

// SingleHookConfig is single configuration for each file hook
type SingleHookConfig struct {
	PostTriggers []string `yaml:"post-triggers"`
	PreTriggers  []string `yaml:"pre-triggers"`
	Links        []string `yaml:"links"`
	Dir          string
}

// NewHookConfig load new TemplateConfig from a template
func NewHookConfig(currentPath string) (HookConfig, error) {
	hookConfig := make(HookConfig)
	absoluteCurrentPath, err := filepath.Abs(currentPath)
	if err != nil {
		return hookConfig, err
	}
	configFile := path.Join(currentPath, config.HookFile)
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return hookConfig, err
	}
	rawHookConfig := make(HookConfig)
	err = yaml.Unmarshal([]byte(data), &rawHookConfig)
	for key, rawSingleHookConfig := range rawHookConfig {
		links := []string{}
		for _, link := range rawSingleHookConfig.Links {
			link, err = filepath.Abs(path.Join(absoluteCurrentPath, link))
			if err != nil {
				return hookConfig, err
			}
			links = append(links, link)
		}
		singleHookConfig := SingleHookConfig{
			PreTriggers:  rawSingleHookConfig.PreTriggers,
			PostTriggers: rawSingleHookConfig.PostTriggers,
			Dir:          absoluteCurrentPath,
			Links:        links,
		}
		log.Printf("%#v", singleHookConfig)
		absoluteKey, err := filepath.Abs(path.Join(absoluteCurrentPath, key))
		if err != nil {
			return hookConfig, err
		}
		hookConfig[absoluteKey] = singleHookConfig
	}
	return hookConfig, err
}
