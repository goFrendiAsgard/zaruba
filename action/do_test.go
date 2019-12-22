package action

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/otiai10/copy"
	"github.com/state-alchemists/zaruba/config"
)

func TestDo(t *testing.T) {
	baseTestPath := config.GetTemplateDir()
	testPath := filepath.Join(baseTestPath, "testDo")
	if err := copy.Copy("../test-resource/testDo.template", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	// Do action should succeed
	err := Do("doWrite", testPath, NewOption())
	if err != nil {
		t.Errorf("[ERROR] Cannot do action: %s", err)
		return
	}

	// After do action, a.txt should be exists and contains "alpha"
	contentB, err := ioutil.ReadFile(filepath.Join(testPath, "a.txt"))
	if err != nil {
		t.Errorf("[ERROR] Cannot read a.txt: %s", err)
	}
	content := strings.Trim(string(contentB), "\n")
	if content != "alpha" {
		t.Errorf("[UNEXPECTED] content should be `alpha`: %s", content)
	}

	// After do action, b.txt should be exists and contains "beta"
	contentB, err = ioutil.ReadFile(filepath.Join(testPath, "subdir/b.txt"))
	if err != nil {
		t.Errorf("[ERROR] Cannot read subdir/b.txt: %s", err)
	}
	content = strings.Trim(string(contentB), "\n")
	if content != "beta" {
		t.Errorf("[UNEXPECTED] content should be `beta`: %s", content)
	}

}
