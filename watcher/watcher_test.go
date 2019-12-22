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
	time.Sleep(time.Millisecond * 500)
	os.Create(filepath.Join(testPath, "lib/b/watchTrigger.txt"))
	time.Sleep(time.Millisecond * 500)
	stopChan <- true
	err := <-errChan
	if err != nil {
		t.Errorf("[ERROR] Cannot watch: %s", err)
	}

	// a.txt
	textFilePath := filepath.Join(testPath, "service/d/controller/c/lib/a/a.txt")
	if _, err := os.Stat(textFilePath); err == nil || !os.IsNotExist(err) {
		t.Errorf("[UNEXPECTED] `a` should not be created because the trigger is on `b` service")
	}

	// b.txt
	textFilePath = filepath.Join(testPath, "service/d/controller/c/lib/b/b.txt")
	contentB, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content := strings.Trim(string(contentB), "\n")
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
