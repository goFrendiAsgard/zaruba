package config

import (
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/modules/file"
)

func TestLoadProjectConfig(t *testing.T) {
	baseTestPath := GetTestDir()
	testPath := filepath.Join(baseTestPath, "testProjectConfig")

	if err := file.Copy("../../test-resource/project", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}
	if err := file.Copy("../../test-resource/testProjectConfig/zaruba.config.yaml", filepath.Join(testPath, "zaruba.config.yaml")); err != nil {
		t.Errorf("[ERROR] Cannot copy zaruba.config.yaml: %s", err)
		return
	}

	config, err := LoadProjectConfig(testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot load config: %s", err)
		return
	}

	sortedLinkSources := config.GetSortedLinkSources()

	expectations := []string{
		filepath.Join(testPath, "./changelog.md"),
		filepath.Join(testPath, "./libraries/greeting-lib"),
		filepath.Join(testPath, "./services/greeter"),
		filepath.Join(testPath, "./services/gateway"),
	}
	for index, expected := range expectations {
		if sortedLinkSources[index] != expected {
			t.Errorf("[UNEXPECTED] sortedLinkSources[%d] should be `%s`, but contains: %s", index, expected, sortedLinkSources[index])
			t.Errorf("[INFO] config.Links: %#v", config.Links)
		}
	}

}
