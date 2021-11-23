package listutil

import "strings"

type ListUtil struct{}

func NewListUtil() *ListUtil {
	return &ListUtil{}
}

func (listUtil *ListUtil) Join(stringList []string, sep string) (joinedString string) {
	return strings.Join(stringList, sep)
}
