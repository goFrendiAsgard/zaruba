package runner

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/monitor"
)

func getRunner(t *testing.T, configFile string, taskNames []string, statusInterval time.Duration) (runner *Runner, err error) {
	decoration := monitor.NewDecoration()
	logger := monitor.NewConsoleLogger(decoration)
	dir := os.ExpandEnv(filepath.Dir(configFile))
	logFile := filepath.Join(dir, "log.zaruba.csv")
	csvLogger := monitor.NewCSVLogWriter(logFile)
	project, err := config.NewProject(logger, csvLogger, decoration, configFile)
	if err != nil {
		t.Error(err)
		return runner, err
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return runner, err
	}
	return NewRunner(logger, decoration, project, taskNames, statusInterval)
}

func readAlembicTxt(t *testing.T) (content string, err error) {
	data, err := ioutil.ReadFile("../test_resource/alchemy/alembic.txt")
	if err != nil {
		t.Error(err)
		return content, err
	}
	return string(data), err
}

func TestValidRunnerAlchemy(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/alchemy/zaruba.yaml", []string{"makeEverything"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err != nil {
		t.Error(err)
		return
	}
	content, err := readAlembicTxt(t)
	if err != nil {
		t.Error(err)
		return
	}
	rows := strings.Split(content, "\n")
	expectations := []string{"NaOH", "HCl", "NaCl", "H2O", "Hot Water", ""}
	if len(rows) != len(expectations) {
		t.Errorf("Rows doesn't have same element as expectations, rows: %#v, expectations: %#v", rows, expectations)
		return
	}
	for index, expectation := range expectations {
		actual := rows[index]
		if actual != expectation {
			t.Errorf("Rows[%d] doesn't meet expectation. Expected: %s, actual: %s", index, expectation, actual)
		}
	}
}

func TestTerminateRunnerLongPreparationBeforeComplete(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/longProcess.yaml", []string{"longProcess"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	startTime := time.Now()
	ch := make(chan error)
	go func() {
		ch <- runner.Run()
	}()
	time.Sleep(1 * time.Second)
	runner.Terminate()
	err = <-ch
	if err == nil {
		t.Errorf("Error expected")
	}
	elapsed := time.Since(startTime)
	if elapsed > 3*time.Second {
		t.Errorf("Process should be ended in approximately 1 second, but currently it need %s", elapsed)
	}
}

func TestTerminateRunnerLongProcessBeforeComplete(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/longProcess.yaml", []string{"longProcess"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	startTime := time.Now()
	ch := make(chan error)
	go func() {
		ch <- runner.Run()
	}()
	time.Sleep(6 * time.Second)
	runner.Terminate()
	err = <-ch
	if err == nil {
		t.Errorf("Error expected")
	}
	elapsed := time.Since(startTime)
	if elapsed > 9*time.Second {
		t.Errorf("Process should be ended in approximately 6 second, but currently it need %s", elapsed)
	}
}

func TestTerminateRunnerLongProcessAfterComplete(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/longProcess.yaml", []string{"longProcess"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	startTime := time.Now()
	ch := make(chan error)
	go func() {
		ch <- runner.Run()
	}()
	time.Sleep(11 * time.Second)
	runner.Terminate()
	err = <-ch
	if err == nil {
		t.Error(err)
	}
	elapsed := time.Since(startTime)
	if elapsed > 14*time.Second {
		t.Errorf("Process should be ended in approximately 11 second, but currently it need %s", elapsed)
	}
}

func TestInvalidLongProcessErrorAfterCheck(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/invalidLongProcessErrorAfterCheck.yaml", []string{"longProcess"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidLongProcessExitedBeforeCheck(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/invalidLongProcessExitedBeforeCheck.yaml", []string{"longProcess"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidNonExistingTask(t *testing.T) {
	_, err := getRunner(t, "../test_resource/alchemy/zaruba.yaml", []string{"makeElixirOfImmortality"}, time.Second)
	if err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidBrokenCommand(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/invalidTaskCommand.yaml", []string{"brokenCommand"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidBrokenProcessStart(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/invalidTaskCommand.yaml", []string{"brokenProcessStart"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidBrokenProcessCheck(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/invalidTaskCommand.yaml", []string{"brokenProcessCheck"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidNonExecutableCommand(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/invalidTaskCommand.yaml", []string{"nonExecutableCommand"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidNonExecutableProcessStart(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/invalidTaskCommand.yaml", []string{"nonExecutableProcessStart"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidNonExecutableProcessCheck(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/invalidTaskCommand.yaml", []string{"nonExecutableProcessCheck"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTimeoutCheck(t *testing.T) {
	runner, err := getRunner(t, "../test_resource/invalidTimeout.yaml", []string{"timeoutTask"}, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}
