package strutil

type ReplacementMapKey []string

func (arr ReplacementMapKey) Len() int {
	return len(arr)
}

func (arr ReplacementMapKey) Less(i, j int) bool {
	// longest key win
	return len(arr[i]) < len(arr[j])
}

func (arr ReplacementMapKey) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
