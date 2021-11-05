package config

import (
	"fmt"

	"github.com/state-alchemists/zaruba/file"
	yaml "gopkg.in/yaml.v3"
)

func SetTaskConfig(task *Task, configMap map[string]string) (err error) {
	if len(configMap) == 0 {
		return nil
	}
	configRefName := GetTaskConfigRefname(task)
	if configRefName == "" {
		// update taskConfig
		return setTaskConfig(task, configMap)
	}
	// update configRef
	return setConfigRef(task.Project.ConfigRefMap[configRefName], configMap)
}

func setTaskConfig(task *Task, configMap map[string]string) (err error) {
	taskName := task.GetName()
	yamlLocation := task.GetFileLocation()
	node, err := file.ReadYaml(yamlLocation)
	if err != nil {
		return err
	}
	docNode := node.Content[0]
	for index := 0; index < len(docNode.Content); index += 2 {
		keyNode := docNode.Content[index]
		valNode := docNode.Content[index+1]
		if keyNode.Value == "tasks" && valNode.ShortTag() == "!!map" {
			for taskNameIndex := 0; taskNameIndex < len(valNode.Content); taskNameIndex += 2 {
				taskNameNode := valNode.Content[taskNameIndex]
				taskNode := valNode.Content[taskNameIndex+1]
				if taskNameNode.Value == taskName && taskNode.ShortTag() == "!!map" {
					for taskPropKeyIndex := 0; taskPropKeyIndex < len(taskNode.Content); taskPropKeyIndex += 2 {
						taskPropKeyNode := taskNode.Content[taskPropKeyIndex]
						taskPropValNode := taskNode.Content[taskPropKeyIndex+1]
						if taskPropKeyNode.Value == "configs" && taskPropValNode.ShortTag() == "!!map" {
							updateConfigMapNode(taskPropValNode, configMap)
							return file.WriteYaml(yamlLocation, node, 0555, []file.YamlLinesPreprocessors{file.YamlTwoSpace, file.YamlFixEmoji, file.YamlAddLineBreakForTwoSpaceIndented})
						}
					}
					// config not found
					taskNode.Style = yaml.LiteralStyle
					taskNode.Content = append(
						taskNode.Content,
						&yaml.Node{Kind: yaml.ScalarNode, Value: "configs"},
						createConfigMapNode(configMap),
					)
					return file.WriteYaml(yamlLocation, node, 0555, []file.YamlLinesPreprocessors{file.YamlTwoSpace, file.YamlFixEmoji, file.YamlAddLineBreakForTwoSpaceIndented})
				}
			}
		}
	}
	return fmt.Errorf("cannot find task %s in %s", taskName, yamlLocation)
}

func setConfigRef(configRef *ConfigRef, configMap map[string]string) (err error) {
	configRefName := configRef.GetName()
	yamlLocation := configRef.GetFileLocation()
	node, err := file.ReadYaml(yamlLocation)
	if err != nil {
		return err
	}
	docNode := node.Content[0]
	for index := 0; index < len(docNode.Content); index += 2 {
		keyNode := docNode.Content[index]
		valNode := docNode.Content[index+1]
		if keyNode.Value == "configs" && valNode.ShortTag() == "!!map" {
			for configRefNameIndex := 0; configRefNameIndex < len(valNode.Content); configRefNameIndex += 2 {
				configRefNameNode := valNode.Content[configRefNameIndex]
				configRefNode := valNode.Content[configRefNameIndex+1]
				if configRefNameNode.Value == configRefName && configRefNode.ShortTag() == "!!map" {
					updateConfigMapNode(configRefNode, configMap)
					return file.WriteYaml(yamlLocation, node, 0555, []file.YamlLinesPreprocessors{file.YamlTwoSpace, file.YamlFixEmoji, file.YamlAddLineBreakForTwoSpaceIndented})
				}
			}
		}
	}
	return fmt.Errorf("cannot find configRef %s in %s", configRefName, yamlLocation)
}

func updateConfigMapNode(configMapNode *yaml.Node, configMap map[string]string) {
	configMapNode.Style = yaml.LiteralStyle
	for configKey, configVal := range configMap {
		configKeyFound := false
		for configKeyIndex := 0; configKeyIndex < len(configMapNode.Content); configKeyIndex += 2 {
			configKeyNode := configMapNode.Content[configKeyIndex]
			configValNode := configMapNode.Content[configKeyIndex+1]
			// "configs" and configKey found, update
			if configKeyNode.Value == configKey {
				configValNode.SetString(configVal)
				configKeyFound = true
				break
			}
		}
		// "configs" found, but configKey not found, create
		if !configKeyFound {
			configMapNode.Content = append(configMapNode.Content, createConfigNode(configKey, configVal)...)
		}
	}
}

func createConfigMapNode(configMap map[string]string) *yaml.Node {
	newConfigNodes := []*yaml.Node{}
	for configKey, configVal := range configMap {
		newConfigNodes = append(
			newConfigNodes,
			createConfigNode(configKey, configVal)...,
		)
	}
	return &yaml.Node{Kind: yaml.MappingNode, Content: newConfigNodes}
}

func createConfigNode(configKey, configVal string) []*yaml.Node {
	return []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: configKey},
		{Kind: yaml.ScalarNode, Value: configVal},
	}
}
