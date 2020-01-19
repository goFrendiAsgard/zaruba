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
	if err := file.Copy("../../test-resource/zaruba.config.megazord.yaml", filepath.Join(testPath, "zaruba.config.yaml")); err != nil {
		t.Errorf("[ERROR] Cannot copy zaruba.config.yaml: %s", err)
		return
	}

	config, err := LoadProjectConfig(testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot load config: %s", err)
		return
	}
	sortedLinkSources := config.GetSortedLinkSources()

	expected := filepath.Join(testPath, "./changelog.md")
	if sortedLinkSources[0] != expected {
		t.Errorf("[UNEXPECTED] sortedLinkSources[0] should be `%s`, but sortedLinksSources contains: %s", expected, sortedLinkSources)
		t.Errorf("[INFO] config.Links: %#v", config.Links)
	}

	expected = filepath.Join(testPath, "./libraries/greeting-lib")
	if sortedLinkSources[1] != expected {
		t.Errorf("[UNEXPECTED] sortedLinkSources[1] should be `%s`, but sortedLinksSources contains: %s", expected, sortedLinkSources)
		t.Errorf("[INFO] config.Links: %#v", config.Links)
	}

}
