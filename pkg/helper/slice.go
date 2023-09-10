package helper

func SliceContain[T comparable](arr []T, x T) bool {
	for _, v := range arr {
		if v == x {
			return true
		}
	}
	return false
}

func LastOfSlice[T comparable](arr []T) T {
	// Get the last element of the slice
	lastIndex := len(arr) - 1
	var lastElement T
	if lastIndex >= 0 {
		lastElement = arr[lastIndex]
	}
	return lastElement
}
