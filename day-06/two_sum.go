package day06

import "sort"

func TwoSumBruteForce(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSumUsingHash(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func TwoSumUsingSortingWithTwoPointer(nums []int, target int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	l, r := 0, len(sorted)-1
	for l < r {
		sum := sorted[l] + sorted[r]
		if sum == target {
			return []int{l, r}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
