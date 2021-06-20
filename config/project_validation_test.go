package config

import (
	"strings"
	"testing"
)

func TestProjectValidateInvalidFile(t *testing.T) {
	_, _, _, err := getProject("inexist.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error reading file") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInvalidYaml(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidYaml.txt")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}
func TestProjectValidateInvalidEnvsValue(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidEnvsValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInvalidInputsValue(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidInputsValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInvalidTasksValue(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidTasksValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInvalidTaskEnvsValue(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidTaskEnvsValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInvalidProjectKey(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidProjectKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInvalidEnvKey(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidEnvKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInvalidInputKey(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidInputKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInvalidTaskKey(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidTaskKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInvalidTasEnvkKey(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidTaskEnvKey.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "invalid key on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateValidKey(t *testing.T) {
	if _, _, _, err := getProject("../test-resources/project/validation/validKey.zaruba.yaml"); err != nil {
		t.Error(err)
	}
}

func TestProjectValidateInvalidValue(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/invalidValue.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error parsing YAML") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInexistInclude(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/inexistInclude/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "error reading file") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateRedundantInput(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/redundantInput/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant input declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateRedundantTask(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/redundantTask/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant task declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateRedundantEnvRef(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/redundantEnvRef/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant envs declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateRedundantConfigRef(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/redundantConfigRef/main.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant configs declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateNonRedundant(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/nonRedundant/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
	}
}

func TestProjectValidateInexistInputs(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/inexistInputs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared input task") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInexistDependencies(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/inexistDependencies.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared task dependency") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateRedundantExtendAndExtends(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/redundantExtendAndExtends.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant key declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInexistExtend(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/inexistExtend.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared parent task") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInexistExtends(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/inexistExtends.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared parent task") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateRedundantEnvRefAndEnvRefs(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/redundantEnvRefAndEnvRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant key declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInexistEnvRef(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/inexistEnvRef.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared envRef") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInexistEnvRefs(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/inexistEnvRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared envRefs") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateRedundantConfigRefAndConfigRefs(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/redundantConfigRefAndConfigRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant key declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInexistConfigRef(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/inexistConfigRef.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared configRef") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateInexistConfigRefs(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/inexistConfigRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared configRefs") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateRecursiveParentTask(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/recursiveParentTask.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "recursive task on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestProjectValidateRecursiveDependencyTask(t *testing.T) {
	_, _, _, err := getProject("../test-resources/project/validation/recursiveDependencyTask.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "recursive task on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}
