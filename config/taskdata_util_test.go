package config

import (
	"path/filepath"
	"testing"
)

func TestTdGetWorkPath(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// absolute
	expected := "/home/gofrendi"
	actual := td.GetWorkPath("/home/gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected, _ = filepath.Abs("../test-resources/taskdata/util/location/gofrendi")
	actual = td.GetWorkPath("./gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdGetRelativePath(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// absolute
	expected := "/home/gofrendi"
	actual := td.GetRelativePath("/home/gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected, _ = filepath.Abs("../test-resources/taskdata/util/zaruba-tasks/gofrendi")
	actual = td.GetRelativePath("./gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdGetTaskExist(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	otherTask, err := td.GetTask("otherTaskName")
	if err != nil {
		t.Error(err)
		return
	}
	if otherTask == nil {
		t.Errorf("otherTask is nil")
	}
}

func TestTdGetTaskInExist(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	otherTask, err := td.GetTask("inexistTask")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	if otherTask != nil {
		t.Errorf("otherTask is not nil")
	}
}
