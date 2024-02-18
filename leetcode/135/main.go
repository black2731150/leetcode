package main

import "fmt"

func candy(ratings []int) int {
	answer := make([]int, len(ratings))

	for i := range answer {
		answer[i] = 1
	}

	for i := 1; i < len(answer); i++ {
		if ratings[i] > ratings[i-1] {
			answer[i] = answer[i-1] + 1
		}
	}

	for i := len(answer) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			answer[i] = max(answer[i+1]+1, answer[i])
		}
	}

	sum := 0
	for _, v := range answer {
		sum += v
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	ratings := []int{3, 2, 1, 1, 4, 3, 3}
	fmt.Println(candy(ratings))
}
