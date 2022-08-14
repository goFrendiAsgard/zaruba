package dsl

import (
	"strings"
	"testing"
)

func TestTaskGetCheckCmdPatterns(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithCheckCmd"]
	expectedList := []string{"sleep", "2"}
	actualList, err := task.GetCheckCmdPatterns()
	if err != nil {
		t.Error(err)
	}
	if !task.IsHavingCheckCmd() {
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
	project, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWhichParentHasCheckCmd"]
	expectedList := []string{"sleep", "1"}
	actualList, err := task.GetCheckCmdPatterns()
	if err != nil {
		t.Error(err)
	}
	if !task.IsHavingCheckCmd() {
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
	project, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithoutCheckCmd"]
	expectedList := []string{}
	actualList, err := task.GetCheckCmdPatterns()
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.Contains(errorMessage, "cannot retrieve check cmd from any parent task of taskWithoutCheckCmd") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
	if task.IsHavingCheckCmd() {
		t.Errorf("pattern should not be exist")
	}
	if len(actualList) != len(expectedList) {
		t.Errorf("expected: %#v, actual %#v", expectedList, actualList)
		return
	}
}

func TestTaskGetCheckCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithCheckCmd"]
	cmd, err := task.GetCheckCmd()
	if err != nil {
		t.Error(err)
	}
	if !task.IsHavingCheckCmd() {
		t.Errorf("cmd should be exist")
	}
	if cmd == nil {
		t.Errorf("cmd is nil")
	}
}

func TestTaskGetCheckCmdFromTaskWhichParentHasCheckCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWhichParentHasCheckCmd"]
	cmd, err := task.GetCheckCmd()
	if err != nil {
		t.Error(err)
	}
	if !task.IsHavingCheckCmd() {
		t.Errorf("cmd should be exist")
	}
	if cmd == nil {
		t.Errorf("cmd is nil")
	}
}

func TestTaskGetCheckCmdFromTaskWithoutCheckCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCheckCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithoutCheckCmd"]
	cmd, err := task.GetCheckCmd()
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.Contains(errorMessage, "cannot retrieve check cmd from any parent task of taskWithoutCheckCmd") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
	if task.IsHavingCheckCmd() {
		t.Errorf("cmd should not be exist")
	}
	if cmd != nil {
		t.Errorf("cmd is not nil")
	}
}
