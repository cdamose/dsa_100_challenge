package day04

// using hash
func FindOnceApperanceUisngHash(arr []int) int {
	m := make(map[int]int)
	//initialize as 1
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0

}

// using XOR
func FindOnceApperanceUisngXOR(arr []int) int {
	res := 0
	for _, v := range arr {
		res = res ^ v
	}
	return res
}
