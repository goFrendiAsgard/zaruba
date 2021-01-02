package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var validConf *ProjectConfig

func setupValidProjectConfig(t *testing.T) (err error) {
	if validConf != nil {
		return err
	}
	validConf, err = NewConfig("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return err
	}
	validConf.AddGlobalEnv("../test_resource/valid/local.env")
	validConf.AddGlobalEnv("foo=bar")
	if err = validConf.AddKwargs("pi=3.14"); err != nil {
		t.Error(err)
		return err
	}
	if err = validConf.AddKwargs("../test_resource/valid/kwargs.yaml"); err != nil {
		t.Error(err)
		return err
	}
	validConf.Init()
	return err
}

func TestValidProjectConfigAddGlobalEnv(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
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

func TestValidProjectConfigAddKwargs(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	if validConf.Kwargs["pi"] != "3.14" {
		t.Error("pi should be 3.14")
	}
	if validConf.Kwargs["g"] != "9.8" {
		t.Error("g should be 9.8")
	}
}

func TestValidProjectConfigInclusion(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	for _, taskName := range []string{"core.runNodeJsService", "core.runShellScript", "core.runBashScript", "core.runPythonScript", "core.runNodeJsScript", "core.runStaticWebService", "runApiGateway", "runIntegrationTest", "serveStaticFiles", "sayPythonHello"} {
		if _, exists := validConf.Tasks[taskName]; !exists {
			t.Errorf(fmt.Sprintf("Task %s is not exist", taskName))
		}
	}
}

func TestValidProjectConfigTaskDirPath(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	if _, exists := validConf.Tasks["runApiGateway"]; !exists {
		t.Errorf("Task runApiGateway is not exist")
		return
	}
	actual := validConf.Tasks["runApiGateway"].BasePath
	expected, err := filepath.Abs("../test_resource/valid/api-gateway")
	if err != nil {
		t.Error(err)
		return
	}
	if expected != actual {
		t.Errorf(fmt.Sprintf("Expected: %s, Actual: %s", expected, actual))
	}
}

func TestValidProjectConfigTaskProject(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	for _, task := range validConf.Tasks {
		if task.Project != validConf {
			t.Errorf("Task's parent is not correctly set")
		}
	}
}

func TestValidProjectConfigEnvTask(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	for _, task := range validConf.Tasks {
		for _, env := range task.Env {
			if env.Task != task {
				t.Errorf("Env's parent is not correctly set")
			}
		}
	}
}

func TestValidProjectConfigName(t *testing.T) {
	conf, err := NewConfig("../test_resource/named.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	expected := "edward"
	actual := conf.GetName()
	if actual != expected {
		t.Errorf(fmt.Sprintf("Expected: %s, Actual: %s", expected, actual))
	}
}

func TestInvalidProjectConfigKwargsNotExist(t *testing.T) {
	conf, err := NewConfig("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.AddKwargs("../test_resource/notExists.yaml"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectConfigKwargsFormat(t *testing.T) {
	conf, err := NewConfig("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.AddKwargs("../test_resource/invalidYaml.txt"); err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectConfigNotExist(t *testing.T) {
	_, err := NewConfig("../test_resource/notExist.yaml")
	if err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectConfigFormat(t *testing.T) {
	_, err := NewConfig("../test_resource/invalidYaml.txt")
	if err == nil {
		t.Error("Error expected")
	}
}

func TestInvalidProjectConfigInclusion(t *testing.T) {
	_, err := NewConfig("../test_resource/invalidInclusion.yaml")
	if err == nil {
		t.Error("Error expected")
	}
}
