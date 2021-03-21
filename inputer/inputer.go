package inputer

import (
	"fmt"
	"os"
	"strings"

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
	logger.Printf("%sLeave blank if you want to keep current values.%s\n\n", d.Yellow, d.Normal)
	for inputIndex, inputName := range inputOrder {
		input := inputs[inputName]
		decoratedIndex := fmt.Sprintf("%s%s%d of %d)%s", d.Bold, d.Blue, inputIndex+1, inputCount, d.Normal)
		decoratedInputName := fmt.Sprintf("%s%s%s%s", d.Bold, d.Yellow, input.GetName(), d.Normal)
		// show number and input title
		fmt.Printf("%s %s\n", decoratedIndex, decoratedInputName)
		if input.Description != "" {
			showInputDescription(input)
		}
		// show current value
		decoratedCurrentValue := fmt.Sprintf("%sempty%s", d.Faint, d.Normal)
		currentValue := project.GetValue(inputName)
		if currentValue != "" {
			decoratedCurrentValue = fmt.Sprintf("%s%s%s", d.Yellow, currentValue, d.Normal)
		}
		fmt.Printf("%s (currently %s): ", decoratedInputName, decoratedCurrentValue)
		userValue := ""
		if input.Secret {
			userValue, err = getPassword()
			if err != nil {
				return err
			}
		} else {
			// handle user input
			fmt.Scanf("%s", &userValue)
		}
		if userValue != "" {
			project.SetValue(inputName, userValue)
		}
		fmt.Println()
	}
	return nil
}

func showInputDescription(input *config.Input) {
	indentation := "  "
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
