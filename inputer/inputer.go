package inputer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/monitor"
)

type Action struct {
	Explain        bool
	RunInteractive bool
	Run            bool
}

type Prompter struct {
	logger  monitor.Logger
	d       *monitor.Decoration
	project *config.Project
}

func NewPrompter(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project) *Prompter {
	return &Prompter{
		logger:  logger,
		d:       decoration,
		project: project,
	}
}

func (prompter *Prompter) GetAction(taskName string) (action *Action, err error) {
	options := []string{"Run Interactively", "Just Run", "Explain"}
	prompt := promptui.Select{
		Label:             fmt.Sprintf("%s What do you want to do with %s?", prompter.d.Skull, taskName),
		Items:             options,
		Size:              3,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			option := options[index]
			return strings.Contains(strings.ToLower(option), strings.ToLower(input))
		},
	}
	action = &Action{Run: false, RunInteractive: false, Explain: false}
	_, chosenItem, err := prompt.Run()
	if err != nil {
		return action, err
	}
	switch chosenItem {
	case options[0]:
		action.RunInteractive = true
	case options[1]:
		action.Run = true
	case options[2]:
		action.Explain = true
	}
	return action, nil
}

func (prompter *Prompter) GetTaskName() (taskName string, err error) {
	sortedTaskNames := prompter.project.GetSortedTaskNames()
	publishedTaskNames := []string{}
	for _, taskName := range sortedTaskNames {
		if task := prompter.project.Tasks[taskName]; !task.Private {
			publishedTaskNames = append(publishedTaskNames, taskName)
		}
	}
	prompt := promptui.Select{
		Label:             fmt.Sprintf("%s Please select task", prompter.d.Skull),
		Items:             publishedTaskNames,
		Size:              10,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			taskName := publishedTaskNames[index]
			return strings.Contains(strings.ToLower(taskName), strings.ToLower(input))
		},
	}
	_, taskName, err = prompt.Run()
	return taskName, err
}

func (prompter *Prompter) SetProjectValuesByTask(taskNames []string) (err error) {
	if prompter.project.IsInitialized {
		return fmt.Errorf("Project is not initialized")
	}
	inputs, inputOrder, err := prompter.project.GetInputs(taskNames)
	if err != nil {
		return err
	}
	inputCount := len(inputOrder)
	for index, inputName := range inputOrder {
		input := inputs[inputName]
		inputPrompt := inputName
		if input.Prompt != "" {
			inputPrompt = input.Prompt
		}
		label := fmt.Sprintf("%s %d of %d) %s", prompter.d.Skull, index+1, inputCount, inputPrompt)
		oldValue := prompter.project.GetValue(inputName)
		newValue := ""
		if input.Secret {
			newValue, err = prompter.askPassword(label)
		} else {
			newValue, err = prompter.askInput(label, input, oldValue)
		}
		if err != nil {
			return err
		}
		prompter.project.SetValue(inputName, newValue)
	}
	return nil
}

func (prompter *Prompter) askPassword(label string) (value string, err error) {
	prompt := promptui.Prompt{
		Label: label,
		Mask:  '*',
	}
	return prompt.Run()
}

func (prompter *Prompter) askInput(label string, input *config.Input, oldValue string) (value string, err error) {
	alternatives := []string{oldValue}
	if input.DefaultValue != oldValue {
		alternatives = append(alternatives, input.DefaultValue)
	}
	for _, option := range input.Options {
		if option != oldValue && option != input.DefaultValue {
			alternatives = append(alternatives, option)
		}
	}
	allowCustom := !boolean.IsFalse(input.AllowCustom)
	if allowCustom {
		alternatives = append(alternatives, fmt.Sprintf("%sLet me type by myself%s", prompter.d.Green, prompter.d.Normal))
	}
	selectPrompt := promptui.Select{
		Label:             label,
		Items:             alternatives,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(userInput string, index int) bool {
			if allowCustom && index == len(alternatives)-1 {
				return true
			}
			option := alternatives[index]
			return strings.Contains(strings.ToLower(option), strings.ToLower(userInput))
		},
	}
	_, value, err = selectPrompt.Run()
	if allowCustom {
		if value == alternatives[len(alternatives)-1] {
			prompt := promptui.Prompt{
				Label: label,
				Validate: func(userInput string) error {
					if input.Validation != "" {
						matched, err := regexp.Match(input.Validation, []byte(userInput))
						if err != nil {
							return err
						}
						if !matched {
							return fmt.Errorf("%s does not match %s", userInput, input.Validation)
						}
					}
					return nil
				},
			}
			value, err = prompt.Run()
		}
	}
	return value, err
}
