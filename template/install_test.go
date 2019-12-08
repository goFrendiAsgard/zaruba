package template

import (
	"os"
	"testing"
)

func TestInstallTemplate(t *testing.T) {
	os.Setenv("ZARUBA_TEMPLATE_DIR", "../templates")

	// Install template should succeed
	err := Install("https://github.com/state-alchemists/zaruba-project-template.git", "project")
	if err != nil {
		t.Errorf("[ERROR] Cannot install template: %s", err)
	}

	// After install template, template signature should be exists
	if _, err := os.Stat("../templates/project/.git"); os.IsNotExist(err) {
		t.Errorf("[ERROR] ../templates/project/.git is not exist: %s", err)
	}

}
