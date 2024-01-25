package main

import "fmt"

func fillCups(amount []int) int {

	answer := 0
	for {
		first, second, _ := help(amount)
		if amount[first] == 0 && amount[second] == 0 {
			break
		} else {
			if amount[first] > 0 {
				amount[first]--
			}
			if amount[second] > 0 {
				amount[second]--
			}
			answer++
		}
	}
	return answer
}

func help(nums []int) (first int, second int, third int) {
	if nums[0] > nums[1] {
		if nums[0] > nums[2] {
			first = 0
			if nums[1] > nums[2] {
				second = 1
				third = 2
			} else {
				second = 2
				third = 1
			}
		} else {
			first = 2
			second = 0
			third = 1
		}
	} else {
		if nums[1] > nums[2] {
			first = 1
			if nums[0] > nums[2] {
				second = 0
				third = 2
			} else {
				second = 2
				third = 0
			}
		} else {
			first = 2
			second = 1
			third = 0
		}
	}
	return
}

func main() {
	amount := []int{5, 4, 4}
	fmt.Println(fillCups(amount))
}
