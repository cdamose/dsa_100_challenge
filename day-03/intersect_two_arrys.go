package day03

// intersection without duplicates
func IntersectTwoArray(a []int, b []int) []int {
	var res []int
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if m[v] {
			res = append(res, v)
			delete(m, v)
		}

	}
	return res

}

// intersection with duplicates

func IntersectTwoArrayWithDuplicates(a []int, b []int) []int {
	var res []int
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}

	}
	return res

}

// intersection with duplicates merge approach
func IntersectTwoArrayWithDuplicatesMerge(a []int, b []int) []int {

	i, j := 0, 0
	var res []int
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			res = append(res, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

// intersection with duplicates merge approach
func IntersectTwoArrayWithoutDuplicatesMerge(a []int, b []int) []int {

	i, j := 0, 0
	var res []int
	for i < len(a) && j < len(b) {
		if i > 0 && a[i] == a[i-1] {
			i++
			continue
		}
		if a[i] == b[j] {
			res = append(res, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
