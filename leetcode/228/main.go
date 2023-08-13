package main

import "fmt"

func summaryRanges(nums []int) []string {
	answer := []string{}

	numsLen := len(nums)
	if numsLen == 0 {
		return answer
	}

	if numsLen == 1 {
		return append(answer, fmt.Sprint(nums[0]))
	}

	start := 0
	for i := 1; i < numsLen; i++ {
		if nums[i-1]+1 == nums[i] {
			continue
		} else {
			answer = append(answer, appendAnswer(nums[start], nums[i-1]))

			start = i
		}
	}

	answer = append(answer, appendAnswer(nums[start], nums[len(nums)-1]))

	return answer
}

func appendAnswer(x, y int) string {
	if x == y {
		return fmt.Sprint(x)
	} else {
		return fmt.Sprint(x) + "->" + fmt.Sprint(y)
	}
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
