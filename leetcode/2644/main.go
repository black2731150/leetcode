package main

import "fmt"

func maxDivScore(nums []int, divisors []int) int {
	answer := make([]int, len(divisors))
	maxNum := 0
	for i, d := range divisors {
		for _, num := range nums {
			if num%d == 0 {
				answer[i]++
			}
		}

		maxNum = max(maxNum, answer[i])
	}

	ra := 0
	for i, ans := range answer {
		if ans == maxNum {
			if ra == 0 {
				ra = divisors[i]
			} else {
				if divisors[i] < ra {
					ra = divisors[i]
				}
			}
		}
	}

	return ra
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{4, 7, 9, 3, 9}
	divisors := []int{5, 2, 3}
	fmt.Println(maxDivScore(nums, divisors))
}
