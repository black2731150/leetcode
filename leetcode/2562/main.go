package main

import "fmt"

func findTheArrayConcVal(nums []int) int64 {
	var answer int64 = 0

	for {
		n := len(nums)

		if n == 0 {
			break
		}

		if n == 1 {
			answer = answer + int64(nums[0])
			break
		} else {
			first := nums[0]
			last := nums[n-1]

			for tmp := last; tmp != 0; {
				first = first * 10
				tmp = tmp / 10
			}
			first = first + last
			answer = answer + int64(first)

			nums = nums[1 : n-1]
		}
	}
	return int64(answer)
}

func main() {
	nums := []int{7, 52, 2, 4}
	fmt.Println(findTheArrayConcVal(nums))
}
