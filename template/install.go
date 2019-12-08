package template

import (
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
	// run install
	err = command.Run(filepath.Join(templateDir, dirName), "./install-template")
	return
}
