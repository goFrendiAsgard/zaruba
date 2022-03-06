package core

import (
	"github.com/state-alchemists/zaruba/output"
)

func getProject(projectFile string) (project *Project, err error) {
	SetDefaultEnv()
	decoration := output.NewDefaultDecoration()
	project, err = NewCustomProject(projectFile, decoration, []string{})
	if err != nil {
		return project, err
	}
	return project, err
}

func getProjectAndInit(projectFile string) (project *Project, err error) {
	project, err = getProject(projectFile)
	if err != nil {
		return project, err
	}
	err = project.Init()
	return project, err
}
