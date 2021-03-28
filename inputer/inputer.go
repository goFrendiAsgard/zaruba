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
	logger.Printf("%sZaruba is running in interactive mode. You will be able to set some values interactively.%s\n", d.Yellow, d.Normal)
	for inputIndex, _ := range inputOrder {
		if err = ask(project, inputs, inputOrder, inputIndex); err != nil {
			return err
		}
	}
	return nil
}

func ask(project *config.Project, inputs map[string]*config.Input, inputOrder []string, inputIndex int) (err error) {
	inputCount := len(inputOrder)
	inputName := inputOrder[inputIndex]
	input := inputs[inputName]
	d := logger.NewDecoration()
	decoratedIndex := fmt.Sprintf("%s%d of %d)%s", d.Blue, inputIndex+1, inputCount, d.Normal)
	decoratedInputName := fmt.Sprintf("%s%s%s", d.Yellow, input.GetName(), d.Normal)
	currentValue := project.GetValue(inputName)
	decoratedCurrentValue := fmt.Sprintf("%sempty%s", d.Faint, d.Normal)
	if currentValue != "" {
		decoratedCurrentValue = fmt.Sprintf("%s%s%s", d.Yellow, currentValue, d.Normal)
	}
	// show number and input title
	fmt.Printf("%s %s\n", decoratedIndex, decoratedInputName)
	if input.Description != "" {
		showInputDescription(input)
	}
	userValue := ""
	if input.Secret {
		fmt.Printf("💀 Please enter new value for %s: ", decoratedInputName)
		userValue, err = getPassword()
		if err != nil {
			return err
		}
		fmt.Println()
	} else {
		// handle user input
		if !getUserOverwriteConfirmation(decoratedInputName, decoratedCurrentValue) {
			userValue = currentValue
		} else {
			fmt.Printf("💀 Please enter new value for %s: ", decoratedInputName)
			fmt.Scanf("%s", &userValue)
		}
	}
	project.SetValue(inputName, userValue)
	fmt.Println()
	return nil
}

func getUserOverwriteConfirmation(decoratedInputName, decoratedCurrentValue string) bool {
	userConfirmation := ""
	d := logger.NewDecoration()
	for !boolean.IsTrue(userConfirmation) && !boolean.IsFalse(userConfirmation) {
		fmt.Printf("   Current value for %s is %s\n", decoratedInputName, decoratedCurrentValue)
		fmt.Printf("💀 Do you want to %soverwrite%s it (Y/n)? ", d.Bold, d.Normal)
		fmt.Scanf("%s", &userConfirmation)
		if userConfirmation == "" {
			userConfirmation = "n"
		}
	}
	return boolean.IsTrue(userConfirmation)
}

func showInputDescription(input *config.Input) {
	indentation := "     "
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
