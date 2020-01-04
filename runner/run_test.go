package runner

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/organizer"
)

func TestRun(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testRun")
	if err := file.Copy("../test-resource/project", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}
	if err := file.Copy("../test-resource/zaruba.config.yaml", filepath.Join(testPath, "zaruba.config.yaml")); err != nil {
		t.Errorf("[ERROR] Cannot copy zaruba.config.yaml: %s", err)
		return
	}

	// Organize project should succeed
	err := organizer.Organize(testPath, organizer.NewOption())
	if err != nil {
		t.Errorf("[ERROR] Cannot organize: %s", err)
	}

	stopChan := make(chan bool)
	errChan := make(chan error)
	executedChan := make(chan bool)
	go Run(testPath, stopChan, executedChan, errChan)
	<-executedChan
	time.Sleep(time.Millisecond * 500)

	// test done
	stopChan <- true
	err = <-errChan
	if err != nil {
		t.Errorf("[ERROR] Error while running: %s", err)
	}
}
