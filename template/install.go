package template

import (
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
)

// Install template
func Install(gitURL, dirName string) (err error) {
	templateDir := config.GetTemplateDir()
	// run git init
	if err = command.Run(templateDir, "git", "clone", gitURL, dirName, "--depth=1"); err != nil {
		return
	}
	// install-template should be exists
	if _, err = os.Stat(filepath.Join(templateDir, dirName, "install-template")); err != nil {
		os.RemoveAll(filepath.Join(templateDir, dirName))
		return
	}
	// create-component should be exists
	if _, err = os.Stat(filepath.Join(templateDir, dirName, "create-component")); err != nil {
		os.RemoveAll(filepath.Join(templateDir, dirName))
		return
	}
	// make the file executable
	os.Chmod(filepath.Join(templateDir, dirName, "install-template"), 0555)
	os.Chmod(filepath.Join(templateDir, dirName, "create-component"), 0555)
	// run install
	err = command.Run(filepath.Join(templateDir, dirName), "./install-template")
	return
}
