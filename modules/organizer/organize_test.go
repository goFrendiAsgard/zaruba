package organizer

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/modules/component"
	"github.com/state-alchemists/zaruba/modules/config"
)

func TestOrganize(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testOrganize")

	// create project and service
	if err := component.Create("project", testPath); err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}
	if err := component.Create("go-service", testPath, "gopher_one"); err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}
	if err := component.Create("go-service", testPath, "gopher_two"); err != nil {
		t.Errorf("[ERROR] Cannot create component: %s", err)
		return
	}
	// load project config
	p, err := config.NewProjectConfig(testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot load config: %s", err)
		return
	}

	// create test.txt
	if err := ioutil.WriteFile(filepath.Join(testPath, "libraries/go/transport/test.txt"), []byte("test"), 0755); err != nil {
		t.Errorf("[ERROR] Cannot create test.txt: %s", err)
	}

	// organize
	if err := Organize(testPath, p); err != nil {
		t.Errorf("[ERROR] Cannot organize: %s", err)
	}

	// check test.txt in gopher_one
	if _, err := os.Stat(filepath.Join(testPath, "services/gopher_one/transport/test.txt")); err != nil && os.IsNotExist(err) {
		t.Errorf("[ERROR] test.txt is not exist on gopher_one: %s", err)
	}

	// check test.txt in gopher_two
	if _, err := os.Stat(filepath.Join(testPath, "services/gopher_two/transport/test.txt")); err != nil && os.IsNotExist(err) {
		t.Errorf("[ERROR] test.txt is not exist on gopher_two: %s", err)
	}

}
