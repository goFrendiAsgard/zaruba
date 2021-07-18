package config

import (
	"github.com/state-alchemists/zaruba/file"
	yaml "gopkg.in/yaml.v2"
)

func SetTaskConfig(task *Task, configMap map[string]string) (err error) {
	configRefName := GetTaskConfigRefname(task)
	if configRefName == "" {
		// update taskConfig
		return setTaskConfig(task, configMap)
	}
	// update configRef
	return setConfigRef(task.Project.ConfigRefMap[configRefName], configMap)
}

func setTaskConfig(task *Task, configMap map[string]string) (err error) {
	if len(configMap) == 0 {
		return nil
	}
	taskName := task.GetName()
	yamlLocation := task.GetFileLocation()
	p, err := loadRawProject(yamlLocation)
	if err != nil {
		return err
	}
	for configKey, configVal := range configMap {
		p.Tasks[taskName].Config[configKey] = configVal
	}
	yamlContentB, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return file.WriteText(yamlLocation, string(yamlContentB), 0555)
}

func setConfigRef(configRef *ConfigRef, configMap map[string]string) (err error) {
	if len(configMap) == 0 {
		return nil
	}
	configRefName := configRef.GetName()
	yamlLocation := configRef.GetFileLocation()
	p, err := loadRawProject(yamlLocation)
	if err != nil {
		return err
	}
	for configKey, configVal := range configMap {
		p.RawConfigRefMap[configRefName][configKey] = configVal
	}
	yamlContentB, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return file.WriteText(yamlLocation, string(yamlContentB), 0555)
}
