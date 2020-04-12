package component

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/modules/config"
)

func TestCreateComponent(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testCreateComponent")

	// Create component should succeed
	if err := Create("project", testPath); err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}

	// After create component, component should be exists
	gitPath := filepath.Join(testPath, ".git")
	if _, err := os.Stat(gitPath); os.IsNotExist(err) {
		t.Errorf("[UNEXPECTED] %s is not exist: %s", gitPath, err)
	}

}
