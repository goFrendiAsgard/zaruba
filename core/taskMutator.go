package core

import (
	"fmt"

	"github.com/state-alchemists/zaruba/fileutil"
	"github.com/state-alchemists/zaruba/yamlstyler"
	yaml "gopkg.in/yaml.v3"
)

func AddTaskDependencies(task *Task, dependencyTaskNames []string) (err error) {
	if len(dependencyTaskNames) == 0 {
		return nil
	}
	for _, dependencyTaskName := range dependencyTaskNames {
		if _, dependencyExist := task.Project.Tasks[dependencyTaskName]; !dependencyExist {
			return fmt.Errorf("dependency task %s is not exist", dependencyTaskName)
		}
	}
	taskName := task.GetName()
	yamlLocation := task.GetFileLocation()
	fileUtil := fileutil.NewFileUtil()
	node, err := fileUtil.ReadYamlNode(yamlLocation)
	if err != nil {
		return err
	}
	// declare new dependencies
	newDependencyVals := []*yaml.Node{}
	for _, dependencyTaskName := range dependencyTaskNames {
		newDependencyVals = append(newDependencyVals, &yaml.Node{Kind: yaml.ScalarNode, Value: dependencyTaskName})
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
						if taskPropKeyNode.Value == "dependencies" && taskPropValNode.ShortTag() == "!!seq" {
							taskPropValNode.Style = yaml.LiteralStyle
							taskPropValNode.Content = append(taskPropValNode.Content, newDependencyVals...)
							return fileUtil.WriteYamlNode(yamlLocation, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
						}
					}
					taskNode.Style = yaml.LiteralStyle
					taskNode.Content = append(
						taskNode.Content,
						&yaml.Node{Kind: yaml.ScalarNode, Value: "dependencies"},
						&yaml.Node{Kind: yaml.SequenceNode, Content: newDependencyVals},
					)
					return fileUtil.WriteYamlNode(yamlLocation, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
				}
			}
		}
	}
	return fmt.Errorf("cannot find task %s in %s", taskName, yamlLocation)

}

func AddTaskParent(task *Task, parentTaskNames []string) (err error) {
	if len(parentTaskNames) == 0 {
		return nil
	}
	for _, parentTaskName := range parentTaskNames {
		if _, parentExist := task.Project.Tasks[parentTaskName]; !parentExist {
			return fmt.Errorf("parent task %s is not exist", parentTaskName)
		}
	}
	taskName := task.GetName()
	yamlLocation := task.GetFileLocation()
	fileUtil := fileutil.NewFileUtil()
	node, err := fileUtil.ReadYamlNode(yamlLocation)
	if err != nil {
		return err
	}
	// declare new parents
	newParentVals := []*yaml.Node{}
	for _, parentTaskName := range parentTaskNames {
		newParentVals = append(newParentVals, &yaml.Node{Kind: yaml.ScalarNode, Value: parentTaskName})
	}
	docNode := node.Content[0]
	for index := 0; index < len(docNode.Content); index += 2 {
		keyNode := docNode.Content[index]
		valNode := docNode.Content[index+1]
		if keyNode.Value == "tasks" && valNode.ShortTag() == "!!map" {
			// look for "taskName"
			for taskNameIndex := 0; taskNameIndex < len(valNode.Content); taskNameIndex += 2 {
				taskNameNode := valNode.Content[taskNameIndex]
				taskNode := valNode.Content[taskNameIndex+1]
				if taskNameNode.Value == taskName && taskNode.ShortTag() == "!!map" {
					// look for "extend", if it is found, add to newParentVals, remove "extend"
					extendFound := false
					for taskPropKeyIndex := 0; taskPropKeyIndex < len(taskNode.Content); taskPropKeyIndex += 2 {
						taskPropKeyNode := taskNode.Content[taskPropKeyIndex]
						taskPropValNode := taskNode.Content[taskPropKeyIndex+1]
						if taskPropKeyNode.Value == "extend" {
							extendFound = true
							newParentVals = append(newParentVals, taskPropValNode)
							newTaskNodeContent := taskNode.Content[0:taskPropKeyIndex]
							if taskPropKeyIndex+2 < len(taskNode.Content) {
								newTaskNodeContent = append(newTaskNodeContent, taskNode.Content[taskPropKeyIndex+2:]...)
							}
							taskNode.Content = newTaskNodeContent
							break
						}
					}
					// look for "extends"
					for taskPropKeyIndex := 0; taskPropKeyIndex < len(taskNode.Content); taskPropKeyIndex += 2 {
						taskPropKeyNode := taskNode.Content[taskPropKeyIndex]
						taskPropValNode := taskNode.Content[taskPropKeyIndex+1]
						// "extends" found, add our new parents to "extends"
						if taskPropKeyNode.Value == "extends" && taskPropValNode.ShortTag() == "!!seq" {
							taskPropValNode.Style = yaml.LiteralStyle
							taskPropValNode.Content = append(taskPropValNode.Content, newParentVals...)
							return fileUtil.WriteYamlNode(yamlLocation, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
						}
					}
					// "extends" and "extend" not found and we only have one new parent, then we set "extend" to new parent
					if !extendFound && len(newParentVals) == 1 {
						taskNode.Style = yaml.LiteralStyle
						taskNode.Content = append(
							taskNode.Content,
							&yaml.Node{Kind: yaml.ScalarNode, Value: "extend"},
							newParentVals[0],
						)
						return fileUtil.WriteYamlNode(yamlLocation, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
					}
					// "extends" not found and we have multiple parents, then create "extends"
					taskNode.Style = yaml.LiteralStyle
					taskNode.Content = append(
						taskNode.Content,
						&yaml.Node{Kind: yaml.ScalarNode, Value: "extends"},
						&yaml.Node{Kind: yaml.SequenceNode, Content: newParentVals},
					)
					return fileUtil.WriteYamlNode(yamlLocation, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
				}
			}
		}
	}
	return fmt.Errorf("cannot find task %s in %s", taskName, yamlLocation)
}
