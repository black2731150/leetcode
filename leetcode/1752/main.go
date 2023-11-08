package main

import "fmt"

func check(nums []int) bool {
	n := len(nums)
	answer := 0
	for i := 0; i < n-1; i++ {
		if nums[i+1]-nums[i] >= 0 {
			continue
		} else {
			answer++
		}
	}

	if nums[0]-nums[n-1] < 0 {
		answer++
	}

	return answer <= 1
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(check(nums))
}
