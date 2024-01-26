package main

import "fmt"

func longestArithSeqLength(nums []int) int {
	n := len(nums)

	answer := 0

	for i := 0; i < n; i++ {

		diffs := make(map[int]struct{})
		for j := i - 1; j >= 0; j-- {
			d := nums[i] - nums[j]
			if _, find := diffs[d]; find {
				continue
			} else {
				diffs[d] = struct{}{}
			}
		}

		for d := range diffs {
			x := 1
			num := nums[i]
			for j := i - 1; j >= 0; j-- {
				if nums[j] == num-d {
					num = nums[j]
					x++
				}
			}

			answer = max(x, answer)
		}
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{9, 4, 7, 2, 10}
	fmt.Println(longestArithSeqLength(nums))
}
