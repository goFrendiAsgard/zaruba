package core

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/state-alchemists/zaruba/core/fileutil"
	"github.com/state-alchemists/zaruba/core/yamlStyler"
	"github.com/state-alchemists/zaruba/env"
	yaml "gopkg.in/yaml.v3"
)

func SyncTaskEnv(task *Task) (err error) {
	if !task.ShouldSyncEnv() {
		return nil
	}
	projectDir := filepath.Dir(task.Project.GetFileLocation())
	taskFileLocation := task.GetFileLocation()
	if !strings.HasPrefix(taskFileLocation, projectDir) {
		return nil
	}
	taskLocation := task.GetSyncEnvLocation()
	if taskLocation == "" || taskLocation == projectDir {
		return nil
	}
	locationEnvMap, err := env.GetEnvByLocation(taskLocation)
	if err != nil {
		return err
	}
	envRefName := task.GetFirstEnvRefName()
	if envRefName == "" {
		// update taskEnv
		newEnvMap := getAdditionalEnvMap(task.Envs, locationEnvMap)
		if len(newEnvMap) == 0 {
			return nil
		}
		if err = setTaskEnv(task, newEnvMap); err != nil {
			return err
		}
		return nil
	}
	// update envRef
	newEnvMap := getAdditionalEnvMap(task.Project.EnvRefMap[envRefName].Map, locationEnvMap)
	if len(newEnvMap) == 0 {
		return nil
	}
	return setEnvRef(task.Project.EnvRefMap[envRefName], newEnvMap)
}

func SetTaskEnv(task *Task, envMap map[string]string) (err error) {
	if len(envMap) == 0 {
		return nil
	}
	envRefName := task.GetFirstEnvRefName()
	if envRefName == "" {
		// update taskEnv
		return setTaskEnv(task, envMap)
	}
	// update envRef
	return setEnvRef(task.Project.EnvRefMap[envRefName], envMap)
}

func setTaskEnv(task *Task, envMap map[string]string) (err error) {
	taskName := task.GetName()
	envPrefix := strings.ToUpper(task.Project.Util.Str.ToSnake(taskName))
	yamlLocation := task.GetFileLocation()
	fileUtil := fileutil.NewFileUtil()
	node, err := fileUtil.ReadYaml(yamlLocation)
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
						if taskPropKeyNode.Value == "envs" && taskPropValNode.ShortTag() == "!!map" {
							updateEnvMapNode(taskPropValNode, envMap, envPrefix)
							return fileUtil.WriteYaml(yamlLocation, node, 0555, []yamlStyler.YamlStyler{yamlStyler.TwoSpaces, yamlStyler.FixEmoji, yamlStyler.AddLineBreak})
						}
					}
					// env not found
					taskNode.Style = yaml.LiteralStyle
					taskNode.Content = append(
						taskNode.Content,
						&yaml.Node{Kind: yaml.ScalarNode, Value: "envs"},
						createEnvMapNode(envMap, envPrefix),
					)
					return fileUtil.WriteYaml(yamlLocation, node, 0555, []yamlStyler.YamlStyler{yamlStyler.TwoSpaces, yamlStyler.FixEmoji, yamlStyler.AddLineBreak})
				}
			}
		}
	}
	return fmt.Errorf("cannot find task %s in %s", taskName, yamlLocation)
}

func setEnvRef(envRef *EnvRef, envMap map[string]string) (err error) {
	util := NewUtil()
	envRefName := envRef.GetName()
	envPrefix := strings.ToUpper(util.Str.ToSnake(envRefName))
	yamlLocation := envRef.GetFileLocation()
	fileUtil := fileutil.NewFileUtil()
	node, err := fileUtil.ReadYaml(yamlLocation)
	if err != nil {
		return err
	}
	docNode := node.Content[0]
	for index := 0; index < len(docNode.Content); index += 2 {
		keyNode := docNode.Content[index]
		valNode := docNode.Content[index+1]
		if keyNode.Value == "envs" && valNode.ShortTag() == "!!map" {
			for envRefNameIndex := 0; envRefNameIndex < len(valNode.Content); envRefNameIndex += 2 {
				envRefNameNode := valNode.Content[envRefNameIndex]
				envRefNode := valNode.Content[envRefNameIndex+1]
				if envRefNameNode.Value == envRefName && envRefNode.ShortTag() == "!!map" {
					updateEnvMapNode(envRefNode, envMap, envPrefix)
					return fileUtil.WriteYaml(yamlLocation, node, 0555, []yamlStyler.YamlStyler{yamlStyler.TwoSpaces, yamlStyler.FixEmoji, yamlStyler.AddLineBreak})
				}
			}
		}
	}
	return fmt.Errorf("cannot find envRef %s in %s", envRefName, yamlLocation)
}

func updateEnvMapNode(envMapNode *yaml.Node, envMap map[string]string, envPrefix string) {
	envMapNode.Style = yaml.LiteralStyle
	envKeys := getSortedEnvMapKeys(envMap)
	for _, envKey := range envKeys {
		envVal := envMap[envKey]
		envKeyFound := false
		for envKeyIndex := 0; envKeyIndex < len(envMapNode.Content); envKeyIndex += 2 {
			envKeyNode := envMapNode.Content[envKeyIndex]
			envValNode := envMapNode.Content[envKeyIndex+1]
			// "envs" and envKey found, update
			if envKeyNode.Value == envKey {
				envFromFound, envDefaultFound := false, false
				envFrom := getEnvFromName(envKey, envPrefix)
				for envPropKeyIndex := 0; envPropKeyIndex < len(envValNode.Content); envPropKeyIndex += 2 {
					envPropKeyNode := envValNode.Content[envPropKeyIndex]
					envPropValNode := envValNode.Content[envPropKeyIndex+1]
					switch envPropKeyNode.Value {
					case "from":
						envPropValNode.SetString(envFrom)
						envFromFound = true
					case "default":
						envPropValNode.SetString(envVal)
						envDefaultFound = true
					}
				}
				if !envFromFound {
					envValNode.Content = append(envValNode.Content, createEnvFromNode(envKey, envPrefix)...)
				}
				if !envDefaultFound {
					envValNode.Content = append(envValNode.Content, createEnvDefaultNode(envVal)...)
				}
				envKeyFound = true
				break
			}
		}
		// "envs" found, but envKey not found, create
		if !envKeyFound {
			envMapNode.Content = append(envMapNode.Content, createEnvNode(envKey, envVal, envPrefix)...)
		}
	}
}

func createEnvMapNode(envMap map[string]string, envPrefix string) *yaml.Node {
	newEnvNodes := []*yaml.Node{}
	envKeys := getSortedEnvMapKeys(envMap)
	for _, envKey := range envKeys {
		envVal := envMap[envKey]
		newEnvNodes = append(
			newEnvNodes,
			createEnvNode(envKey, envVal, envPrefix)...,
		)
	}
	return &yaml.Node{Kind: yaml.MappingNode, Content: newEnvNodes}
}

func getSortedEnvMapKeys(envMap map[string]string) (envKeys []string) {
	envKeys = []string{}
	for envKey := range envMap {
		envKeys = append(envKeys, envKey)
	}
	sort.Strings(envKeys)
	return envKeys
}

func createEnvNode(envKey, envVal, envPrefix string) []*yaml.Node {
	return []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: envKey},
		{
			Kind: yaml.MappingNode,
			Content: append(
				createEnvFromNode(envKey, envPrefix),
				createEnvDefaultNode(envVal)...,
			),
		},
	}
}

func createEnvFromNode(envKey, envPrefix string) []*yaml.Node {
	envFrom := getEnvFromName(envKey, envPrefix)
	return []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "from"},
		{Kind: yaml.ScalarNode, Value: envFrom},
	}
}

func createEnvDefaultNode(envVal string) []*yaml.Node {
	return []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "default"},
		{Kind: yaml.ScalarNode, Value: envVal},
	}
}

func getEnvFromName(envKey, envPrefix string) string {
	if !strings.HasPrefix(envKey, envPrefix) {
		return fmt.Sprintf("%s_%s", envPrefix, envKey)
	}
	return envKey
}

func getAdditionalEnvMap(existingEnvMap map[string]*Env, locationEnvMap map[string]string) (additionalEnvMap map[string]string) {
	additionalEnvMap = map[string]string{}
	for envKey, envVal := range locationEnvMap {
		if _, exist := existingEnvMap[envKey]; exist {
			continue
		}
		additionalEnvMap[envKey] = envVal
	}
	return additionalEnvMap
}
