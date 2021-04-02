package inputer

import (
	"fmt"
	"os"
	"strings"

	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/monitor"
	"golang.org/x/crypto/ssh/terminal"
)

func Ask(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project, taskNames []string) (err error) {
	if project.IsInitialized {
		return fmt.Errorf("Cannot ask for input because project has been initialized")
	}
	inputs, inputOrder, err := project.GetInputs(taskNames)
	if err != nil {
		return err
	}
	inputCount := len(inputOrder)
	if inputCount == 0 {
		logger.DPrintf("%sZaruba is running in interactive mode. But no value you can set interactively.%s\n", decoration.Yellow, decoration.Normal)
		return nil
	}
	logger.DPrintf("%sZaruba is running in interactive mode. You will be able to set some values interactively.%s\n\n", decoration.Yellow, decoration.Normal)
	for inputIndex, _ := range inputOrder {
		if err = askInputByIndex(logger, decoration, project, inputs, inputOrder, inputIndex); err != nil {
			return err
		}
	}
	return nil
}

func askInputByIndex(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project, inputs map[string]*config.Input, inputOrder []string, inputIndex int) (err error) {
	inputCount := len(inputOrder)
	inputName := inputOrder[inputIndex]
	input := inputs[inputName]
	// show input title
	logger.Println(getInputTitleText(decoration, input, inputIndex, inputCount))
	if input.Description != "" {
		showInputDescription(logger, decoration, input)
	}
	userValue := ""
	if input.Secret {
		logger.Print(getEnterNewValueText(decoration, input))
		userValue, err = getPassword()
		if err != nil {
			return err
		}
		project.SetValue(inputName, userValue)
		logger.Println()
		logger.Println()
		return nil
	}
	if !getKeepValueConfirmation(logger, decoration, project, input) {
		logger.Print(getEnterNewValueText(decoration, input))
		fmt.Scanf("%s", &userValue)
		project.SetValue(inputName, userValue)
	}
	logger.Println()
	return nil
}

func getInputTitleText(decoration *monitor.Decoration, input *config.Input, inputIndex int, inputCount int) string {
	inputNumber := fmt.Sprintf("%d of %d)", inputIndex+1, inputCount)
	return fmt.Sprintf("💀 %s%s%s %s%s", decoration.Bold, decoration.Yellow, inputNumber, input.GetName(), decoration.Normal)
}

func getEnterNewValueText(decoration *monitor.Decoration, input *config.Input) string {
	decoratedInputName := getDecoratedInputName(decoration, input)
	return fmt.Sprintf("💀 Please enter new value for %s: ", decoratedInputName)
}

func getDecoratedInputName(decoration *monitor.Decoration, input *config.Input) string {
	return fmt.Sprintf("%s%s%s", decoration.Yellow, input.GetName(), decoration.Normal)
}

func getDecoratedValue(d *monitor.Decoration, value string) string {
	if value != "" {
		return fmt.Sprintf("%s%s%s", d.Yellow, value, d.Normal)
	}
	return fmt.Sprintf("%sempty%s", d.Faint, d.Normal)
}

func getKeepValueConfirmation(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project, input *config.Input) bool {
	decoratedInputName := getDecoratedInputName(decoration, input)
	currentValue := project.GetValue(input.GetName())
	decoratedCurrentValue := getDecoratedValue(decoration, currentValue)
	defaultValue := input.DefaultValue
	decoratedDefaultValue := getDecoratedValue(decoration, defaultValue)
	userConfirmation := ""
	for !boolean.IsTrue(userConfirmation) && !boolean.IsFalse(userConfirmation) {
		if defaultValue == currentValue {
			logger.Printf("   Default/Current value for %s is %s\n", decoratedInputName, decoratedDefaultValue)
		} else {
			logger.Printf("   Default value for %s is %s\n", decoratedInputName, decoratedDefaultValue)
			logger.Printf("   Current value for %s is %s\n", decoratedInputName, decoratedCurrentValue)
		}
		logger.Printf("💀 Do you want to keep it %s (Y/n)? ", decoratedCurrentValue)
		fmt.Scanf("%s", &userConfirmation)
		if userConfirmation == "" {
			userConfirmation = "y"
		}
	}
	return boolean.IsTrue(userConfirmation)
}

func showInputDescription(logger monitor.Logger, decoration *monitor.Decoration, input *config.Input) {
	indentation := "      "
	descriptionRows := strings.Split(strings.Trim(input.Description, "\n "), "\n")
	for _, row := range descriptionRows {
		logger.Printf("%s%s\n", indentation, row)
	}
}

func getPassword() (passwd string, err error) {
	passwdB, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	// Type cast byte slice to string.
	return string(passwdB), err
}
