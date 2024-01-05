package main

import "fmt"

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	sets := [4]int{0, 0, 0, 0}
	SetsIsndex := 0
	mod := 4
	xuanzhuan := 1
	allConstommers := customers[0]
	heveCostomers := 0

	answer := -1
	costomerIndex := 1
	MaxNum := 0

	for allConstommers > 0 || costomerIndex < len(customers) {
		if sets[SetsIsndex] > 0 {
			sets[SetsIsndex] = 0
		}

		if allConstommers > 4 {
			sets[SetsIsndex] = 4
			allConstommers = allConstommers - 4
			heveCostomers += 4
		} else {
			sets[SetsIsndex] = allConstommers
			heveCostomers += allConstommers
			allConstommers = 0
		}

		if heveCostomers*boardingCost-xuanzhuan*runningCost > MaxNum {
			answer = xuanzhuan
			MaxNum = heveCostomers*boardingCost - xuanzhuan*runningCost
		}

		if costomerIndex < len(customers) {
			allConstommers = allConstommers + customers[costomerIndex]
			costomerIndex++
		}
		SetsIsndex = (SetsIsndex + 1) % mod
		xuanzhuan++
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
	customers := []int{2, 3}
	boardingCost := 3
	runningCost := 5
	fmt.Println(minOperationsMaxProfit(customers, boardingCost, runningCost))
}
