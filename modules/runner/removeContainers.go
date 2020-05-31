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
	for componentName, component := range p.GetComponents() {
		if component.GetType() != "container" {
			continue
		}
		logger.Info("🚮 Remove %s container", componentName)
		_, err = command.RunSilently(projectDir, "docker", "rm", component.GetRuntimeContainerName())
		if err != nil {
			logger.Error("Cannot remove container %s: %s", componentName, err)
		}
	}
	return nil
}
