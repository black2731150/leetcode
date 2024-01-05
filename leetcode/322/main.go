package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Ints(coins)

	bp := make([]int, amount+1)
	for i := range bp {
		bp[i] = amount + 1 // 初始化为大于amount的值
	}
	bp[0] = 0 // 兑换0元需要0个硬币

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				bp[i] = min(bp[i], bp[i-coin]+1)
			}
		}
	}

	if bp[amount] > amount { // 如果bp[amount]仍然是初始化的值，则说明无解
		return -1
	}
	return bp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	// Test the function with a simple example
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println("Minimum coins needed:", coinChange(coins, amount))
}
