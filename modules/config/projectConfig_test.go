package config

import (
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/modules/file"
)

func TestLoadProjectConfig(t *testing.T) {
	baseTestPath := GetTestDir()
	testPath := filepath.Join(baseTestPath, "testProjectConfig")
	if err := file.Copy("../../test-resource/testProjectConfig", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	// load project config
	p, err := NewProjectConfig(testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot load config: %s", err)
		return
	}

	// test cascaded config
	expected := filepath.Join(testPath, "./services/gopher")
	component, err := p.GetComponentByName("gopher")
	if err != nil {
		t.Errorf("[ERROR] Cannot get component: %s", err)
	} else if component.GetLocation() != expected {
		t.Errorf("[UNEXPECTED] config.Components[\"gopher\"].Location should be `%s`, but contains `%s`", expected, component.GetLocation())
	}

	// test component runtime environment
	component, err = p.GetComponentByName("gopher")
	if err != nil {
		t.Errorf("[ERROR] Cannot get component: %s", err)
	}
	expectedEnv := map[string]string{
		"GOPHER_HTTP_PORT": "3011",
		"RMQ_HOST":         "0.0.0.0",
		"RMQ_PASSWORD":     "toor",
		"RMQ_PORT":         "5672",
		"RMQ_USER":         "root",
		"RMQ_VHOST":        "/",
		"gopher":           "0.0.0.0",
		"rmq":              "0.0.0.0",
	}
	env := component.GetRuntimeEnv()
	for name, expectedValue := range expectedEnv {
		if value := env[name]; value != expectedValue {
			t.Errorf("[UNEXPECTED] env `%s` should be `%s`, but contains `%s`", name, expectedValue, value)
		}
	}

	// test subRepoPrefixMap
	expected = "services/gopher"
	if subRepoPrefixMap := p.GetSubrepoPrefixMap(testPath); subRepoPrefixMap["gopher"] != expected {
		t.Errorf("[UNEXPECTED] subRepoPrefixMap[\"gopher\"] should be `%s`, but contains: `%s`", expected, subRepoPrefixMap["gopher"])
	}

	// test YAML conversion
	yaml, err := p.ToYaml()
	if err != nil {
		t.Errorf("[ERROR] Cannot convert config into YAML: %s", err)
	} else if yaml == "" {
		t.Errorf("[UNEXPECTED] YAML is empty")
	}

}
