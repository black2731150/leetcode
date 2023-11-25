package main

import "fmt"

func predictTheWinner(nums []int) bool {
	return work(nums, 1, 0, len(nums)-1) >= 0
}

func work(nums []int, player int, start, end int) int {
	if start == end {
		return nums[start] * player
	}
	startNum := work(nums, -player, start+1, end) + nums[start]*player
	endNum := work(nums, -player, start, end-1) + nums[end]*player
	return max(startNum*player, endNum*player) * player
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(predictTheWinner(nums))
}
