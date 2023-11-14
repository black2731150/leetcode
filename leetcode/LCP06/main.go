package main

import "fmt"

func minCount(coins []int) int {
	answer := 0
	for i := 0; i < len(coins); i++ {
		if coins[i]%2 == 0 {
			answer = answer + coins[i]/2
		} else {
			answer = answer + coins[i]/2 + 1
		}
	}
	return answer
}

func main() {
	coint := []int{4, 2, 1}
	fmt.Println(minCount(coint))
}
