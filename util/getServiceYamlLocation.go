package util

import (
	"fmt"
	"path/filepath"

	"github.com/state-alchemists/zaruba/config"
)

func GetServiceYamlLocation(project *config.Project, serviceName string) (yamlLocation string) {
	projectDir := filepath.Dir(project.GetFileLocation())
	return filepath.Join(projectDir, "zaruba-tasks", fmt.Sprintf("%s.zaruba.yaml", serviceName))
}
