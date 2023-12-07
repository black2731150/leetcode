package main

import "fmt"

func diagonalSum(mat [][]int) int {
	n := len(mat)

	answer := 0
	left, right := 0, n-1
	i := 0
	for left < n {
		if left != right {
			answer += mat[i][left] + mat[i][right]
		} else {
			answer += mat[i][left]
		}
		i++
		left++
		right--
	}
	return answer
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(diagonalSum(mat))
}
