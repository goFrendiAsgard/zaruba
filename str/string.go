package str

import "strings"

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
