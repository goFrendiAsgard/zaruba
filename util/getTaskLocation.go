package util

import (
	"fmt"

	"github.com/state-alchemists/zaruba/config"
)

func GetTaskLocation(project *config.Project, taskName string) (location string, err error) {
	task, taskExist := project.Tasks[taskName]
	if !taskExist {
		return "", fmt.Errorf("task %s doesn't exist", taskName)
	}
	return task.GetWorkPath(), nil
}
