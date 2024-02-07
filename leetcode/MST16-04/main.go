package main

import "fmt"

func tictactoe(board []string) string {
	N := len(board)

	XWin, OWin := 0, 0
	HangSum, LieSum := make([]int, N), make([]int, N)
	HasKong := false
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if x := board[i][j]; x == ' ' {
				HasKong = true
			} else {
				if x == 'X' {
					HangSum[i]++
					LieSum[j]++
				}

				if x == 'O' {
					HangSum[i]--
					LieSum[j]--
				}
			}
		}
	}

	zs := board[0][0]
	ztrue := true
	ys := board[0][N-1]
	ytrue := true
	for i := 1; i < N; i++ {
		if zs != board[i][i] {
			ztrue = false
		}

		if ys != board[i][N-i-1] {
			ytrue = false
		}

	}

	if ztrue {
		if zs == 'X' {
			XWin++
		}
		if zs == 'O' {
			OWin++
		}
	}

	if ytrue {
		if ys == 'X' {
			XWin++
		}
		if ys == 'O' {
			OWin++
		}
	}

	for _, v := range HangSum {
		if v == N {
			XWin++
		}

		if v == -N {
			OWin++
		}
	}
	for _, v := range LieSum {
		if v == N {
			XWin++
		}

		if v == -N {
			OWin++
		}
	}

	fmt.Println(HangSum, LieSum)

	if XWin > OWin {
		return "X"
	}

	if XWin < OWin {
		return "O"
	}

	if HasKong {
		return "Pending"
	}

	if XWin == OWin {
		return "Draw"
	}

	return ""
}

func main() {
	borad := []string{"OX ", "OX ", "O  "}
	fmt.Println(tictactoe(borad))
}
