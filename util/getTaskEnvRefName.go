package util

import "github.com/state-alchemists/zaruba/config"

func GetTaskEnvRefname(task *config.Task) (envRefName string) {
	if task.EnvRef != "" {
		return task.EnvRef
	}
	if len(task.EnvRefs) > 0 {
		return task.EnvRefs[0]
	}
	return ""
}
