package config

import (
	"github.com/state-alchemists/zaruba/output"
)

func getProject(projectFile string) (project *Project, logger output.Logger, recordLogger output.RecordLogger, err error) {
	decoration := output.NewDecoration()
	logger = output.NewMockLogger()
	recordLogger = output.NewMockRecordLogger()
	project, err = NewProject(logger, recordLogger, decoration, projectFile, []string{})
	return project, logger, recordLogger, err
}

func getProjectAndInit(projectFile string) (project *Project, logger output.Logger, recordLogger output.RecordLogger, err error) {
	project, logger, recordLogger, err = getProject(projectFile)
	if err != nil {
		return project, logger, recordLogger, err
	}
	err = project.Init()
	return project, logger, recordLogger, err
}
