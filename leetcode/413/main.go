package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	answer := 0

	if len(nums) < 3 {
		return answer
	}

	first, seconde, thred := 0, 1, 2
	max := 2

	for thred < len(nums) {
		if nums[seconde]-nums[first] == nums[thred]-nums[seconde] {
			max++
			first++
			seconde++
			thred++
		} else {
			x := max - 2
			answer += x * (x + 1) / 2
			max = 2
			first++
			seconde++
			thred++
		}
	}

	x := max - 2
	answer += x * (x + 1) / 2

	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(numberOfArithmeticSlices(nums))
}
