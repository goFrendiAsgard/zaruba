package core

import "testing"

func TestTplGetConfig(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/getConfig.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	// key::subkey1
	expected := "value1"
	actual, err := tpl.GetConfig("key::subKey1")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// key::subkey2
	expected = "value2"
	actual, err = tpl.GetConfig("key::subKey2")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTplGetSubConfigKeys(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/getConfig.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	// key::subkey1
	expectedKeys := []string{"subKey1", "subKey2"}
	actualKeys := tpl.GetSubConfigKeys("key")
	if err != nil {
		t.Error(err)
	}
	if len(actualKeys) != len(expectedKeys) {
		t.Errorf("expected: %#v, actual %#v", expectedKeys, actualKeys)
	}
	for _, expected := range expectedKeys {
		keyFound := false
		for _, actual := range actualKeys {
			if actual == expected {
				keyFound = true
				break
			}
		}
		if !keyFound {
			t.Errorf("key not found: %s, keys: %#v", expected, actualKeys)
		}
	}
}

func TestTplGetValue(t *testing.T) {
	project, err := getProject("../test-resources/taskdata/getValue/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	project.AddValue("../test-resources/taskdata/getValue/default.values.yaml")
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	// key::subkey1
	expected := "value1"
	actual, err := tpl.GetValue("key::subKey1")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// key::subkey2
	expected = "value2"
	actual, err = tpl.GetValue("key::subKey2")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTplGetSubValueKeys(t *testing.T) {
	project, err := getProject("../test-resources/taskdata/getValue/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	project.AddValue("../test-resources/taskdata/getValue/default.values.yaml")
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	// key::subkey1
	expectedKeys := []string{"subKey1", "subKey2"}
	actualKeys := tpl.GetSubValueKeys("key")
	if err != nil {
		t.Error(err)
	}
	if len(actualKeys) != len(expectedKeys) {
		t.Errorf("expected: %#v, actual %#v", expectedKeys, actualKeys)
	}
	for _, expected := range expectedKeys {
		keyFound := false
		for _, actual := range actualKeys {
			if actual == expected {
				keyFound = true
				break
			}
		}
		if !keyFound {
			t.Errorf("key not found: %s, keys: %#v", expected, actualKeys)
		}
	}
}

func TestTplGetEnv(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/getEnv.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	// key1
	expected := "VALUE1"
	actual, err := tpl.GetEnv("KEY1")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// key2
	expected = "VALUE2"
	actual, err = tpl.GetEnv("KEY2")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTplGetEnvs(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/getEnv.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	// key1
	expectedEnvs := map[string]string{"KEY1": "VALUE1", "KEY2": "VALUE2"}
	actualEnvs, err := tpl.GetEnvs()
	if err != nil {
		t.Error(err)
	}
	for key, expected := range expectedEnvs {
		actual := actualEnvs[key]
		if actual != expected {
			t.Errorf("key: %s, expected: %s, actual: %s", key, expected, actual)
		}
	}
}
