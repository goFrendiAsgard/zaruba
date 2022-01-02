package explainer

import (
	"fmt"
	"sort"
	"strings"

	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/dictutil"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/strutil"
)

type TaskExplanationVariable struct {
	DefaultValue string   `yaml:"Default Value"`
	Description  string   `yaml:"Description"`
	Prompt       string   `yaml:"Prompt"`
	Secret       bool     `yaml:"Secret"`
	Validation   string   `yaml:"Validation"`
	Options      []string `yaml:"Options"`
}

type TaskExplanationEnv struct {
	From    string
	Default string
}

type TaskExplanation struct {
	d               *output.Decoration
	Icon            string                             `yaml:"Icon"`
	Name            string                             `yaml:"Name"`
	FileLocation    string                             `yaml:"File Location"`
	Location        string                             `yaml:"Location"`
	ShouldSyncEnv   bool                               `yaml:"Should Sync Env"`
	SyncEnvLocation string                             `yaml:"Sync Env Location"`
	Type            string                             `yaml:"Type"`
	Description     string                             `yaml:"Description"`
	Extends         []string                           `yaml:"Extends"`
	Dependencies    []string                           `yaml:"Dependencies"`
	Start           []string                           `yaml:"Start"`
	Check           []string                           `yaml:"Check"`
	Inputs          map[string]TaskExplanationVariable `yaml:"Inputs"`
	Configs         map[string]string                  `yaml:"Configs"`
	Envs            map[string]TaskExplanationEnv      `yaml:"Envs"`
}

func NewTaskExplanation(decoration *output.Decoration, task *core.Task) (taskExplanation *TaskExplanation) {
	// start, check,
	startPattern, startExist, _ := task.GetStartCmdPatterns()
	checkPattern, checkExist, _ := task.GetCheckCmdPatterns()
	taskType := "wrapper"
	if startExist && checkExist {
		taskType = "service"
	} else if startExist {
		taskType = "command"
	}
	// inputs
	inputNames := task.Inputs
	taskInputs := map[string]TaskExplanationVariable{}
	for _, inputName := range inputNames {
		variable := task.Project.Inputs[inputName]
		taskInputs[inputName] = TaskExplanationVariable{
			DefaultValue: variable.DefaultValue,
			Description:  variable.Description,
			Prompt:       variable.Prompt,
			Secret:       variable.Secret,
			Validation:   variable.Validation,
			Options:      variable.Options,
		}
	}
	// configs
	configKeys := task.GetConfigKeys()
	sort.Strings(configKeys)
	taskConfigs := map[string]string{}
	for _, configKey := range configKeys {
		configVal, _ := task.GetConfigPattern(configKey)
		taskConfigs[configKey] = configVal
	}
	// envs
	envKeys := task.GetEnvKeys()
	sort.Strings(envKeys)
	taskEnvs := map[string]TaskExplanationEnv{}
	for _, envKey := range envKeys {
		env, _ := task.GetEnvObject(envKey)
		taskEnvs[envKey] = TaskExplanationEnv{
			From:    env.From,
			Default: env.Default,
		}
	}
	// task explanation
	return &TaskExplanation{
		d:               decoration,
		Icon:            task.Icon,
		Name:            task.GetName(),
		FileLocation:    task.GetFileLocation(),
		Location:        task.GetLocation(),
		ShouldSyncEnv:   task.ShouldSyncEnv(),
		SyncEnvLocation: task.GetSyncEnvLocation(),
		Type:            taskType,
		Description:     task.Description,
		Extends:         task.GetParentTaskNames(),
		Dependencies:    task.GetDependencies(),
		Start:           startPattern,
		Check:           checkPattern,
		Inputs:          taskInputs,
		Configs:         taskConfigs,
		Envs:            taskEnvs,
	}
}

func (t *TaskExplanation) ToString() string {
	strUtil := strutil.NewStrUtil()
	lines := []string{}
	lines = append(lines, t.h1(fmt.Sprintf("%s %s", t.Icon, strUtil.ToPascal(t.Name))))
	lines = append(lines, t.prop("File Location", t.FileLocation))
	if t.Location != "" {
		lines = append(lines, t.prop("Location", t.Location))
	}
	lines = append(lines, t.prop("Should Sync Env", fmt.Sprintf("%t", t.ShouldSyncEnv)))
	if t.SyncEnvLocation != "" {
		lines = append(lines, t.prop("Sync Env Location", t.SyncEnvLocation))
	}
	lines = append(lines, t.prop("Type", t.Type))
	if t.Description != "" {
		lines = append(lines, t.prop("Description", t.Description))
	}
	if len(t.Extends) > 0 {
		lines = append(lines, t.h2("Extends"))
		lines = append(lines, t.ul(t.Extends))
	}
	if len(t.Dependencies) > 0 {
		lines = append(lines, t.h2("Dependencies"))
		lines = append(lines, t.ul(t.Dependencies))
	}
	if len(t.Start) > 0 {
		lines = append(lines, t.h2("Start"))
		lines = append(lines, t.ul(t.Start))
	}
	if len(t.Check) > 0 {
		lines = append(lines, t.h2("Check"))
		lines = append(lines, t.ul(t.Check))
	}
	if len(t.Inputs) > 0 {
		lines = append(lines, t.h2("Inputs"))
		inputNames, _ := dictutil.DictGetSortedKeys(t.Inputs)
		for _, inputName := range inputNames {
			variable := t.Inputs[inputName]
			lines = append(lines, t.h3(fmt.Sprintf("Inputs.%s", inputName)))
			if variable.Description != "" {
				lines = append(lines, t.prop("Description", variable.Description))
			}
			if variable.Prompt != "" {
				lines = append(lines, t.prop("Prompt", variable.Prompt))
			}
			if variable.DefaultValue != "" {
				lines = append(lines, t.prop("Default Value", variable.DefaultValue))
			}
			lines = append(lines, t.prop("Secret", fmt.Sprintf("%t", variable.Secret)))
			if variable.Validation != "" {
				lines = append(lines, t.prop("Validation", variable.Validation))
			}
			if len(variable.Options) > 0 {
				lines = append(lines, t.prop("Options", strings.Join(variable.Options, "; ")))
			}
		}
	}
	if len(t.Configs) > 0 {
		lines = append(lines, t.h2("Configs"))
		configNames, _ := dictutil.DictGetSortedKeys(t.Configs)
		for _, configName := range configNames {
			configValue := t.Configs[configName]
			lines = append(lines, t.h3(fmt.Sprintf("Configs.%s", configName)))
			if configValue != "" {
				lines = append(lines, t.prop("Value", configValue))
			}
		}
	}
	if len(t.Envs) > 0 {
		lines = append(lines, t.h2("Envs"))
		envNames, _ := dictutil.DictGetSortedKeys(t.Envs)
		for _, envName := range envNames {
			env := t.Envs[envName]
			lines = append(lines, t.h3(fmt.Sprintf("Envs.%s", envName)))
			if env.From != "" {
				lines = append(lines, t.prop("From", env.From))
			}
			if env.Default != "" {
				lines = append(lines, t.prop("Default", env.Default))
			}
		}
	}
	return strings.Join(lines, "\n")
}

func (t *TaskExplanation) ul(items []string) string {
	lines := []string{}
	for _, item := range items {
		itemLines := strings.Split(item, "\n")
		if len(itemLines) < 2 {
			lines = append(lines, fmt.Sprintf("* `%s`", item))
			continue
		}
		lines = append(lines, "*")
		lines = append(lines, strutil.StrFullIndent(fmt.Sprintf("```\n%s\n```", item), "    "))
	}
	return strings.Join(lines, "\n") + "\n"
}

func (t *TaskExplanation) h1(header string) string {
	return fmt.Sprintf("\n%s%s# %s%s\n", t.d.Bold, t.d.Yellow, header, t.d.Normal)
}

func (t *TaskExplanation) h2(header string) string {
	return fmt.Sprintf("\n%s%s## %s%s\n", t.d.Bold, t.d.Yellow, header, t.d.Normal)
}

func (t *TaskExplanation) h3(header string) string {
	return fmt.Sprintf("\n%s%s### %s%s\n", t.d.Bold, t.d.Yellow, header, t.d.Normal)
}

func (t *TaskExplanation) prop(propertyName, value string) string {
	caption := fmt.Sprintf("%s%s%s%s:", t.d.Bold, t.d.Blue, propertyName, t.d.Normal)
	if value == "" {
		return fmt.Sprintf("%s\n", caption)
	}
	multiLineValue := strutil.StrFullIndent(fmt.Sprintf("\n%s", value), "    ")
	return fmt.Sprintf("%s\n%s\n", caption, multiLineValue)
}
