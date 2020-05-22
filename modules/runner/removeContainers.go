package runner

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// RemoveContainers stop and remove all containers in this project
func RemoveContainers(projectDir string, p *config.ProjectConfig) (err error) {
	// stop containers
	err = StopContainers(projectDir, p)
	if err != nil {
		return err
	}
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
			logger.Info("Remove %s container", componentName)
			_, err = command.Run(projectDir, "docker", "rm", component.GetRuntimeContainerName())
			if err != nil {
				logger.Error("Cannot stop container %s: %s", componentName, err)
			}
		}
	}
	return nil
}
