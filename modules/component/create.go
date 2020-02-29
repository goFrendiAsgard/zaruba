package component

import (
	"log"
	"path/filepath"

	"github.com/state-alchemists/zaruba/modules/action"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/stringformat"
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
	log.Printf("[INFO] Create component from `%s` into `%s` %s", templateDir, projectDir, stringformat.SprintArgs(createComponentArgs))
	return action.Do(
		"create-component",
		action.NewOption().
			SetScriptDir(templateDir).
			SetWorkDir(templateDir).
			SetIsRecursiveWorkDir(false),
		createComponentArgs...,
	)
}
