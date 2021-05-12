package config

import (
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
