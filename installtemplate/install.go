package installtemplate

import (
	"fmt"
	"path/filepath"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
)

// Install template
func Install(gitURL, dirName string) (err error) {
	templateDir := config.GetTemplateDir()
	shell := config.GetShell()
	// run git init
	if err = command.Run(shell, templateDir, fmt.Sprintf("git clone %s %s --depth=1", gitURL, dirName)); err != nil {
		return
	}
	// run install
	err = command.Run(shell, filepath.Join(templateDir, dirName), "./install-template.sh")
	return
}
