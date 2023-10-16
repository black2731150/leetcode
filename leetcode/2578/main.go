package main

import (
	"fmt"
	"sort"
)

func splitNum(num int) int {
	nums := []int{}
	for num != 0 {
		nums = append(nums, num%10)
		num = num / 10
	}

	sort.Ints(nums)

	num1 := 0
	for i := 0; i < len(nums); i = i + 2 {
		num1 = num1*10 + nums[i]
	}

	num2 := 0
	for i := 1; i < len(nums); i = i + 2 {
		num2 = num2*10 + nums[i]
	}

	return num1 + num2
}

func main() {
	num := 4325
	fmt.Println(splitNum(num))
}
