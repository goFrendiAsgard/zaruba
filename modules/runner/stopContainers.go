package runner

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// StopContainers stop all containers in this project
func StopContainers(projectDir string, p *config.ProjectConfig) (err error) {
	for componentName, component := range p.GetComponents() {
		if component.GetType() != "container" {
			continue
		}
		logger.Info("Stop %s container", componentName)
		_, err = command.Run(projectDir, "docker", "stop", component.GetRuntimeContainerName())
		if err != nil {
			logger.Error("Cannot stop container %s: %s", componentName, err)
		}
	}
	return nil
}
