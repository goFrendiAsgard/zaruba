package strutil

type ReplacementMapKey []string

func (arr ReplacementMapKey) Len() int {
	return len(arr)
}

func (arr ReplacementMapKey) Less(i, j int) bool {
	// longest key win
	first, second := arr[i], arr[j]
	return len(second) < len(first)
}

func (arr ReplacementMapKey) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
