package input

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/pathutil"
)

func fileMustExist(filePath string) (err error) {
	absFilePath := filePath
	if !filepath.IsAbs(absFilePath) {
		absFilePath, err = filepath.Abs(absFilePath)
		if err != nil {
			return nil
		}
	}
	if _, err = os.Stat(absFilePath); err != nil {
		return err
	}
	return nil
}

type Action struct {
	Explain        bool
	RunInteractive bool
	Run            bool
}

type Prompter struct {
	logger  output.Logger
	d       *output.Decoration
	project *dsl.Project
	util    *dsl.DSLUtil
}

func NewPrompter(logger output.Logger, decoration *output.Decoration, project *dsl.Project) *Prompter {
	return &Prompter{
		logger:  logger,
		d:       decoration,
		project: project,
		util:    project.Util,
	}
}

func (prompter *Prompter) GetAdditionalValue() (err error) {
	return prompter.getAdditionalValue("Do you want to load additional value file?")
}

func (prompter *Prompter) getAdditionalValue(label string) (err error) {
	captions := []string{"🏁 No", "📝 Yes"}
	options := []string{"no", "file"}
	prompter.logger.Println(fmt.Sprintf("%s Load additional value file", prompter.d.ZarubaIcon))
	selectPrompt := promptui.Select{
		Label:             label,
		Items:             captions,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(userInput string, index int) bool {
			option := captions[index]
			return strings.Contains(strings.ToLower(option), strings.ToLower(userInput))
		},
	}
	optionIndex, _, err := selectPrompt.Run()
	if err != nil {
		return err
	}
	selectedOption := options[optionIndex]
	switch selectedOption {
	case "no":
		return nil
	case "file":
		if err = prompter.getAdditionalFileValue(); err != nil {
			return err
		}
	}
	return prompter.getAdditionalValue("Do you want to load another value file?")
}

func (prompter *Prompter) getAdditionalFileValue() (err error) {
	valueFileList, err := prompter.getValueFileList()
	if err != nil {
		return err
	}
	if len(valueFileList) > 0 {
		// input by options
		captions := append(valueFileList, fmt.Sprintf("%sLet me type it!%s", prompter.d.Green, prompter.d.Normal))
		options := append(valueFileList, "")
		prompter.logger.Println(fmt.Sprintf("%s Value file", prompter.d.ZarubaIcon))
		selectPrompt := promptui.Select{
			Label:             fmt.Sprintf("%s Value file", prompter.d.ZarubaIcon),
			Items:             captions,
			Stdout:            &bellSkipper{},
			StartInSearchMode: true,
			Searcher: func(userInput string, index int) bool {
				option := options[index]
				return strings.Contains(strings.ToLower(option), strings.ToLower(userInput))
			},
		}
		optionIndex, _, err := selectPrompt.Run()
		if err != nil {
			return err
		}
		if optionIndex < len(options)-1 {
			value := options[optionIndex]
			return prompter.project.AddValue(value)
		}
	}
	// manual input
	prompter.logger.Println(fmt.Sprintf("%s Value file", prompter.d.ZarubaIcon))
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("%s Value file", prompter.d.ZarubaIcon),
		Validate: fileMustExist,
	}
	value, err := prompt.Run()
	if err != nil {
		return err
	}
	return prompter.project.AddValue(value)
}

func (prompter *Prompter) getValueFileList() (valueFileList []string, err error) {
	dir, err := os.Open(".")
	if err != nil {
		return valueFileList, err
	}
	defer dir.Close()
	fileList, err := dir.Readdirnames(0)
	if err != nil {
		return valueFileList, err
	}
	valueFileList = []string{}
	for _, fileName := range fileList {
		if strings.HasSuffix(fileName, ".values.yaml") && fileName != "default.values.yaml" && fileName != ".previous.values.yaml" {
			valueFileList = append(valueFileList, fileName)
		}
	}
	return valueFileList, err
}

func (prompter *Prompter) GetAdditionalEnv(taskNames []string) (err error) {
	return prompter.getAdditionalEnv("Do you want to load additional env?", taskNames)
}

func (prompter *Prompter) getAdditionalEnv(label string, taskNames []string) (err error) {
	captions := []string{"🏁 No", "📝 Yes, from file", "📝 Yes, manually"}
	options := []string{"no", "file", "manual"}
	prompter.logger.Println(fmt.Sprintf("%s Load additional env", prompter.d.ZarubaIcon))
	selectPrompt := promptui.Select{
		Label:             label,
		Items:             captions,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(userInput string, index int) bool {
			option := captions[index]
			return strings.Contains(strings.ToLower(option), strings.ToLower(userInput))
		},
	}
	optionIndex, _, err := selectPrompt.Run()
	if err != nil {
		return err
	}
	selectedOption := options[optionIndex]
	switch selectedOption {
	case "no":
		return nil
	case "file":
		if err = prompter.getAdditionalFileEnv(taskNames); err != nil {
			return err
		}
	case "manual":
		if err = prompter.getAdditionalManualEnv(taskNames); err != nil {
			return err
		}
	}
	return prompter.getAdditionalEnv("Do you want to load another env?", taskNames)
}

func (prompter *Prompter) getAdditionalFileEnv(taskNames []string) (err error) {
	envFileList, err := pathutil.PathGetEnvFileList(".")
	if err != nil {
		return err
	}
	if len(envFileList) > 0 {
		// input by options
		captions := append(envFileList, fmt.Sprintf("%sLet me type it!%s", prompter.d.Green, prompter.d.Normal))
		options := append(envFileList, "")
		prompter.logger.Println(fmt.Sprintf("%s Environment file", prompter.d.ZarubaIcon))
		selectPrompt := promptui.Select{
			Label:             fmt.Sprintf("%s Environment file", prompter.d.ZarubaIcon),
			Items:             captions,
			Stdout:            &bellSkipper{},
			StartInSearchMode: true,
			Searcher: func(userInput string, index int) bool {
				option := options[index]
				return strings.Contains(strings.ToLower(option), strings.ToLower(userInput))
			},
		}
		optionIndex, _, err := selectPrompt.Run()
		if err != nil {
			return err
		}
		if optionIndex < len(options)-1 {
			value := options[optionIndex]
			return prompter.project.AddEnv(value)
		}
	}
	// manual input
	prompter.logger.Println(fmt.Sprintf("%s Environment file", prompter.d.ZarubaIcon))
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("%s Environment file", prompter.d.ZarubaIcon),
		Validate: fileMustExist,
	}
	value, err := prompt.Run()
	if err != nil {
		return err
	}
	return prompter.project.AddEnv(value)
}

func (prompter *Prompter) getAdditionalManualEnv(taskNames []string) (err error) {
	envMap, err := prompter.getEnvMap(taskNames)
	if err != nil {
		return err
	}
	options := []string{}
	captions := []string{}
	for envName := range envMap {
		options = append(options, envName)
		captions = append(captions, fmt.Sprintf("%s (current value: %s)", envName, envMap[envName]))
	}
	prompter.logger.Println(fmt.Sprintf("%s Environment variable", prompter.d.ZarubaIcon))
	selectPrompt := promptui.Select{
		Label:             fmt.Sprintf("%s Environment variable name", prompter.d.ZarubaIcon),
		Items:             captions,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(userInput string, index int) bool {
			option := options[index]
			return strings.Contains(strings.ToLower(option), strings.ToLower(userInput))
		},
	}
	selectedIndex, _, err := selectPrompt.Run()
	if err != nil {
		return err
	}
	prompter.logger.Println(fmt.Sprintf("%s %s", prompter.d.ZarubaIcon, options[selectedIndex]))
	prompt := promptui.Prompt{
		Label: fmt.Sprintf("New value for %s", captions[selectedIndex]),
	}
	value, err := prompt.Run()
	if err != nil {
		return err
	}
	selectedOption := options[selectedIndex]
	return prompter.project.AddEnv(fmt.Sprintf("%s=%s", selectedOption, value))
}

func (prompter *Prompter) getEnvMap(taskNames []string) (envMap map[string]string, err error) {
	envMap = map[string]string{}
	for _, taskName := range taskNames {
		task := prompter.project.Tasks[taskName]
		envKeys := task.GetEnvKeys()
		for _, envKey := range envKeys {
			env, declared := task.GetEnvObject(envKey)
			if !declared {
				continue
			}
			from := env.From
			value, err := task.GetEnv(envKey)
			if err != nil {
				return envMap, err
			}
			envMap[from] = value
		}
	}
	return envMap, nil
}

func (prompter *Prompter) GetAutoTerminate(taskNames []string) (autoTerminate bool, err error) {
	if prompter.project.GetAutoTerminate(taskNames) {
		return true, nil
	}
	captions := []string{"🏁 No", "🔪 Yes"}
	options := []string{"no", "yes"}
	prompter.logger.Println(fmt.Sprintf("%s Auto terminate", prompter.d.ZarubaIcon))
	selectPrompt := promptui.Select{
		Label:             fmt.Sprintf("%s Do you want to terminate tasks once completed?", prompter.d.ZarubaIcon),
		Items:             captions,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(userInput string, index int) bool {
			option := options[index]
			return strings.Contains(strings.ToLower(option), strings.ToLower(userInput))
		},
	}
	optionIndex, _, err := selectPrompt.Run()
	selectedOption := options[optionIndex]
	if err == nil {
		return prompter.util.Bool.IsTrue(selectedOption), nil
	}
	return false, err
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
	prompter.logger.Println(fmt.Sprintf("%s Action", prompter.d.ZarubaIcon))
	selectPrompt := promptui.Select{
		Label:             fmt.Sprintf("%s What do you want to do with %s?", prompter.d.ZarubaIcon, taskName),
		Items:             options,
		Size:              3,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			option := options[index]
			return strings.Contains(strings.ToLower(option), strings.ToLower(input))
		},
	}
	_, chosenItem, err := selectPrompt.Run()
	if err != nil {
		return nil, err
	}
	return action_map[chosenItem], nil
}

func (prompter *Prompter) GetTaskName() (taskName string, err error) {
	sortedTaskNames := prompter.project.GetSortedTaskNames()
	publicTaskOptions := []string{}
	privateTaskOptions := []string{}
	publicTaskCaptions := []string{}
	privateTaskCaptions := []string{}
	for _, taskName := range sortedTaskNames {
		task := prompter.project.Tasks[taskName]
		taskCaption := fmt.Sprintf("%s %s", task.GetIcon(), taskName)
		if task.Private {
			privateTaskOptions = append(privateTaskOptions, taskName)
			privateTaskCaptions = append(privateTaskCaptions, taskCaption)
			continue
		}
		publicTaskOptions = append(publicTaskOptions, taskName)
		publicTaskCaptions = append(publicTaskCaptions, taskCaption)
	}
	options := append(publicTaskOptions, privateTaskOptions...)
	captions := append(publicTaskCaptions, privateTaskCaptions...)
	prompter.logger.Println(fmt.Sprintf("%s Task Name", prompter.d.ZarubaIcon))
	selectPrompt := promptui.Select{
		Label:             fmt.Sprintf("%s Please select task", prompter.d.ZarubaIcon),
		Items:             captions,
		Size:              10,
		Stdout:            &bellSkipper{},
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			taskName := options[index]
			return strings.Contains(strings.ToLower(taskName), strings.ToLower(input))
		},
	}
	optionIndex, _, err := selectPrompt.Run()
	return options[optionIndex], err
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
		oldValue := prompter.project.GetValue(inputName)
		newValue := ""
		prompter.logger.Println(fmt.Sprintf("%s %d of %d) %s", prompter.d.ZarubaIcon, index+1, inputCount, inputName))
		if input.Description != "" {
			indentation := prompter.d.EmptyIcon + " "
			prompter.logger.Println(indentation + prompter.util.Str.Indent(input.Description, indentation))
		}
		if input.Secret {
			newValue, err = prompter.askPassword(inputPrompt)
		} else {
			newValue, err = prompter.askInput(inputPrompt, input, oldValue)
		}
		if err != nil {
			return err
		}
		if err = prompter.project.SetValue(inputName, newValue); err != nil {
			return err
		}
	}
	return nil
}

func (prompter *Prompter) askPassword(inputPrompt string) (value string, err error) {
	prompt := promptui.Prompt{
		Label: inputPrompt,
		Mask:  '*',
	}
	return prompt.Run()
}

func (prompter *Prompter) askInput(inputPrompt string, input *dsl.Variable, oldValue string) (value string, err error) {
	options, captions := prompter.getInputOptions(input, oldValue)
	allowCustom := !prompter.util.Bool.IsFalse(input.AllowCustom)
	if allowCustom {
		// Directly ask user input in case of no available option
		if len(options) == 0 {
			return prompter.askUserInput(inputPrompt, input)
		}
		options = append(options, "")
		captions = append(captions, fmt.Sprintf("%sLet me type it!%s", prompter.d.Green, prompter.d.Normal))
	}
	selectPrompt := promptui.Select{
		Label:             inputPrompt,
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
	if err == nil && allowCustom && selectedIndex == len(options)-1 {
		return prompter.askUserInput(inputPrompt, input)
	}
	return options[selectedIndex], err
}

func (prompter *Prompter) askUserInput(inputPrompt string, input *dsl.Variable) (value string, err error) {
	prompt := promptui.Prompt{
		Label: inputPrompt,
		Validate: func(userInput string) error {
			return input.Validate(os.ExpandEnv(userInput))
		},
	}
	return prompt.Run()
}

func (prompter *Prompter) getInputOptions(input *dsl.Variable, oldValue string) (options []string, captions []string) {
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
