package main

import "fmt"

func game(guess []int, answer []int) int {
	x := 0
	for i := range guess {
		if guess[i] == answer[i] {
			x++
		}
	}

	return x
}

func main() {
	guess := []int{1, 2, 3}
	answer := []int{2, 2, 3}
	fmt.Println(game(guess, answer))
}
