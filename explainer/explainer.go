package explainer

import (
	"fmt"
	"strings"

	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

type Explainer struct {
	logger  output.Logger
	d       *output.Decoration
	project *core.Project
	util    *core.CoreUtil
}

func NewExplainer(logger output.Logger, decoration *output.Decoration, project *core.Project) *Explainer {
	return &Explainer{
		logger:  logger,
		d:       decoration,
		project: project,
		util:    core.NewCoreUtil(),
	}
}

func (e *Explainer) Explain(taskNames ...string) (err error) {
	for _, taskName := range taskNames {
		if _, exist := e.project.Tasks[taskName]; !exist {
			return fmt.Errorf("task %s does not exist", taskName)
		}
		e.explainTask(taskName)
	}
	return nil
}

func (e *Explainer) explainTask(taskName string) {
	task := e.project.Tasks[taskName]
	taskExplanation := NewTaskExplanation(e.d, task)
	fmt.Println(taskExplanation.ToString())
}

func (e *Explainer) GetZarubaCommand(taskNames []string, autoTerminate bool, autoTerminateDelayInterval string) (command string) {
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
			inputArgs = append(inputArgs, fmt.Sprintf("%s=%s", inputName, e.util.Str.EscapeShellArg(inputValue)))
		}
	}
	if len(inputArgs) != 0 {
		command = fmt.Sprintf("%s %s", command, strings.Join(inputArgs, " "))
	}
	if autoTerminate {
		command = fmt.Sprintf("%s -t -w %s", command, autoTerminateDelayInterval)
	}
	return command
}
