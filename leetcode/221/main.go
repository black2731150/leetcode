package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	bp := make([][]int, m)
	for i := range bp {
		bp[i] = make([]int, n)
	}

	answer := 0

	for i, v := range bp {
		for j := range v {
			if i == 0 || j == 0 {
				bp[i][j] = int(matrix[i][j] - '0')
				if bp[i][j] > answer {
					answer = bp[i][j]
				}
			} else {
				if matrix[i][j] == '0' {
					continue
				} else {
					bp[i][j] = min(bp[i-1][j], bp[i-1][j-1], bp[i][j-1]) + 1
					if bp[i][j] > answer {
						answer = bp[i][j]
					}
				}

			}
		}
	}
	return answer * answer
}

func min(x, y, z int) int {
	if x < y {
		if z < x {
			return z
		} else {
			return x
		}
	} else {
		if z < y {
			return z
		} else {
			return y
		}
	}
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}
