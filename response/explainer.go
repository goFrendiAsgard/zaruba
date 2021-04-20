package response

import (
	"fmt"
	"strings"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/monitor"
)

type Explainer struct {
	logger  monitor.Logger
	d       *monitor.Decoration
	project *config.Project
}

func NewExplainer(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project) *Explainer {
	return &Explainer{
		logger:  logger,
		d:       decoration,
		project: project,
	}
}

func (e *Explainer) Explain(taskName string) {
	task := e.project.Tasks[taskName]
	indentation := strings.Repeat(" ", 21)
	parentTasks := task.Extends
	if len(parentTasks) == 0 && task.Extend != "" {
		parentTasks = []string{task.Extend}
	}
	e.printField("TASK NAME   ", taskName, indentation)
	e.printField("LOCATION    ", task.GetFileLocation(), indentation)
	e.printField("DESCRIPTION ", task.Description, indentation)
	e.printField("INPUTS      ", e.getInputString(task), indentation)
	e.printField("DEPENDENCIES", strings.Join(task.Dependencies, ", "), indentation)
	e.printField("PARENT TASKS", strings.Join(parentTasks, ", "), indentation)
}

func (e *Explainer) getInputString(task *config.Task) (inputString string) {
	inputNames := task.Inputs
	inputCount := len(inputNames)
	if inputCount > 0 {
		paramIndentation := strings.Repeat(" ", 2)
		for _, inputName := range inputNames {
			input := e.project.Inputs[inputName]
			rawInputLines := []string{
				fmt.Sprintf("- %s", inputName),
				e.getInputFieldString("DESCRIPTION", input.Description, paramIndentation),
				e.getInputFieldString("PROMPT     ", input.Prompt, paramIndentation),
				e.getInputFieldString("OPTIONS    ", strings.Join(input.Options, ", "), paramIndentation),
				e.getInputFieldString("DEFAULT    ", input.DefaultValue, paramIndentation),
				e.getInputFieldString("VALIDATION ", input.Validation, paramIndentation),
			}
			inputLines := []string{}
			for _, inputLine := range rawInputLines {
				if inputLine != "" {
					inputLines = append(inputLines, inputLine)
				}
			}
			inputString += "\n" + strings.Trim(strings.Join(inputLines, "\n"), "\n")
		}
	}
	return inputString
}

func (e *Explainer) getInputFieldString(inputFieldName string, inputFieldValue string, paramIndentation string) (inputFieldString string) {
	inputFieldValue = strings.Trim(inputFieldValue, "\n")
	if inputFieldValue == "" {
		return ""
	}
	paramFieldIndentation := paramIndentation + strings.Repeat(" ", 12)
	inputFieldLines := strings.Split(strings.Trim(inputFieldValue, "\n"), "\n")
	inputFieldValueStr := strings.Join(inputFieldLines, "\n  "+paramFieldIndentation)
	return fmt.Sprintf("%s%s%s :%s %s", paramIndentation, e.d.Yellow, inputFieldName, e.d.Normal, inputFieldValueStr)
}

func (e *Explainer) printField(fieldName string, value string, indentation string) {
	if value == "" {
		return
	}
	trimmedValue := strings.Trim(value, "\n ")
	valueLines := strings.Split(trimmedValue, "\n")
	indentedValue := strings.Join(valueLines, "\n"+indentation)
	e.logger.DPrintf("%s%s :%s %s\n", e.d.Yellow, fieldName, e.d.Normal, indentedValue)
}

func (e *Explainer) GetCommand(taskNames []string) (command string) {
	command = fmt.Sprintf("zaruba please %s", strings.Join(taskNames, " "))
	inputArgs := []string{}
	for _, taskName := range taskNames {
		task := e.project.Tasks[taskName]
		for _, inputName := range task.Inputs {
			inputValue := ""
			if input := e.project.Inputs[inputName]; input.Secret {
				inputValue = "[HIDDEN_VALUE]"
			} else {
				inputValue = e.project.GetValue(inputName)
			}
			inputArgs = append(inputArgs, fmt.Sprintf("%s=%s", inputName, inputValue))
		}
	}
	if len(inputArgs) != 0 {
		command = fmt.Sprintf("%s %s", command, strings.Join(inputArgs, " "))
	}
	return command
}
