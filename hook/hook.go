package hook

import (
	"io/ioutil"
	"path"
	"sort"
	"strings"

	"github.com/otiai10/copy"
	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

// SingleConfig is single configuration for each file hook
type SingleConfig struct {
	PostTriggers []string `yaml:"post-triggers"`
	PreTriggers  []string `yaml:"pre-triggers"`
	Links        []string `yaml:"links"`
	Dir          string
}

// Config is project hook configuration map
type Config map[string]SingleConfig

func (hc Config) getAllLinksByKey(key string, maxDepth int) []string {
	links := hc[key].Links
	if maxDepth == 0 {
		return links
	}
	visited := make(map[string]bool)
	for _, link := range links {
		visited[link] = true
	}
	// browse for all other keys
	for otherKey := range hc {
		// skip current key
		if otherKey == key {
			continue
		}
		for _, link := range links {
			// if otherKey is part of current links, fetch it
			if strings.HasPrefix(link, otherKey) {
				// add otherKey
				if _, ok := visited[otherKey]; !ok {
					links = append(links, otherKey)
					visited[otherKey] = true
				}
				// add subLinks
				subLinks := hc.getAllLinksByKey(otherKey, maxDepth-1)
				for _, subLink := range subLinks {
					// if subLink is not visited, add it to links
					if _, ok := visited[subLink]; !ok {
						links = append(links, subLink)
						visited[subLink] = true
					}
				}
			}
		}
	}
	return links
}

// GetSortedKeys sort keys based on dependency tree
func (hc Config) GetSortedKeys() []string {
	keys := []string{}
	for key := range hc {
		keys = append(keys, key)
	}
	linksMemo := make(map[string][]string)
	sort.SliceStable(keys, func(i int, j int) bool {
		firstKey, secondKey := keys[i], keys[j]
		// get firstLinks
		if _, ok := linksMemo[firstKey]; !ok {
			linksMemo[firstKey] = hc.getAllLinksByKey(firstKey, 100)
		}
		firstLinks := linksMemo[firstKey]
		// do comparison
		for _, link := range firstLinks {
			// if links of i contain j, then i should preceed j
			if strings.HasPrefix(secondKey, link) {
				return true
			}
		}
		return false
	})
	return keys
}

// RunAction run a single hook
func (hc Config) RunAction(shell []string, environ []string, key string) (err error) {
	singleHookConfig := hc[key]
	// run pre-triggers
	if err = command.RunMultiple(shell, singleHookConfig.Dir, environ, singleHookConfig.PreTriggers); err != nil {
		return
	}
	// process links
	for _, link := range singleHookConfig.Links {
		if err = copy.Copy(key, link); err != nil {
			return
		}
	}
	// run post-triggers
	err = command.RunMultiple(shell, singleHookConfig.Dir, environ, singleHookConfig.PostTriggers)
	return
}

// NewConfig load new TemplateConfig from a template
func NewConfig(currentPath string) (Config, error) {
	hc := make(Config)
	absoluteCurrentPath, err := filepath.Abs(currentPath)
	if err != nil {
		return hc, err
	}
	configFile := path.Join(currentPath, config.HookFile)
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return hc, err
	}
	rawHookConfig := make(Config)
	err = yaml.Unmarshal([]byte(data), &rawHookConfig)
	for key, rawSingleHookConfig := range rawHookConfig {
		links := []string{}
		visited := make(map[string]bool)
		for _, link := range rawSingleHookConfig.Links {
			link, err = filepath.Abs(path.Join(absoluteCurrentPath, link))
			if err != nil {
				return hc, err
			}
			if _, ok := visited[link]; !ok {
				links = append(links, link)
				visited[link] = true
			}
		}
		singleHookConfig := SingleConfig{
			PreTriggers:  rawSingleHookConfig.PreTriggers,
			PostTriggers: rawSingleHookConfig.PostTriggers,
			Dir:          absoluteCurrentPath,
			Links:        links,
		}
		absoluteKey, err := filepath.Abs(path.Join(absoluteCurrentPath, key))
		if err != nil {
			return hc, err
		}
		hc[absoluteKey] = singleHookConfig
	}
	return hc, err
}

// NewCascadedConfig create cascaded config
func NewCascadedConfig(allDirPaths []string) (hc Config, err error) {
	hc = make(Config)
	for _, dirPath := range allDirPaths {
		currentHookConfig, err := NewConfig(dirPath)
		if err != nil {
			continue
		}
		for key, val := range currentHookConfig {
			if _, ok := hc[key]; !ok {
				hc[key] = SingleConfig{}
			}
			singleConfig := hc[key]
			singleConfig.Dir = val.Dir
			for _, preTrigger := range val.PreTriggers {
				singleConfig.PreTriggers = append(singleConfig.PreTriggers, preTrigger)
			}
			for _, postTrigger := range val.PostTriggers {
				singleConfig.PostTriggers = append(singleConfig.PostTriggers, postTrigger)
			}
			for _, link := range val.Links {
				singleConfig.Links = append(singleConfig.Links, link)
			}
			hc[key] = singleConfig
		}
	}
	return
}
