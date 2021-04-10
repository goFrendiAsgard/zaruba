package boolean

import "testing"

func TestIsTrueLowerYes(t *testing.T) {
	if !IsTrue("yes") {
		t.Errorf("yes should be true")
	}
}

func TestIsTrueUpperYes(t *testing.T) {
	if !IsTrue("YES") {
		t.Errorf("YES should be true")
	}
}

func TestIsTrueCapitalYes(t *testing.T) {
	if !IsTrue("Yes") {
		t.Errorf("Yes should be true")
	}
}

func TestIsTrueLowerY(t *testing.T) {
	if !IsTrue("y") {
		t.Errorf("y should be true")
	}
}

func TestIsTrueUpperY(t *testing.T) {
	if !IsTrue("Y") {
		t.Errorf("Y should be true")
	}
}

func TestIsTrueLowerTrue(t *testing.T) {
	if !IsTrue("true") {
		t.Errorf("true should be true")
	}
}

func TestIsTrueUpperTrue(t *testing.T) {
	if !IsTrue("TRUE") {
		t.Errorf("TRUE should be true")
	}
}

func TestIsTrueCapitalTrue(t *testing.T) {
	if !IsTrue("True") {
		t.Errorf("True should be true")
	}
}

func TestIsTrueLowerOn(t *testing.T) {
	if !IsTrue("on") {
		t.Errorf("on should be true")
	}
}

func TestIsTrueUpperOn(t *testing.T) {
	if !IsTrue("ON") {
		t.Errorf("ON should be true")
	}
}

func TestIsTrueCapitalOn(t *testing.T) {
	if !IsTrue("On") {
		t.Errorf("On should be true")
	}
}

func TestIsTrueLowerNo(t *testing.T) {
	if IsTrue("no") {
		t.Errorf("no should not be true")
	}
}

func TestIsTrueUpperNo(t *testing.T) {
	if IsTrue("NO") {
		t.Errorf("NO should not be true")
	}
}

func TestIsTrueCapitalNo(t *testing.T) {
	if IsTrue("No") {
		t.Errorf("No should not be true")
	}
}

func TestIsTrueLowerN(t *testing.T) {
	if IsTrue("n") {
		t.Errorf("n should not be true")
	}
}

func TestIsTrueUpperN(t *testing.T) {
	if IsTrue("N") {
		t.Errorf("N should not be true")
	}
}

func TestIsTrueLowerFalse(t *testing.T) {
	if IsTrue("false") {
		t.Errorf("false should not be true")
	}
}

func TestIsTrueUpperFalse(t *testing.T) {
	if IsTrue("FALSE") {
		t.Errorf("FALSE should not be true")
	}
}

func TestIsTrueCapitalFalse(t *testing.T) {
	if IsTrue("False") {
		t.Errorf("False should not be true")
	}
}

func TestIsTrueLowerOff(t *testing.T) {
	if IsTrue("off") {
		t.Errorf("off should not be true")
	}
}

func TestIsTrueUpperOff(t *testing.T) {
	if IsTrue("OFF") {
		t.Errorf("OFF should not be true")
	}
}

func TestIsTrueCapitalOff(t *testing.T) {
	if IsTrue("Off") {
		t.Errorf("Off should not be true")
	}
}

func TestIsTrueOther(t *testing.T) {
	if IsTrue("Other") {
		t.Errorf("Other should not not be true")
	}
}
