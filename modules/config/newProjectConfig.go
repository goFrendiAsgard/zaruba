package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"github.com/state-alchemists/zaruba/modules/file"
	"github.com/state-alchemists/zaruba/modules/logger"
	"gopkg.in/yaml.v2"
)

// CreateProjectConfig load project configuration from project directory
func CreateProjectConfig(projectDir string) (p *ProjectConfig, err error) {
	// get absolute project dir
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return p, err
	}
	// load main projectConfig (zaruba.yaml)
	p, err = loadMainProjectConfig(projectDir)
	if err != nil {
		return p, err
	}
	// fetch sub projectConfigs
	configFiles, err := getConfigFiles(projectDir, p.ignores)
	if err != nil {
		return p, err
	}
	for _, configFile := range configFiles {
		subP, err := loadSingleProjectConfig(configFile)
		if err != nil {
			return p, err
		}
		p = mergeEnvironment(p, subP)
		p = mergeComponents(p, subP)
		p = mergeLinks(p, subP)
	}
	// inject project to components (for back-linking)
	for componentName := range p.components {
		p.components[componentName].project = p
		p.components[componentName].name = componentName
	}
	str, _ := p.ToYaml()
	logger.Info("Project Config Loaded:\n%s", str)
	return p, err
}

func loadMainProjectConfig(projectDir string) (p *ProjectConfig, err error) {
	p, err = loadSingleProjectConfig(filepath.Join(projectDir, "zaruba.yaml"))
	if err != nil {
		return p, err
	}
	p.dirName = projectDir
	// set projectName if not exists
	if p.name == "" {
		p.name = filepath.Base(projectDir)
	}
	return p, err
}

func getConfigFiles(parentDir string, ignores []string) (configFiles []string, err error) {
	configFiles = []string{}
	allFiles, err := file.GetAllFiles(
		parentDir,
		file.CreateOption().
			SetIsOnlyDir(false).
			SetIgnores(ignores))
	if err != nil {
		return configFiles, err
	}
	for _, fileName := range allFiles {
		if strings.HasSuffix(fileName, ".zaruba.yaml") {
			configFiles = append(configFiles, fileName)
		}
	}
	return configFiles, err
}

// newEmptyProjectConfig create new ProjectConfig
func newEmptyProjectConfig() (p *ProjectConfig) {
	return &ProjectConfig{
		dirName:                      "",
		ignores:                      []string{},
		name:                         "",
		env:                          make(map[string]string),
		components:                   make(map[string]*Component),
		links:                        make(map[string][]string),
		sortedLinkSources:            []string{},
		isSortedLinkSourcesCached:    false,
		lastGeneratedSymbolIndexLock: &sync.RWMutex{},
		lastGeneratedSymbolIndex:     0,
		lastGeneratedColorIndexLock:  &sync.RWMutex{},
		lastGeneratedColorIndex:      0,
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
func loadSingleProjectConfig(configFile string) (p *ProjectConfig, err error) {
	p = newEmptyProjectConfig()
	pYaml := &ProjectConfigYaml{
		Ignores:    []string{},
		Name:       "",
		Env:        make(map[string]string),
		Components: make(map[string]ComponentYaml),
		Links:      make(map[string][]string),
	}
	configFile, err = filepath.Abs(configFile)
	if err != nil {
		return p, err
	}
	// read file's content
	b, err := ioutil.ReadFile(configFile)
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
	p.fromProjectConfigYaml(pYaml, configFile)
	// adjust location
	p = adjustLocation(p, filepath.Dir(configFile))
	return p, err
}

func adjustLocation(p *ProjectConfig, absDirPath string) *ProjectConfig {
	// adjust component's ignores
	for index, ignore := range p.ignores {
		p.ignores[index] = file.GetAbsoluteLocation(absDirPath, ignore)
	}
	// adjust component's location
	for componentName, component := range p.components {
		component.location = file.GetAbsoluteLocation(absDirPath, component.location)
		volumes := map[string]string{}
		for hostLoc, containerLoc := range component.volumes {
			volumes[file.GetAbsoluteLocation(component.location, hostLoc)] = containerLoc
		}
		component.volumes = volumes
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
