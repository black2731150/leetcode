package main

import "fmt"

func sumOfSquares(nums []int) int {
	answer := 0
	n := len(nums)

	for i := 1; i <= len(nums); i++ {
		num := nums[i-1]
		if n%i == 0 {
			answer = answer + num*num
		}
	}

	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(sumOfSquares(nums))
}
