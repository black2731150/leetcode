package main

import "fmt"

func numIdenticalPairs(nums []int) int {

	answrr := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				answrr++
			}
		}
	}
	return answrr
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums))
}
