package dsl

import (
	"path/filepath"
	"testing"
)

func TestTplGetWorkPath(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	// absolute
	expected := "/home/gofrendi"
	actual := tpl.GetWorkPath("/home/gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected, _ = filepath.Abs("../test-resources/taskdata/util/location/gofrendi")
	actual = tpl.GetWorkPath("./gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTplGetPath(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	// absolute
	expected := "/home/gofrendi"
	actual := tpl.GetTaskPath("/home/gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected, _ = filepath.Abs("../test-resources/taskdata/util/zaruba-tasks/gofrendi")
	actual = tpl.GetTaskPath("./gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTplParseFile(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	expected := "value"
	actual, err := tpl.ParseFile("../gotmpl/good.gotmpl")
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTplParseFileInvalid(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	_, err = tpl.ParseFile("../gotmpl/invalid.gotmpl")
	if err == nil {
		t.Errorf("error expected")
		return
	}
}

func TestTplParseFileError(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	_, err = tpl.ParseFile("../gotmpl/error.gotmpl")
	if err == nil {
		t.Errorf("error expected")
		return
	}
}

func TestTplParseFileInexist(t *testing.T) {
	project, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	tpl := NewTpl(task)
	_, err = tpl.ParseFile("../gotmpl/inexist.gotmpl")
	if err == nil {
		t.Errorf("error expected")
		return
	}
}
