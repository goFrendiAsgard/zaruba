package config

import (
	"strings"
	"testing"
)

func TestInvalidFile(t *testing.T) {
	_, err := getProject("inexist.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error reading file") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInvalidYaml(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidYaml.txt")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}
func TestInvalidEnvsValue(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidEnvsValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInvalidInputsValue(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidInputsValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInvalidTasksValue(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidTasksValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInvalidTaskEnvsValue(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidTaskEnvsValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInvalidProjectKey(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidProjectKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInvalidEnvKey(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidEnvKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInvalidInputKey(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidInputKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInvalidTaskKey(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidTaskKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInvalidTasEnvkKey(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidTaskEnvKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestValidKey(t *testing.T) {
	if _, err := getProject("../test-resources/validation/validKey.zaruba.yaml"); err != nil {
		t.Error(err)
	}
}

func TestInvalidValue(t *testing.T) {
	_, err := getProject("../test-resources/validation/invalidValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistInclude(t *testing.T) {
	_, err := getProject("../test-resources/validation/inexistInclude/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error reading file") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantInput(t *testing.T) {
	_, err := getProject("../test-resources/validation/redundantInput/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant input declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantTask(t *testing.T) {
	_, err := getProject("../test-resources/validation/redundantTask/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant task declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantEnvRef(t *testing.T) {
	_, err := getProject("../test-resources/validation/redundantEnvRef/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant envs declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantConfigRef(t *testing.T) {
	_, err := getProject("../test-resources/validation/redundantConfigRef/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant configs declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantLConfigRef(t *testing.T) {
	_, err := getProject("../test-resources/validation/redundantLConfigRef/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant lconfigs declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestNonRedundant(t *testing.T) {
	_, err := getProject("../test-resources/validation/nonRedundant/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
	}
}

func TestInexistInput(t *testing.T) {
	_, err := getProject("../test-resources/validation/inexistInput.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared input task") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantExtendAndExtends(t *testing.T) {
	_, err := getProject("../test-resources/validation/redundantExtendAndExtends.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant key declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistExtend(t *testing.T) {
	_, err := getProject("../test-resources/validation/inexistExtend.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared parent task") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistExtends(t *testing.T) {
	_, err := getProject("../test-resources/validation/inexistExtends.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared parent task") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}
