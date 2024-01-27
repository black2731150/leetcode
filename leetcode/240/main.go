package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])

	for i := range matrix {
		if matrix[i][0] <= target && matrix[i][n-1] >= target {
			left, right := 0, n-1
			for left <= right {
				mid := (left + right) / 2

				if matrix[i][mid] == target {
					return true
				}

				if matrix[i][mid] > target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}

	return false
}

func main() {
	mayrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 20
	fmt.Println(searchMatrix(mayrix, target))
}
