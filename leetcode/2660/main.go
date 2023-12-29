package main

import "fmt"

func isWinner(player1 []int, player2 []int) int {
	n := len(player1)
	if n == 1 {
		if player1[0] > player2[0] {
			return 1
		}

		if player1[0] < player2[0] {
			return 2
		}

		return 0
	}

	p1 := 0
	p1 += player1[0]

	p2 := 0
	p2 += player2[0]

	if n == 2 {
		if player1[0] == 10 {
			p1 += player1[1] * 2
		} else {
			p1 += player1[1]
		}

		if player2[0] == 10 {
			p2 += player2[1] * 2
		} else {
			p2 += player2[1]
		}

		if p1 > p2 {
			return 1
		}

		if p1 < p2 {
			return 2
		}

		return 0
	}

	if player1[0] == 10 {
		p1 += player1[1] * 2
	} else {
		p1 += player1[1]
	}

	if player2[0] == 10 {
		p2 += player2[1] * 2
	} else {
		p2 += player2[1]
	}

	for i := 2; i < n; i++ {
		if player1[i-2] == 10 || player1[i-1] == 10 {
			p1 += player1[i] * 2
		} else {
			p1 += player1[i]
		}

		if player2[i-2] == 10 || player2[i-1] == 10 {
			p2 += player2[i] * 2
		} else {
			p2 += player2[i]
		}
	}

	if p1 > p2 {
		return 1
	}

	if p1 < p2 {
		return 2
	}

	return 0
}

func main() {
	player1 := []int{10, 2, 2, 3}
	player2 := []int{3, 8, 4, 5}
	fmt.Println(isWinner(player1, player2))
}
