package runner

import (
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

func getProject(projectFile string) (project *config.Project, logger output.Logger, decoration *output.Decoration, err error) {
	decoration = output.NewDecoration()
	logger = output.NewMockLogger()
	recordLogger := output.NewMockRecordLogger()
	project, err = config.NewProject(logger, recordLogger, decoration, projectFile, []string{})
	return project, logger, decoration, err
}

func getProjectAndInit(projectFile string) (project *config.Project, logger output.Logger, decoration *output.Decoration, err error) {
	project, logger, decoration, err = getProject(projectFile)
	if err != nil {
		return project, logger, decoration, err
	}
	err = project.Init()
	return project, logger, decoration, err
}
