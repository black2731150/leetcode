package main

import "fmt"

func maxOperations(nums []int) int {
	sum := nums[0] + nums[1]

	answer := 1
	for i := 2; i+1 < len(nums); i += 2 {
		if sum == nums[i]+nums[i+1] {
			answer++
		} else {
			break
		}
	}

	return answer
}

func main() {
	nums := []int{3, 2, 1, 4, 5}
	fmt.Println(maxOperations(nums))
}
