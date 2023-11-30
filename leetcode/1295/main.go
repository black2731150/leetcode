package main

import "fmt"

func findNumbers(nums []int) int {
	answer := 0
	for _, v := range nums {
		x := 0
		for v > 0 {
			x++
			v = v / 10
		}
		if x%2 == 0 {
			answer++
		}
	}
	return answer
}

func main() {
	nums := []int{3, 56, 78, 999, 32424}
	fmt.Println(findNumbers(nums))
}
