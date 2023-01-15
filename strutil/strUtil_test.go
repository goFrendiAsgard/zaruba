package strutil

import (
	"testing"
)

func TestIsUpperTrue(t *testing.T) {
	strutil := NewStrUtil()
	expected := false
	actual := strutil.IsUpper("Some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsUpperFalse(t *testing.T) {
	strutil := NewStrUtil()
	expected := true
	actual := strutil.IsUpper("SOME RANDOM STRING")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsLowerTrue(t *testing.T) {
	strutil := NewStrUtil()
	expected := false
	actual := strutil.IsLower("Some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsLowerFalse(t *testing.T) {
	strutil := NewStrUtil()
	expected := true
	actual := strutil.IsLower("some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestToCamelCase(t *testing.T) {
	strutil := NewStrUtil()
	expected := "someRandomString"
	actual := strutil.ToCamel("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToCamelCaseEmpty(t *testing.T) {
	strutil := NewStrUtil()
	expected := ""
	actual := strutil.ToCamel("")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToPascalCase(t *testing.T) {
	strutil := NewStrUtil()
	expected := "SomeRandomString"
	actual := strutil.ToPascal("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToPascalCaseEmpty(t *testing.T) {
	strutil := NewStrUtil()
	expected := ""
	actual := strutil.ToPascal("")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToTitle(t *testing.T) {
	strutil := NewStrUtil()
	expected := "Some Random String"
	actual := strutil.ToTitle("Some randomString")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToSnakeCase(t *testing.T) {
	strutil := NewStrUtil()
	expected := "some_random_string"
	actual := strutil.ToSnake("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToKebabCase(t *testing.T) {
	strutil := NewStrUtil()
	expected := "some-random-string"
	actual := strutil.ToKebab("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToUrlPath(t *testing.T) {
	strutil := NewStrUtil()
	expected := "/some/random-string"
	actual := strutil.ToUrlPath("Some/random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestIndent(t *testing.T) {
	strutil := NewStrUtil()
	expected := "def add(a, b):\n    return a + b"
	actual := strutil.Indent("def add(a, b):\n  return a + b", "  ")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestStrReplace(t *testing.T) {
	expected := "orange, egg, grape, grape"
	actual := StrReplace("jeruk, egg, anggur, anggur", map[string]string{"jeruk": "orange", "anggur": "grape"})
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestRepeat(t *testing.T) {
	strutil := NewStrUtil()
	expected := "ora umumora umumora umum"
	actual := strutil.Repeat("ora umum", 3)
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestGetSingleIndentationNoSpaceOrTab(t *testing.T) {
	strutil := NewStrUtil()
	expected := ""
	actual, err := strutil.GetIndentation("no space or tab", 3)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestGetSingleIndentationIrregularIndentation(t *testing.T) {
	strutil := NewStrUtil()
	if _, err := strutil.GetIndentation("\t something", 3); err == nil {
		t.Errorf("error expected")
	}
}

func TestGetSingleIndentationValidIndentation(t *testing.T) {
	strutil := NewStrUtil()
	expected := "  "
	actual, err := strutil.GetIndentation("      something", 3)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestGetFirstMatchInvalidFirstPattern(t *testing.T) {
	if _, _, err := StrGetLineSubmatch([]string{}, []string{"[[^"}, 0); err == nil {
		t.Errorf("error expected")
	}
}

func TestGetFirstMatchInvalidNonFirstPattern(t *testing.T) {
	lines := []string{"something", "really", "interesting", "", ""}
	if _, _, err := StrGetLineSubmatch(lines, []string{"something", "interesting", "[[^"}, 2); err == nil {
		t.Errorf("error expected")
	}
}

func TestGetFirstMatchFound(t *testing.T) {
	lines := []string{"something", "really", "interesting", "", "name: garo", ""}
	actualIndex, actualSubmatch, err := StrGetLineSubmatch(lines, []string{"something", "interesting", "^name: (.+)$"}, 2)
	if err != nil {
		t.Error(err)
	}
	expectedIndex, expectedSubmatch := 4, []string{"name: garo", "garo"}
	if actualIndex != expectedIndex {
		t.Errorf("expected index: %d, actual index: %d", expectedIndex, actualIndex)
	}
	if len(actualSubmatch) != len(expectedSubmatch) {
		t.Errorf("expected submatch: %#v, actual submatch: %#v", expectedSubmatch, actualSubmatch)
		return
	}
	for index, expectedStr := range expectedSubmatch {
		actualStr := actualSubmatch[index]
		if expectedStr != actualStr {
			t.Errorf("expected submatch: %#v, actual submatch: %#v", expectedSubmatch, actualSubmatch)
			return
		}
	}
}

func TestGetFirstMatchNotFound(t *testing.T) {
	lines := []string{"something", "really", "interesting", "", "", ""}
	actualIndex, actualSubmatch, err := StrGetLineSubmatch(lines, []string{"something", "interesting", "^name: (.+)$"}, 2)
	if err != nil {
		t.Error(err)
	}
	expectedIndex, expectedSubmatch := -1, []string{}
	if actualIndex != expectedIndex {
		t.Errorf("expected index: %d, actual index: %d", expectedIndex, actualIndex)
	}
	if len(actualSubmatch) != len(expectedSubmatch) {
		t.Errorf("expected submatch: %#v, actual submatch: %#v", expectedSubmatch, actualSubmatch)
		return
	}
}

func TestReplaceLineNegativeIndex(t *testing.T) {
	lines := []string{"something", "to be replaced", "is here"}
	if _, err := StrReplaceLineAtIndex(lines, -1, []string{"new", "and interesting"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestReplaceLineOutOfBoundIndex(t *testing.T) {
	lines := []string{"something", "to be replaced", "is here"}
	if _, err := StrReplaceLineAtIndex(lines, 3, []string{"new", "and interesting"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestReplaceLineFirst(t *testing.T) {
	lines := []string{"something to be replaced", "is here"}
	actual, err := StrReplaceLineAtIndex(lines, 0, []string{"something new", "and interesting"})
	if err != nil {
		t.Error(err)
	}
	expected := []string{"something new", "and interesting", "is here"}
	if len(actual) != len(expected) {
		t.Errorf("expected: %#v, actual: %#v", expected, actual)
		return
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected: %#v, actual: %#v", expected, actual)
			return
		}
	}
}

func TestReplaceLineMidsle(t *testing.T) {
	lines := []string{"something", "to be replaced", "is here"}
	actual, err := StrReplaceLineAtIndex(lines, 1, []string{"new", "and interesting"})
	if err != nil {
		t.Error(err)
	}
	expected := []string{"something", "new", "and interesting", "is here"}
	if len(actual) != len(expected) {
		t.Errorf("expected: %#v, actual: %#v", expected, actual)
		return
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected: %#v, actual: %#v", expected, actual)
			return
		}
	}
}

func TestReplaceLineLast(t *testing.T) {
	lines := []string{"something", "to be replaced"}
	actual, err := StrReplaceLineAtIndex(lines, 1, []string{"new", "and interesting"})
	if err != nil {
		t.Error(err)
	}
	expected := []string{"something", "new", "and interesting"}
	if len(actual) != len(expected) {
		t.Errorf("expected: %#v, actual: %#v", expected, actual)
		return
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected: %#v, actual: %#v", expected, actual)
			return
		}
	}
}
