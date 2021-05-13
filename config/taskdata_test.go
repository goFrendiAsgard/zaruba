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

func TestTdGetLConfig(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/getLConfig.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// key::subkey1
	expectedList := []string{"value1", "value2"}
	actualList, err := td.GetLConfig("key::subKey1")
	if err != nil {
		t.Error(err)
	}
	if len(actualList) != len(expectedList) {
		t.Errorf("expected: %#v, actual: %#v", expectedList, actualList)
	}
	for index, expected := range expectedList {
		actual := actualList[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
	// key::subkey2
	expectedList = []string{"value3", "value4"}
	actualList, err = td.GetLConfig("key::subKey2")
	if err != nil {
		t.Error(err)
	}
	if len(actualList) != len(expectedList) {
		t.Errorf("expected: %#v, actual: %#v", expectedList, actualList)
	}
	for index, expected := range expectedList {
		actual := actualList[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
}

func TestTdGetSubLConfigKeys(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/getLConfig.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// key::subkey1
	expectedKeys := []string{"subKey1", "subKey2"}
	actualKeys := td.GetSubLConfigKeys("key")
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

func TestTdGetValue(t *testing.T) {
	project, _, _, err := getProject("../test-resources/taskdata/getValue/main.zaruba.yaml")
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
	td := NewTaskData(task)
	// key::subkey1
	expected := "value1"
	actual, err := td.GetValue("key::subKey1")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// key::subkey2
	expected = "value2"
	actual, err = td.GetValue("key::subKey2")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdGetSubValueKeys(t *testing.T) {
	project, _, _, err := getProject("../test-resources/taskdata/getValue/main.zaruba.yaml")
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
	td := NewTaskData(task)
	// key::subkey1
	expectedKeys := []string{"subKey1", "subKey2"}
	actualKeys := td.GetSubValueKeys("key")
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

func TestTdGetEnv(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/getEnv.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// key1
	expected := "VALUE1"
	actual, err := td.GetEnv("KEY1")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// key2
	expected = "VALUE2"
	actual, err = td.GetEnv("KEY2")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdGetEnvs(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/getEnv.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// key1
	expectedEnvs := map[string]string{"KEY1": "VALUE1", "KEY2": "VALUE2"}
	actualEnvs, err := td.GetEnvs()
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
