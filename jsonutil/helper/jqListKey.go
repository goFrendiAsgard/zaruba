package helper

import (
	"regexp"
	"strconv"
)

var listKeyPattern = regexp.MustCompile(`\[([\-]?[0-9]+)\]`)

func isListKey(key string) bool {
	return listKeyPattern.MatchString(key)
}

func getIndexFromKey(key string) int {
	matches := listKeyPattern.FindStringSubmatch(key)
	if len(matches) > 0 {
		index, err := strconv.Atoi(matches[1])
		if err != nil {
			return -1
		}
		return index
	}
	return -1
}
