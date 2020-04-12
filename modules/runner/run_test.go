package runner

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/modules/component"
	"github.com/state-alchemists/zaruba/modules/config"
)

func TestRun(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testRun")

	// create project and service
	if err := component.Create("project", testPath); err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}
	if err := component.Create("go-service", testPath, "gopher"); err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}
	// load project config
	p, err := config.NewProjectConfig(testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot load config: %s", err)
		return
	}

	stopChan := make(chan bool)
	errChan := make(chan error)
	executedChan := make(chan bool)
	go Run(testPath, p, []string{}, stopChan, executedChan, errChan)
	<-executedChan

	res, err := http.Get("http://localhost:3011/hello/Tony")
	if err != nil {
		t.Errorf("[ERROR] Cannot send request: %s", err)
	} else {
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("[ERROR] Cannot parse response: %s", err)
		}
		expected := "Hello Tony"
		actual := string(content)
		if actual != expected {
			t.Errorf("[UNEXPECTED] expecting response to be `%s`, get: %s", expected, actual)
		}
	}

	res, err = http.Get("http://localhost:3011/hello-rpc/Tony")
	if err != nil {
		t.Errorf("[ERROR] Cannot send request: %s", err)
	} else {
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("[ERROR] Cannot parse response: %s", err)
		}
		expected := "Hello Tony"
		actual := string(content)
		if actual != expected {
			t.Errorf("[UNEXPECTED] expecting response to be `%s`, get: %s", expected, actual)
		}
	}

	// test done
	stopChan <- true
	err = <-errChan
	if err != nil {
		t.Errorf("[ERROR] Error while running: %s", err)
	}
}
