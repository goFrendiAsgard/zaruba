package runner

import (
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/file"
)

func TestRun(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testOrganize")
	if err := file.Copy("../test-resource/project", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}
	if err := file.Copy("../test-resource/zaruba.config.megazord.yaml", filepath.Join(testPath, "zaruba.config.yaml")); err != nil {
		t.Errorf("[ERROR] Cannot copy zaruba.config.yaml: %s", err)
		return
	}
}
