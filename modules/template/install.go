package template

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Install template
func Install(gitURL, newTemplateName string) (err error) {
	baseTemplateDir := config.GetTemplateDir()
	templateDir := filepath.Join(baseTemplateDir, newTemplateName)
	logger.Info("☀️ Install template from `%s` to `%s`", gitURL, templateDir)
	// run git clone
	if err = command.RunInteractively(baseTemplateDir, "git", "clone", gitURL, newTemplateName, "--depth=1"); err != nil {
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
	return command.RunInteractively(templateDir, filepath.Join(templateDir, "install-template.zaruba"))
}

func isScriptExists(templateDir, actionName string) (exist bool) {
	// imperative
	if _, err := os.Stat(filepath.Join(templateDir, actionName+".zaruba")); err == nil {
		return true
	}
	return false
}
