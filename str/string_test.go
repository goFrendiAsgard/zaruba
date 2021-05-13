package str

import "testing"

func TestStrReplaceAllWith(t *testing.T) {
	expected := "fruit fruit fruit vegetable"
	actual := ReplaceAllWith("strawberry grape orange vegetable", "strawberry", "grape", "orange", "fruit")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestStrReplaceAllWithSingleParameter(t *testing.T) {
	expected := "vegetable"
	actual := ReplaceAllWith("vegetable")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestStrEscapeShellValueSingleQuote(t *testing.T) {
	expected := "\\\\\\\\t \\n \\`backtick\\` \\'quoted\\' \"quoted\""
	actual := EscapeShellValue("\\t \n `backtick` 'quoted' \"quoted\"", "'")
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestStrEscapeShellValueDoubleQuote(t *testing.T) {
	expected := "\\\\\\\\t \\n \\`backtick\\` 'quoted' \\\"quoted\\\""
	actual := EscapeShellValue("\\t \n `backtick` 'quoted' \"quoted\"", "\"")
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
