package inputer

import (
	"fmt"
	"os"
	"strings"

	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/logger"
	"golang.org/x/crypto/ssh/terminal"
)

func Ask(project *config.Project, taskNames []string) (err error) {
	if project.IsInitialized {
		return fmt.Errorf("Cannot ask for input because project has been initialized")
	}
	inputs, inputOrder, err := project.GetInputs(taskNames)
	if err != nil {
		return err
	}
	d := logger.NewDecoration()
	inputCount := len(inputOrder)
	if inputCount == 0 {
		logger.Printf("%sZaruba is running in interactive mode. But no value you can set interactively.%s\n", d.Yellow, d.Normal)
		return nil
	}
	logger.Printf("%sZaruba is running in interactive mode. You will be able to set some values interactively.%s\n\n", d.Yellow, d.Normal)
	for inputIndex, _ := range inputOrder {
		if err = askInputByIndex(project, inputs, inputOrder, inputIndex); err != nil {
			return err
		}
	}
	return nil
}

func askInputByIndex(project *config.Project, inputs map[string]*config.Input, inputOrder []string, inputIndex int) (err error) {
	inputCount := len(inputOrder)
	inputName := inputOrder[inputIndex]
	input := inputs[inputName]
	// show input title
	fmt.Println(getInputTitleText(input, inputIndex, inputCount))
	if input.Description != "" {
		showInputDescription(input)
	}
	userValue := ""
	if input.Secret {
		fmt.Print(getEnterNewValueText(input))
		userValue, err = getPassword()
		if err != nil {
			return err
		}
		project.SetValue(inputName, userValue)
		fmt.Println()
		fmt.Println()
		return nil
	}
	if !getKeepValueConfirmation(project, input) {
		fmt.Print(getEnterNewValueText(input))
		fmt.Scanf("%s", &userValue)
		project.SetValue(inputName, userValue)
	}
	fmt.Println()
	return nil
}

func getInputTitleText(input *config.Input, inputIndex int, inputCount int) string {
	d := logger.NewDecoration()
	inputNumber := fmt.Sprintf("%d of %d)", inputIndex+1, inputCount)
	return fmt.Sprintf("💀 %s%s%s %s%s", d.Bold, d.Yellow, inputNumber, input.GetName(), d.Normal)
}

func getEnterNewValueText(input *config.Input) string {
	decoratedInputName := getDecoratedInputName(input)
	return fmt.Sprintf("💀 Please enter new value for %s: ", decoratedInputName)
}

func getDecoratedInputName(input *config.Input) string {
	d := logger.NewDecoration()
	return fmt.Sprintf("%s%s%s", d.Yellow, input.GetName(), d.Normal)
}

func getDecoratedValue(value string) string {
	d := logger.NewDecoration()
	if value != "" {
		return fmt.Sprintf("%s%s%s", d.Yellow, value, d.Normal)
	}
	return fmt.Sprintf("%sempty%s", d.Faint, d.Normal)
}

func getKeepValueConfirmation(project *config.Project, input *config.Input) bool {
	decoratedInputName := getDecoratedInputName(input)
	currentValue := project.GetValue(input.GetName())
	decoratedCurrentValue := getDecoratedValue(currentValue)
	defaultValue := input.DefaultValue
	decoratedDefaultValue := getDecoratedValue(defaultValue)
	userConfirmation := ""
	for !boolean.IsTrue(userConfirmation) && !boolean.IsFalse(userConfirmation) {
		if defaultValue == currentValue {
			fmt.Printf("   Default/Current value for %s is %s\n", decoratedInputName, decoratedDefaultValue)
		} else {
			fmt.Printf("   Default value for %s is %s\n", decoratedInputName, decoratedDefaultValue)
			fmt.Printf("   Current value for %s is %s\n", decoratedInputName, decoratedCurrentValue)
		}
		fmt.Printf("💀 Do you want to keep it %s (Y/n)? ", decoratedCurrentValue)
		fmt.Scanf("%s", &userConfirmation)
		if userConfirmation == "" {
			userConfirmation = "y"
		}
	}
	return boolean.IsTrue(userConfirmation)
}

func showInputDescription(input *config.Input) {
	indentation := "      "
	descriptionRows := strings.Split(strings.Trim(input.Description, "\n "), "\n")
	for _, row := range descriptionRows {
		fmt.Printf("%s%s\n", indentation, row)
	}
}

func getPassword() (passwd string, err error) {
	// https://godoc.org/golang.org/x/crypto/ssh/terminal#ReadPassword
	// terminal.ReadPassword accepts file descriptor as argument, returns byte slice and error.
	passwdB, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	// Type cast byte slice to string.
	return string(passwdB), err
}
