package main

import "fmt"

func returnToBoundaryCount(nums []int) int {
	x := 0
	answer := 0
	for _, v := range nums {
		x += v
		if x == 0 {
			answer++
		}
	}

	return answer
}

func main() {
	nums := []int{2, 3, -5}
	fmt.Println(returnToBoundaryCount(nums))
}
