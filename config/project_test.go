package config

import (
	"os"
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
	if err = project.AddValue("input2=value2"); err != nil {
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

func TestProjectGetValue(t *testing.T) {
	project, err := getProject("../test-resources/project/getValue/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("../test-resources/project/getValue/default.values.yaml"); err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("input2=value2"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	expected := "value1"
	actual := project.GetValue("input1")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected = "value2"
	actual = project.GetValue("input2")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected = ""
	actual = project.GetValue("input3")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected = ""
	actual = project.GetValue("input4")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestProjectIsValueExist(t *testing.T) {
	project, err := getProject("../test-resources/project/getValue/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("../test-resources/project/getValue/default.values.yaml"); err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("input2=value2"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	expected := true
	actual := project.IsValueExist("input1")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
	expected = true
	actual = project.IsValueExist("input2")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
	expected = true
	actual = project.IsValueExist("input3")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
	expected = false
	actual = project.IsValueExist("input4")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestProjectGetValues(t *testing.T) {
	project, err := getProject("../test-resources/project/getValue/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("../test-resources/project/getValue/default.values.yaml"); err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("input2=value2"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	expectedValues := map[string]string{
		"input1": "value1",
		"input2": "value2",
		"input3": "",
	}
	actualValues := project.GetValues()
	for key, expected := range expectedValues {
		actual := actualValues[key]
		if actual != expected {
			t.Errorf("%s, expected: %s, actual: %s", key, expected, actual)
		}
	}
}

func TestProjectAddGlobalEnvAfterInit(t *testing.T) {
	project, err := getProject("../test-resources/project/addGlobalEnv/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	if err = project.AddGlobalEnv("ENV2=value2"); err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "cannot AddGlobalEnv") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectAddGlobalEnv(t *testing.T) {
	project, err := getProject("../test-resources/project/addGlobalEnv/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddGlobalEnv("../test-resources/project/addGlobalEnv/template.env"); err != nil {
		t.Error(err)
		return
	}
	if err = project.AddGlobalEnv("ENV2=value2"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err != nil {
		t.Error(err)
		return
	}
	expected := "value1"
	actual := os.Getenv("ENV1")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
	expected = "value2"
	actual = os.Getenv("ENV2")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestProjectInitInvalidInputValue(t *testing.T) {
	project, err := getProject("../test-resources/project/init/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	project.SetValue("input1", "012")
	if err := project.Init(); err == nil {
		t.Errorf("error expected")
	}
}

func TestProjectInit(t *testing.T) {
	project, err := getProject("../test-resources/project/init/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	project.SetValue("input1", "value")
	if err := project.Init(); err != nil {
		t.Error(err)
		return
	}
	expected := "value"
	actual := os.Getenv("ZARUBA_INPUT_INPUT1")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestProjectGetInputsInexistTask(t *testing.T) {
	project, err := getProject("../test-resources/project/getInputsInexistTask.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	_, _, err = project.GetInputs([]string{"inexistTask"})
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "task 'inexistTask' is not exist" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectGetInputs(t *testing.T) {
	project, err := getProject("../test-resources/project/getInputs.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	_, actualInputOrder, err := project.GetInputs([]string{"task"})
	if err != nil {
		t.Error(err)
		return
	}
	expectedInputOrder := []string{"input4", "input2", "input3", "input1"}
	if len(actualInputOrder) != len(expectedInputOrder) {
		t.Errorf("expected: %#v, actual: %#v", expectedInputOrder, actualInputOrder)
	}
	for index, expected := range expectedInputOrder {
		actual := actualInputOrder[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
}

func TestProjectValidateByTaskNamesInexistTask(t *testing.T) {
	project, err := getProject("../test-resources/project/validateByTaskNamesInexistTask.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	err = project.ValidateByTaskNames([]string{"inexistTask"})
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "task 'inexistTask' is not exist" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateByTaskNamesInvalidInput(t *testing.T) {
	project, err := getProject("../test-resources/project/validateByTaskNames.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	project.SetValue("input", "invalidValue")
	err = project.ValidateByTaskNames([]string{"task"})
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if errorMessage != "value of input variable 'input' does not match '^[0-9]+$': invalidValue" {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateByTaskNames(t *testing.T) {
	project, err := getProject("../test-resources/project/validateByTaskNames.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	project.SetValue("input", "012")
	err = project.ValidateByTaskNames([]string{"task"})
	if err != nil {
		t.Error(err)
		return
	}
}
