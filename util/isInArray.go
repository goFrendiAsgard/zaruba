package util

import (
	"strings"
)

func IsInArray(needle, separator, haystackStr string) (found bool) {
	haystackList := strings.Split(haystackStr, separator)
	for _, haystack := range haystackList {
		if strings.Trim(haystack, " ") == needle {
			return true
		}
	}
	return false
}
