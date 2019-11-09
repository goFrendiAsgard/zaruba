package watcher

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/otiai10/copy"
	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

// HookConfig is project hook configuration map
type HookConfig map[string]SingleHookConfig

func (hc HookConfig) getAllLinksByKey(key string, maxDepth int) []string {
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
func (hc HookConfig) GetSortedKeys() []string {
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
func (hc HookConfig) RunAction(shell []string, key string) {
	environ := os.Environ()
	singleHookConfig := hc[key]
	// run pre-triggers
	if err := command.RunMultiple(shell, singleHookConfig.Dir, environ, singleHookConfig.PreTriggers); err != nil {
		log.Println(err)
		return
	}
	// process links
	for _, link := range singleHookConfig.Links {
		if err := copy.Copy(key, link); err != nil {
			log.Println(err)
			return
		}
	}
	// run post-triggers
	if err := command.RunMultiple(shell, singleHookConfig.Dir, environ, singleHookConfig.PostTriggers); err != nil {
		log.Println(err)
		return
	}
}

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
		visited := make(map[string]bool)
		for _, link := range rawSingleHookConfig.Links {
			link, err = filepath.Abs(path.Join(absoluteCurrentPath, link))
			if err != nil {
				return hookConfig, err
			}
			if _, ok := visited[link]; !ok {
				links = append(links, link)
				visited[link] = true
			}
		}
		singleHookConfig := SingleHookConfig{
			PreTriggers:  rawSingleHookConfig.PreTriggers,
			PostTriggers: rawSingleHookConfig.PostTriggers,
			Dir:          absoluteCurrentPath,
			Links:        links,
		}
		absoluteKey, err := filepath.Abs(path.Join(absoluteCurrentPath, key))
		if err != nil {
			return hookConfig, err
		}
		hookConfig[absoluteKey] = singleHookConfig
	}
	return hookConfig, err
}
