package listutil

import "strings"

type ListUtil struct{}

func NewListUtil() *ListUtil {
	return &ListUtil{}
}

func (listUtil *ListUtil) Join(stringList []string, sep string) (joinedString string) {
	return strings.Join(stringList, sep)
}

func (listUtil *ListUtil) Contains(stringList []string, str string) bool {
	for _, value := range stringList {
		if value == str {
			return true
		}
	}
	return false
}
