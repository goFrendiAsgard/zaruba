package explainer

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

func (e *Explainer) printField(fieldName string, value string, indentation string) {
	if value != "" {
		trimmedValue := strings.Trim(value, "\n ")
		valueLines := strings.Split(trimmedValue, "\n")
		indentedValue := strings.Join(valueLines, "\n"+indentation)
		e.logger.DPrintf("%s%s :%s %s\n", e.d.Yellow, fieldName, e.d.Normal, indentedValue)
	}
}

func (e *Explainer) Explain(taskName string) {
	task := e.project.Tasks[taskName]
	indentation := strings.Repeat(" ", 21)
	parentTasks := task.Extends
	if len(parentTasks) == 0 && task.Extend != "" {
		parentTasks = []string{task.Extend}
	}
	inputString := ""
	inputNames := task.Inputs
	inputCount := len(inputNames)
	if inputCount > 0 {
		paramIndentation := strings.Repeat(" ", 2)
		paramFieldIndentation := paramIndentation + strings.Repeat(" ", 12)
		for _, inputName := range inputNames {
			input := e.project.Inputs[inputName]
			inputOptionStr := strings.Join(input.Options, ", ")
			inputDesc := strings.Trim(input.Description, "\n")
			inputDescLines := strings.Split(inputDesc, "\n")
			inputDescStr := strings.Join(inputDescLines, "\n  "+paramFieldIndentation)
			inputLines := []string{
				fmt.Sprintf("- %s", inputName),
				fmt.Sprintf("%s%sOPTIONS     :%s %s", e.d.Yellow, paramIndentation, e.d.Normal, inputOptionStr),
				fmt.Sprintf("%s%sDESCRIPTION :%s %s", e.d.Yellow, paramIndentation, e.d.Normal, inputDescStr),
			}
			inputString += "\n" + strings.Trim(strings.Join(inputLines, "\n"), "\n")
		}
	}
	e.printField("TASK NAME   ", taskName, indentation)
	e.printField("LOCATION    ", task.GetFileLocation(), indentation)
	e.printField("DESCRIPTION ", task.Description, indentation)
	e.printField("INPUTS      ", inputString, indentation)
	e.printField("DEPENDENCIES", strings.Join(task.Dependencies, ", "), indentation)
	e.printField("PARENT TASKS", strings.Join(parentTasks, ", "), indentation)
}
