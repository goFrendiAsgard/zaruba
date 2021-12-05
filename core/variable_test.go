package core

import (
	"testing"
)

func TestGetVariableName(t *testing.T) {
	project, err := getProject("../test-resources/variable/getName.zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	expected := "variableName"
	actual := project.Inputs["variableName"].GetName()
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
