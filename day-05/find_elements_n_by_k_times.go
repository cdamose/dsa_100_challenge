package day05

func FindElementsNByKTimesUsingHashing(arr []int, k int) []int {
	var res []int
	n := len(arr)
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		if v > n/k {
			res = append(res, k)
		}
	}
	return res

}
