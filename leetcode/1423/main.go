package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	left := make([]int, len(cardPoints)+1)

	for i := 0; i < len(cardPoints); i++ {
		left[i+1] = left[i] + cardPoints[i]
	}

	answer := 0
	for i := 0; i <= k; i++ {
		answer = max(answer, left[i]+(left[len(left)-1]-left[len(left)-1-k+i]))
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
	cardPoints := []int{1, 2, 3, 4, 5}
	k := 3
	fmt.Println(maxScore(cardPoints, k))
}
