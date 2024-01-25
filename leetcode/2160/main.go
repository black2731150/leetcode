package main

import (
	"fmt"
	"sort"
)

func minimumSum(num int) int {
	nums := []int{}
	for num > 0 {
		nums = append(nums, num%10)
		num = num / 10
	}
	sort.Ints(nums)
	n1 := nums[0]*10 + nums[2]
	n2 := nums[1]*10 + nums[3]
	return n1 + n2
}

func main() {
	num := 1234
	fmt.Println(minimumSum(num))
}
