package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	bp := make([]int, len(cost)+1)

	for i := 2; i <= len(cost); i++ {
		bp[i] = min(bp[i-1]+cost[i-1], bp[i-2]+cost[i-2])
	}
	return bp[len(bp)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}
