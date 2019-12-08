package component

import (
	"os"
	"testing"
)

func TestCreateComponent(t *testing.T) {
	os.Setenv("ZARUBA_TEMPLATE_DIR", "../templates")

	// Create component should succeed
	err := Create("project", "../test-playground/testCreateComponent")
	if err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
	}

	// After create component, component should be exists
	if _, err := os.Stat("../test-playground/testCreateComponent/.git"); os.IsNotExist(err) {
		t.Errorf("[ERROR] ../test-playground/testCreateComponent/.git is not exist: %s", err)
	}

}
