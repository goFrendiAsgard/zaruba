package config

import (
	"strings"
	"testing"
)

func TestTaskGetCheckCmdPatterns(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithCheckCmd"]
	expectedList := []string{"sleep", "2"}
	actualList, exist, err := task.GetCheckCmdPatterns()
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Errorf("pattern should be exist")
	}
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

func TestTaskGetCheckCmdPatternsFromTaskWhichParentHasCheckCmd(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWhichParentHasCheckCmd"]
	expectedList := []string{"sleep", "1"}
	actualList, exist, err := task.GetCheckCmdPatterns()
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Errorf("pattern should be exist")
	}
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

func TestTaskGetCheckCmdPatternsFromTaskWithoutCheckCmd(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithoutCheckCmd"]
	expectedList := []string{}
	actualList, exist, err := task.GetCheckCmdPatterns()
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.Contains(errorMessage, "cannot retrieve check cmd from any parent task of taskWithoutCheckCmd") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
	if exist {
		t.Errorf("pattern should not be exist")
	}
	if len(actualList) != len(expectedList) {
		t.Errorf("expected: %#v, actual %#v", expectedList, actualList)
		return
	}
}

func TestTaskGetCheckCmd(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithCheckCmd"]
	cmd, exist, err := task.GetCheckCmd(make(chan error))
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Errorf("cmd should be exist")
	}
	if cmd == nil {
		t.Errorf("cmd is nil")
	}
}

func TestTaskGetCheckCmdFromTaskWhichParentHasCheckCmd(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWhichParentHasCheckCmd"]
	cmd, exist, err := task.GetCheckCmd(make(chan error))
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Errorf("cmd should be exist")
	}
	if cmd == nil {
		t.Errorf("cmd is nil")
	}
}

func TestTaskGetCheckCmdFromTaskWithoutCheckCmd(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithoutCheckCmd"]
	cmd, exist, err := task.GetCheckCmd(make(chan error))
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.Contains(errorMessage, "cannot retrieve check cmd from any parent task of taskWithoutCheckCmd") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
	if exist {
		t.Errorf("cmd should not be exist")
	}
	if cmd != nil {
		t.Errorf("cmd is not nil")
	}
}
