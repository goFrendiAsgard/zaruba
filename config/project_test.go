package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var validProject *Project

func setupValidProject(t *testing.T) (err error) {
	if validProject != nil {
		return err
	}
	validProject, err = NewProject("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return err
	}
	if err = validProject.AddGlobalEnv("../test_resource/valid/local.env"); err != nil {
		t.Error(err)
		return err
	}
	if err = validProject.AddGlobalEnv("foo=bar"); err != nil {
		t.Error(err)
		return err
	}
	if err = validProject.AddValue("pi=3.14"); err != nil {
		t.Error(err)
		return err
	}
	if err = validProject.AddValue("../test_resource/valid/values.yaml"); err != nil {
		t.Error(err)
		return err
	}
	validProject.Init()
	return err
}

func TestValidProjectAddGlobalEnv(t *testing.T) {
	if err := setupValidProject(t); err != nil {
		return
	}
	if os.Getenv("foo") != "bar" {
		t.Errorf("foo value should be bar, get %s", os.Getenv("foo"))
	}
	if os.Getenv("API_GATEWAY_HTTP_PORT") != "8080" {
		t.Errorf("API_GATEWAY_HTTP_PORT value should be 8080, get %s", os.Getenv("API_GATEWAY_HTTP_PORT"))
	}
	if os.Getenv("API_GATEWAY_PROMETHEUS_PORT") != "8081" {
		t.Errorf("API_GATEWAY_PROMTHEUS_PORT value should be 8081, get %s", os.Getenv("API_GATEWAY_PROMETHEUS_PORT"))
	}
	if os.Getenv("ZARUBA_HOME") == "" {
		t.Error("ZARUBA_HOME should be automatically set")
	}
}

func TestValidProjectAddValues(t *testing.T) {
	if err := setupValidProject(t); err != nil {
		return
	}
	if validProject.GetValue("pi") != "3.14" {
		t.Error("pi should be 3.14")
	}
	if validProject.GetValue("g") != "9.8" {
		t.Error("g should be 9.8")
	}
}

func TestValidProjectInclusion(t *testing.T) {
	if err := setupValidProject(t); err != nil {
		return
	}
	for _, taskName := range []string{"core.runNodeJsService", "core.runShellScript", "core.runBashScript", "core.runPythonScript", "core.runNodeJsScript", "core.runStaticWebService", "runApiGateway", "runIntegrationTest", "serveStaticFiles", "sayPythonHello"} {
		if _, exists := validProject.Tasks[taskName]; !exists {
			t.Errorf(fmt.Sprintf("Task %s is not exist", taskName))
		}
		fmt.Println(taskName, validProject.Tasks[taskName].Extend)
	}
}

func TestValidProjectInputs(t *testing.T) {
	if err := setupValidProject(t); err != nil {
		return
	}
	expectedInputs := map[string]*Input{
		"testName": {
			DefaultValue: "myTest",
			Description:  "Test name",
		},
		"taskName": {
			DefaultValue: "myTask",
			Description:  "Task name",
		},
		"host": {
			DefaultValue: "localhost",
			Description:  "Host",
		},
		"font": {
			DefaultValue: "cascadia",
			Description:  "Font",
		},
	}
	actualInputs := validProject.Inputs
	actualInputCount := 0
	for inputName, actualInput := range actualInputs {
		expectedInput, inputExist := expectedInputs[inputName]
		if !inputExist {
			t.Errorf(fmt.Sprintf("Input %s is not expected", inputName))
		}
		if actualInput.DefaultValue != expectedInput.DefaultValue {
			t.Errorf(fmt.Sprintf("Expected default value of %s: %s, Actual: %s", inputName, expectedInput.DefaultValue, actualInput.DefaultValue))
		}
		if actualInput.Description != expectedInput.Description {
			t.Errorf(fmt.Sprintf("Expected description of %s: %s, Actual: %s", inputName, expectedInput.Description, actualInput.Description))
		}
		actualInputCount++
	}
	if actualInputCount != 4 {
		t.Errorf(fmt.Sprintf("There should be 4 inputs, currently %#v", actualInputs))
	}
	// expected inputs should also populate project's value
	for inputName, input := range expectedInputs {
		actualValue, valueExist := validProject.GetValue(inputName), validProject.IsValueExist(inputName)
		if !valueExist {
			t.Errorf(fmt.Sprintf("Value %s is expected", inputName))
		}
		if actualValue != input.DefaultValue {
			t.Errorf(fmt.Sprintf("Expected %s to contains default value from input: %s, Actual: %s", inputName, actualValue, input.DefaultValue))
		}
	}
}

func TestValidProjectGetInputs(t *testing.T) {
	if err := setupValidProject(t); err != nil {
		return
	}
	expectedInputs := map[string]*Input{
		"testName": {
			DefaultValue: "myTest",
			Description:  "Test name",
		},
		"taskName": {
			DefaultValue: "myTask",
			Description:  "Task name",
		},
		"host": {
			DefaultValue: "localhost",
			Description:  "Host",
		},
	}
	expectedInputOrder := []string{"taskName", "host", "testName"}
	actualInputs, actualInputOrder, err := validProject.GetInputs([]string{"runIntegrationTest"})
	if err != nil {
		t.Error(err)
	}
	actualInputCount := 0
	for inputName, actualInput := range actualInputs {
		expectedInput, exist := expectedInputs[inputName]
		if !exist {
			t.Errorf(fmt.Sprintf("Input %s is not expected", inputName))
		}
		if actualInput.DefaultValue != expectedInput.DefaultValue {
			t.Errorf(fmt.Sprintf("Expected default value of %s: %s, Actual: %s", inputName, expectedInput.DefaultValue, actualInput.DefaultValue))
		}
		if actualInput.Description != expectedInput.Description {
			t.Errorf(fmt.Sprintf("Expected description of %s: %s, Actual: %s", inputName, expectedInput.Description, actualInput.Description))
		}
		actualInputCount++
	}
	if actualInputCount != 3 {
		t.Errorf(fmt.Sprintf("There should be 3 inputs, currently %#v", actualInputs))
	}
	if len(actualInputOrder) != actualInputCount {
		t.Errorf(fmt.Sprintf("Expected inputOrder to contains %d elements, but actualInputOrder is: %#v", actualInputCount, actualInputOrder))
		return
	}
	for orderIndex := range expectedInputOrder {
		expected := expectedInputOrder[orderIndex]
		actual := actualInputOrder[orderIndex]
		if expected != actual {
			t.Errorf("Expected inputOrder[%d] to be %s, but get %s", orderIndex, expected, actual)
		}
	}
}

func TestValidProjectTaskDirPath(t *testing.T) {
	if err := setupValidProject(t); err != nil {
		return
	}
	if _, exists := validProject.Tasks["runApiGateway"]; !exists {
		t.Errorf("Task runApiGateway is not exist")
		return
	}
	actual := validProject.Tasks["runApiGateway"].basePath
	expected, err := filepath.Abs("../test_resource/valid/api-gateway")
	if err != nil {
		t.Error(err)
		return
	}
	if expected != actual {
		t.Errorf(fmt.Sprintf("Expected: %s, Actual: %s", expected, actual))
	}
}

func TestValidProjectTaskProject(t *testing.T) {
	if err := setupValidProject(t); err != nil {
		return
	}
	for _, task := range validProject.Tasks {
		if task.Project != validProject {
			t.Errorf("Task's parent is not correctly set")
		}
	}
}

func TestValidProjectEnvTask(t *testing.T) {
	if err := setupValidProject(t); err != nil {
		return
	}
	for _, task := range validProject.Tasks {
		for _, env := range task.Env {
			if env.Task != task {
				t.Errorf("Env's parent is not correctly set")
			}
		}
	}
}

func TestValidProjectName(t *testing.T) {
	project, err := NewProject("../test_resource/named.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	expected := "edward"
	actual := project.GetName()
	if actual != expected {
		t.Errorf(fmt.Sprintf("Expected: %s, Actual: %s", expected, actual))
	}
}

func TestValidProjectWithNonExistValueFile(t *testing.T) {
	project, err := NewProject("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("../test_resource/notExists.yaml"); err == nil {
		t.Error("Error expected")
	}
}

func TestValidProjectGetInputsFromNonExistingTask(t *testing.T) {
	project, err := NewProject("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if _, _, err = project.GetInputs([]string{"nonExistingTask"}); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectTaskRedeclared(t *testing.T) {
	if _, err := NewProject("../test_resource/invalidTaskRedeclared/task.yaml"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectInputRedeclared(t *testing.T) {
	if _, err := NewProject("../test_resource/invalidInputRedeclared/input.yaml"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectValuesFormat(t *testing.T) {
	project, err := NewProject("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValue("../test_resource/invalidYaml.txt"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectNotExist(t *testing.T) {
	if _, err := NewProject("../test_resource/notExist.yaml"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectUndeclaredInput(t *testing.T) {
	if _, err := NewProject("../test_resource/invalidUndefinedInput/task.yaml"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectFormat(t *testing.T) {
	if _, err := NewProject("../test_resource/invalidYaml.txt"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectInclusion(t *testing.T) {
	if _, err := NewProject("../test_resource/invalidInclusion.yaml"); err == nil {
		t.Error("Error expected")
	}
}
