package main

import "fmt"

func firstCompleteIndex(arr []int, mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	hangMap := make(map[int]int)
	lieMap := make(map[int]int)

	indexMap := make(map[int][2]int, m*n)

	for i, v := range mat {
		for j, v2 := range v {
			indexMap[v2] = [2]int{i, j}
		}
	}

	for i, v := range arr {
		index := indexMap[v]
		hangindex := index[0]
		lieindex := index[1]

		hangMap[hangindex]++
		if hangMap[hangindex] == n {
			return i
		}

		lieMap[lieindex]++
		if lieMap[lieindex] == m {
			return i
		}
	}
	return 0
}

func main() {
	arr := []int{2, 8, 7, 4, 1, 3, 5, 6, 9}
	mat := [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}
	fmt.Println(firstCompleteIndex(arr, mat))
}
