package config

import (
	"github.com/state-alchemists/zaruba/file"
	yaml "gopkg.in/yaml.v3"
)

func IncludeFileToProject(mainP *Project, fileName string) (err error) {
	yamlLocation := mainP.GetFileLocation()
	node, err := file.ReadYaml(yamlLocation)
	if err != nil {
		return err
	}
	docNode := node.Content[0]
	// declare new includeVal
	newIncludeVal := &yaml.Node{Kind: yaml.ScalarNode, Value: fileName}
	// look for "includes"
	for index := 0; index < len(docNode.Content); index += 2 {
		keyNode := docNode.Content[index]
		valNode := docNode.Content[index+1]
		if keyNode.Value == "includes" && valNode.ShortTag() == "!!seq" {
			valNode.Style = yaml.LiteralStyle
			valNode.Content = append(valNode.Content, newIncludeVal)
			return file.WriteYaml(yamlLocation, node, 0555, []file.YamlLinesPreprocessors{file.YamlTwoSpace, file.YamlFixEmoji, file.YamlAddLineBreakForTwoSpaceIndented})
		}
	}
	includesKey := &yaml.Node{Kind: yaml.ScalarNode, Value: "includes"}
	includesVal := &yaml.Node{Kind: yaml.SequenceNode, Content: []*yaml.Node{newIncludeVal}}
	docNode.Style = yaml.LiteralStyle
	docNode.Content = append(
		[]*yaml.Node{includesKey, includesVal},
		docNode.Content...,
	)
	return file.WriteYaml(yamlLocation, node, 0555, []file.YamlLinesPreprocessors{file.YamlTwoSpace, file.YamlFixEmoji, file.YamlAddLineBreakForTwoSpaceIndented})
}

func CreateTaskIfNotExist(mainP *Project, taskName string) (err error) {
	yamlLocation := mainP.GetFileLocation()
	node, err := file.ReadYaml(yamlLocation)
	if err != nil {
		return err
	}
	docNode := node.Content[0]
	// declare newTaskName node and newTask node
	newTaskName := &yaml.Node{Kind: yaml.ScalarNode, Value: taskName}
	newTask := &yaml.Node{Kind: yaml.MappingNode}
	// look for "tasks"
	for index := 0; index < len(docNode.Content); index += 2 {
		keyNode := docNode.Content[index]
		valNode := docNode.Content[index+1]
		if keyNode.Value == "tasks" && valNode.ShortTag() == "!!map" {
			for taskNameIndex := 0; taskNameIndex < len(valNode.Content); taskNameIndex += 2 {
				if valNode.Content[taskNameIndex].Value == taskName {
					return nil
				}
			}
			valNode.Style = yaml.LiteralStyle
			valNode.Content = append(valNode.Content, newTaskName, newTask)
			return file.WriteYaml(yamlLocation, node, 0555, []file.YamlLinesPreprocessors{file.YamlTwoSpace, file.YamlFixEmoji, file.YamlAddLineBreakForTwoSpaceIndented})
		}
	}
	// "tasks" not found, add it
	docNode.Style = yaml.LiteralStyle
	docNode.Content = append(
		[]*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "tasks"},
			{
				Kind: yaml.MappingNode, Content: []*yaml.Node{
					newTaskName, newTask,
				},
			},
		},
		docNode.Content...,
	)
	return file.WriteYaml(yamlLocation, node, 0555, []file.YamlLinesPreprocessors{file.YamlTwoSpace, file.YamlFixEmoji, file.YamlAddLineBreakForTwoSpaceIndented})
}
