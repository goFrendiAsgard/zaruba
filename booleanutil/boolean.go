package booleanutil

import "strings"

type BooleanUtil struct{}

func NewBooleanUtil() *BooleanUtil {
	return &BooleanUtil{}
}

// IsTrue check whether str is semantically equal to "true"
func (booleanUtil *BooleanUtil) IsTrue(str string) bool {
	lowerStr := strings.ToLower(str)
	return lowerStr == "yes" || lowerStr == "y" || lowerStr == "true" || lowerStr == "on" || lowerStr == "1"
}

// IsFalse check whether str is semantically equal to "true"
func (booleanUtil *BooleanUtil) IsFalse(str string) bool {
	lowerStr := strings.ToLower(str)
	return lowerStr == "no" || lowerStr == "n" || lowerStr == "false" || lowerStr == "off" || lowerStr == "0"
}
