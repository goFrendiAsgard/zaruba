package core

import (
	"strings"
	"testing"
)

func TestNoValidationVariable(t *testing.T) {
	variable := Variable{}
	if err := variable.Validate("anyValue"); err != nil {
		t.Error(err)
	}
}

func TestInvalidValidationVariable(t *testing.T) {
	variable := Variable{Validation: "[["}
	if err := variable.Validate("anyValue"); err == nil {
		t.Errorf("error expected")
	}
}

func TestNotMatchValidationVariable(t *testing.T) {
	variable := Variable{Validation: "^[a-z]+$"}
	err := variable.Validate("012")
	if err == nil {
		t.Errorf("error expected")
	}
	errorMessage := err.Error()
	if !strings.HasPrefix(errorMessage, "value of input variable") {
		t.Errorf("invalid error message: %s", errorMessage)
	}
}

func TestMatchValidationVariable(t *testing.T) {
	variable := Variable{Validation: "^[a-z]+$"}
	if err := variable.Validate("abc"); err != nil {
		t.Error(err)
	}
}
