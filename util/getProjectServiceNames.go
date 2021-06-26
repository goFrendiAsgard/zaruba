package util

import (
	"strings"

	"github.com/state-alchemists/zaruba/config"
)

func GetProjectServiceNames(project *config.Project) (serviceNames []string) {
	serviceNames = []string{}
	for taskName, task := range project.Tasks {
		if !strings.HasPrefix(taskName, "run") {
			continue
		}
		if task.Location == "" {
			continue
		}
		serviceName := strings.TrimPrefix(taskName, "run")
		if _, envRefExist := project.EnvRefMap[serviceName]; !envRefExist {
			continue
		}
		serviceNames = append(serviceNames, serviceName)
	}
	return serviceNames
}
