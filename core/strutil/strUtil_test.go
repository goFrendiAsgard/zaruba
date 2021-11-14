package strutil

import (
	"testing"
)

func TestIsUpperTrue(t *testing.T) {
	strutil := NewStrutil()
	expected := false
	actual := strutil.IsUpper("Some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsUpperFalse(t *testing.T) {
	strutil := NewStrutil()
	expected := true
	actual := strutil.IsUpper("SOME RANDOM STRING")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsLowerTrue(t *testing.T) {
	strutil := NewStrutil()
	expected := false
	actual := strutil.IsLower("Some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsLowerFalse(t *testing.T) {
	strutil := NewStrutil()
	expected := true
	actual := strutil.IsLower("some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestToCamelCase(t *testing.T) {
	strutil := NewStrutil()
	expected := "someRandomString"
	actual := strutil.ToCamel("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToCamelCaseEmpty(t *testing.T) {
	strutil := NewStrutil()
	expected := ""
	actual := strutil.ToCamel("")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToPascalCase(t *testing.T) {
	strutil := NewStrutil()
	expected := "SomeRandomString"
	actual := strutil.ToPascal("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToPascalCaseEmpty(t *testing.T) {
	strutil := NewStrutil()
	expected := ""
	actual := strutil.ToPascal("")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToSnakeCase(t *testing.T) {
	strutil := NewStrutil()
	expected := "some_random_string"
	actual := strutil.ToSnake("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToKebabCase(t *testing.T) {
	strutil := NewStrutil()
	expected := "some-random-string"
	actual := strutil.ToKebab("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestEscapeShellArg(t *testing.T) {
	strutil := NewStrutil()
	expected := "'a friend\\'s home'"
	actual := strutil.EscapeShellArg("a friend's home")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestStrGetUniqueElement(t *testing.T) {
	strutil := NewStrutil()
	expectedList := []string{"a", "b", "c", "d"}
	actualList := strutil.GetUniqueElements([]string{"a", "a", "b", "c", "b", "c", "d", "a", "d"})
	if len(actualList) != len(expectedList) {
		t.Errorf("expected: %#v, actual: %#v", expectedList, actualList)
	}
	for index, expected := range expectedList {
		actual := actualList[index]
		if actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
}

func TestStrGetSubKeys(t *testing.T) {
	strutil := NewStrutil()
	expectedList := []string{"a", "b"}
	actualList := strutil.GetSubKeys(
		[]string{"key::a::name", "key::b::name", "key::a::address", "key::b::address", "key", "otherKey", "key::"},
		[]string{"key"},
	)
	if len(actualList) != len(expectedList) {
		t.Errorf("expected: %#v, actual: %#v", expectedList, actualList)
	}
	for _, expected := range expectedList {
		actualFound := false
		for _, actual := range actualList {
			if actual == expected {
				actualFound = true
			}
		}
		if !actualFound {
			t.Errorf("cannot find key %s, on: %#v", expected, actualList)
		}
	}
}

func TestIndent(t *testing.T) {
	strutil := NewStrutil()
	expected := "def add(a, b):\n    return a + b"
	actual := strutil.Indent("def add(a, b):\n  return a + b", "  ")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestReplaceByMap(t *testing.T) {
	strutil := NewStrutil()
	expected := "orange, egg, grape, grape"
	actual := strutil.Replace("jeruk, egg, anggur, anggur", map[string]string{"jeruk": "orange", "anggur": "grape"})
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestRepeat(t *testing.T) {
	strutil := NewStrutil()
	expected := "ora umumora umumora umum"
	actual := strutil.Repeat("ora umum", 3)
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestGetSingleIndentationNoSpaceOrTab(t *testing.T) {
	strutil := NewStrutil()
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
	strutil := NewStrutil()
	if _, err := strutil.GetIndentation("\t something", 3); err == nil {
		t.Errorf("error expected")
	}
}

func TestGetSingleIndentationValidIndentation(t *testing.T) {
	strutil := NewStrutil()
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
	strutil := NewStrutil()
	if _, _, err := strutil.GetLineSubmatch([]string{}, []string{"[[^"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestGetFirstMatchInvalidNonFirstPattern(t *testing.T) {
	strutil := NewStrutil()
	lines := []string{"something", "really", "interesting", "", ""}
	if _, _, err := strutil.GetLineSubmatch(lines, []string{"something", "interesting", "[[^"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestGetFirstMatchFound(t *testing.T) {
	strutil := NewStrutil()
	lines := []string{"something", "really", "interesting", "", "name: garo", ""}
	actualIndex, actualSubmatch, err := strutil.GetLineSubmatch(lines, []string{"something", "interesting", "^name: (.+)$"})
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
	strutil := NewStrutil()
	lines := []string{"something", "really", "interesting", "", "", ""}
	actualIndex, actualSubmatch, err := strutil.GetLineSubmatch(lines, []string{"something", "interesting", "^name: (.+)$"})
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
	strutil := NewStrutil()
	lines := []string{"something", "to be replaced", "is here"}
	if _, err := strutil.ReplaceLineAtIndex(lines, -1, []string{"new", "and interesting"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestReplaceLineOutOfBoundIndex(t *testing.T) {
	strutil := NewStrutil()
	lines := []string{"something", "to be replaced", "is here"}
	if _, err := strutil.ReplaceLineAtIndex(lines, 3, []string{"new", "and interesting"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestReplaceLineFirst(t *testing.T) {
	strutil := NewStrutil()
	lines := []string{"something to be replaced", "is here"}
	actual, err := strutil.ReplaceLineAtIndex(lines, 0, []string{"something new", "and interesting"})
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

func TestReplaceLineMiddle(t *testing.T) {
	strutil := NewStrutil()
	lines := []string{"something", "to be replaced", "is here"}
	actual, err := strutil.ReplaceLineAtIndex(lines, 1, []string{"new", "and interesting"})
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
	strutil := NewStrutil()
	lines := []string{"something", "to be replaced"}
	actual, err := strutil.ReplaceLineAtIndex(lines, 1, []string{"new", "and interesting"})
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

func TestCompleteLinesDifferentPatternAndSuplementLength(t *testing.T) {
	strutil := NewStrutil()
	_, err := strutil.CompleteLines(
		[]string{},
		[]string{
			"^task:(.*)$",
		},
		[]string{},
	)
	if err == nil {
		t.Errorf("error expected")
	}
}

func TestCompleteLinesInvalidPattern(t *testing.T) {
	strutil := NewStrutil()
	_, err := strutil.CompleteLines(
		[]string{},
		[]string{
			"[[^",
		},
		[]string{
			"something",
		},
	)
	if err == nil {
		t.Errorf("error expected")
	}
}

func TestCompleteLinesUnmatchPattern(t *testing.T) {
	strutil := NewStrutil()
	_, err := strutil.CompleteLines(
		[]string{},
		[]string{
			"ab",
		},
		[]string{
			"something",
		},
	)
	if err == nil {
		t.Errorf("error expected")
	}
}

func TestCompleteLinesNoMatchAtAll(t *testing.T) {
	strutil := NewStrutil()
	actual, err := strutil.CompleteLines(
		[]string{"includes: []"},
		[]string{
			"^tasks:(.*)$",
			"^ +myTask:(.*$)",
			"^ +configs:(.*$)",
			"^ +ports:(.*$)",
		},
		[]string{
			"tasks:",
			"  myTask:",
			"    configs:",
			"      ports: 8080",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	expected := []string{
		"includes: []",
		"tasks:",
		"  myTask:",
		"    configs:",
		"      ports: 8080",
	}
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

func TestCompleteLinesNoMatchPartial(t *testing.T) {
	strutil := NewStrutil()
	actual, err := strutil.CompleteLines(
		[]string{
			"includes: []",
			"tasks: # list of task",
			"  someTask: {}",
		},
		[]string{
			"^tasks:(.*)$",
			"^ +myTask:(.*)$",
			"^ +configs:(.*)$",
			"^ +ports:(.*)$",
		},
		[]string{
			"tasks:",
			"  myTask:",
			"    configs:",
			"      ports: 8080",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	expected := []string{
		"includes: []",
		"tasks: # list of task",
		"  myTask:",
		"    configs:",
		"      ports: 8080",
		"  someTask: {}",
	}
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
