package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	answer := 0

	nums := make([]int, 100)
	for _, domino := range dominoes {
		if domino[0] > domino[1] {
			domino[0], domino[1] = domino[1], domino[0]
		}

		x := domino[0]*10 + domino[1]
		answer = answer + nums[x]
		nums[x]++
	}

	return answer
}

func main() {
	dominoes := [][]int{{1, 2}, {2, 1}, {4, 5}, {7, 9}}
	fmt.Println(numEquivDominoPairs(dominoes))
}
