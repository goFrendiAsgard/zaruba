package runner

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/state-alchemists/zaruba/config"
)

func createConf(t *testing.T, path string) (conf *config.ProjectConfig, err error) {
	conf, err = config.NewConfig(path)
	if err != nil {
		t.Error(err)
		return conf, err
	}
	if err = conf.Init(); err != nil {
		t.Error(err)
	}
	return conf, err
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
	conf, err := createConf(t, "../test_resource/alchemy/zaruba.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"makeEverything"}, time.Second)
	if err = runner.Run(); err != nil {
		t.Error(err)
	}
	content, err := readAlembicTxt(t)
	if err != nil {
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
	conf, err := createConf(t, "../test_resource/longProcess.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"longProcess"}, time.Second)
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
	conf, err := createConf(t, "../test_resource/longProcess.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"longProcess"}, time.Second)
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
	if elapsed > 8*time.Second {
		t.Errorf("Process should be ended in approximately 6 second, but currently it need %s", elapsed)
	}
}

func TestTerminateRunnerLongProcessAfterComplete(t *testing.T) {
	conf, err := createConf(t, "../test_resource/longProcess.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"longProcess"}, time.Second)
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
	if elapsed > 13*time.Second {
		t.Errorf("Process should be ended in approximately 11 second, but currently it need %s", elapsed)
	}
}

func TestInvalidLongProcessErrorAfterCheck(t *testing.T) {
	conf, err := createConf(t, "../test_resource/invalidLongProcessErrorAfterCheck.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"longProcess"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidLongProcessExitedBeforeCheck(t *testing.T) {
	conf, err := createConf(t, "../test_resource/invalidLongProcessExitedBeforeCheck.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"longProcess"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidNonExistingTask(t *testing.T) {
	conf, err := createConf(t, "../test_resource/alchemy/zaruba.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"makeElixirOfImmortality"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidBrokenCommand(t *testing.T) {
	conf, err := createConf(t, "../test_resource/invalidTaskCommand.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"brokenCommand"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidBrokenProcessStart(t *testing.T) {
	conf, err := createConf(t, "../test_resource/invalidTaskCommand.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"brokenProcessStart"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidBrokenProcessCheck(t *testing.T) {
	conf, err := createConf(t, "../test_resource/invalidTaskCommand.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"brokenProcessCheck"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidNonExecutableCommand(t *testing.T) {
	conf, err := createConf(t, "../test_resource/invalidTaskCommand.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"nonExecutableCommand"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidNonExecutableProcessStart(t *testing.T) {
	conf, err := createConf(t, "../test_resource/invalidTaskCommand.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"nonExecutableProcessStart"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidNonExecutableProcessCheck(t *testing.T) {
	conf, err := createConf(t, "../test_resource/invalidTaskCommand.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"nonExecutableProcessCheck"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTimeoutCheck(t *testing.T) {
	conf, err := createConf(t, "../test_resource/invalidTimeout.yaml")
	if err != nil {
		return
	}
	runner := NewRunner(conf, []string{"timeoutTask"}, time.Second)
	if err = runner.Run(); err == nil {
		t.Errorf("Error expected")
	}
}
