package input

import (
	"fmt"
	"os"
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
	caption_interactive := "🏁 Run interactively"
	caption_run := "🏁 Run"
	caption_explain := "🔍 Explain"
	action_map := map[string]*Action{
		caption_interactive: {Run: false, RunInteractive: true, Explain: false},
		caption_run:         {Run: true, RunInteractive: false, Explain: false},
		caption_explain:     {Run: false, RunInteractive: false, Explain: true},
	}
	task := prompter.project.Tasks[taskName]
	options := []string{}
	if task.Private {
		options = []string{caption_explain}
	} else if len(prompter.project.Tasks[taskName].Inputs) > 0 {
		options = []string{caption_interactive, caption_run, caption_explain}
	} else {
		options = []string{caption_run, caption_explain}
	}
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
	_, chosenItem, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	return action_map[chosenItem], nil
}

func (prompter *Prompter) GetTaskName() (taskName string, err error) {
	sortedTaskNames := prompter.project.GetSortedTaskNames()
	publicTasks := []string{}
	privateTasks := []string{}
	for _, taskName := range sortedTaskNames {
		task := prompter.project.Tasks[taskName]
		if task.Private {
			privateTasks = append(privateTasks, taskName)
			continue
		}
		publicTasks = append(publicTasks, taskName)
	}
	options := append(publicTasks, privateTasks...)
	prompt := promptui.Select{
		Label:             fmt.Sprintf("%s Please select task", prompter.d.Skull),
		Items:             options,
		Size:              10,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			taskName := options[index]
			return strings.Contains(strings.ToLower(taskName), strings.ToLower(input))
		},
	}
	_, taskName, err = prompt.Run()
	return taskName, err
}

func (prompter *Prompter) SetProjectValuesByTask(taskNames []string) (err error) {
	if prompter.project.IsInitialized {
		return fmt.Errorf("project is not initialized")
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

func (prompter *Prompter) askInput(label string, input *config.Variable, oldValue string) (value string, err error) {
	options, captions := prompter.getInputOptions(input, oldValue)
	allowCustom := !boolean.IsFalse(input.AllowCustom)
	if allowCustom {
		// Directly ask user input in case ofno available option
		if len(options) == 0 {
			return prompter.askUserInput(label, input)
		}
		options = append(options, "")
		captions = append(captions, fmt.Sprintf("%sLet me type it!%s", prompter.d.Green, prompter.d.Normal))
	}
	selectPrompt := promptui.Select{
		Label:             label,
		Items:             captions,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(userInput string, index int) bool {
			if allowCustom && index == len(options)-1 {
				return true
			}
			option := options[index]
			return strings.Contains(strings.ToLower(option), strings.ToLower(userInput))
		},
	}
	selectedIndex, _, err := selectPrompt.Run()
	if allowCustom && selectedIndex == len(options)-1 {
		value, err = prompter.askUserInput(label, input)
	} else {
		value = options[selectedIndex]
	}
	return value, err
}

func (prompter *Prompter) askUserInput(label string, input *config.Variable) (value string, err error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(userInput string) error {
			return input.Validate(os.ExpandEnv(userInput))
		},
	}
	return prompt.Run()
}

func (prompter *Prompter) getInputOptions(input *config.Variable, oldValue string) (options []string, captions []string) {
	options = []string{}
	captions = []string{}
	if err := input.Validate(os.ExpandEnv(oldValue)); err == nil {
		options = append(options, oldValue)
		captions = append(captions, prompter.getOptionCaption(oldValue))
	}
	if oldValue != input.DefaultValue {
		if err := input.Validate(os.ExpandEnv(input.DefaultValue)); err == nil {
			options = append(options, input.DefaultValue)
			captions = append(captions, prompter.getOptionCaption(input.DefaultValue))
		}
	}
	for _, option := range input.Options {
		if option == oldValue || option == input.DefaultValue {
			continue
		}
		if err := input.Validate(os.ExpandEnv(option)); err == nil {
			options = append(options, option)
			captions = append(captions, prompter.getOptionCaption(option))
		}
	}
	return options, captions
}

func (prompter *Prompter) getOptionCaption(option string) (caption string) {
	if option == "" {
		return fmt.Sprintf("%sBlank%s", prompter.d.Green, prompter.d.Normal)
	}
	return option
}
