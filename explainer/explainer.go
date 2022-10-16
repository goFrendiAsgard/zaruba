package explainer

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

type Explainer struct {
	logger  output.Logger
	d       *output.Decoration
	project *dsl.Project
	util    *dsl.DSLUtil
}

func NewExplainer(logger output.Logger, decoration *output.Decoration, project *dsl.Project) *Explainer {
	return &Explainer{
		logger:  logger,
		d:       decoration,
		project: project,
		util:    dsl.NewDSLUtil(),
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
