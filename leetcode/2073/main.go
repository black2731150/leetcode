package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	answer := 0
	for tickets[k] > 0 {
		for i, v := range tickets {
			if v == 0 {
				continue
			} else {
				tickets[i] = tickets[i] - 1
				answer++
				if i == k && tickets[i] == 0 {
					break
				}
			}
		}
	}
	return answer
}

func main() {
	tickets := []int{2, 3, 2}
	k := 2
	fmt.Println(timeRequiredToBuy(tickets, k))
}
