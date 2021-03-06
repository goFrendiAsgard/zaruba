package boolean

import "strings"

// IsTrue check whether str is semantically equal to "true"
func IsTrue(str string) bool {
	lowerStr := strings.ToLower(str)
	return lowerStr == "yes" || lowerStr == "true" || lowerStr == "on"
}

// IsFalse check whether str is semantically equal to "true"
func IsFalse(str string) bool {
	lowerStr := strings.ToLower(str)
	return lowerStr == "no" || lowerStr == "false" || lowerStr == "off"
}
