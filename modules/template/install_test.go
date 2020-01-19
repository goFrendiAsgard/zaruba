package template

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/modules/config"
)

func TestInstallTemplate(t *testing.T) {
	templatePath := config.GetTemplateDir()
	os.MkdirAll(templatePath, 0777)
	templateGitPath := filepath.Join(templatePath, "testInstallTemplate/.git")

	// Install template should succeed
	err := Install("https://github.com/state-alchemists/zaruba-project-template.git", "testInstallTemplate")
	if err != nil {
		t.Errorf("[ERROR] Cannot install template: %s", err)
	}

	// After install template, template signature should be exists
	if _, err := os.Stat(templateGitPath); os.IsNotExist(err) {
		t.Errorf("[ERROR] .git is not exist: %s", err)
	}

}
