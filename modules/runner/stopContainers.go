package runner

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// StopContainers stop all containers in this project
func StopContainers(projectDir string, p *config.ProjectConfig) (err error) {
	componentNames, err := getExecutableComponentNames(p, []string{})
	if err != nil {
		return err
	}
	for _, componentName := range componentNames {
		component, err := p.GetComponentByName(componentName)
		if err != nil {
			return err
		}
		if component.GetType() == "container" {
			logger.Info("Stop %s container", componentName)
			_, err = command.Run(projectDir, "docker", "stop", component.GetRuntimeContainerName())
			if err != nil {
				logger.Error("Cannot stop container %s: %s", componentName, err)
			}
		}
	}
	return nil
}
