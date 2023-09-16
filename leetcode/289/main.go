package main

import "fmt"

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])

	answer := make([][]int, m)
	for i := range answer {
		answer[i] = make([]int, n)

	}

	for i := range board {
		for j := range board[i] {
			numOfOne := check(board, m, n, i, j)
			if board[i][j] == 1 {
				if numOfOne < 2 {
					answer[i][j] = 0
					continue
				} else if numOfOne == 2 || numOfOne == 3 {
					answer[i][j] = 1
					continue
				} else {
					answer[i][j] = 0
					continue
				}
			} else {
				if numOfOne == 3 {
					answer[i][j] = 1
					continue
				}
			}

			answer[i][j] = board[i][j]
		}
	}

	for i := range answer {
		for j := range answer[i] {
			board[i][j] = answer[i][j]
		}
	}
}

func check(board [][]int, m, n, x, y int) int {
	answer := 0
	if x-1 >= 0 {
		if y-1 >= 0 {
			answer = answer + board[x-1][y-1]
		}

		if y >= 0 {
			answer = answer + board[x-1][y]
		}

		if y+1 < n {
			answer = answer + board[x-1][y+1]
		}
	}

	if y-1 >= 0 {
		answer = answer + board[x][y-1]
	}

	if y+1 < n {
		answer = answer + board[x][y+1]
	}

	if x+1 < m {
		if y-1 >= 0 {
			answer = answer + board[x+1][y-1]
		}

		if y >= 0 {
			answer = answer + board[x+1][y]
		}

		if y+1 < n {
			answer = answer + board[x+1][y+1]
		}
	}

	return answer
}

func main() {
	borad := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	gameOfLife(borad)
	fmt.Println(borad)
}
