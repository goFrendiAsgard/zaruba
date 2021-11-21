package strutil

type ByLenDescending []string

func (arr ByLenDescending) Len() int {
	return len(arr)
}

func (arr ByLenDescending) Less(i, j int) bool {
	return len(arr[i]) < len(arr[j])
}

func (arr ByLenDescending) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
