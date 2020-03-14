package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/modules/file"
	"gopkg.in/yaml.v2"
)

// NewProjectConfig load project configuration from project directory
func NewProjectConfig(args ...string) (p *ProjectConfig, err error) {
	allDirs, err := getAllDirs(args...)
	if err != nil {
		return p, err
	}
	p = newEmptyProjectConfig()
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
		p = mergeExecutions(p, subP)
		p = mergeLinks(p, subP)
	}
	str, _ := p.ToYaml()
	log.Printf("[INFO] Project Config Loaded:\n%s", str)
	return p, err
}

func getAllDirs(args ...string) (allDirs []string, err error) {
	allDirs = []string{}
	if len(args) > 0 {
		projectDir := args[0]
		allDirs, err = file.GetAllFiles(projectDir, file.NewOption().SetIsOnlyDir(true))
		return allDirs, err
	}
	return allDirs, nil
}

// newEmptyProjectConfig create new ProjectConfig
func newEmptyProjectConfig() (p *ProjectConfig) {
	return &ProjectConfig{
		projectName: "",
		environments: &Environments{
			general:  make(map[string]string),
			services: make(map[string]map[string]string),
		},
		components: make(map[string]*Component),
		executions: []string{},
		links:      make(map[string][]string),
	}
}

func mergeEnvironment(p, subP *ProjectConfig) *ProjectConfig {
	// merge general environment
	for generalSubEnvName, generalSubEnv := range subP.environments.general {
		if _, exists := p.environments.general[generalSubEnvName]; !exists {
			p.environments.general[generalSubEnvName] = generalSubEnv
		}
	}
	// merge service environment
	for serviceName, serviceEnvMap := range subP.environments.services {
		// if p doesn't have any environment for the service, add it
		if _, exists := p.environments.services[serviceName]; !exists {
			p.environments.services[serviceName] = serviceEnvMap
			continue
		}
		// p already has environment for the service, cascade it
		for serviceSubEnvName, serviceSubEnv := range serviceEnvMap {
			if _, exists := p.environments.services[serviceName][serviceSubEnvName]; !exists {
				p.environments.services[serviceName][serviceSubEnvName] = serviceSubEnv
			}
		}
	}
	return p
}

func mergeComponents(p, subP *ProjectConfig) *ProjectConfig {
	// merge component
	for componentName, component := range subP.components {
		if _, exists := p.components[componentName]; !exists {
			p.components[componentName] = component
		}
	}
	return p
}

func mergeExecutions(p, subP *ProjectConfig) *ProjectConfig {
	// merge component
	for _, subExecution := range subP.executions {
		exists := false
		for _, execution := range p.executions {
			if execution == subExecution {
				exists = true
				break
			}
		}
		if !exists {
			p.executions = append(p.executions, subExecution)
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
		ProjectName: "",
		Environments: EnvironmentsYaml{
			General:  make(map[string]string),
			Services: make(map[string]map[string]string),
		},
		Components: make(map[string]ComponentYaml),
		Executions: []string{},
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
