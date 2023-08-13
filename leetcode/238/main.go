package main

import "fmt"

func productExceptSelf(nums []int) []int {
	answer := []int{}
	left := []int{}
	right := []int{}

	x := 1
	left = append(left, 1)
	for i := 0; i < len(nums)-1; i++ {
		x = x * nums[i]
		left = append(left, x)
	}

	y := 1
	right = append(right, 1)
	for j := len(nums) - 1; j >= 0; j-- {
		y = y * nums[j]
		right = append(right, y)
	}

	for i := range nums {
		answer = append(answer, left[i]*right[len(nums)-1-i])
	}

	return answer
}

// func chengji(nums []int) int {
// 	answer := 1
// 	for _, v := range nums {
// 		answer = answer * v
// 	}
// 	return answer
// }

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
