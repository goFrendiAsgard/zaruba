package config

import (
	"strings"
	"testing"

	"github.com/state-alchemists/zaruba/output"
)

func TestTaskGetCmdLog(t *testing.T) {
	project, logger, _, err := getProjectAndInit("../test-resources/task/getCmdLog.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	logDone := make(chan error)
	cmd, exist, err := task.GetStartCmd(logDone)
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
	if err = cmd.Run(); err != nil {
		t.Error(err)
		return
	}
	err = <-logDone
	if err != nil {
		t.Error(err)
		return
	}
	outputFound := false
	mockLogger := logger.(*output.MockLogger)
	for _, logData := range mockLogger.Data {
		if strings.Contains(logData.Str, "hello world") {
			outputFound = true
			break
		}
	}
	if !outputFound {
		t.Errorf("expected output not found: %#v", mockLogger.Data)
	}
}

func TestTaskGetCmdLogWithInvalidRecordLogger(t *testing.T) {
	decoration := output.NewDecoration()
	mockLogger := output.NewMockLogger()
	mockInvalidRecordLogger := output.NewMockInvalidRecordLogger()
	project, err := NewProject(mockLogger, mockInvalidRecordLogger, decoration, "../test-resources/task/getCmdLog.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	logDone := make(chan error)
	cmd, exist, err := task.GetStartCmd(logDone)
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
	if err = cmd.Run(); err != nil {
		t.Error(err)
		return
	}
	logErr := <-logDone
	if logErr == nil {
		t.Errorf("log error expected")
		return
	}
	logErrorMessage := logErr.Error()
	if logErrorMessage != "cannot write" {
		t.Errorf("invalid error message: %s", logErrorMessage)
	}
	outputFound := false
	for _, logData := range mockLogger.Data {
		if strings.Contains(logData.Str, "hello world") {
			outputFound = true
			break
		}
	}
	if !outputFound {
		t.Errorf("expected output not found: %#v", mockLogger.Data)
	}
}
