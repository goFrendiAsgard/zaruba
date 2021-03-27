package boolean

import "testing"

func TestIsTrueYesLower(t *testing.T) {
	if !IsTrue("yes") {
		t.Errorf("yes should be true")
	}
}

func TestIsTrueYesUpper(t *testing.T) {
	if !IsTrue("YES") {
		t.Errorf("YES should be true")
	}
}

func TestIsTrueYesCapital(t *testing.T) {
	if !IsTrue("Yes") {
		t.Errorf("Yes should be true")
	}
}

func TestIsTrueYLower(t *testing.T) {
	if !IsTrue("y") {
		t.Errorf("y should be true")
	}
}

func TestIsTrueYUpper(t *testing.T) {
	if !IsTrue("Y") {
		t.Errorf("Y should be true")
	}
}

func TestIsTrueTrueLower(t *testing.T) {
	if !IsTrue("true") {
		t.Errorf("true should be true")
	}
}

func TestIsTrueTrueUpper(t *testing.T) {
	if !IsTrue("TRUE") {
		t.Errorf("TRUE should be true")
	}
}

func TestIsTrueTrueCapital(t *testing.T) {
	if !IsTrue("True") {
		t.Errorf("True should be true")
	}
}

func TestIsTrueOnLower(t *testing.T) {
	if !IsTrue("on") {
		t.Errorf("on should be true")
	}
}

func TestIsTrueOnUpper(t *testing.T) {
	if !IsTrue("ON") {
		t.Errorf("ON should be true")
	}
}

func TestIsTrueOnCapital(t *testing.T) {
	if !IsTrue("On") {
		t.Errorf("On should be true")
	}
}

func TestIsTrueOther(t *testing.T) {
	if IsTrue("Other") {
		t.Errorf("Other should not be true")
	}
}

func TestIsFalseNoLower(t *testing.T) {
	if !IsFalse("no") {
		t.Errorf("no should be false")
	}
}

func TestIsFalseNoUpper(t *testing.T) {
	if !IsFalse("NO") {
		t.Errorf("NO should be false")
	}
}

func TestIsFalseNoCapital(t *testing.T) {
	if !IsFalse("No") {
		t.Errorf("No should be false")
	}
}

func TestIsFalseNLower(t *testing.T) {
	if !IsFalse("n") {
		t.Errorf("n should be false")
	}
}

func TestIsFalseNUpper(t *testing.T) {
	if !IsFalse("N") {
		t.Errorf("N should be false")
	}
}

func TestIsFalseFalseLower(t *testing.T) {
	if !IsFalse("false") {
		t.Errorf("false should be false")
	}
}

func TestIsFalseFalseUpper(t *testing.T) {
	if !IsFalse("FALSE") {
		t.Errorf("FALSE should be false")
	}
}

func TestIsFalseFalseCapital(t *testing.T) {
	if !IsFalse("False") {
		t.Errorf("False should be false")
	}
}

func TestIsFalseOffLower(t *testing.T) {
	if !IsFalse("off") {
		t.Errorf("off should be false")
	}
}

func TestIsFalseOffUpper(t *testing.T) {
	if !IsFalse("OFF") {
		t.Errorf("OFF should be false")
	}
}

func TestIsFalseOffCapital(t *testing.T) {
	if !IsFalse("Off") {
		t.Errorf("Off should be false")
	}
}

func TestIsFalseOther(t *testing.T) {
	if IsFalse("Other") {
		t.Errorf("Other should not be false")
	}
}
