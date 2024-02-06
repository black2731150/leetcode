package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	for i := 0; i < m; i++ {
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
