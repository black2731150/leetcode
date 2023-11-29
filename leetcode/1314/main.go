package main

import "fmt"

func matrixBlockSum(mat [][]int, k int) [][]int {
	erWerQianZhuiHe := make([][]int, len(mat)+1)
	for i := range erWerQianZhuiHe {
		erWerQianZhuiHe[i] = make([]int, len(mat[0])+1)
	}

	// 构建二维前缀和
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			erWerQianZhuiHe[i+1][j+1] = erWerQianZhuiHe[i][j+1] + erWerQianZhuiHe[i+1][j] - erWerQianZhuiHe[i][j] + mat[i][j]
		}
	}

	answer := make([][]int, len(mat))
	for i := range answer {
		answer[i] = make([]int, len(mat[0]))
	}

	for i := range answer {
		for j := range answer[i] {
			left, right := max(i-k, 0), min(i+k, len(mat)-1)
			up, down := max(j-k, 0), min(j+k, len(mat[0])-1)

			answer[i][j] = erWerQianZhuiHe[right+1][down+1] - erWerQianZhuiHe[right+1][up] - erWerQianZhuiHe[left][down+1] + erWerQianZhuiHe[left][up]
		}
	}

	return answer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 1
	for _, v := range matrixBlockSum(mat, k) {
		fmt.Println(v)
	}
}
