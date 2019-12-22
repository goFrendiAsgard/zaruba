package organizer

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/file"
)

func TestOrganize(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testOrganize")
	if err := file.Copy("../test-resource/testOrganize.template", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	// Organize project should succeed
	err := Organize(testPath, NewOption())
	if err != nil {
		t.Errorf("[ERROR] Cannot organize: %s", err)
	}

	// test content (result of pre-organize, organize, and post-organize)
	testFilePath := filepath.Join(testPath, "test")
	testByteContent, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", testFilePath, err)
	} else {
		testContent := strings.Trim(string(testByteContent), "\n")
		if testContent != "pre-organize\norganize\npost-organize" {
			t.Errorf("[UNEXPECTED] content should be `pre-organize\norganize\npost-organize`: %s", testContent)
		}
	}

	// a.txt content
	aFilePath := filepath.Join(testPath, "service/d/controller/c/lib/a/a.txt")
	aByteContent, err := ioutil.ReadFile(aFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", aFilePath, err)
	} else {
		aContent := strings.Trim(string(aByteContent), "\n")
		if aContent != "a" {
			t.Errorf("[UNEXPECTED] content should be `a`: %s", aContent)
		}
	}

	// b.txt content
	bFilePath := filepath.Join(testPath, "service/d/controller/c/lib/b/b.txt")
	bByteContent, err := ioutil.ReadFile(bFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", bFilePath, err)
	} else {
		bContent := strings.Trim(string(bByteContent), "\n")
		if bContent != "b" {
			t.Errorf("[UNEXPECTED] content should be `b`: %s", bContent)
		}
	}

	// c.txt content
	cFilePath := filepath.Join(testPath, "service/d/controller/c/c.txt")
	cByteContent, err := ioutil.ReadFile(cFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", cFilePath, err)
	} else {
		cContent := strings.Trim(string(cByteContent), "\n")
		if cContent != "c" {
			t.Errorf("[UNEXPECTED] content should be `c`: %s", cContent)
		}
	}

}
