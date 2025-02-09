package day05

import "sort"

func MostFrequentElementArrayBruteForce(arr []int) int {
	n := len(arr)
	maxFreq := 0
	maxFreqElement := 0
	for i := 0; i < n; i++ {
		count := 1
		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count > maxFreq {
			maxFreq = count
			maxFreqElement = arr[i]
		}
	}
	return maxFreqElement
}

func MostFrequentElementArray(arr []int) int {
	m := make(map[int]int)
	maxFreq := 0
	maxFreqElement := 0
	for _, v := range arr {
		m[v]++
		if m[v] > maxFreq {
			maxFreq = m[v]
			maxFreqElement = v
		}
	}
	return maxFreqElement
}

func MostFrequentElementArrayUsingSorting(arr []int) int {
	sort.Ints(arr)
	maxFreq := 0
	count := 1
	element := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			maxFreq++
		} else {
			if maxFreq > count {
				count = maxFreq
				element = arr[i-1]
			}
			maxFreq = 0
		}
	}
	return element
}
