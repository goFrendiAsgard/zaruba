package runner

import (
	"fmt"
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
	if err := component.Create("go-service", testPath, "alpha"); err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}
	if err := component.Create("nodejs-service", testPath, "beta"); err != nil {
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

	testRequest(t, 3011, "hello/Tony", "Hello Tony")
	testRequest(t, 3012, "hello/Tony", "Hello Tony")
	testRequest(t, 3011, "hello-rpc/Tony", "Hello Tony")
	testRequest(t, 3012, "hello-rpc/Tony", "Hello Tony")

	// test done
	stopChan <- true
	err = <-errChan
	if err != nil {
		t.Errorf("[ERROR] Error while running: %s", err)
	}
}

func testRequest(t *testing.T, port int, url, expected string) {
	res, err := http.Get(fmt.Sprintf("http://localhost:%d/%s", port, url))
	if err != nil {
		t.Errorf("[ERROR] Cannot send request: %s", err)
	} else {
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("[ERROR] Cannot parse response: %s", err)
		}
		actual := string(content)
		if actual != expected {
			t.Errorf("[UNEXPECTED] expecting response to be `%s`, get: %s", expected, actual)
		}
	}
}
