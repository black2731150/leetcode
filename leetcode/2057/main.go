package main

import "fmt"

func smallestEqual(nums []int) int {
	for i, num := range nums {
		if num >= 10 {
			continue
		}
		if i%10 == num {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Print(smallestEqual(nums))
}
