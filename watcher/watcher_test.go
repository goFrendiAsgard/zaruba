package watcher

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/otiai10/copy"
	"github.com/state-alchemists/zaruba/config"
)

func TestWatch(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testWatch")
	if err := copy.Copy("../test-resource/testOrganize.template", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	// Watch project should succeed
	stopChan := make(chan bool)
	errChan := make(chan error)
	go Watch(testPath, errChan, stopChan)
	os.Create(filepath.Join(testPath, "lib/b/watchTrigger.txt"))
	time.Sleep(time.Second)
	stopChan <- true
	err := <-errChan
	if err != nil {
		t.Errorf("[ERROR] Cannot watch: %s", err)
	}

	// a.txt
	textFilePath := filepath.Join(testPath, "service/d/controller/c/lib/a/a.txt")
	contentB, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content := strings.Trim(string(contentB), "\n")
	if content != "a" {
		t.Errorf("[UNEXPECTED] content should be `a`: %s", content)
	}

	// b.txt
	textFilePath = filepath.Join(testPath, "service/d/controller/c/lib/b/b.txt")
	contentB, err = ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content = strings.Trim(string(contentB), "\n")
	if content != "b" {
		t.Errorf("[UNEXPECTED] content should be `b`: %s", content)
	}

	// c.txt
	textFilePath = filepath.Join(testPath, "service/d/controller/c/c.txt")
	contentB, err = ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content = strings.Trim(string(contentB), "\n")
	if content != "c" {
		t.Errorf("[UNEXPECTED] content should be `c`: %s", content)
	}

}
