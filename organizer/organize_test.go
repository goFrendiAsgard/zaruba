package organizer

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestOrganize(t *testing.T) {
	// Organize project should succeed
	err := Organize("../test-playground/testOrganize")
	if err != nil {
		t.Errorf("[ERROR] Cannot organize: %s", err)
	}

	// a.txt
	textFilePath := "../test-playground/testOrganize/service/d/controller/c/lib/a/a.txt"
	contentB, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content := strings.Trim(string(contentB), "\n")
	if content != "a" {
		t.Errorf("[UNEXPECTED] content should be `a`: %s", content)
	}

	// b.txt
	textFilePath = "../test-playground/testOrganize/service/d/controller/c/lib/b/b.txt"
	contentB, err = ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content = strings.Trim(string(contentB), "\n")
	if content != "b" {
		t.Errorf("[UNEXPECTED] content should be `b`: %s", content)
	}

	// c.txt
	textFilePath = "../test-playground/testOrganize/service/d/controller/c/c.txt"
	contentB, err = ioutil.ReadFile(textFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", textFilePath, err)
	}
	content = strings.Trim(string(contentB), "\n")
	if content != "c" {
		t.Errorf("[UNEXPECTED] content should be `c`: %s", content)
	}

}
