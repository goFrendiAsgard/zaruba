package component

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/file"
)

func TestCreateComponent(t *testing.T) {
	templatePath := config.GetTemplateDir()
	if err := file.Copy("../test-resource/project.template", filepath.Join(templatePath, "project")); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testCreateComponent")

	// Create component should succeed
	err := Create("project", testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}

	// After create component, component should be exists
	if _, err := os.Stat(filepath.Join(testPath, ".git")); os.IsNotExist(err) {
		t.Errorf("[UNEXPECTED] ../test-playground/testCreateComponent/.git is not exist: %s", err)
	}

}
