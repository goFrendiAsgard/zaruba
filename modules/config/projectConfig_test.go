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
	if err := file.Copy("../../test-resource/testProjectConfig/services/zaruba.config.yaml", filepath.Join(testPath, "services", "zaruba.config.yaml")); err != nil {
		t.Errorf("[ERROR] Cannot copy services/zaruba.config.yaml: %s", err)
		return
	}

	config, err := LoadProjectConfig(testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot load config: %s", err)
		return
	}

	// test cascaded config
	expected := filepath.Join(testPath, "./services/gateway")
	if config.Components["gateway"].Location != expected {
		t.Errorf("[UNEXPECTED] config.Components[\"gateway\"].Location should be `%s`, but contains `%s`", expected, config.Components["gateway"].Location)
	}

	// test sorted Link sources
	sortedLinkSources := config.GetSortedLinkSources()
	expectations := []string{
		filepath.Join(testPath, "./changelog.md"),
		filepath.Join(testPath, "./libraries/greeting-lib"),
	}
	for index, expected := range expectations {
		if sortedLinkSources[index] != expected {
			t.Errorf("[UNEXPECTED] sortedLinkSources[%d] should be `%s`, but contains: %s", index, expected, sortedLinkSources[index])
			t.Errorf("[INFO] config.Links: %#v", config.Links)
		}
	}
	// the length of sortedLinkSources should be 4
	if len(sortedLinkSources) != 4 {
		t.Errorf("[UNEXPECTED] len(sortedLinkSources) should be 4, but contains %d", len(sortedLinkSources))
	}

	// test subRepoPrefixMap
	subRepoPrefixMap := config.GetSubrepoPrefixMap(testPath)
	expected = "services/gateway"
	if subRepoPrefixMap["gateway"] != expected {
		t.Errorf("[UNEXPECTED] subRepoPrefixMap[\"gateway\"] should be `%s`, but contains: `%s`", subRepoPrefixMap["gateway"], expected)
	}

	// test YAML conversion
	yaml, err := config.ToYaml()
	if err != nil {
		t.Errorf("[ERROR] Cannot convert config into YAML: %s", err)
	} else if yaml == "" {
		t.Errorf("[UNEXPECTED] YAML is empty")
	}

}
