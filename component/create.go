package component

import (
	"path/filepath"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
)

// Create component
func Create(template, projectDir string, args ...string) (err error) {
	templateDir := config.GetTemplateDir()
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	// run create-component.sh
	createComponentArgs := append([]string{projectDir}, args...)
	err = command.Run(filepath.Join(templateDir, template), "./create-component.zaruba", createComponentArgs...)
	return
}
