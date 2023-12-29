package main

import "fmt"

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 != 0 {
		return []int{}
	}

	if x, y := tomatoSlices/2-cheeseSlices, 2*cheeseSlices-tomatoSlices/2; x >= 0 && y >= 0 {
		return []int{x, y}
	} else {
		return []int{}
	}
}

func main() {
	fmt.Println(numOfBurgers(8, 2))
}
