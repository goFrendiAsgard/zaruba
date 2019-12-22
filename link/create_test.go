package link

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/file"
)

func TestCreateLink(t *testing.T) {
	var err error

	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testLink")
	if err = file.Copy("../test-resource/testLink.template", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	a := filepath.Join(testPath, "a")
	b := filepath.Join(testPath, "b")
	c := filepath.Join(testPath, "c")
	d := filepath.Join(testPath, "d")

	// Add dependency link
	if err = Create(testPath, a, b); err != nil {
		t.Errorf("[ERROR] Cannot create link: %s", err)
		return
	}
	if err = Create(testPath, a, c); err != nil {
		t.Errorf("[ERROR] Cannot create link: %s", err)
		return
	}
	if err = Create(testPath, b, d); err != nil {
		t.Errorf("[ERROR] Cannot create link: %s", err)
		return
	}

	// Read dependency file
	depFileName := filepath.Join(testPath, "zaruba.dependency.json")
	jsonB, err := ioutil.ReadFile(depFileName)
	if err != nil {
		t.Errorf("[UNEXPECTED] Cannot read zaruba.dependency.json: %s", err)
	}

	// unmarshal
	dep := map[string][]string{}
	if err = json.Unmarshal(jsonB, &dep); err != nil {
		t.Errorf("[ERROR] Cannot unmarshal JSON: %s", err)
		return
	}

	// check a
	aLink, aExists := dep[a]
	if aExists {
		if len(aLink) != 2 || aLink[0] != b || aLink[1] != c {
			t.Errorf("[UNEXPECTED] a should has b and c: %#v", dep)
		}
	} else {
		t.Errorf("[UNEXPECTED] a is not exists: %#v", dep)
	}

	// check b
	bLink, bExists := dep[b]
	if bExists {
		if len(bLink) != 1 || bLink[0] != d {
			t.Errorf("[UNEXPECTED] b should has b and c: %#v", dep)
		}
	} else {
		t.Errorf("[UNEXPECTED] b is not exists: %#v", dep)
	}

}
