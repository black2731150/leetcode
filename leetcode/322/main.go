package main

import "sort"

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	sort.Ints(coins)

	bp := make([]int, len(coins))
	if amount%coins[0] == 0 {
		bp[0] = amount / coins[0]
	} else {
		bp[0] = -1
	}

	for i := 1; i < len(coins); i++ {

	}
}
