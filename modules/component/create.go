package component

import (
	"path/filepath"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
	"github.com/state-alchemists/zaruba/modules/strutil"
)

// Create component
func Create(template, projectDir string, args ...string) (err error) {
	templateDir := filepath.Join(config.GetTemplateDir(), template)
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return err
	}
	// run create-component script
	createComponentArgs := append([]string{projectDir}, args...)
	logger.Info("☀️ Create component `%s` at `%s` %s", templateDir, projectDir, strutil.SprintArgs(createComponentArgs))
	// create component
	return command.RunInteractively(templateDir, filepath.Join(templateDir, "create-component.zaruba"), createComponentArgs...)
}
