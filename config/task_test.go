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

func TestTaskGetBasePath(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getBasePath.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	basePath := task.GetBasePath()
	if !strings.HasSuffix(basePath, "test-resources/task") {
		t.Errorf("unexpected basepath: %s", basePath)
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
		t.Errorf("unexpected basepath: %s", workPath)
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
		t.Errorf("unexpected basepath: %s", workPath)
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
		t.Errorf("unexpected basepath: %s", workPath)
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
	for index, expected := range expectedKeys {
		actual := actualKeys[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
}

func TestTaskGetValue(t *testing.T) {
	project, err := getProject("../test-resources/task/getValue.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	project.SetValue("key::subKey1", "value1")
	project.SetValue("key::subKey2", "value2")
	project.Init()
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
	// nonExistKey
	expected = ""
	actual, declared = task.GetConfigPattern("nonExistKey")
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
	// nonExistKey
	expected = ""
	actual, err = task.GetConfig("nonExistKey")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
