package component

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/otiai10/copy"
)

func TestCreateComponent(t *testing.T) {

	templatePath, err := filepath.Abs(os.Getenv("ZARUBA_TEMPLATE_DIR"))
	if err != nil {
		t.Errorf("[ERROR] Cannot fetch templatePath from envvar: %s", err)
	}
	if err = copy.Copy("../test-resource/project.template", filepath.Join(templatePath, "project")); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	baseTestPath, err := filepath.Abs(os.Getenv("ZARUBA_TEST_DIR"))
	if err != nil {
		t.Errorf("[ERROR] Cannot fetch testPath from envvar: %s", err)
		return
	}
	testPath := filepath.Join(baseTestPath, "testCreateComponent")

	// Create component should succeed
	err = Create("project", testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}

	// After create component, component should be exists
	if _, err := os.Stat(filepath.Join(testPath, ".git")); os.IsNotExist(err) {
		t.Errorf("[UNEXPECTED] ../test-playground/testCreateComponent/.git is not exist: %s", err)
	}

}
