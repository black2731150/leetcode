package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	n := len(nums)
	for even < n && odd < n {
		for even < n && nums[even]%2 == 0 {
			even += 2
		}

		for odd < n && nums[odd]%2 != 0 {
			odd += 2
		}

		if even < n && odd < n {
			nums[odd], nums[even] = nums[even], nums[odd]
		}

	}
	return nums
}

func main() {
	nums := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(nums))
}
