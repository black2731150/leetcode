package main

import "fmt"

func lengthOfLIS(nums []int) int {
	bp := make([]int, len(nums))

	bp[0] = 1
	maxAnswer := 1
	for i := 1; i < len(nums); i++ {
		bp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				bp[i] = max(bp[j]+1, bp[i])
			}
		}
		maxAnswer = max(maxAnswer, bp[i])
	}

	return maxAnswer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}
