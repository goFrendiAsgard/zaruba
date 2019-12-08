package template

import (
	"os"
	"testing"
)

func TestInstallTemplate(t *testing.T) {
	os.Setenv("ZARUBA_TEMPLATE_DIR", "../templates")

	// Install template should succeed
	err := Install("https://github.com/state-alchemists/zaruba-project-template.git", "testInstallTemplate")
	if err != nil {
		t.Errorf("[ERROR] Cannot install template: %s", err)
	}

	// After install template, template signature should be exists
	if _, err := os.Stat("../templates/testInstallTemplate/.git"); os.IsNotExist(err) {
		t.Errorf("[ERROR] ../templates/testInstallTemplate/.git is not exist: %s", err)
	}

}
