package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/modules/file"
	"github.com/state-alchemists/zaruba/modules/logger"
	"gopkg.in/yaml.v2"
)

// NewProjectConfig load project configuration from project directory
func NewProjectConfig(projectDir string) (p *ProjectConfig, err error) {
	allDirs, err := getAllDirs(projectDir)
	if err != nil {
		return p, err
	}
	p = newEmptyProjectConfig()
	p.dirName = projectDir
	for _, directory := range allDirs {
		subP, loadSubErr := loadSingleProjectConfig(directory)
		if loadSubErr != nil {
			if os.IsNotExist(loadSubErr) {
				continue
			}
			err = loadSubErr
			break
		}
		p = mergeEnvironment(p, subP)
		p = mergeComponents(p, subP)
		p = mergeLinks(p, subP)
	}
	// set projectName if not exists
	if p.name == "" {
		p.name = filepath.Base(projectDir)
	}
	// inject project to components
	for componentName := range p.components {
		p.components[componentName].project = p
		p.components[componentName].name = componentName
	}
	str, _ := p.ToYaml()
	logger.Info("Project Config Loaded: %s", str)
	return p, err
}

func getAllDirs(parentDir string) (allDirs []string, err error) {
	allDirs = []string{}
	allDirs, err = file.GetAllFiles(parentDir, file.NewOption().SetIsOnlyDir(true))
	return allDirs, err
}

// newEmptyProjectConfig create new ProjectConfig
func newEmptyProjectConfig() (p *ProjectConfig) {
	return &ProjectConfig{
		name:                      "",
		env:                       make(map[string]string),
		components:                make(map[string]*Component),
		links:                     make(map[string][]string),
		sortedLinkSources:         []string{},
		isSortedLinkSourcesCached: false,
		lastGeneratedSymbolIndex:  0,
		lastGeneratedColorIndex:   0,
	}
}

func mergeEnvironment(p, subP *ProjectConfig) *ProjectConfig {
	for name, value := range subP.env {
		if _, exists := p.env[name]; !exists {
			p.env[name] = value
		}
	}
	return p
}

func mergeComponents(p, subP *ProjectConfig) *ProjectConfig {
	for componentName, component := range subP.components {
		if _, exists := p.components[componentName]; !exists {
			p.components[componentName] = component
		}
	}
	return p
}

func mergeLinks(p, subP *ProjectConfig) *ProjectConfig {
	// merge links
	for libPath, subLinks := range subP.links {
		// if p doesn't have any link for libPath, add it
		if _, exists := p.links[libPath]; !exists {
			p.links[libPath] = subLinks
			continue
		}
		for _, subLink := range subLinks {
			exists := false
			for _, link := range p.links[libPath] {
				if subLink == link {
					exists = true
					break
				}
			}
			if !exists {
				p.links[libPath] = append(p.links[libPath], subLink)
			}
		}
	}
	return p
}

// loadSingleProjectConfig load project configuration from a directory
func loadSingleProjectConfig(directory string) (p *ProjectConfig, err error) {
	p = newEmptyProjectConfig()
	pYaml := &ProjectConfigYaml{
		Name:       "",
		Env:        make(map[string]string),
		Components: make(map[string]ComponentYaml),
		Links:      make(map[string][]string),
	}
	directory, err = filepath.Abs(directory)
	if err != nil {
		return p, err
	}
	// read file's content
	b, err := ioutil.ReadFile(filepath.Join(directory, "zaruba.config.yaml"))
	if err != nil {
		return p, err
	}
	str := string(b)
	// create new ProjectConfig and unmarshal
	err = yaml.Unmarshal([]byte(str), pYaml)
	if err != nil {
		return p, err
	}
	// load pYaml into p
	p.fromProjectConfigYaml(pYaml, directory)
	// adjust location
	p = adjustLocation(p, directory)
	return p, err
}

func adjustLocation(p *ProjectConfig, absDirPath string) *ProjectConfig {
	// adjust component's location
	for componentName, component := range p.components {
		component.location = file.GetAbsoluteLocation(absDirPath, component.location)
		p.components[componentName] = component
	}
	// adjust component's link
	newLinks := make(map[string][]string)
	for source, destinations := range p.links {
		newSource := file.GetAbsoluteLocation(absDirPath, source)
		newLinks[newSource] = []string{}
		for _, destination := range destinations {
			newLinks[newSource] = append(newLinks[newSource], file.GetAbsoluteLocation(absDirPath, destination))
		}
	}
	p.links = newLinks
	return p
}
