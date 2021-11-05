package utility

import (
	"testing"
)

func TestIsUpperTrue(t *testing.T) {
	util := NewUtil()
	expected := false
	actual := util.Str.IsUpper("Some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsUpperFalse(t *testing.T) {
	util := NewUtil()
	expected := true
	actual := util.Str.IsUpper("SOME RANDOM STRING")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsLowerTrue(t *testing.T) {
	util := NewUtil()
	expected := false
	actual := util.Str.IsLower("Some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsLowerFalse(t *testing.T) {
	util := NewUtil()
	expected := true
	actual := util.Str.IsLower("some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestToCamelCase(t *testing.T) {
	util := NewUtil()
	expected := "someRandomString"
	actual := util.Str.ToCamel("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToCamelCaseEmpty(t *testing.T) {
	util := NewUtil()
	expected := ""
	actual := util.Str.ToCamel("")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToPascalCase(t *testing.T) {
	util := NewUtil()
	expected := "SomeRandomString"
	actual := util.Str.ToPascal("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToPascalCaseEmpty(t *testing.T) {
	util := NewUtil()
	expected := ""
	actual := util.Str.ToPascal("")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToSnakeCase(t *testing.T) {
	util := NewUtil()
	expected := "some_random_string"
	actual := util.Str.ToSnake("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToKebabCase(t *testing.T) {
	util := NewUtil()
	expected := "some-random-string"
	actual := util.Str.ToKebab("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestEscapeShellArg(t *testing.T) {
	util := NewUtil()
	expected := "'a friend\\'s home'"
	actual := util.Str.EscapeShellArg("a friend's home")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestStrGetUniqueElement(t *testing.T) {
	util := NewUtil()
	expectedList := []string{"a", "b", "c", "d"}
	actualList := util.Str.GetUniqueElements([]string{"a", "a", "b", "c", "b", "c", "d", "a", "d"})
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
	util := NewUtil()
	expectedList := []string{"a", "b"}
	actualList := util.Str.GetSubKeys(
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
	util := NewUtil()
	expected := "def add(a, b):\n    return a + b"
	actual := util.Str.Indent("def add(a, b):\n  return a + b", "  ")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestReplaceByMap(t *testing.T) {
	util := NewUtil()
	expected := "orange, egg, grape, grape"
	actual := util.Str.Replace("jeruk, egg, anggur, anggur", map[string]string{"jeruk": "orange", "anggur": "grape"})
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestRepeat(t *testing.T) {
	util := NewUtil()
	expected := "ora umumora umumora umum"
	actual := util.Str.Repeat("ora umum", 3)
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestGetSingleIndentationNoSpaceOrTab(t *testing.T) {
	util := NewUtil()
	expected := ""
	actual, err := util.Str.GetIndentation("no space or tab", 3)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestGetSingleIndentationIrregularIndentation(t *testing.T) {
	util := NewUtil()
	if _, err := util.Str.GetIndentation("\t something", 3); err == nil {
		t.Errorf("error expected")
	}
}

func TestGetSingleIndentationValidIndentation(t *testing.T) {
	util := NewUtil()
	expected := "  "
	actual, err := util.Str.GetIndentation("      something", 3)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestGetFirstMatchInvalidFirstPattern(t *testing.T) {
	util := NewUtil()
	if _, _, err := util.Str.GetLineSubmatch([]string{}, []string{"[[^"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestGetFirstMatchInvalidNonFirstPattern(t *testing.T) {
	util := NewUtil()
	lines := []string{"something", "really", "interesting", "", ""}
	if _, _, err := util.Str.GetLineSubmatch(lines, []string{"something", "interesting", "[[^"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestGetFirstMatchFound(t *testing.T) {
	util := NewUtil()
	lines := []string{"something", "really", "interesting", "", "name: garo", ""}
	actualIndex, actualSubmatch, err := util.Str.GetLineSubmatch(lines, []string{"something", "interesting", "^name: (.+)$"})
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
	util := NewUtil()
	lines := []string{"something", "really", "interesting", "", "", ""}
	actualIndex, actualSubmatch, err := util.Str.GetLineSubmatch(lines, []string{"something", "interesting", "^name: (.+)$"})
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
	util := NewUtil()
	lines := []string{"something", "to be replaced", "is here"}
	if _, err := util.Str.ReplaceLineAtIndex(lines, -1, []string{"new", "and interesting"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestReplaceLineOutOfBoundIndex(t *testing.T) {
	util := NewUtil()
	lines := []string{"something", "to be replaced", "is here"}
	if _, err := util.Str.ReplaceLineAtIndex(lines, 3, []string{"new", "and interesting"}); err == nil {
		t.Errorf("error expected")
	}
}

func TestReplaceLineFirst(t *testing.T) {
	util := NewUtil()
	lines := []string{"something to be replaced", "is here"}
	actual, err := util.Str.ReplaceLineAtIndex(lines, 0, []string{"something new", "and interesting"})
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
	util := NewUtil()
	lines := []string{"something", "to be replaced", "is here"}
	actual, err := util.Str.ReplaceLineAtIndex(lines, 1, []string{"new", "and interesting"})
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
	util := NewUtil()
	lines := []string{"something", "to be replaced"}
	actual, err := util.Str.ReplaceLineAtIndex(lines, 1, []string{"new", "and interesting"})
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
	util := NewUtil()
	_, err := util.Str.CompleteLines(
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
	util := NewUtil()
	_, err := util.Str.CompleteLines(
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
	util := NewUtil()
	_, err := util.Str.CompleteLines(
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
	util := NewUtil()
	actual, err := util.Str.CompleteLines(
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
	util := NewUtil()
	actual, err := util.Str.CompleteLines(
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
