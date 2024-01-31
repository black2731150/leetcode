package main

import "fmt"

func change(amount int, coins []int) int {
	N := len(coins)

	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < N; i++ {
		for v := coins[i]; v <= amount; v++ {
			dp[v] += dp[v-coins[i]]
		}
	}

	return dp[amount]
}

func main() {
	amount := 2
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}
