package link

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestCreateLink(t *testing.T) {
	var err error

	a, _ := filepath.Abs("../test-playground/testLink/a")
	b, _ := filepath.Abs("../test-playground/testLink/b")
	c, _ := filepath.Abs("../test-playground/testLink/c")
	d, _ := filepath.Abs("../test-playground/testLink/d")

	// Add dependency link
	if err = Create("../test-playground/testLink", a, b); err != nil {
		t.Errorf("[ERROR] Cannot create link: %s", err)
	}
	if err = Create("../test-playground/testLink", a, c); err != nil {
		t.Errorf("[ERROR] Cannot create link: %s", err)
	}
	if err = Create("../test-playground/testLink", b, d); err != nil {
		t.Errorf("[ERROR] Cannot create link: %s", err)
	}

	// Read dependency file
	depFileName := "../test-playground/testLink/zaruba.dependency.json"
	jsonB, err := ioutil.ReadFile(depFileName)
	if err != nil {
		t.Errorf("[ERROR] Cannot read zaruba.dependency.json: %s", err)
	}

	// unmarshal
	dep := map[string][]string{}
	if err = json.Unmarshal(jsonB, &dep); err != nil {
		t.Errorf("[ERROR] Cannot unmarshal JSON: %s", err)
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
