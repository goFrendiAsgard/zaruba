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
	_, err := getProject("../test-resources/project/validation/invalidYaml.txt")
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
	_, err := getProject("../test-resources/project/validation/invalidEnvsValue.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/invalidInputsValue.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/invalidTasksValue.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/invalidTaskEnvsValue.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/invalidProjectKey.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/invalidEnvKey.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/invalidInputKey.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/invalidTaskKey.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/invalidTaskEnvKey.zaruba.yaml")
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
	if _, err := getProject("../test-resources/project/validation/validKey.zaruba.yaml"); err != nil {
		t.Error(err)
	}
}

func TestInvalidValue(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/invalidValue.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/inexistInclude/main.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/redundantInput/main.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/redundantTask/main.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/redundantEnvRef/main.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/redundantConfigRef/main.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/redundantLConfigRef/main.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/nonRedundant/main.zaruba.yaml")
	if err != nil {
		t.Error(err)
	}
}

func TestInexistInputs(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/inexistInputs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared input task") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistDependencies(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/inexistDependencies.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared task dependency") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantExtendAndExtends(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/redundantExtendAndExtends.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/inexistExtend.zaruba.yaml")
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
	_, err := getProject("../test-resources/project/validation/inexistExtends.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared parent task") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantEnvRefAndEnvRefs(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/redundantEnvRefAndEnvRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant key declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistEnvRef(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/inexistEnvRef.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared envRef") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistEnvRefs(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/inexistEnvRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared envRefs") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantConfigRefAndConfigRefs(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/redundantConfigRefAndConfigRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant key declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistConfigRef(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/inexistConfigRef.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared configRef") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistConfigRefs(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/inexistConfigRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared configRefs") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRedundantLConfigRefAndLConfigRefs(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/redundantLConfigRefAndLConfigRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "redundant key declaration") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistLConfigRef(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/inexistLConfigRef.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared lconfig") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestInexistLConfigRefs(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/inexistLConfigRefs.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "undeclared lconfig") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRecursiveParentTask(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/recursiveParentTask.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "recursive task on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestRecursiveDependencyTask(t *testing.T) {
	_, err := getProject("../test-resources/project/validation/recursiveDependencyTask.zaruba.yaml")
	if err == nil {
		t.Errorf("error expected")
		return
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "recursive task on") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}
