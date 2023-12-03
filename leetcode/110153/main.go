package main

import "fmt"

func minimumAddedCoins(coins []int, target int) int {
	need := 0
	for target > 0 {
		need++
		target = target / 2
	}

	var x uint32

	x = 0
	for _, v := range coins {
		x = x | uint32(v)
	}

	answer := 0

	for i := 0; i <= need; i++ {
		if x&(1<<i) != 0 {
			continue
		} else {
			answer++
		}
	}
	return answer
}

func main() {
	coins := []int{1, 1, 1}
	target := 19
	fmt.Println(minimumAddedCoins(coins, target))
}
