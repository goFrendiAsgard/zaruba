package runner

import (
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

func getProject(projectFile string) (project *dsl.Project, err error) {
	dsl.SetDefaultEnv()
	decoration := output.NewDefaultDecoration()
	return dsl.NewCustomProject(projectFile, decoration, false, []string{})
}

func getProjectAndInit(projectFile string) (project *dsl.Project, err error) {
	project, err = getProject(projectFile)
	if err != nil {
		return project, err
	}
	err = project.Init()
	return project, err
}

func getRunner(project *dsl.Project, taskNames []string, statusIntervalStr string, autoTerminate bool, autoTerminateDelayStr string) (runner *Runner, logger *output.MockLogger, recordLogger *output.MockRecordLogger, err error) {
	logger = output.NewMockLogger()
	recordLogger = output.NewMockRecordLogger()
	runner, err = NewRunner(logger, recordLogger, project, taskNames, statusIntervalStr, 0, autoTerminate, autoTerminateDelayStr)
	return runner, logger, recordLogger, err
}
