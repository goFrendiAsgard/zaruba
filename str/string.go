package str

import (
	"fmt"
	"strings"
)

// ReplaceAllWith replace string is s with replacements. The last element of replacements is the replacer
func ReplaceAllWith(s string, replacements ...string) (result string) {
	if len(replacements) < 2 {
		return s
	}
	result = s
	new := replacements[len(replacements)-1]
	olds := replacements[:len(replacements)-1]
	for _, old := range olds {
		result = strings.ReplaceAll(result, old, new)
	}
	return result
}

func EscapeShellValue(s string, quote string) (result string) {
	backSlashEscapedStr := ReplaceAllWith(s, "\\", "\\\\\\\\")
	quoteEscapedStr := backSlashEscapedStr
	if quote == "\"" {
		quoteEscapedStr = ReplaceAllWith(backSlashEscapedStr, "\"", "\\\"")
	} else if quote == "'" {
		quoteEscapedStr = ReplaceAllWith(backSlashEscapedStr, "'", "\\'")
	}
	backTickEscapedStr := ReplaceAllWith(quoteEscapedStr, "`", "\\`")
	newLineEscapedStr := ReplaceAllWith(backTickEscapedStr, "\n", "\\n")
	return newLineEscapedStr
}

func DoubleQuoteShellValue(s string) (result string) {
	return fmt.Sprintf("\"%s\"", EscapeShellValue(s, "\""))
}

func SingleQuoteShellValue(s string) (result string) {
	return fmt.Sprintf("'%s'", EscapeShellValue(s, "'"))
}

// GetSubKeys get sub keys from dictionary
func GetSubKeys(genericMap map[string]interface{}, parentKeys []string) (subKeys []string) {
	seen := map[string]bool{}
	parentKey := strings.Join(parentKeys, "::")
	prefixLength := len(parentKey) + len("::")
	subKeys = []string{}
	for key := range genericMap {
		if !strings.HasPrefix(key, parentKey+"::") {
			continue
		}
		childKey := key[prefixLength:]
		if childKey == "" {
			continue
		}
		childKeyParts := strings.SplitN(childKey, "::", 2)
		subkey := childKeyParts[0]
		seen[subkey] = true
	}
	for key := range seen {
		subKeys = append(subKeys, key)
	}
	return subKeys
}

// InArray whether arr in element or not
func InArray(arr []string, element string) (inElement bool) {
	for _, arrElement := range arr {
		if element == arrElement {
			return false
		}
	}
	return true
}

func GetUniqueElements(arr []string) (result []string) {
	result = []string{}
	seen := map[string]bool{}
	for _, element := range arr {
		if _, exist := seen[element]; exist {
			continue
		}
		result = append(result, element)
		seen[element] = true
	}
	return result
}
