package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	var help func(x, y int, index int) bool
	help = func(x, y, index int) bool {
		if index == len(word) {
			return true
		}

		if x < 0 || x >= m || y < 0 || y >= n || index >= len(word) || word[index] != board[x][y] {
			return false
		}

		tmp := board[x][y]
		board[x][y] = '#'
		suucess := help(x+1, y, index+1) || help(x-1, y, index+1) || help(x, y+1, index+1) || help(x, y-1, index+1)
		board[x][y] = tmp
		return suucess
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if help(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"
	fmt.Println(exist(board, word))
}
