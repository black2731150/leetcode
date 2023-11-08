package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		x := target - numbers[i]

		for j := n - 1; j > i; j-- {
			if numbers[j] == x {
				return []int{i, j}
			}
		}

	}
	return []int{}
}

func main() {
	numbers := []int{1, 2, 3}
	target := 5
	fmt.Println(twoSum(numbers, target))
}
