package boolean

import "testing"

func TestIsFalseLowerYes(t *testing.T) {
	if IsFalse("yes") {
		t.Errorf("yes should not be false")
	}
}

func TestIsFalseUpperYes(t *testing.T) {
	if IsFalse("YES") {
		t.Errorf("YES should not be false")
	}
}

func TestIsFalseCapitalYes(t *testing.T) {
	if IsFalse("Yes") {
		t.Errorf("Yes should not be false")
	}
}

func TestIsFalseLowerY(t *testing.T) {
	if IsFalse("y") {
		t.Errorf("y should not be false")
	}
}

func TestIsFalseUpperY(t *testing.T) {
	if IsFalse("Y") {
		t.Errorf("Y should not be false")
	}
}

func TestIsFalseLowerTrue(t *testing.T) {
	if IsFalse("true") {
		t.Errorf("true should not be false")
	}
}

func TestIsFalseUpperTrue(t *testing.T) {
	if IsFalse("TRUE") {
		t.Errorf("TRUE should not be false")
	}
}

func TestIsFalseCapitalTrue(t *testing.T) {
	if IsFalse("True") {
		t.Errorf("True should not be false")
	}
}

func TestIsFalseLowerOn(t *testing.T) {
	if IsFalse("on") {
		t.Errorf("on should not be false")
	}
}

func TestIsFalseUpperOn(t *testing.T) {
	if IsFalse("ON") {
		t.Errorf("ON should not be false")
	}
}

func TestIsFalseCapitalOn(t *testing.T) {
	if IsFalse("On") {
		t.Errorf("On should not be false")
	}
}

func TestIsFalseLowerNo(t *testing.T) {
	if !IsFalse("no") {
		t.Errorf("no should be false")
	}
}

func TestIsFalseUpperNo(t *testing.T) {
	if !IsFalse("NO") {
		t.Errorf("NO should be false")
	}
}

func TestIsFalseCapitalNo(t *testing.T) {
	if !IsFalse("No") {
		t.Errorf("No should be false")
	}
}

func TestIsFalseLowerN(t *testing.T) {
	if !IsFalse("n") {
		t.Errorf("n should be false")
	}
}

func TestIsFalseUpperN(t *testing.T) {
	if !IsFalse("N") {
		t.Errorf("N should be false")
	}
}

func TestIsFalseLowerFalse(t *testing.T) {
	if !IsFalse("false") {
		t.Errorf("false should be false")
	}
}

func TestIsFalseUpperFalse(t *testing.T) {
	if !IsFalse("FALSE") {
		t.Errorf("FALSE should be false")
	}
}

func TestIsFalseCapitalFalse(t *testing.T) {
	if !IsFalse("False") {
		t.Errorf("False should be false")
	}
}

func TestIsFalseLowerOff(t *testing.T) {
	if !IsFalse("off") {
		t.Errorf("off should be false")
	}
}

func TestIsFalseUpperOff(t *testing.T) {
	if !IsFalse("OFF") {
		t.Errorf("OFF should be false")
	}
}

func TestIsFalseCapitalOff(t *testing.T) {
	if !IsFalse("Off") {
		t.Errorf("Off should be false")
	}
}

func TestIsFalseOther(t *testing.T) {
	if IsFalse("Other") {
		t.Errorf("Other should not be false")
	}
}
