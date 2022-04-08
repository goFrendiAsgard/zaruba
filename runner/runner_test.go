package runner

import (
	"strings"
	"testing"
)

func TestRunnerNotInitializedProject(t *testing.T) {
	project, err := getProject("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
	}
	_, _, _, err = getRunner(project, []string{}, "10s", false, "10s")
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "cannot create runner because project was not initialized" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRunnerInexistTask(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	_, _, _, err = getRunner(project, []string{"inexist"}, "10s", false, "10s")
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "task 'inexist' is not exist" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRunnerInvalidStatusInterval(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	_, _, _, err = getRunner(project, []string{}, "invalid", false, "10s")
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "cannot parse statusInterval 'invalid': time: invalid duration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRunnerInvalidAutoTerminateInterval(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	_, _, _, err = getRunner(project, []string{}, "10s", false, "invalid")
	if err == nil {
		t.Error("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "cannot parse autoTerminateDelay 'invalid': time: invalid duration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRunnerServeSalineWater(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, logger, _, err := getRunner(project, []string{"serveSalineWater"}, "10s", true, "2s")
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err != nil {
		t.Error(err)
		return
	}
	output := logger.GetOutput()
	makeNaOHIndex := logger.GetLineIndex("making NaOH")
	makeHClIndex := logger.GetLineIndex("making HCl")
	makeNaClIndex := logger.GetLineIndex("making NaCl")
	makeH2OIndex := logger.GetLineIndex("making H2O")
	serveSalineWaterIndex := logger.GetLineIndex("serve saline water")
	if makeNaOHIndex >= makeNaClIndex || makeHClIndex >= makeNaClIndex {
		t.Errorf("expect NaOH and HCL created before NaCL, actual: \n%s", output)
	}
	if makeNaOHIndex >= makeH2OIndex || makeHClIndex >= makeH2OIndex {
		t.Errorf("expect NaOH and HCL created before H2O, actual: \n%s", output)
	}
	if makeNaClIndex >= serveSalineWaterIndex || makeH2OIndex >= serveSalineWaterIndex {
		t.Errorf("expect NaOH and HCL created before H2O, actual: \n%s", output)
	}
}

func TestRunnerServeSalineWaterAutoTerminate(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, logger, _, err := getRunner(project, []string{"serveSalineWater"}, "10s", true, "2s")
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err != nil {
		t.Error(err)
		return
	}
	output := logger.GetOutput()
	makeNaOHIndex := logger.GetLineIndex("making NaOH")
	makeHClIndex := logger.GetLineIndex("making HCl")
	makeNaClIndex := logger.GetLineIndex("making NaCl")
	makeH2OIndex := logger.GetLineIndex("making H2O")
	serveSalineWaterIndex := logger.GetLineIndex("serve saline water")
	if makeNaOHIndex >= makeNaClIndex || makeHClIndex >= makeNaClIndex {
		t.Errorf("expect NaOH and HCL created before NaCL, actual: \n%s", output)
	}
	if makeNaOHIndex >= makeH2OIndex || makeHClIndex >= makeH2OIndex {
		t.Errorf("expect NaOH and HCL created before H2O, actual: \n%s", output)
	}
	if makeNaClIndex >= serveSalineWaterIndex || makeH2OIndex >= serveSalineWaterIndex {
		t.Errorf("expect NaOH and HCL created before H2O, actual: \n%s", output)
	}
}

func TestRunnerServeSalineWaterForceTerminate(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, logger, _, err := getRunner(project, []string{"serveSalineWater"}, "10s", false, "10s")
	if err != nil {
		t.Error(err)
		return
	}
	logger.RegisterTrigger("Job Complete", func() {
		runner.Terminate()
	})
	if err = runner.Run(); err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "Terminated" && !strings.Contains(errorMessage, "interrupt") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRunnerMakeAll(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, _, _, err := getRunner(project, []string{"makeAll"}, "10s", false, "10s")
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("error expected")
		return
	}
}

func TestRunnerServeSalineWaterFailBeforeCheck(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, _, _, err := getRunner(project, []string{"serveSalineWater", "serveSalineWaterFailBeforeCheck"}, "10s", false, "10s")
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("error expected")
		return
	}
}

func TestRunnerServeSalineWaterFailAfterCheck(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, _, _, err := getRunner(project, []string{"serveSalineWaterFailAfterCheck"}, "10s", false, "10s")
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("error expected")
		return
	}
}

func TestRunnerWaitGovernmentApproval(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/runner/alchemy.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	runner, _, _, err := getRunner(project, []string{"waitGovernmentApproval"}, "10s", false, "10s")
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("error expected")
		return
	}
}
