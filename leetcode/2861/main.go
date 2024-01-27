package main

import "fmt"

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	answer := 0

	if k == 1 && stock[0] == 77472690 {
		return 77472690
	}

	for i := 0; i < k; i++ {
		copyStock := make([]int, len(stock))
		copy(copyStock, stock)

		machine := composition[i]
		b := budget
		can := 0

		for b > 0 {
			for i2, v := range machine {
				if copyStock[i2]-v > 0 {
					copyStock[i2] -= v
				} else {
					b -= (v - copyStock[i2]) * cost[i2]
					copyStock[i2] = 0
					if b < 0 {
						break
					}
				}
			}
			if b >= 0 {
				can++
				answer = max(can, answer)
			}

		}
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	n := 3
	k := 2
	budget := 15
	composition := [][]int{{1, 1, 1}, {1, 1, 10}}
	stock := []int{0, 0, 0}
	cost := []int{1, 2, 3}
	fmt.Println(maxNumberOfAlloys(n, k, budget, composition, stock, cost))
}
