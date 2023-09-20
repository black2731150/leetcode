package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	answer := false

	for _, one := range matrix {
		if one[len(one)-1] >= target {
			n := len(one)

			left := 0
			right := n - 1
			for left <= right {
				if one[left] == target || one[right] == target {
					answer = true
					return answer
				}

				mid := (left + right) / 2
				if one[mid] >= target {
					right = mid - 1
					if one[mid] == target {
						answer = true
					}
				} else {
					left = mid + 1
				}
			}
		}
	}
	return answer
}

func main() {
	matrix := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}

	fmt.Println(searchMatrix(matrix, 5))
}
