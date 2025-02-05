package day02

func MoveZeroEnd(arr []int) []int {
	ptr := 0
	for pos, val := range arr {
		if arr[pos] != 0 {
			arr[ptr] = val
			ptr++
		}
	}

	for ptr < len(arr) {
		arr[ptr] = 0
		ptr++

	}
	return arr
}

func MoveZeroEndUsingSwap(arr []int) []int {
	ptr := 0
	for pos, _ := range arr {
		if arr[pos] != 0 {
			arr[ptr], arr[pos] = arr[pos], arr[ptr]
			ptr++
		}
	}

	return arr
}


