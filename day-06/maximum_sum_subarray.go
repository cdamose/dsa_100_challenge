package day06

func MaximumSumSubArray(arr []int) int {
	max_sum := arr[0]
	curr_sum := arr[0]
	for i := 1; i < len(arr); i++ {
		curr_sum = max(arr[i], curr_sum+arr[i])
		max_sum = max(max_sum, curr_sum)
	}
	return max_sum
}

func MaximumSumSubArrayWithIndex(arr []int) (int, int, int) {
	max_sum := arr[0]
	curr_sum := arr[0]
	start := 0
	end := 0
	s := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > curr_sum+arr[i] {
			s = i
		}
		curr_sum = max(arr[i], curr_sum+arr[i])
		if curr_sum > max_sum {
			start = s
			end = i
		}
		max_sum = max(max_sum, curr_sum)
	}
	return max_sum, start, end
}
