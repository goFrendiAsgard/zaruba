package template

import (
	"os"
	"testing"
)

func TestInstall(t *testing.T) {
	os.Setenv("ZARUBA_TEMPLATE_DIR", "../templates")
	err := Install("https://github.com/state-alchemists/zaruba-project-template.git", "project")
	if err != nil {
		t.Errorf("[ERROR] Cannot install template: %s", err)
	}
}
