package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	answer := 1
	pDiff := nums[1] - nums[0]
	if pDiff != 0 {
		answer++
	}

	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if (diff > 0 && pDiff <= 0) || (diff < 0 && pDiff >= 0) {
			answer++
			pDiff = diff
		}
	}

	return answer
}

func main() {
	nums := []int{1, 7, 4, 9, 2, 5}
	fmt.Println(wiggleMaxLength(nums))
}
