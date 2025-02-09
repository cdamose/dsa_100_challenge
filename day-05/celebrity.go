package day05

func FindClebrityUsingBruteForce(matrix [][]int) int {
	n := len(matrix)
	for i := 0; i < n; i++ {
		isCelebrity := true
		for j := 0; j < n; j++ {
			if i != j && (matrix[i][j] == 1 || matrix[j][i] == 0) {
				isCelebrity = false
				break
			}
		}
		if isCelebrity {
			return i
		}
	}
	return -1
}

func FindClebrityUsingTwoPOinter(mat [][]int) int {
	n := len(mat)
	i := 0
	j := n - 1
	for i < j {
		if mat[j][i] == 1 {
			j++
		} else {
			i++
		}

	}
	candidate := i

	for i := 0; i < n; i++ {
		if i != candidate && (mat[candidate][i] == 1 || mat[i][candidate] == 0) {
			return -1
		}

	}
	return candidate

}

func FindClebrityUsingStack(mat [][]int) int {
	n := len(mat)
	var stack []int
	for i := 0; i < n; i++ {
		stack = append(stack, i)
	}
	for len(stack) > 1 {
		a := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		b := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if mat[a][b] == 1 {
			stack = append(stack, b)
		} else {
			stack = append(stack, a)
		}
	}
	candidate := stack[0]
	for i := 0; i < n; i++ {
		if i != candidate && (mat[candidate][i] == 1 || mat[i][candidate] == 0) {
			return -1
		}
	}
	return candidate
}
