package day02

func RotateByOneLeft(arr []int) []int {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp

	return arr

}

func RotateByKInLeft(arr []int, k int) []int {
	for i := 0; i < k; i++ {
		arr = RotateByOneLeft(arr)
	}
	return arr
}

func RotateByTempArray(arr []int, k int) []int {
	if k > len(arr) {
		k = k % len(arr)
	}
	temp := make([]int, len(arr))
	for i := k; i < len(arr); i++ {
		temp[i-k] = arr[i]
	}
	for i := 0; i < k; i++ {
		temp[len(arr)-k+i] = arr[i]

	}
	return temp

}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func RotateByKUisngJuggling(arr []int, k int) []int {
	cycles := GCD(len(arr), k)
	for i := 0; i < cycles; i++ {
		currIndex := i
		startEle := arr[i]
		for {
			next_index := (currIndex + k) % len(arr)
			if next_index == i {
				break
			}
			arr[currIndex] = arr[next_index]
			currIndex = next_index
		}
		arr[currIndex] = startEle

	}
	return arr

}

func Reverese(arr []int, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

}

func RotateByReversal(arr []int, k int) []int {
	n := len(arr)
	k = k % n
	Reverese(arr, 0, k-1)
	Reverese(arr, k, n-1)
	Reverese(arr, 0, n-1)
	return arr
}
