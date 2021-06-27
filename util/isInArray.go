package util

import (
	"strings"
)

func IsInArray(needle, haystackStr, separator string) (found bool) {
	haystackList := strings.Split(haystackStr, separator)
	for _, haystack := range haystackList {
		if strings.Trim(haystack, " ") == needle {
			return true
		}
	}
	return false
}
