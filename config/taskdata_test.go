package config

import "testing"

func TestTdGetConfig(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/getConfig.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// key::subkey1
	expected := "value1"
	actual, err := td.GetConfig("key::subKey1")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// key::subkey2
	expected = "value2"
	actual, err = td.GetConfig("key::subKey2")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdGetSubConfigKeys(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/getConfig.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// key::subkey1
	expectedKeys := []string{"subKey1", "subKey2"}
	actualKeys := td.GetSubConfigKeys("key")
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
