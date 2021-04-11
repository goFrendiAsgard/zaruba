package config

import "testing"

func getValidInitiatedInputProject(t *testing.T) (project *Project, err error) {
	project, err = getProject("../test_resource/input.yaml")
	if err != nil {
		t.Error(err)
		return project, err
	}
	if err = project.SetValue("goodValidation", "012"); err != nil {
		t.Error(err)
		return project, err
	}
	if err = project.Init(); err != nil {
		t.Error(err)
	}
	return project, err
}

func TestInputName(t *testing.T) {
	project, err := getValidInitiatedInputProject(t)
	if err != nil {
		return
	}
	expectedNames := []string{"noValidation", "goodValidation"}
	for _, expectedName := range expectedNames {
		input, exist := project.Inputs[expectedName]
		if !exist {
			t.Errorf("Input %s is not exist", expectedName)
		}
		actualName := input.GetName()
		if actualName != expectedName {
			t.Errorf("Invalid name. Expected: %s, get %s", expectedName, actualName)
		}
	}
}

func TestInputNoValidation(t *testing.T) {
	project, err := getValidInitiatedInputProject(t)
	if err != nil {
		return
	}
	inputName := "noValidation"
	input, exist := project.Inputs[inputName]
	if !exist {
		t.Errorf("Input %s is not exist", inputName)
		return
	}
	if validationErr := input.Validate("anyValue"); validationErr != nil {
		t.Errorf("Input with noValidation should never throw error")
		t.Error(validationErr)
	}
}

func TestInputGoodValidation(t *testing.T) {
	project, err := getValidInitiatedInputProject(t)
	if err != nil {
		return
	}
	inputName := "goodValidation"
	input, exist := project.Inputs[inputName]
	if !exist {
		t.Errorf("Input %s is not exist", inputName)
		return
	}
	if validationErr := input.Validate("0123"); validationErr != nil {
		t.Errorf("Input with goodValidation should not throw error when value match")
		t.Error(validationErr)
	}
	if validationErr := input.Validate("abc"); validationErr == nil {
		t.Errorf("Input with goodValidation should throw error when value doesn't match")
	}
}

func TestInputBadValidation(t *testing.T) {
	project, err := getProject("../test_resource/invalidInput.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.SetValue("badValidation", "012"); err != nil {
		t.Error(err)
		return
	}
	if err = project.Init(); err == nil {
		t.Errorf("Project containing badValidation input should throw error when initialized")
	}
}
