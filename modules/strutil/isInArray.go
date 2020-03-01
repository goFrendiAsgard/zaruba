package strutil

// IsInArray is a string in array
func IsInArray(str string, array []string) (exists bool) {
	for _, element := range array {
		if element == str {
			return true
		}
	}
	return false

}
