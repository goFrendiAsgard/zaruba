package runner

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
	"time"

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
	if err := component.Create("mysql", testPath); err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}
	if err := component.Create("redis", testPath); err != nil {
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
	if err := component.Create("python-service", testPath, "gamma"); err != nil {
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

	for _, port := range []int{3011, 3012, 3013} {
		testRequest(t, port, "hello/Tony", "Hello Tony")
		testRequest(t, port, "hello-rpc/Tony", "Hello Tony")
		testRequest(t, port, "hello-all", "Hello everyone !!!")
		testRequest(t, port, "hello-pub/Dono", "Message sent")
		time.Sleep(1 * time.Second) // Consuming is somehow slow in node.js
		testRequest(t, port, "hello-all", "Hello Dono, and everyone")
	}

	// test done
	stopChan <- true
	err = <-errChan
	if err != nil {
		t.Errorf("[ERROR] Error while running: %s", err)
	}

	if err = StopContainers(testPath, p); err != nil {
		t.Errorf("[ERROR] Cannot stop containers: %s", err)
	}

	if err = RemoveContainers(testPath, p); err != nil {
		t.Errorf("[ERROR] Cannot remove containers: %s", err)
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
