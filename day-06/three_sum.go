package day06

func TripletSumBruteForce(arr []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == target {
					result = append(result, []int{arr[i], arr[j], arr[k]})
				}
			}
		}
	}
	return result
}

func TripletSumUsingHashing(arr []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(arr)-1; i++ {
		hash := make(map[int]int)
		currSum := target - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if _, ok := hash[currSum-arr[j]]; ok {
				result = append(result, []int{arr[i], arr[j], currSum - arr[j]})
			}
			hash[arr[j]] = j
		}
	}
	return result
}

func TripletSumUsingSorting(arr []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			currSum := arr[i] + arr[left] + arr[right]
			if currSum == target {
				result = append(result, []int{arr[i], arr[left], arr[right]})
				left++
				right--
			} else if currSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
