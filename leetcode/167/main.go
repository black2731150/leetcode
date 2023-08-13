package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}

		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}

	}
	return []int{}
}

func main() {
	numbers := []int{1, 2, 3}
	target := 4
	fmt.Println(twoSum(numbers, target))
}
