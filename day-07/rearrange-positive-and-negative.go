package day07

func RearrangePositiveAndNegative(arr []int) []int {
	positive_array := []int{}
	negative_array := []int{}
	positive_array_pointer := 0
	negative_array_pointer := 0
	//iterating through the array

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative_array = append(negative_array, arr[i])
		} else {
			positive_array = append(positive_array, arr[i])
		}
	}
	//merging the positive and negative arrays
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 && positive_array_pointer < len(positive_array) {
			arr[i] = positive_array[positive_array_pointer]
			positive_array_pointer++
		} else if negative_array_pointer < len(negative_array) {
			arr[i] = negative_array[negative_array_pointer]
			negative_array_pointer++
		} else if positive_array_pointer < len(positive_array) {
			arr[i] = positive_array[positive_array_pointer]
			positive_array_pointer++
		}
	}

	return arr
}

func RearrangePositiveAndNegativeUsingTwoPointers(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		if arr[left] < 0 && arr[right] < 0 {
			left++
		} else if arr[left] > 0 && arr[right] < 0 {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		} else if arr[left] > 0 && arr[right] > 0 {
			right--
		} else {
			left++
			right--
		}
	}
	return arr
}
