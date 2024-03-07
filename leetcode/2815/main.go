package main

import "fmt"

func maxSum(nums []int) int {
	answer := -1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if maxWei(nums[i]) == maxWei(nums[j]) {
				answer = max(answer, nums[i]+nums[j])
			}
		}
	}

	return answer
}

func maxWei(num int) int {
	answer := num % 10
	for num != 0 {
		answer = max(answer, num%10)
		num = num / 10
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
	nums := []int{51, 71, 17, 24, 42}
	fmt.Println(maxSum(nums))
}
