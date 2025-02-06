package day02

import "fmt"

func RotateByOneRight(arr []int) []int {
	last_val := arr[len(arr)-1]
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = last_val
	return arr
}

func RotateByKInRight(arr []int, k int) []int {
	k = k % len(arr)
	for i := 0; i < k; i++ {
		arr = RotateByOneRight(arr)
		fmt.Println(arr)
	}
	return arr
}

func RotateByTempRightArray(arr []int, k int) []int {
	k = k % len(arr)
	temp := make([]int, len(arr))
	for i := 0; i < len(arr)-k; i++ {
		temp[i+k] = arr[i]
	}
	for i := len(arr)-k; i < len(arr); i++ {
		temp[i-(len(arr)-k)] = arr[i]
	}

	return temp
}
