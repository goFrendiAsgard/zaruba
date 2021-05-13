package config

import (
	"path/filepath"
	"testing"
)

func TestTdGetWorkPath(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// absolute
	expected := "/home/gofrendi"
	actual := td.GetWorkPath("/home/gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected, _ = filepath.Abs("../test-resources/taskdata/util/location/gofrendi")
	actual = td.GetWorkPath("./gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdGetRelativePath(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// absolute
	expected := "/home/gofrendi"
	actual := td.GetRelativePath("/home/gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected, _ = filepath.Abs("../test-resources/taskdata/util/zaruba-tasks/gofrendi")
	actual = td.GetRelativePath("./gofrendi")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdGetTaskExist(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	otherTask, err := td.GetTask("otherTaskName")
	if err != nil {
		t.Error(err)
		return
	}
	if otherTask == nil {
		t.Errorf("otherTask is nil")
	}
}

func TestTdGetTaskInExist(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	otherTask, err := td.GetTask("inexistTask")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	if otherTask != nil {
		t.Errorf("otherTask is not nil")
	}
}

func TestTdIsTrue(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// true
	expected := true
	actual := td.IsTrue("true")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
	// yes
	expected = true
	actual = td.IsTrue("yes")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
	// false
	expected = false
	actual = td.IsTrue("false")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestTdIsFalse(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	// true
	expected := true
	actual := td.IsFalse("false")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
	// yes
	expected = true
	actual = td.IsFalse("no")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
	// false
	expected = false
	actual = td.IsFalse("true")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestTdTrim(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	expected := "value"
	actual := td.Trim("\n   value\n\n  ", " \n")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdReplaceAllWith(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	expected := "fruit fruit fruit vegetable"
	actual := td.ReplaceAllWith("strawberry grape orange vegetable", "strawberry", "grape", "orange", "fruit")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdReplaceAllWithSingleParameter(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	expected := "vegetable"
	actual := td.ReplaceAllWith("vegetable")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdEscapeShellValueSingleQuote(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	expected := "\\\\\\\\t \\n \\`backtick\\` \\'quoted\\' \"quoted\""
	actual := td.EscapeShellValue("\\t \n `backtick` 'quoted' \"quoted\"", "'")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdEscapeShellValueDoubleQuote(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	expected := "\\\\\\\\t \\n \\`backtick\\` 'quoted' \\\"quoted\\\""
	actual := td.EscapeShellValue("\\t \n `backtick` 'quoted' \"quoted\"", "\"")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdSingleQuoteShellValue(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	expected := "'\\\\\\\\t \\n \\`backtick\\` \\'quoted\\' \"quoted\"'"
	actual := td.SingleQuoteShellValue("\\t \n `backtick` 'quoted' \"quoted\"")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdDoubleQuoteShellValue(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	expected := "\"\\\\\\\\t \\n \\`backtick\\` 'quoted' \\\"quoted\\\"\""
	actual := td.DoubleQuoteShellValue("\\t \n `backtick` 'quoted' \"quoted\"")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestTdListDir(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	expectedList := []string{"file1", "file2"}
	actualList, err := td.ListDir(".")
	if err != nil {
		t.Error(err)
		return
	}
	if len(actualList) != len(expectedList) {
		t.Errorf("expected: %#v, actual: %#v", expectedList, actualList)
	}
	for _, expected := range expectedList {
		actualFound := false
		for _, actual := range actualList {
			if actual == expected {
				actualFound = true
			}
		}
		if !actualFound {
			t.Errorf("cannot find key %s, on: %#v", expected, actualList)
		}
	}
}

func TestTdListDirInexist(t *testing.T) {
	project, _, _, err := getProjectAndInit("../test-resources/taskdata/util/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := project.Tasks["taskName"]
	td := NewTaskData(task)
	_, err = td.ListDir("./inexist")
	if err == nil {
		t.Errorf("error expected")
		return
	}
}
