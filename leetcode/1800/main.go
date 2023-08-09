package main

import "fmt"

func maxAscendingSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max := 0
	tmpmax := 0
	last := -1
	for i := 0; i < len(nums); i++ {
		if last < nums[i] {
			tmpmax = tmpmax + nums[i]
		} else {
			if max < tmpmax {
				max = tmpmax
			}
			tmpmax = nums[i]
		}

		last = nums[i]
	}

	if max < tmpmax {
		max = tmpmax
	}

	return max
}

func main() {
	nums := []int{12, 17, 15, 13, 10, 11, 12}
	fmt.Println(maxAscendingSum(nums))
}
