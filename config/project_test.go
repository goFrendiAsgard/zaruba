package config

import (
	"strings"
	"testing"
)

func TestProjectGetNameDeclared(t *testing.T) {
	project, err := getProject("../test-resources/project/getNameDeclared.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	expected := "projectName"
	actual := project.GetName()
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestProjectGetNameUndeclared(t *testing.T) {
	project, err := getProject("../test-resources/project/getNameUndeclared.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	expected := "project"
	actual := project.GetName()
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestProjectGetBasePath(t *testing.T) {
	project, err := getProject("../test-resources/project/getBasePath.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	basePath := project.GetBasePath()
	if !strings.HasSuffix(basePath, "test-resources/project") {
		t.Errorf("unexpected basePath: %s", basePath)
	}
}

func TestProjectGetSortedInputNames(t *testing.T) {
	project, err := getProject("../test-resources/project/getSortedInputNames.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	sortedInputNames := project.GetSortedInputNames()
	if len(sortedInputNames) != 3 {
		t.Errorf("unexpected sortedInputNames count: %#v", sortedInputNames)
		return
	}
	expectedList := []string{"a", "b", "c"}
	for index, expected := range expectedList {
		actual := sortedInputNames[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
}

func TestProjectGetSortedTaskNames(t *testing.T) {
	project, err := getProject("../test-resources/project/getSortedTaskNames.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	sortedTaskNames := project.GetSortedTaskNames()
	if len(sortedTaskNames) != 3 {
		t.Errorf("unexpected sortedInputNames count: %#v", sortedTaskNames)
		return
	}
	expectedList := []string{"a", "b", "c"}
	for index, expected := range expectedList {
		actual := sortedTaskNames[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
}

func TestProjectAddValueAfterInit(t *testing.T) {
	project, err := getProject("../test-resources/project/addValueFromPair/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("input1=value1"); err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "cannot AddValue") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectSetValueAfterInit(t *testing.T) {
	project, err := getProject("../test-resources/project/setValueFromPair/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	if err = project.SetValue("input1", "value1"); err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "cannot SetValue") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectAddValueFromPair(t *testing.T) {
	project, err := getProject("../test-resources/project/addValueFromPair/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("input1=value1"); err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("input1=value2"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
}

func TestProjectAddValueFromFile(t *testing.T) {
	project, err := getProject("../test-resources/project/addValueFromFile/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("../test-resources/project/addValueFromFile/default.values.yaml"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
}

func TestProjectAddValueFromInexistFile(t *testing.T) {
	project, err := getProject("../test-resources/project/addValueFromFile/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("inexist.values.yaml"); err == nil {
		t.Errorf("error expected")
		return
	}
}

func TestProjectAddValueFromInvalidYamlFile(t *testing.T) {
	project, err := getProject("../test-resources/project/addValueFromFile/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("../test-resources/project/addValueFromFile/invalid.values.txt"); err == nil {
		t.Errorf("error expected")
		return
	}
}
