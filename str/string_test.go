package str

import "testing"

func TestIsUpperTrue(t *testing.T) {
	expected := false
	actual := IsUpper("Some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsUpperFalse(t *testing.T) {
	expected := true
	actual := IsUpper("SOME RANDOM STRING")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsLowerTrue(t *testing.T) {
	expected := false
	actual := IsLower("Some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestIsLowerFalse(t *testing.T) {
	expected := true
	actual := IsLower("some random string")
	if actual != expected {
		t.Errorf("expected: %t, actual: %t", expected, actual)
	}
}

func TestToCamelCase(t *testing.T) {
	expected := "someRandomString"
	actual := ToCamelCase("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToCamelCaseEmpty(t *testing.T) {
	expected := ""
	actual := ToCamelCase("")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToPascalCase(t *testing.T) {
	expected := "SomeRandomString"
	actual := ToPascalCase("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToPascalCaseEmpty(t *testing.T) {
	expected := ""
	actual := ToPascalCase("")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToSnakeCase(t *testing.T) {
	expected := "some_random_string"
	actual := ToSnakeCase("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestToKebabCase(t *testing.T) {
	expected := "some-random-string"
	actual := ToKebabCase("Some random string")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestEscapeShellArg(t *testing.T) {
	expected := "'a friend\\'s home'"
	actual := EscapeShellArg("a friend's home")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestStrSingleQuoteShellValue(t *testing.T) {
	expected := "'\\\\\\\\t \\n \\`backtick\\` \\'quoted\\' \"quoted\"'"
	actual := SingleQuoteShellValue("\\t \n `backtick` 'quoted' \"quoted\"")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestStrDoubleQuoteShellValue(t *testing.T) {
	expected := "\"\\\\\\\\t \\n \\`backtick\\` 'quoted' \\\"quoted\\\"\""
	actual := DoubleQuoteShellValue("\\t \n `backtick` 'quoted' \"quoted\"")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestStrGetUniqueElement(t *testing.T) {
	expectedList := []string{"a", "b", "c", "d"}
	actualList := GetUniqueElements([]string{"a", "a", "b", "c", "b", "c", "d", "a", "d"})
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
	expectedList := []string{"a", "b"}
	actualList := GetSubKeys(
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
	expected := "def add(a, b):\n    return a + b"
	actual := Indent("def add(a, b):\n  return a + b", "  ")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
