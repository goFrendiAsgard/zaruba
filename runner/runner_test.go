package runner

import (
	"testing"

	"github.com/state-alchemists/zaruba/output"
)

func TestRunnerNotInitializedProject(t *testing.T) {
	project, logger, decoration, err := getProject("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
	}
	if _, err = NewRunner(logger, decoration, project, []string{}, "10s", false, "10s"); err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "cannot create runner because project was not initialized" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRunnerInexistTask(t *testing.T) {
	project, logger, decoration, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if _, err = NewRunner(logger, decoration, project, []string{"inexist"}, "10s", false, "10s"); err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "task 'inexist' is not exist" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRunnerInvalidStatusInterval(t *testing.T) {
	project, logger, decoration, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if _, err = NewRunner(logger, decoration, project, []string{}, "invalid", false, "10s"); err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "cannot parse statusInterval 'invalid': time: invalid duration invalid" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRunnerInvalidAutoTerminateInterval(t *testing.T) {
	project, logger, decoration, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if _, err = NewRunner(logger, decoration, project, []string{}, "10s", false, "invalid"); err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "cannot parse autoTerminateDelay 'invalid': time: invalid duration invalid" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRunnerServeSalineWater(t *testing.T) {
	project, logger, decoration, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, err := NewRunner(logger, decoration, project, []string{"serveSalineWater"}, "10s", false, "10s")
	if err != nil {
		t.Error(err)
		return
	}
	mockLogger := logger.(*output.MockLogger)
	if err = runner.Run(); err != nil {
		t.Error(err)
		return
	}
	output := mockLogger.GetOutput()
	makeNaOHIndex := mockLogger.GetLineIndex("making NaOH")
	makeHClIndex := mockLogger.GetLineIndex("making HCl")
	makeNaClIndex := mockLogger.GetLineIndex("making NaCl")
	makeH2OIndex := mockLogger.GetLineIndex("making H2O")
	serveSalineWaterIndex := mockLogger.GetLineIndex("serve saline water")
	if makeNaOHIndex >= makeNaClIndex || makeHClIndex >= makeNaClIndex {
		t.Errorf("expect NaOH and HCL created befor NaCL, actual: \n%s", output)
	}
	if makeNaOHIndex >= makeH2OIndex || makeHClIndex >= makeH2OIndex {
		t.Errorf("expect NaOH and HCL created befor H2O, actual: \n%s", output)
	}
	if makeNaClIndex >= serveSalineWaterIndex || makeH2OIndex >= serveSalineWaterIndex {
		t.Errorf("expect NaOH and HCL created befor H2O, actual: \n%s", output)
	}
}

func TestRunnerServeSalineWaterAutoTerminate(t *testing.T) {
	project, logger, decoration, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, err := NewRunner(logger, decoration, project, []string{"serveSalineWater"}, "10s", true, "2s")
	if err != nil {
		t.Error(err)
		return
	}
	mockLogger := logger.(*output.MockLogger)
	if err = runner.Run(); err != nil {
		t.Error(err)
		return
	}
	output := mockLogger.GetOutput()
	makeNaOHIndex := mockLogger.GetLineIndex("making NaOH")
	makeHClIndex := mockLogger.GetLineIndex("making HCl")
	makeNaClIndex := mockLogger.GetLineIndex("making NaCl")
	makeH2OIndex := mockLogger.GetLineIndex("making H2O")
	serveSalineWaterIndex := mockLogger.GetLineIndex("serve saline water")
	if makeNaOHIndex >= makeNaClIndex || makeHClIndex >= makeNaClIndex {
		t.Errorf("expect NaOH and HCL created befor NaCL, actual: \n%s", output)
	}
	if makeNaOHIndex >= makeH2OIndex || makeHClIndex >= makeH2OIndex {
		t.Errorf("expect NaOH and HCL created befor H2O, actual: \n%s", output)
	}
	if makeNaClIndex >= serveSalineWaterIndex || makeH2OIndex >= serveSalineWaterIndex {
		t.Errorf("expect NaOH and HCL created befor H2O, actual: \n%s", output)
	}
}

func TestRunnerServeSalineWaterForceTerminate(t *testing.T) {
	project, logger, decoration, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, err := NewRunner(logger, decoration, project, []string{"serveSalineWater"}, "10s", false, "10s")
	if err != nil {
		t.Error(err)
		return
	}
	mockLogger := logger.(*output.MockLogger)
	mockLogger.RegisterTrigger("Job Complete", func() {
		runner.Terminate()
	})
	if err = runner.Run(); err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "terminated" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}
