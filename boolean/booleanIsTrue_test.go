package boolean

import "testing"

func TestBooleanIsTrueLowerYes(t *testing.T) {
	if !IsTrue("yes") {
		t.Errorf("yes should be true")
	}
}

func TestBooleanIsTrueUpperYes(t *testing.T) {
	if !IsTrue("YES") {
		t.Errorf("YES should be true")
	}
}

func TestBooleanIsTrueCapitalYes(t *testing.T) {
	if !IsTrue("Yes") {
		t.Errorf("Yes should be true")
	}
}

func TestBooleanIsTrueLowerY(t *testing.T) {
	if !IsTrue("y") {
		t.Errorf("y should be true")
	}
}

func TestBooleanIsTrueUpperY(t *testing.T) {
	if !IsTrue("Y") {
		t.Errorf("Y should be true")
	}
}

func TestBooleanIsTrueLowerTrue(t *testing.T) {
	if !IsTrue("true") {
		t.Errorf("true should be true")
	}
}

func TestBooleanIsTrueUpperTrue(t *testing.T) {
	if !IsTrue("TRUE") {
		t.Errorf("TRUE should be true")
	}
}

func TestBooleanIsTrueCapitalTrue(t *testing.T) {
	if !IsTrue("True") {
		t.Errorf("True should be true")
	}
}

func TestBooleanIsTrueLowerOn(t *testing.T) {
	if !IsTrue("on") {
		t.Errorf("on should be true")
	}
}

func TestBooleanIsTrueUpperOn(t *testing.T) {
	if !IsTrue("ON") {
		t.Errorf("ON should be true")
	}
}

func TestBooleanIsTrueCapitalOn(t *testing.T) {
	if !IsTrue("On") {
		t.Errorf("On should be true")
	}
}

func TestBooleanIsTrueLowerNo(t *testing.T) {
	if IsTrue("no") {
		t.Errorf("no should not be true")
	}
}

func TestBooleanIsTrueUpperNo(t *testing.T) {
	if IsTrue("NO") {
		t.Errorf("NO should not be true")
	}
}

func TestBooleanIsTrueCapitalNo(t *testing.T) {
	if IsTrue("No") {
		t.Errorf("No should not be true")
	}
}

func TestBooleanIsTrueLowerN(t *testing.T) {
	if IsTrue("n") {
		t.Errorf("n should not be true")
	}
}

func TestBooleanIsTrueUpperN(t *testing.T) {
	if IsTrue("N") {
		t.Errorf("N should not be true")
	}
}

func TestBooleanIsTrueLowerFalse(t *testing.T) {
	if IsTrue("false") {
		t.Errorf("false should not be true")
	}
}

func TestBooleanIsTrueUpperFalse(t *testing.T) {
	if IsTrue("FALSE") {
		t.Errorf("FALSE should not be true")
	}
}

func TestBooleanIsTrueCapitalFalse(t *testing.T) {
	if IsTrue("False") {
		t.Errorf("False should not be true")
	}
}

func TestBooleanIsTrueLowerOff(t *testing.T) {
	if IsTrue("off") {
		t.Errorf("off should not be true")
	}
}

func TestBooleanIsTrueUpperOff(t *testing.T) {
	if IsTrue("OFF") {
		t.Errorf("OFF should not be true")
	}
}

func TestBooleanIsTrueCapitalOff(t *testing.T) {
	if IsTrue("Off") {
		t.Errorf("Off should not be true")
	}
}

func TestBooleanIsTrueOther(t *testing.T) {
	if IsTrue("Other") {
		t.Errorf("Other should not not be true")
	}
}
