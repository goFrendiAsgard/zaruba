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
	pYaml, _ := p.ToYaml()
	log.Printf("[INFO] Project Config Loaded:\n%s", pYaml)
	return p, err
}

// newEmptyProjectConfig create new ProjectConfig
func newEmptyProjectConfig() (p *ProjectConfig) {
	return &ProjectConfig{
		ProjectName: "",
		Environments: Environments{
			General:  make(map[string]string),
			Services: make(map[string]map[string]string),
		},
		Components: make(map[string]Component),
		Executions: []string{},
		Links:      make(map[string][]string),
	}
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

func mergeEnvironment(p, subP *ProjectConfig) *ProjectConfig {
	// merge general environment
	for generalSubEnvName, generalSubEnv := range subP.Environments.General {
		if _, exists := p.Environments.General[generalSubEnvName]; !exists {
			p.Environments.General[generalSubEnvName] = generalSubEnv
		}
	}
	// merge service environment
	for serviceName, serviceEnvMap := range subP.Environments.Services {
		// if p doesn't have any environment for the service, add it
		if _, exists := p.Environments.Services[serviceName]; !exists {
			p.Environments.Services[serviceName] = serviceEnvMap
			continue
		}
		// p already has environment for the service, cascade it
		for serviceSubEnvName, serviceSubEnv := range serviceEnvMap {
			if _, exists := p.Environments.Services[serviceName][serviceSubEnvName]; !exists {
				p.Environments.Services[serviceName][serviceSubEnvName] = serviceSubEnv
			}
		}
	}
	return p
}

func mergeComponents(p, subP *ProjectConfig) *ProjectConfig {
	// merge component
	for componentName, component := range subP.Components {
		if _, exists := p.Components[componentName]; !exists {
			p.Components[componentName] = component
		}
	}
	return p
}

func mergeExecutions(p, subP *ProjectConfig) *ProjectConfig {
	// merge component
	for _, subExecution := range subP.Executions {
		exists := false
		for _, execution := range p.Executions {
			if execution == subExecution {
				exists = true
				break
			}
		}
		if !exists {
			p.Executions = append(p.Executions, subExecution)
		}
	}
	return p
}

func mergeLinks(p, subP *ProjectConfig) *ProjectConfig {
	// merge links
	for libPath, subLinks := range subP.Links {
		// if p doesn't have any link for libPath, add it
		if _, exists := p.Links[libPath]; !exists {
			p.Links[libPath] = subLinks
			continue
		}
		for _, subLink := range subLinks {
			exists := false
			for _, link := range p.Links[libPath] {
				if subLink == link {
					exists = true
					break
				}
			}
			if !exists {
				p.Links[libPath] = append(p.Links[libPath], subLink)
			}
		}
	}
	return p
}

// loadSingleProjectConfig load project configuration from a directory
func loadSingleProjectConfig(directory string) (p *ProjectConfig, err error) {
	p = newEmptyProjectConfig()
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
	err = yaml.Unmarshal([]byte(str), p)
	if err != nil {
		return p, err
	}
	p = adjustLocation(p, directory)
	return p, err
}

func adjustLocation(p *ProjectConfig, absDirPath string) *ProjectConfig {
	// adjust component's location
	for componentName, component := range p.Components {
		component.Location = file.GetAbsoluteLocation(absDirPath, component.Location)
		p.Components[componentName] = component
	}
	// adjust component's link
	newLinks := make(map[string][]string)
	for source, destinations := range p.Links {
		newSource := file.GetAbsoluteLocation(absDirPath, source)
		newLinks[newSource] = []string{}
		for _, destination := range destinations {
			newLinks[newSource] = append(newLinks[newSource], file.GetAbsoluteLocation(absDirPath, destination))
		}
	}
	p.Links = newLinks
	return p
}
