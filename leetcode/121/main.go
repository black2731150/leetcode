package main

import "fmt"

func maxProfit(prices []int) int {
	lenPrices := len(prices)
	if lenPrices == 1 {
		return 0
	}

	min := prices[0]
	max := 0

	for i := 1; i < lenPrices; i++ {
		min = Min(min, prices[i])
		max = Max(prices[i]-min, max)
	}

	return max

}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
