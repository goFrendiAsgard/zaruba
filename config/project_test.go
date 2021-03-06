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
	validProject.AddGlobalEnv("../test_resource/valid/local.env")
	validProject.AddGlobalEnv("foo=bar")
	if err = validProject.AddValues("pi=3.14"); err != nil {
		t.Error(err)
		return err
	}
	if err = validProject.AddValues("../test_resource/valid/values.yaml"); err != nil {
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
	if validProject.Values["pi"] != "3.14" {
		t.Error("pi should be 3.14")
	}
	if validProject.Values["g"] != "9.8" {
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
	actual := validProject.Tasks["runApiGateway"].BasePath
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
	if err = project.AddValues("../test_resource/notExists.yaml"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectTaskRedeclared(t *testing.T) {
	_, err := NewProject("../test_resource/invalidTaskRedeclared/task.yaml")
	if err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectValuesFormat(t *testing.T) {
	project, err := NewProject("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = project.AddValues("../test_resource/invalidYaml.txt"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectNotExist(t *testing.T) {
	_, err := NewProject("../test_resource/notExist.yaml")
	if err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectFormat(t *testing.T) {
	_, err := NewProject("../test_resource/invalidYaml.txt")
	if err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectInclusion(t *testing.T) {
	_, err := NewProject("../test_resource/invalidInclusion.yaml")
	if err == nil {
		t.Error("Error expected")
	}
}
