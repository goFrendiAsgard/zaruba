package dsl

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/fileutil"
	"github.com/state-alchemists/zaruba/jsonutil"
	"github.com/state-alchemists/zaruba/yamlstyler"
	yaml "gopkg.in/yaml.v3"
)

type ProjectUtil struct {
	file               *fileutil.FileUtil
	Task               *TaskUtil
	locationProjectMap map[string]*Project
}

func NewProjectUtil(fileUtil *fileutil.FileUtil, jsonUtil *jsonutil.JsonUtil) *ProjectUtil {
	taskUtil := NewTaskUtil(fileUtil, jsonUtil)
	projectUtil := &ProjectUtil{
		file:               fileUtil,
		Task:               taskUtil,
		locationProjectMap: map[string]*Project{},
	}
	projectUtil.Task.project = projectUtil
	return projectUtil
}

func (projectUtil *ProjectUtil) getProject(projectFile string) (project *Project, err error) {
	project, exist := projectUtil.locationProjectMap[projectFile]
	if exist {
		return project, nil
	}
	project, err = NewDefaultProject(projectFile)
	projectUtil.locationProjectMap[projectFile] = project
	if err != nil {
		return nil, err
	}
	err = project.Init()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (projectUtil *ProjectUtil) SetValue(key, newValue, projectFile string) (err error) {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if newValue == "" {
		return fmt.Errorf("value cannot be empty")
	}
	project, err := projectUtil.getProject(projectFile)
	if err != nil {
		return err
	}
	oldValue := project.GetValue(key)
	valueFileNames, err := project.GetValueFileNames()
	if err != nil {
		return err
	}
	for _, valueFileName := range valueFileNames {
		fileContentB, err := ioutil.ReadFile(valueFileName)
		if err != nil {
			return err
		}
		valueMap := map[string]string{}
		if err := yaml.Unmarshal(fileContentB, &valueMap); err != nil {
			return err
		}
		currentValue, exist := valueMap[key]
		if !exist || currentValue == oldValue || strings.HasSuffix(valueFileName, "default.values.yaml") {
			valueMap[key] = newValue
		}
		newFileContentB, err := yaml.Marshal(valueMap)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(valueFileName, newFileContentB, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (projectUtil *ProjectUtil) IncludeFile(fileName string, projectFilePath string) (err error) {
	node, err := projectUtil.file.ReadYamlNode(projectFilePath)
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
			return projectUtil.file.WriteYamlNode(projectFilePath, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
		}
	}
	includesKey := &yaml.Node{Kind: yaml.ScalarNode, Value: "includes"}
	includesVal := &yaml.Node{Kind: yaml.SequenceNode, Content: []*yaml.Node{newIncludeVal}}
	docNode.Style = yaml.LiteralStyle
	docNode.Content = append(
		[]*yaml.Node{includesKey, includesVal},
		docNode.Content...,
	)
	return projectUtil.file.WriteYamlNode(projectFilePath, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
}

func (projectUtil *ProjectUtil) AddTaskIfNotExist(taskName string, taskFilePath string) (err error) {
	node, err := projectUtil.file.ReadYamlNode(taskFilePath)
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
			return projectUtil.file.WriteYamlNode(taskFilePath, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
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
	return projectUtil.file.WriteYamlNode(taskFilePath, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
}

func (projectUtil *ProjectUtil) SyncEnvFiles(projectFile string) (err error) {
	project, err := projectUtil.getProject(projectFile)
	if err != nil {
		return err
	}
	envFileNames, err := project.GetEnvFileNames()
	if err != nil {
		return err
	}
	for _, envFileName := range envFileNames {
		fileEnvMap, err := godotenv.Read(envFileName)
		if err != nil {
			return err
		}
		for _, task := range project.Tasks {
			for _, envKey := range task.GetEnvKeys() {
				envObj, declared := task.GetEnvObject(envKey)
				if !declared {
					continue
				}
				envFrom := envObj.From
				if envFrom == "" {
					continue
				}
				if _, keyExist := fileEnvMap[envFrom]; keyExist {
					continue
				}
				envDefault := envObj.Default
				fileEnvMap[envFrom] = envDefault
			}
		}
		godotenv.Write(fileEnvMap, envFileName)
	}
	return nil
}

func (projectUtil *ProjectUtil) SyncTasksEnv(projectFile string) (err error) {
	project, err := projectUtil.getProject(projectFile)
	if err != nil {
		return err
	}
	for _, task := range project.Tasks {
		taskName := task.GetName()
		err := projectUtil.Task.Env.Sync(taskName, projectFile)
		if err != nil {
			return err
		}
	}
	return nil
}
