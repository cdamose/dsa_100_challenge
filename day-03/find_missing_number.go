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
	return 0
}

// FindMissingNumberUsingFormula find missing number using formula
func FindMissingNumberUsingFormula(arr []int) int {
	n := len(arr) + 1
	sum := n * (n + 1) / 2
	for _, v := range arr {
		sum -= v
	}
	return sum
}

// FindMIssingNumber using Hash
func FindMissingNumberUsingHash(arr []int) int {
	m := make(map[int]bool)
	for _, v := range arr {
		m[v] = true
	}
	for i := 1; i < len(arr)+2; i++ {
		if !m[i] {
			return i
		}
	}
	return 0
}
