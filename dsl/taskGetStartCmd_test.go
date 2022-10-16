package dsl

import (
	"strings"
	"testing"
)

func TestTaskGetStartCmdPatterns(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getStartCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithStartCmd"]
	expectedList := []string{"sleep", "2"}
	actualList, err := task.GetStartCmdPatterns()
	if err != nil {
		t.Error(err)
	}
	if !task.IsHavingStartCmd() {
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

func TestTaskGetStartCmdPatternsFromTaskWhichParentHasStartCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getStartCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWhichParentHasStartCmd"]
	expectedList := []string{"sleep", "1"}
	actualList, err := task.GetStartCmdPatterns()
	if err != nil {
		t.Error(err)
	}
	if !task.IsHavingStartCmd() {
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

func TestTaskGetStartCmdPatternsFromTaskWithoutStartCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getStartCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithoutStartCmd"]
	expectedList := []string{}
	actualList, err := task.GetStartCmdPatterns()
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.Contains(errorMessage, "cannot retrieve start cmd from any parent task of taskWithoutStartCmd") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
	if task.IsHavingStartCmd() {
		t.Errorf("pattern should not be exist")
	}
	if len(actualList) != len(expectedList) {
		t.Errorf("expected: %#v, actual %#v", expectedList, actualList)
		return
	}
}

func TestTaskGetStartCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getStartCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithStartCmd"]
	cmd, err := task.GetStartCmd()
	if err != nil {
		t.Error(err)
	}
	if !task.IsHavingStartCmd() {
		t.Errorf("cmd should be exist")
	}
	if cmd == nil {
		t.Errorf("cmd is nil")
	}
}

func TestTaskGetStartCmdFromTaskWhichParentHasStartCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getStartCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWhichParentHasStartCmd"]
	cmd, err := task.GetStartCmd()
	if err != nil {
		t.Error(err)
	}
	if !task.IsHavingStartCmd() {
		t.Errorf("cmd should be exist")
	}
	if cmd == nil {
		t.Errorf("cmd is nil")
	}
}

func TestTaskGetStartCmdFromTaskWithoutStartCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getStartCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithoutStartCmd"]
	cmd, err := task.GetStartCmd()
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.Contains(errorMessage, "cannot retrieve start cmd from any parent task of taskWithoutStartCmd") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
	if task.IsHavingStartCmd() {
		t.Errorf("cmd should not be exist")
	}
	if cmd != nil {
		t.Errorf("cmd is not nil")
	}
}
