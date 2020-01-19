package organizer

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/file"
)

func TestOrganize(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testOrganize")
	if err := file.Copy("../../test-resource/project", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}
	if err := file.Copy("../../test-resource/zaruba.config.megazord.yaml", filepath.Join(testPath, "zaruba.config.yaml")); err != nil {
		t.Errorf("[ERROR] Cannot copy zaruba.config.yaml: %s", err)
		return
	}

	// Organize project should succeed
	err := Organize(testPath, NewOption())
	if err != nil {
		t.Errorf("[ERROR] Cannot organize: %s", err)
	}

	// `log.txt` should contains `pre-organize`, `organize`, and `post-organize`
	testFilePath := filepath.Join(testPath, "log.txt")
	testByteContent, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", testFilePath, err)
	} else {
		testContent := strings.Trim(string(testByteContent), "\n")
		expected := "pre-organize\norganize\npost-organize"
		if testContent != expected {
			t.Errorf("[UNEXPECTED] content should be `%s`, get: `%s`", expected, testContent)
		}
	}

	// `megazord/gateway/greeting-lib/changelog.md` should contains `init`
	testFilePath = filepath.Join(testPath, "megazord/gateway/greeting-lib/parent-changelog.md")
	testByteContent, err = ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", testFilePath, err)
	} else {
		testContent := strings.Trim(string(testByteContent), "\n")
		expected := "init"
		if testContent != expected {
			t.Errorf("[UNEXPECTED] content should be `%s`, get: `%s`", expected, testContent)
		}
	}

}
