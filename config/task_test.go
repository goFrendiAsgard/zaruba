package config

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestTaskGetName(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getName.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expected := "taskName"
	actual := task.GetName()
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTaskGetTimeoutDuration(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getTimeoutDuration.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expected := time.Hour
	actual := task.GetTimeoutDuration()
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTaskGetWorkPathByLocation(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getWorkPathByLocation.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	workPath := task.GetWorkPath()
	if !strings.HasSuffix(workPath, "/someLocation") {
		t.Errorf("unexpected workPath: %s", workPath)
	}
}

func TestTaskGetWorkPathByParentLocation(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getWorkPathByParentLocation.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	workPath := task.GetWorkPath()
	if !strings.HasSuffix(workPath, "/someLocation") {
		t.Errorf("unexpected workPath: %s", workPath)
	}
}

func TestTaskGetWorkPathWithoutLocation(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getWorkPathWithoutLocation.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expectedWorkPath, _ := os.Getwd()
	workPath := task.GetWorkPath()
	if workPath != expectedWorkPath {
		t.Errorf("unexpected workPath: %s", workPath)
	}
}

func TestTaskHaveStartCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/haveStartCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expected := true
	actual := task.HaveStartCmd()
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestTaskHaveStartCmdByParent(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/haveStartCmdByParent.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expected := true
	actual := task.HaveStartCmd()
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestTaskHaveNoStartCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/haveNoStartCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expected := false
	actual := task.HaveStartCmd()
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestTaskHaveCheckCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/haveCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expected := true
	actual := task.HaveCheckCmd()
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestTaskHaveCheckCmdByParent(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/haveCheckCmdByParent.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expected := true
	actual := task.HaveCheckCmd()
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestTaskHaveNoCheckCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/haveNoCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expected := false
	actual := task.HaveCheckCmd()
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestTaskGetValueKeys(t *testing.T) {
	project, err := getProject("../test-resources/task/getValue.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	project.SetValue("key::subKey1", "value1")
	project.SetValue("key::subKey2", "value2")
	project.Init()
	task := project.Tasks["taskName"]
	expectedKeys := []string{"key::subKey1", "key::subKey2"}
	actualKeys := task.GetValueKeys()
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

func TestTaskGetValue(t *testing.T) {
	project, err := getProject("../test-resources/task/getValue.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.SetValue("key::subKey1", "value1"); err != nil {
		t.Error(err)
		return
	}
	if err = project.SetValue("key::subKey2", "value2"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expected := "value1"
	actual, err := task.GetValue("key", "subKey1")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected = "value2"
	actual, err = task.GetValue("key", "subKey2")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected = ""
	actual, err = task.GetValue("key", "subKey3")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTaskGetConfigKeys(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getConfig.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expectedKeys := []string{"key", "refKey", "parentKey"}
	actualKeys := task.GetConfigKeys()
	if len(actualKeys) != len(expectedKeys) {
		t.Errorf("expected: %#v, actual %#v", expectedKeys, actualKeys)
		return
	}
	for index, expected := range expectedKeys {
		actual := actualKeys[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
}

func TestTaskGetConfigPattern(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getConfig.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	// key
	expected := "value"
	actual, declared := task.GetConfigPattern("key")
	if !declared {
		t.Error("declared is false")
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// refKey
	expected = "refValue"
	actual, declared = task.GetConfigPattern("refKey")
	if !declared {
		t.Error("declared is false")
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// parentKey
	expected = "{{ .GetConfig \"key\" }}"
	actual, declared = task.GetConfigPattern("parentKey")
	if !declared {
		t.Error("declared is false")
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// inExistKey
	expected = ""
	actual, declared = task.GetConfigPattern("inExistKey")
	if declared {
		t.Error("declared is true")
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTaskGetConfig(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getConfig.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	// key
	expected := "value"
	actual, err := task.GetConfig("key")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// refKey
	expected = "refValue"
	actual, err = task.GetConfig("refKey")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// parentKey
	expected = "value"
	actual, err = task.GetConfig("parentKey")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// inExistKey
	expected = ""
	actual, err = task.GetConfig("inExistKey")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTaskGetConfigBrokenTemplate(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getConfigBrokenTemplate.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	// key
	if _, err = task.GetConfig("key"); err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "template:") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestTaskGetEnvKeys(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getEnv.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expectedKeys := []string{"KEY", "REF_KEY", "PARENT_KEY"}
	actualKeys := task.GetEnvKeys()
	if len(actualKeys) != len(expectedKeys) {
		t.Errorf("expected: %#v, actual %#v", expectedKeys, actualKeys)
		return
	}
	for index, expected := range expectedKeys {
		actual := actualKeys[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
}

func TestTaskGetEnv(t *testing.T) {
	project, err := getProject("../test-resources/task/getEnv.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddGlobalEnv("MY_KEY=MY_VALUE"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	// key
	expected := "MY_VALUE"
	actual, err := task.GetEnv("KEY")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// refKey
	expected = "REF_VALUE"
	actual, err = task.GetEnv("REF_KEY")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// parentKey
	expected = "PARENT_VALUE"
	actual, err = task.GetEnv("PARENT_KEY")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// inExistKey
	expected = ""
	actual, err = task.GetEnv("inExistKey")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTaskGetEnvs(t *testing.T) {
	project, err := getProject("../test-resources/task/getEnv.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddGlobalEnv("MY_KEY=MY_VALUE"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expectedEnvs := map[string]string{
		"KEY":        "MY_VALUE",
		"REF_KEY":    "REF_VALUE",
		"PARENT_KEY": "PARENT_VALUE",
	}
	actualEnvs, err := task.GetEnvs()
	if err != nil {
		t.Error(err)
		return
	}
	for key, expected := range expectedEnvs {
		actual := actualEnvs[key]
		if actual != expected {
			t.Errorf("key: %s, expected: %s, actual: %s", key, expected, actual)
		}
	}
}

func TestTaskGetEnvBrokenTemplate(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getEnvBrokenTemplate.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	// key
	if _, err = task.GetEnvs(); err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "template:") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestTaskRecursiveTemplate(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/recursiveTemplate.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	// key
	if _, err = task.GetConfig("key"); err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.Contains(errorMessage, "max recursive parsing on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestTaskMultiLineTemplate(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/multiLineTemplate.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	// keyTwoLine
	expected := "value"
	actual, err := task.GetConfig("keyTwoLine")
	if err != nil {
		t.Error(err)
		return
	}
	if strings.Trim(actual, "\n") != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	// keyTwoLine
	expected = "value1\nvalue2\nvalue3"
	actual, err = task.GetConfig("keyMultiLine")
	if err != nil {
		t.Error(err)
		return
	}
	if strings.Trim(actual, "\n") != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTaskGetDependenciesTemplate(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getDependencies.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	expectedList := []string{"dependencyTaskName1", "dependencyTaskName2", "dependencyTaskName3"}
	actualList := task.GetDependencies()
	if len(actualList) != len(expectedList) {
		t.Errorf("expected: %#v, actual %#v", expectedList, actualList)
		return
	}
	for index, expected := range expectedList {
		actual := actualList[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
}
