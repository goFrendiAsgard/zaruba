package config

import (
	"github.com/state-alchemists/zaruba/file"
	yaml "gopkg.in/yaml.v2"
)

func IncludeFileToProject(mainP *Project, fileName string) (err error) {
	yamlLocation := mainP.GetFileLocation()
	p, err := loadRawProject(yamlLocation)
	if err != nil {
		return err
	}
	p.Includes = append(p.Includes, fileName)
	yamlContentB, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return file.WriteText(yamlLocation, string(yamlContentB), 0555)
}

func CreateTaskIfNotExist(mainP *Project, taskName string) (err error) {
	yamlLocation := mainP.GetFileLocation()
	p, err := loadRawProject(yamlLocation)
	if err != nil {
		return err
	}
	if _, exist := mainP.Tasks[taskName]; exist {
		return nil
	}
	p.Tasks[taskName] = &Task{}
	yamlContentB, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return file.WriteText(yamlLocation, string(yamlContentB), 0555)
}
