package main

import "fmt"

func isValidSudoku(board [][]byte) bool {

	hang := make([][]int, 9)
	shu := make([][]int, 9)
	gezi := make([][]int, 9)

	for i := 0; i <= 8; i++ {
		hang[i] = make([]int, 9)
		shu[i] = make([]int, 9)
		gezi[i] = make([]int, 9)
	}

	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			num := byteToNum(board[i][j])
			if num == 0 {
				continue
			} else {
				num = num - 1
				hang[i][num]++
				shu[j][num]++
				gezi[i/3+3*(j/3)][num]++

				if hang[i][num] > 1 || shu[j][num] > 1 || gezi[i/3+3*(j/3)][num] > 1 {
					return false
				}

			}
		}
	}

	return true
}

func byteToNum(num byte) int {
	if num == '.' {
		return 0
	} else {
		return int(num) - '0'
	}
}

func main() {
	board := [][]byte{}
	fmt.Println(isValidSudoku(board))
}
