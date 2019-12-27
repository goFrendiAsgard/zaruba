package action

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/file"
)

func TestDo(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testDo")
	if err := file.Copy("../test-resource/testDo.template", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	// Do action should succeed
	err := Do("doWrite", NewOption().SetWorkDir(testPath))
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
