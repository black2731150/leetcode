package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	answer := []int{}
	lastMaxIndex := 0

	for i := 0; i < k; i++ {
		if nums[i] >= nums[lastMaxIndex] {
			lastMaxIndex = i
		}
	}
	answer = append(answer, nums[lastMaxIndex])

	for i := k; i < len(nums); i++ {

		if nums[i] >= nums[lastMaxIndex] {
			lastMaxIndex = i
		}

		if lastMaxIndex > i-k {
			answer = append(answer, nums[lastMaxIndex])
		} else {
			lastMaxIndex = i - k + 1
			for j := i - k + 1; j <= i; j++ {
				if nums[j] >= nums[lastMaxIndex] {
					lastMaxIndex = j
				}
			}
			answer = append(answer, nums[lastMaxIndex])
		}
	}

	return answer
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
