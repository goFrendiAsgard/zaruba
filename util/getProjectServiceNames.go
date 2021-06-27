package util

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/config"
)

func GetProjectServiceNames(project *config.Project) (serviceNames []string) {
	projectDir := filepath.Dir(project.GetFileLocation())
	serviceNames = []string{}
	for taskName, task := range project.Tasks {
		// taskName should be started with "run"
		if !strings.HasPrefix(taskName, "run") || taskName == "run" {
			continue
		}
		upperServiceName := strings.TrimPrefix(taskName, "run")
		serviceName := strings.ToLower(string(upperServiceName[0])) + upperServiceName[1:]
		// service's envRef should be exist
		if _, envRefExist := project.EnvRefMap[serviceName]; !envRefExist {
			continue
		}
		// task should be located at project's zaruba-tasks
		if task.GetFileLocation() != filepath.Join(projectDir, "zaruba-tasks", fmt.Sprintf("%s.zaruba.yaml", serviceName)) {
			continue
		}
		serviceNames = append(serviceNames, serviceName)
	}
	return serviceNames
}
