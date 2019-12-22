package component

import (
	"log"
	"path/filepath"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/format"
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
	log.Printf("[INFO] Create component from `%s` into `%s` %s", templateDir, projectDir, format.SprintArgs(createComponentArgs))
	err = command.Run(filepath.Join(templateDir, template), "./create-component.zaruba", createComponentArgs...)
	return
}
