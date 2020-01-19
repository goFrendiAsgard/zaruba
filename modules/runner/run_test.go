package runner

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/file"
	"github.com/state-alchemists/zaruba/modules/organizer"
)

func TestRun(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testRun")
	if err := file.Copy("../../test-resource/project", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}
	if err := file.Copy("../../test-resource/zaruba.config.yaml", filepath.Join(testPath, "zaruba.config.yaml")); err != nil {
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
	time.Sleep(time.Second * 10)

	res, err := http.Get("http://localhost:3000/go/frendi")
	if err != nil {
		t.Errorf("[ERROR] Cannot send request: %s", err)
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("[ERROR] Cannot parse response: %s", err)
	}
	expected := "Hello go frendi"
	actual := string(content)
	if actual != expected {
		t.Errorf("[UNEXPECTED] expecting response to be `%s`, get: %s", expected, actual)
	}

	res, err = http.Get("http://localhost:3000")
	if err != nil {
		t.Errorf("[ERROR] Cannot send request: %s", err)
	}
	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("[ERROR] Cannot parse response: %s", err)
	}
	expected = "Hello world"
	actual = string(content)
	if actual != expected {
		t.Errorf("[UNEXPECTED] expecting response to be `%s`, get: %s", expected, actual)
	}

	// test done
	stopChan <- true
	err = <-errChan
	if err != nil {
		t.Errorf("[ERROR] Error while running: %s", err)
	}
}
