package utility

type ByLenDesc []string

func (arr ByLenDesc) Len() int {
	return len(arr)
}

func (arr ByLenDesc) Less(i, j int) bool {
	return len(arr[i]) < len(arr[j])
}

func (arr ByLenDesc) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
