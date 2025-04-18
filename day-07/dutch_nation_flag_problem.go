package day07

import "sort"

func DutchNationalFlagProblem(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func DutchNationalFlagProblemUsingHashing(arr []int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		hash[arr[i]]++
	}

	keys := make([]int, 0, len(hash))
	for key := range hash {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	i := 0
	for _, key := range keys {
		value := hash[key]
		for j := 0; j < value; j++ {
			arr[i] = key
			i++
		}
	}
	return arr
}
