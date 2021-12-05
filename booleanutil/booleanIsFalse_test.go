package booleanutil

import "testing"

func TestIsFalseLowerYes(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("yes") {
		t.Errorf("yes should not be false")
	}
}

func TestIsFalseUpperYes(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("YES") {
		t.Errorf("YES should not be false")
	}
}

func TestIsFalseCapitalYes(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("Yes") {
		t.Errorf("Yes should not be false")
	}
}

func TestIsFalseLowerY(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("y") {
		t.Errorf("y should not be false")
	}
}

func TestIsFalseUpperY(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("Y") {
		t.Errorf("Y should not be false")
	}
}

func TestIsFalseLowerTrue(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("true") {
		t.Errorf("true should not be false")
	}
}

func TestIsFalseUpperTrue(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("TRUE") {
		t.Errorf("TRUE should not be false")
	}
}

func TestIsFalseCapitalTrue(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("True") {
		t.Errorf("True should not be false")
	}
}

func TestIsFalseLowerOn(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("on") {
		t.Errorf("on should not be false")
	}
}

func TestIsFalseUpperOn(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("ON") {
		t.Errorf("ON should not be false")
	}
}

func TestIsFalseCapitalOn(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("On") {
		t.Errorf("On should not be false")
	}
}

func TestIsFalseLowerNo(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("no") {
		t.Errorf("no should be false")
	}
}

func TestIsFalseUpperNo(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("NO") {
		t.Errorf("NO should be false")
	}
}

func TestIsFalseCapitalNo(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("No") {
		t.Errorf("No should be false")
	}
}

func TestIsFalseLowerN(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("n") {
		t.Errorf("n should be false")
	}
}

func TestIsFalseUpperN(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("N") {
		t.Errorf("N should be false")
	}
}

func TestIsFalseLowerFalse(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("false") {
		t.Errorf("false should be false")
	}
}

func TestIsFalseUpperFalse(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("FALSE") {
		t.Errorf("FALSE should be false")
	}
}

func TestIsFalseCapitalFalse(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("False") {
		t.Errorf("False should be false")
	}
}

func TestIsFalseLowerOff(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("off") {
		t.Errorf("off should be false")
	}
}

func TestIsFalseUpperOff(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("OFF") {
		t.Errorf("OFF should be false")
	}
}

func TestIsFalseCapitalOff(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsFalse("Off") {
		t.Errorf("Off should be false")
	}
}

func TestIsFalseOther(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsFalse("Other") {
		t.Errorf("Other should not be false")
	}
}
