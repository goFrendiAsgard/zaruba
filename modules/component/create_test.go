package component

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/file"
)

func TestCreateComponent(t *testing.T) {
	templatePath := config.GetTemplateDir()
	if err := file.Copy("../../test-resource/template", filepath.Join(templatePath, "empty-project")); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testCreateComponent")

	// Create component should succeed
	err := Create("empty-project", testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}

	// After create component, component should be exists
	gitPath := filepath.Join(testPath, ".git")
	if _, err := os.Stat(gitPath); os.IsNotExist(err) {
		t.Errorf("[UNEXPECTED] %s is not exist: %s", gitPath, err)
	}

}
