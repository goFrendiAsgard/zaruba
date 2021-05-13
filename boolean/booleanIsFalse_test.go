package boolean

import "testing"

func TestBooleanIsFalseLowerYes(t *testing.T) {
	if IsFalse("yes") {
		t.Errorf("yes should not be false")
	}
}

func TestBooleanIsFalseUpperYes(t *testing.T) {
	if IsFalse("YES") {
		t.Errorf("YES should not be false")
	}
}

func TestBooleanIsFalseCapitalYes(t *testing.T) {
	if IsFalse("Yes") {
		t.Errorf("Yes should not be false")
	}
}

func TestBooleanIsFalseLowerY(t *testing.T) {
	if IsFalse("y") {
		t.Errorf("y should not be false")
	}
}

func TestBooleanIsFalseUpperY(t *testing.T) {
	if IsFalse("Y") {
		t.Errorf("Y should not be false")
	}
}

func TestBooleanIsFalseLowerTrue(t *testing.T) {
	if IsFalse("true") {
		t.Errorf("true should not be false")
	}
}

func TestBooleanIsFalseUpperTrue(t *testing.T) {
	if IsFalse("TRUE") {
		t.Errorf("TRUE should not be false")
	}
}

func TestBooleanIsFalseCapitalTrue(t *testing.T) {
	if IsFalse("True") {
		t.Errorf("True should not be false")
	}
}

func TestBooleanIsFalseLowerOn(t *testing.T) {
	if IsFalse("on") {
		t.Errorf("on should not be false")
	}
}

func TestBooleanIsFalseUpperOn(t *testing.T) {
	if IsFalse("ON") {
		t.Errorf("ON should not be false")
	}
}

func TestBooleanIsFalseCapitalOn(t *testing.T) {
	if IsFalse("On") {
		t.Errorf("On should not be false")
	}
}

func TestBooleanIsFalseLowerNo(t *testing.T) {
	if !IsFalse("no") {
		t.Errorf("no should be false")
	}
}

func TestBooleanIsFalseUpperNo(t *testing.T) {
	if !IsFalse("NO") {
		t.Errorf("NO should be false")
	}
}

func TestBooleanIsFalseCapitalNo(t *testing.T) {
	if !IsFalse("No") {
		t.Errorf("No should be false")
	}
}

func TestBooleanIsFalseLowerN(t *testing.T) {
	if !IsFalse("n") {
		t.Errorf("n should be false")
	}
}

func TestBooleanIsFalseUpperN(t *testing.T) {
	if !IsFalse("N") {
		t.Errorf("N should be false")
	}
}

func TestBooleanIsFalseLowerFalse(t *testing.T) {
	if !IsFalse("false") {
		t.Errorf("false should be false")
	}
}

func TestBooleanIsFalseUpperFalse(t *testing.T) {
	if !IsFalse("FALSE") {
		t.Errorf("FALSE should be false")
	}
}

func TestBooleanIsFalseCapitalFalse(t *testing.T) {
	if !IsFalse("False") {
		t.Errorf("False should be false")
	}
}

func TestBooleanIsFalseLowerOff(t *testing.T) {
	if !IsFalse("off") {
		t.Errorf("off should be false")
	}
}

func TestBooleanIsFalseUpperOff(t *testing.T) {
	if !IsFalse("OFF") {
		t.Errorf("OFF should be false")
	}
}

func TestBooleanIsFalseCapitalOff(t *testing.T) {
	if !IsFalse("Off") {
		t.Errorf("Off should be false")
	}
}

func TestBooleanIsFalseOther(t *testing.T) {
	if IsFalse("Other") {
		t.Errorf("Other should not be false")
	}
}
