package config

import (
	"fmt"

	"github.com/state-alchemists/zaruba/file"
	yaml "gopkg.in/yaml.v2"
)

func AddTaskDependencies(task *Task, dependencyTaskNames []string) (err error) {
	for _, dependencyTaskName := range dependencyTaskNames {
		if _, dependencyExist := task.Project.Tasks[dependencyTaskName]; !dependencyExist {
			return fmt.Errorf("task %s is not exist", dependencyTaskName)
		}
	}
	taskName := task.GetName()
	yamlLocation := task.GetFileLocation()
	p, err := loadRawProject(yamlLocation)
	if err != nil {
		return err
	}
	p.Tasks[taskName].Dependencies = append(p.Tasks[taskName].Dependencies, dependencyTaskNames...)
	yamlContentB, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return file.WriteText(yamlLocation, string(yamlContentB), 0555)
}

func AddTaskParent(task *Task, parentTaskName string) (err error) {
	if _, dependencyExist := task.Project.Tasks[parentTaskName]; !dependencyExist {
		return fmt.Errorf("task %s is not exist", parentTaskName)
	}
	taskName := task.GetName()
	yamlLocation := task.GetFileLocation()
	p, err := loadRawProject(yamlLocation)
	if err != nil {
		return err
	}
	if len(p.Tasks[taskName].Extends) > 0 {
		p.Tasks[taskName].Extends = append(p.Tasks[taskName].Extends, parentTaskName)
	} else if p.Tasks[taskName].Extend == "" {
		p.Tasks[taskName].Extend = parentTaskName
	} else {
		p.Tasks[taskName].Extends = []string{
			p.Tasks[taskName].Extend,
			parentTaskName,
		}
		p.Tasks[taskName].Extend = ""
	}
	yamlContentB, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return file.WriteText(yamlLocation, string(yamlContentB), 0555)
}
