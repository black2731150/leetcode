package main

import (
	"fmt"
	"sort"
)

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	if x := prices[0] + prices[1]; x <= money {
		return money - x
	}
	return money
}

func main() {
	prices := []int{1, 2, 2}
	money := 3
	fmt.Println(buyChoco(prices, money))
}
