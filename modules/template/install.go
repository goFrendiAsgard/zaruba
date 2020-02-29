package template

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/modules/action"
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
)

// Install template
func Install(gitURL, newTemplateName string) (err error) {
	baseTemplateDir := config.GetTemplateDir()
	templateDir := filepath.Join(baseTemplateDir, newTemplateName)
	log.Printf("[INFO] Install template from `%s` to `%s`", gitURL, templateDir)
	// run git clone
	if err = command.Run(baseTemplateDir, "git", "clone", gitURL, newTemplateName, "--depth=1"); err != nil {
		return err
	}
	// install-template should be exists
	if !isScriptExists(templateDir, "install-template") {
		os.RemoveAll(templateDir)
		err = errors.New("Cannot find `install-template` script")
		return err
	}
	// create-component should be exists
	if !isScriptExists(templateDir, "create-component") {
		os.RemoveAll(templateDir)
		err = errors.New("Cannot find `create-component` script")
		return err
	}
	// make the file executable
	os.Chmod(filepath.Join(templateDir, "install-template.zaruba"), 0555)
	os.Chmod(filepath.Join(templateDir, "create-component.zaruba"), 0555)
	// run install
	return action.Do(
		"install-template",
		action.NewOption().
			SetScriptDir(templateDir).
			SetWorkDir(templateDir).
			SetIsRecursiveWorkDir(false),
		templateDir,
	)
}

func isScriptExists(templateDir, actionName string) (exist bool) {
	// imperative
	if _, err := os.Stat(filepath.Join(templateDir, actionName+".zaruba")); err == nil {
		return true
	}
	// declarative yml
	if _, err := os.Stat(filepath.Join(templateDir, actionName+".zaruba.yml")); err == nil {
		return true
	}
	// declarative yaml
	if _, err := os.Stat(filepath.Join(templateDir, actionName+".zaruba.yml")); err == nil {
		return true
	}
	return false
}
