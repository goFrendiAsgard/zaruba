package dsl

import (
	"strings"
	"testing"
)

func TestTaskGetCmdFromTaskWithConfig(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithConfig"]
	startCmdExist := task.IsHavingStartCmd()
	if !startCmdExist {
		t.Errorf("startCmd should be exist")
	}
	startCmd, err := task.GetStartCmd()
	if err != nil {
		t.Error(err)
	}
	if startCmd == nil {
		t.Errorf("cmd is nil")
		return
	}
	envFound := false
	for _, env := range startCmd.Env {
		if env == "ZARUBA_CONFIG_SOME_KEY=value" {
			envFound = true
			break
		}
	}
	if !envFound {
		t.Errorf("env ZARUBA_CONFIG_SOME_KEY=value not found in %#v", startCmd.Env)
	}
}

func TestTaskGetCmdFromTaskWithEnv(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithEnv"]
	startCmdExist := task.IsHavingStartCmd()
	if !startCmdExist {
		t.Errorf("startCmd should be exist")
	}
	startCmd, err := task.GetStartCmd()
	if err != nil {
		t.Error(err)
	}
	if startCmd == nil {
		t.Errorf("cmd is nil")
		return
	}
	envFound := false
	for _, env := range startCmd.Env {
		if env == "KEY=VALUE" {
			envFound = true
			break
		}
	}
	if !envFound {
		t.Errorf("env KEY=VALUE not found in %#v", startCmd.Env)
	}
}

func TestTaskGetCmdFromTaskWithBrokenEnv(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithBrokenEnv"]
	startCmdExist := task.IsHavingStartCmd()
	if !startCmdExist {
		t.Errorf("startCmd should be exist")
	}
	_, err = task.GetStartCmd()
	if err == nil {
		t.Error("error expected")
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "template:") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestTaskGetCmdFromTaskWithBrokenCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithBrokenCmd"]
	startCmdExist := task.IsHavingStartCmd()
	if !startCmdExist {
		t.Errorf("startCmd should be exist")
	}
	_, err = task.GetStartCmd()
	if err == nil {
		t.Error("error expected")
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "template:") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}
