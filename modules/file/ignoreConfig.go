package file

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// IgnoreConfig describe ignores
type IgnoreConfig struct {
	ignores    []string
	absIgnores []string
}

// ShouldIgnore determine whether absFileName (or directory name) should be ignored or not
func (i *IgnoreConfig) ShouldIgnore(absFileName string) (ignored bool) {
	for _, prefix := range i.absIgnores {
		if strings.HasPrefix(absFileName, prefix) {
			return true
		}
	}
	return false
}

// NewIgnoreConfig load ignore config
func NewIgnoreConfig(dirName string) (i *IgnoreConfig) {
	i = &IgnoreConfig{
		ignores:    []string{},
		absIgnores: []string{},
	}
	// initiate ignoreConfig yaml container
	iYaml := &ignoreConfigYaml{
		Ignores: []string{},
	}
	// read file's content
	b, err := ioutil.ReadFile(filepath.Join(dirName, "zaruba.config.yaml"))
	if err != nil {
		return i
	}
	err = yaml.Unmarshal(b, iYaml)
	if err != nil {
		return i
	}
	i.ignores = iYaml.Ignores
	i.ignores = append(i.ignores, ".git")
	// set absIgnores
	for _, ignore := range i.ignores {
		absIgnore := GetAbsoluteLocation(dirName, ignore)
		i.absIgnores = append(i.absIgnores, absIgnore)
	}
	return i
}
