package organizer

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/otiai10/copy"
)

func TestOrganize(t *testing.T) {
	baseTestPath, err := filepath.Abs(os.Getenv("ZARUBA_TEST_DIR"))
	if err != nil {
		t.Errorf("[ERROR] Cannot fetch testPath from envvar: %s", err)
		return
	}
	testPath := filepath.Join(baseTestPath, "testOrganize")
	if err = copy.Copy("../test-resource/testOrganize.template", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	// Organize project should succeed
	err = Organize(testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot organize: %s", err)
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
