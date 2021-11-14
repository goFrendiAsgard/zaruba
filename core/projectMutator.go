package core

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/fileutil"
	"github.com/state-alchemists/zaruba/yamlstyler"
	yaml "gopkg.in/yaml.v3"
)

func SetProjectValue(valueFilePath, key, value string) (err error) {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if value == "" {
		return fmt.Errorf("value cannot be empty")
	}
	fileContentB, err := ioutil.ReadFile(valueFilePath)
	if err != nil {
		return err
	}
	configMap := map[string]string{}
	if err := yaml.Unmarshal(fileContentB, &configMap); err != nil {
		return err
	}
	configMap[key] = value
	newFileContentB, err := yaml.Marshal(configMap)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(valueFilePath, newFileContentB, 0755)
}

func IncludeFileToProject(projectFilePath string, fileName string) (err error) {
	fileUtil := fileutil.NewFileUtil()
	node, err := fileUtil.ReadYamlNode(projectFilePath)
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
			return fileUtil.WriteYamlNode(projectFilePath, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
		}
	}
	includesKey := &yaml.Node{Kind: yaml.ScalarNode, Value: "includes"}
	includesVal := &yaml.Node{Kind: yaml.SequenceNode, Content: []*yaml.Node{newIncludeVal}}
	docNode.Style = yaml.LiteralStyle
	docNode.Content = append(
		[]*yaml.Node{includesKey, includesVal},
		docNode.Content...,
	)
	return fileUtil.WriteYamlNode(projectFilePath, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
}

func AddTaskIfNotExist(taskFilePath string, taskName string) (err error) {
	fileUtil := fileutil.NewFileUtil()
	node, err := fileUtil.ReadYamlNode(taskFilePath)
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
			return fileUtil.WriteYamlNode(taskFilePath, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
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
	return fileUtil.WriteYamlNode(taskFilePath, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
}

func SyncProjectEnvFiles(project *Project) (err error) {
	projectDir := filepath.Dir(project.GetFileLocation())
	files, err := ioutil.ReadDir(projectDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		isDir := file.IsDir()
		if isDir {
			continue
		}
		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".env") && !strings.HasSuffix(fileName, ".env.template") {
			continue
		}
		fileEnvMap, err := godotenv.Read(filepath.Join(projectDir, fileName))
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
		godotenv.Write(fileEnvMap, filepath.Join(projectDir, fileName))
	}
	return nil
}

func SyncProjectEnv(project *Project) (err error) {
	for _, task := range project.Tasks {
		if err := SyncTaskEnv(task); err != nil {
			return err
		}
	}
	return nil
}
