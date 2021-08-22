package config

import (
	"strings"
	"testing"
)

func TestTaskGetCmdFromTaskWithEnv(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithEnv"]
	cmd, exist, err := task.GetStartCmd()
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Errorf("cmd should be exist")
	}
	if cmd == nil {
		t.Errorf("cmd is nil")
		return
	}
	envFound := false
	for _, env := range cmd.Env {
		if env == "KEY=VALUE" {
			envFound = true
			break
		}
	}
	if !envFound {
		t.Errorf("env KEY=VALUE not found in %#v", cmd.Env)
	}
}

func TestTaskGetCmdFromTaskWithBrokenEnv(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithBrokenEnv"]
	_, exist, err := task.GetStartCmd()
	if err == nil {
		t.Error("error expected")
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "template:") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
	if !exist {
		t.Errorf("cmd should be exist")
	}
}

func TestTaskGetCmdFromTaskWithBrokenCmd(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCmd.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskWithBrokenCmd"]
	_, exist, err := task.GetStartCmd()
	if err == nil {
		t.Error("error expected")
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "template:") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
	if !exist {
		t.Errorf("cmd should be exist")
	}
}
