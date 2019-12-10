package watcher

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestWatch(t *testing.T) {
	// Watch project should succeed
	stopChan := make(chan bool)
	errChan := make(chan error)
	go Watch("../test-playground/testWatch", errChan, stopChan)
	os.Create("../test-playground/testWatch/lib/b/watchTrigger.txt")
	time.Sleep(time.Second)
	stopChan <- true
	err := <-errChan
	if err != nil {
		t.Errorf("[ERROR] Cannot watch: %s", err)
	}

	// a.txt
	textFilePath := "../test-playground/testWatch/service/d/controller/c/lib/a/a.txt"
	contentB, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content := strings.Trim(string(contentB), "\n")
	if content != "a" {
		t.Errorf("[UNEXPECTED] content should be `a`: %s", content)
	}

	// b.txt
	textFilePath = "../test-playground/testWatch/service/d/controller/c/lib/b/b.txt"
	contentB, err = ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content = strings.Trim(string(contentB), "\n")
	if content != "b" {
		t.Errorf("[UNEXPECTED] content should be `b`: %s", content)
	}

	// c.txt
	textFilePath = "../test-playground/testWatch/service/d/controller/c/c.txt"
	contentB, err = ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content = strings.Trim(string(contentB), "\n")
	if content != "c" {
		t.Errorf("[UNEXPECTED] content should be `c`: %s", content)
	}

}
