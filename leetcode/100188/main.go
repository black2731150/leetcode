package main

import "fmt"

func countOfPairs(n int, x int, y int) []int {
	if x > y {
		return countOfPairs(n, y, x)
	}

	anwer := make([]int, n)

	for i, _ := range anwer {
		anwer[i] = 2 * (n - 1 - i)
	}

	if x == y {
		return anwer
	}

	for i := 0; i < x; i++ {
		anwer[i] += 2
	}

	for i := x; i < n; i++ {
		if anwer[i] > 0 {
			anwer[i] -= 2
		}
	}

	return anwer
}

func main() {
	n := 3
	x := 1
	y := 3
	fmt.Println(countOfPairs(n, x, y))
}
