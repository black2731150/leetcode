package main

import "fmt"

func alternatingSubarray(nums []int) int {
	n := len(nums)
	maxNum := -1

	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] == 1 {
			dangqian := 1
			answer := 2
			maxNum = max(maxNum, answer)
			for j := i + 1; j < n; j++ {
				if nums[j]-nums[j-1] == -dangqian {
					answer++
					maxNum = max(maxNum, answer)
					dangqian = -dangqian
				} else {
					break
				}
			}
		}
	}
	return maxNum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{64, 65, 64, 65, 64, 65, 66, 65, 66, 65}
	fmt.Println(alternatingSubarray(nums))
}
