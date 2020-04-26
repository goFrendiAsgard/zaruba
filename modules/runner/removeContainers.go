package runner

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// RemoveContainers stop and remove all containers in this project
func RemoveContainers(projectDir string, p *config.ProjectConfig) (err error) {
	err = StopContainers(projectDir, p)
	if err != nil {
		return err
	}
	serviceNames := getServiceNames(p)
	for _, serviceName := range serviceNames {
		component, err := p.GetComponentByName(serviceName)
		if err != nil {
			return err
		}
		if component.GetType() == "container" {
			logger.Info("Remove %s container", serviceName)
			err = command.RunAndRedirect(projectDir, "docker", "rm", component.GetRuntimeContainerName())
			if err != nil {
				logger.Error("Cannot stop container %s: %s", serviceName, err)
			}
		}
	}
	return nil
}
