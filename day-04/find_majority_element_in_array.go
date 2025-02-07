package day04

import (
	"sort"
)

// using nested loop o(n^2)
func FindMajorityElementNaive(arr []int) int {
	for _, v := range arr {
		count := 0
		for _, val := range arr {
			if v == val {
				count++
			}
		}
		if count > len(arr)/2 {
			return v
		}
	}
	return 0
}

// using hash o(n) o(n)
func FindMajorityElement(arr []int) int {
	m := make(map[int]int)
	k := len(arr) / 2
	for _, v := range arr {
		m[v]++
	}
	for _, v := range m {
		if v > k {
			return k
		}
	}
	return 0

}

// using sorting o(nlogn) o(1)
func FindMajorityElementUsingSorting(arr []int) int {
	sort.Ints(arr)
	count := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
		} else {
			if count > len(arr)/2 {
				return arr[i-1]
			}
			count = 0
		}
	}
	if count > len(arr)/2 {
		return arr[len(arr)-1]
	}
	return 0
}

// using moore voting algorithm o(n) o(1)
