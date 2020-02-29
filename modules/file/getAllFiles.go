package file

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/zealic/xignore"
	"gopkg.in/yaml.v2"
)

// IgnoreConfig describe ignores
type IgnoreConfig struct {
	Ignores []string `yaml:"ignores"`
}

// GetIgnores get ignores from IgnoreList
func (i *IgnoreConfig) GetIgnores() (ignores []string) {
	return i.Ignores
}

// NewIgnoreConfig create new IgnoreConfig
func NewIgnoreConfig() (i *IgnoreConfig) {
	return &IgnoreConfig{
		Ignores: []string{},
	}
}

// LoadIgnoreConfig load ignore config
func LoadIgnoreConfig(dirName string) (i *IgnoreConfig) {
	i = NewIgnoreConfig()
	// read file's content
	b, err := ioutil.ReadFile(filepath.Join(dirName, "zaruba.config.yaml"))
	if err == nil {
		err = yaml.Unmarshal(b, i)
	}
	// add .git
	i.Ignores = append(i.Ignores, ".git")
	// adjust ignores
	for index, ignore := range i.Ignores {
		i.Ignores[index] = GetAbsoluteLocation(dirName, ignore)
	}
	return i
}

// GetAllFiles fetch sub-files and sub-directories recursively
func GetAllFiles(dirName string, option *Option) (allFiles []string, err error) {
	allFiles = []string{dirName}
	// create pattern
	result, err := xignore.DirMatches(dirName, &xignore.MatchesOptions{
		Ignorefile: ".gitignore",
		Nested:     true, // Handle nested ignorefile
	})
	if err != nil {
		return allFiles, err
	}
	// load ignoreConfig
	i := LoadIgnoreConfig(dirName)
	// add all sub-directories that doesn't match gitignore
	for _, subDirName := range result.UnmatchedDirs {
		absSubDirName := path.Join(dirName, subDirName)
		ignored := false
		for _, prefix := range i.GetIgnores() {
			if strings.HasPrefix(absSubDirName, prefix) {
				ignored = true
				break
			}
		}
		if !ignored {
			allFiles = append(allFiles, absSubDirName)
		}
	}
	if !option.GetIsOnlyDir() {
		// add all sub-files that doesn't match gitignore
		for _, subFileName := range result.UnmatchedFiles {
			absFileName := path.Join(dirName, subFileName)
			ignored := false
			for _, prefix := range i.GetIgnores() {
				if strings.HasPrefix(absFileName, prefix) {
					ignored = true
					break
				}
			}
			if !ignored {
				allFiles = append(allFiles, absFileName)
			}
		}
	}
	return allFiles, err
}
