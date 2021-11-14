package booleanutil

import "testing"

func TestIsTrueLowerYes(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("yes") {
		t.Errorf("yes should be true")
	}
}

func TestIsTrueUpperYes(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("YES") {
		t.Errorf("YES should be true")
	}
}

func TestIsTrueCapitalYes(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("Yes") {
		t.Errorf("Yes should be true")
	}
}

func TestIsTrueLowerY(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("y") {
		t.Errorf("y should be true")
	}
}

func TestIsTrueUpperY(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("Y") {
		t.Errorf("Y should be true")
	}
}

func TestIsTrueLowerTrue(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("true") {
		t.Errorf("true should be true")
	}
}

func TestIsTrueUpperTrue(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("TRUE") {
		t.Errorf("TRUE should be true")
	}
}

func TestIsTrueCapitalTrue(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("True") {
		t.Errorf("True should be true")
	}
}

func TestIsTrueLowerOn(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("on") {
		t.Errorf("on should be true")
	}
}

func TestIsTrueUpperOn(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("ON") {
		t.Errorf("ON should be true")
	}
}

func TestIsTrueCapitalOn(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if !booleanUtil.IsTrue("On") {
		t.Errorf("On should be true")
	}
}

func TestIsTrueLowerNo(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("no") {
		t.Errorf("no should not be true")
	}
}

func TestIsTrueUpperNo(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("NO") {
		t.Errorf("NO should not be true")
	}
}

func TestIsTrueCapitalNo(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("No") {
		t.Errorf("No should not be true")
	}
}

func TestIsTrueLowerN(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("n") {
		t.Errorf("n should not be true")
	}
}

func TestIsTrueUpperN(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("N") {
		t.Errorf("N should not be true")
	}
}

func TestIsTrueLowerFalse(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("false") {
		t.Errorf("false should not be true")
	}
}

func TestIsTrueUpperFalse(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("FALSE") {
		t.Errorf("FALSE should not be true")
	}
}

func TestIsTrueCapitalFalse(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("False") {
		t.Errorf("False should not be true")
	}
}

func TestIsTrueLowerOff(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("off") {
		t.Errorf("off should not be true")
	}
}

func TestIsTrueUpperOff(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("OFF") {
		t.Errorf("OFF should not be true")
	}
}

func TestIsTrueCapitalOff(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("Off") {
		t.Errorf("Off should not be true")
	}
}

func TestIsTrueOther(t *testing.T) {
	booleanUtil := NewBooleanUtil()
	if booleanUtil.IsTrue("Other") {
		t.Errorf("Other should not not be true")
	}
}
