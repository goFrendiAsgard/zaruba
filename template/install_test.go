package template

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInstallTemplate(t *testing.T) {
	templatePath, err := filepath.Abs(os.Getenv("ZARUBA_TEMPLATE_DIR"))
	if err != nil {
		t.Errorf("[ERROR] Cannot fetch templatePath from envvar: %s", err)
	}
	os.MkdirAll(templatePath, 0777)
	templateGitPath := filepath.Join(templatePath, "testInstallTemplate/.git")

	// Install template should succeed
	err = Install("https://github.com/state-alchemists/zaruba-project-template.git", "testInstallTemplate")
	if err != nil {
		t.Errorf("[ERROR] Cannot install template: %s", err)
	}

	// After install template, template signature should be exists
	if _, err := os.Stat(templateGitPath); os.IsNotExist(err) {
		t.Errorf("[ERROR] .git is not exist: %s", err)
	}

}
