package main

import "fmt"

func findMaximumXOR(nums []int) int {
	answer := 0

	for i := 30; i >= 0; i-- {
		theMap := make(map[int]bool, 2)

		for _, num := range nums {
			theMap[num>>i] = true
		}

		answerNext := answer*2 + 1
		found := false

		for _, num := range nums {
			if theMap[num>>i^answerNext] {
				found = true
				break
			}
		}

		if found {
			answer = answerNext
		} else {
			answer = answerNext - 1
		}

	}
	return answer
}

func main() {
	nums := []int{3, 4, 5, 6, 7, 8}
	fmt.Println(findMaximumXOR(nums))
}
