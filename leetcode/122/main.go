package main

import "fmt"

func maxProfit(prices []int) int {
	var left, right int

	var answer int = 0

	for left, right = 0, 1; right < len(prices); {
		if prices[right] > prices[left] {
			answer += prices[right] - prices[left]
		}

		left++
		right++
	}

	return answer
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
