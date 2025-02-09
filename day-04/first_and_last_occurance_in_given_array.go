package day04

// using two pointer
func FindFirstandLastOccurence(arr []int, k int) (int, int) {
	startPos := -1
	endPos := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == k {
			if startPos == -1 {
				startPos = i
			}
			endPos = i
		}
	}
	return startPos, endPos
}

// FindFirstandLastOccurenceUsingBinarySearch finds the first and last occurence of k in arr using binary search
func FindFirstandLastOccurenceUsingBinarySearch(arr []int, k int) (int, int) {
	startPos := findFirstOccurence(arr, k)
	endPos := findLastOccurence(arr, k)
	return startPos, endPos
}

func findFirstOccurence(arr []int, k int) int {
	low, high := 0, len(arr)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == k {
			result = mid
			high = mid - 1
		} else if arr[mid] < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

func findLastOccurence(arr []int, k int) int {
	low, high := 0, len(arr)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == k {
			result = mid
			low = mid + 1
		} else if arr[mid] < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}
