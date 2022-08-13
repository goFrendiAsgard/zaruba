package dsl

import (
	"strings"
	"testing"
)

func TestTaskGetCmdLog(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/task/getCmdLog.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
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
	if err = startCmd.Run(); err != nil {
		t.Error(err)
		return
	}
	outputFound := false
	output := <-project.StdoutChan
	if strings.Contains(output, "hello world") {
		outputFound = true
	}
	if !outputFound {
		t.Errorf("expected output not found: %#v", output)
	}
}
