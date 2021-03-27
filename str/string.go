package str

import "strings"

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
