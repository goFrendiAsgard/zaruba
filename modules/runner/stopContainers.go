package runner

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// StopContainers stop all containers in this project
func StopContainers(projectDir string, p *config.ProjectConfig) (err error) {
	serviceNames := getServiceNames(p)
	for _, serviceName := range serviceNames {
		component, err := p.GetComponentByName(serviceName)
		if err != nil {
			return err
		}
		if component.GetType() == "container" {
			logger.Info("Stop %s container", serviceName)
			err = command.RunAndRedirect(projectDir, "docker", "stop", component.GetRuntimeContainerName())
			if err != nil {
				logger.Error("Cannot stop container %s: %s", serviceName, err)
			}
		}
	}
	return nil
}
