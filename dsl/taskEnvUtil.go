package dsl

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/dictutil"
	"github.com/state-alchemists/zaruba/pathutil"
	"github.com/state-alchemists/zaruba/yamlstyler"
	yaml "gopkg.in/yaml.v3"
)

type TaskEnvUtil struct {
	task *TaskUtil
}

func NewTaskEnvUtil(taskUtil *TaskUtil) *TaskEnvUtil {
	return &TaskEnvUtil{
		task: taskUtil,
	}
}

func (envUtil *TaskEnvUtil) Sync(taskName, projectFile string) (err error) {
	task, err := envUtil.task.getTaskByProjectFile(projectFile, taskName)
	if err != nil {
		return err
	}
	if !task.GetShouldSyncEnv() {
		return nil
	}
	projectDir := filepath.Dir(task.Project.GetFileLocation())
	taskFileLocation := task.GetFileLocation()
	if !strings.HasPrefix(taskFileLocation, projectDir) {
		return nil
	}
	syncEnvLocation := task.GetSyncEnvLocation()
	if syncEnvLocation == "" || syncEnvLocation == projectDir {
		return nil
	}
	locationEnvMap, err := pathutil.PathGetEnvByLocation(syncEnvLocation)
	if err != nil {
		return err
	}
	envRefName := task.GetFirstEnvRefName()
	if envRefName == "" {
		// update taskEnv
		newEnvMap := envUtil.getAdditionalEnvMap(task.Envs, locationEnvMap)
		if len(newEnvMap) == 0 {
			return nil
		}
		if err = envUtil.set(task, newEnvMap); err != nil {
			return err
		}
		return nil
	}
	// update envRef
	newEnvMap := envUtil.getAdditionalEnvMap(task.Project.EnvRefMap[envRefName].Map, locationEnvMap)
	if len(newEnvMap) == 0 {
		return nil
	}
	return envUtil.setEnvsRef(task.Project.EnvRefMap[envRefName], newEnvMap)
}

func (envUtil *TaskEnvUtil) Set(taskName string, envMap map[string]string, projectFile string) (err error) {
	if len(envMap) == 0 {
		return nil
	}
	task, err := envUtil.task.getTaskByProjectFile(projectFile, taskName)
	if err != nil {
		return err
	}
	// adjust envFiles
	if err = envUtil.SyncEnvFiles(task, envMap, projectFile); err != nil {
		return err
	}
	envRefName := task.GetFirstEnvRefName()
	if envRefName == "" {
		// update taskEnv
		return envUtil.set(task, envMap)
	}
	// update envRef
	return envUtil.setEnvsRef(task.Project.EnvRefMap[envRefName], envMap)
}

func (envUtil *TaskEnvUtil) SyncEnvFiles(task *Task, envMap map[string]string, projectFile string) (err error) {
	envFileNames, err := task.Project.GetEnvFileNames()
	if err != nil {
		return err
	}
	for _, envFileName := range envFileNames {
		fileEnvMap, err := godotenv.Read(envFileName)
		if err != nil {
			return err
		}
		for envKey := range envMap {
			newEnvValue := envMap[envKey]
			envObj, declared := task.GetEnvObject(envKey)
			if !declared {
				continue
			}
			envFrom := envObj.From
			if envFrom == "" {
				continue
			}
			if _, keyExist := fileEnvMap[envFrom]; !keyExist {
				continue
			}
			envDefault := envObj.Default
			fileEnvValue, exist := fileEnvMap[envFrom]
			if exist && fileEnvValue == envDefault {
				fileEnvMap[envFrom] = newEnvValue
			}
		}
		godotenv.Write(fileEnvMap, envFileName)
	}
	return nil
}

func (envUtil *TaskEnvUtil) set(task *Task, envMap map[string]string) (err error) {
	taskName := task.GetName()
	envPrefix := strings.ToUpper(task.Project.Util.Str.ToSnake(taskName))
	yamlLocation := task.GetFileLocation()
	node, err := envUtil.task.file.ReadYamlNode(yamlLocation)
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
							envUtil.updateEnvMapNode(taskPropValNode, envMap, envPrefix)
							return envUtil.task.file.WriteYamlNode(yamlLocation, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
						}
					}
					// env not found
					taskNode.Style = yaml.LiteralStyle
					taskNode.Content = append(
						taskNode.Content,
						&yaml.Node{Kind: yaml.ScalarNode, Value: "envs"},
						envUtil.createEnvMapNode(envMap, envPrefix),
					)
					return envUtil.task.file.WriteYamlNode(yamlLocation, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
				}
			}
		}
	}
	return fmt.Errorf("cannot find task %s in %s", taskName, yamlLocation)
}

func (envUtil *TaskEnvUtil) setEnvsRef(envRef *EnvRef, envMap map[string]string) (err error) {
	util := NewDSLUtil()
	envRefName := envRef.GetName()
	envPrefix := strings.ToUpper(util.Str.ToSnake(envRefName))
	yamlLocation := envRef.GetFileLocation()
	node, err := envUtil.task.file.ReadYamlNode(yamlLocation)
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
					envUtil.updateEnvMapNode(envRefNode, envMap, envPrefix)
					return envUtil.task.file.WriteYamlNode(yamlLocation, node, 0555, []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji, yamlstyler.AddLineBreak})
				}
			}
		}
	}
	return fmt.Errorf("cannot find envRef %s in %s", envRefName, yamlLocation)
}

func (envUtil *TaskEnvUtil) updateEnvMapNode(envMapNode *yaml.Node, envMap map[string]string, envPrefix string) {
	envMapNode.Style = yaml.LiteralStyle
	envKeys, _ := dictutil.DictGetSortedKeys(envMap)
	for _, envKey := range envKeys {
		envVal := envMap[envKey]
		envKeyFound := false
		for envKeyIndex := 0; envKeyIndex < len(envMapNode.Content); envKeyIndex += 2 {
			envKeyNode := envMapNode.Content[envKeyIndex]
			envValNode := envMapNode.Content[envKeyIndex+1]
			// "envs" and envKey found, update
			if envKeyNode.Value == envKey {
				envFromFound, envDefaultFound := false, false
				envFrom := envUtil.getEnvFromName(envKey, envPrefix)
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
					envValNode.Content = append(envValNode.Content, envUtil.createEnvFromNode(envKey, envPrefix)...)
				}
				if !envDefaultFound {
					envValNode.Content = append(envValNode.Content, envUtil.createEnvDefaultNode(envVal)...)
				}
				envKeyFound = true
				break
			}
		}
		// "envs" found, but envKey not found, create
		if !envKeyFound {
			envMapNode.Content = append(envMapNode.Content, envUtil.createEnvNode(envKey, envVal, envPrefix)...)
		}
	}
}

func (envUtil *TaskEnvUtil) createEnvMapNode(envMap map[string]string, envPrefix string) *yaml.Node {
	newEnvNodes := []*yaml.Node{}
	envKeys, _ := dictutil.DictGetSortedKeys(envMap)
	for _, envKey := range envKeys {
		envVal := envMap[envKey]
		newEnvNodes = append(
			newEnvNodes,
			envUtil.createEnvNode(envKey, envVal, envPrefix)...,
		)
	}
	return &yaml.Node{Kind: yaml.MappingNode, Content: newEnvNodes}
}

func (envUtil *TaskEnvUtil) createEnvNode(envKey, envVal, envPrefix string) []*yaml.Node {
	return []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: envKey},
		{
			Kind: yaml.MappingNode,
			Content: append(
				envUtil.createEnvFromNode(envKey, envPrefix),
				envUtil.createEnvDefaultNode(envVal)...,
			),
		},
	}
}

func (envUtil *TaskEnvUtil) createEnvFromNode(envKey, envPrefix string) []*yaml.Node {
	envFrom := envUtil.getEnvFromName(envKey, envPrefix)
	return []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "from"},
		{Kind: yaml.ScalarNode, Value: envFrom},
	}
}

func (envUtil *TaskEnvUtil) createEnvDefaultNode(envVal string) []*yaml.Node {
	return []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "default"},
		{Kind: yaml.ScalarNode, Value: envVal},
	}
}

func (envUtil *TaskEnvUtil) getEnvFromName(envKey, envPrefix string) string {
	if !strings.HasPrefix(envKey, envPrefix) {
		return fmt.Sprintf("%s_%s", envPrefix, envKey)
	}
	return envKey
}

func (envUtil *TaskEnvUtil) getAdditionalEnvMap(existingEnvMap map[string]*Env, locationEnvMap map[string]string) (additionalEnvMap map[string]string) {
	additionalEnvMap = map[string]string{}
	for envKey, envVal := range locationEnvMap {
		if _, exist := existingEnvMap[envKey]; exist {
			continue
		}
		additionalEnvMap[envKey] = envVal
	}
	return additionalEnvMap
}
