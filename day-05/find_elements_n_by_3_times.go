package day05

func FindElementsNBy3Times(arr []int) []int {
	var res []int
	n := len(arr)
	count1, count2 := 0, 0
	first, second := 0, 1
	for i := 0; i < n; i++ {
		if arr[i] == arr[first] {
			count1++
		} else if arr[i] == arr[second] {
			count2++
		} else if count1 == 0 {
			count1 = 1
			first = i
		} else if count2 == 0 {
			count2 = 1
			second = i
		} else {
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	for i := 0; i < n; i++ {
		if arr[i] == arr[first] {
			count1++
		} else if arr[i] == arr[second] {
			count2++
		}
	}
	if count1 > n/3 {
		res = append(res, arr[first])
	}
	if count2 > n/3 {
		res = append(res, arr[second])
	}
	return res
}

func FindElementsNBy3TimesHash(arr []int) []int {
	var res []int
	n := len(arr)
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		if v > n/3 {
			res = append(res, k)
		}
	}
	return res
}
