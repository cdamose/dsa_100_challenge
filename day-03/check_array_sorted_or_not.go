package day03

func IsArraySorted(arr []int) bool {
	if len(arr) == 0 {
		return true
	}
	for pos := 0; pos < len(arr)-1; pos++ {
		if arr[pos] > arr[pos+1] {
			return false
		}
	}
	return true
}
