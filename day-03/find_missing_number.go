package day03

// using formula n+1 * n+2 / 2

func FindMissingNumber(arr []int) int {
	for i := 1; i < len(arr)+2; i++ {
		found := false
		for j := 0; j < len(arr); j++ {
			if arr[j] == i {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return -1
}
