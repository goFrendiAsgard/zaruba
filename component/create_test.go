package component

import (
	"os"
	"testing"

	"github.com/state-alchemists/zaruba/template"
)

func TestCreateComponent(t *testing.T) {
	os.Setenv("ZARUBA_TEMPLATE_DIR", "../templates")

	// Install template should succeed
	err := template.Install("https://github.com/state-alchemists/zaruba-project-template.git", "templateForComponentTest")
	if err != nil {
		t.Errorf("[ERROR] Cannot install template: %s", err)
	}

	// Create component should succeed
	err = Create("templateForComponentTest", "../test-playground/project")
	if err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
	}

	// After create component, component should be exists
	if _, err := os.Stat("../test-playground/project/.git"); os.IsNotExist(err) {
		t.Errorf("[ERROR] ../test-playground/project/.git is not exist: %s", err)
	}

}
