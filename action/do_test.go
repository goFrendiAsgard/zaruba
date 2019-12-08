package action

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestDo(t *testing.T) {

	// Create component should succeed
	err := Do("../test-playground/testDo", "doWrite")
	if err != nil {
		t.Errorf("[ERROR] Cannot do action: %s", err)
	}

	// After do action, a.txt should be exists and contains "alpha"
	contentB, err := ioutil.ReadFile("../test-playground/testDo/a.txt")
	if err != nil {
		t.Errorf("[ERROR] Cannot read ../test-playground/testDo/a.txt: %s", err)
	}
	content := strings.Trim(string(contentB), "\n")
	if content != "alpha" {
		t.Errorf("[UNEXPECTED] content should be `alpha`: %s", content)
	}

	// After do action, b.txt should be exists and contains "beta"
	contentB, err = ioutil.ReadFile("../test-playground/testDo/subdir/b.txt")
	if err != nil {
		t.Errorf("[ERROR] Cannot read ../test-playground/testDo/subdir/b.txt: %s", err)
	}
	content = strings.Trim(string(contentB), "\n")
	if content != "beta" {
		t.Errorf("[UNEXPECTED] content should be `beta`: %s", content)
	}

}
