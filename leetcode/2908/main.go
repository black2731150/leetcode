package main

import "fmt"

func minimumSum(nums []int) int {
	answer := 200

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[j] > nums[i] && nums[j] > nums[k] {
					if x := nums[i] + nums[j] + nums[k]; x < answer {
						answer = x
					}
				}
			}
		}
	}

	if answer == 200 {
		return -1
	}

	return answer
}

func main() {
	nums := []int{6, 5, 4, 3, 4, 5}
	fmt.Println(minimumSum(nums))
}
