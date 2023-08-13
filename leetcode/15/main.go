package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	numnsLen := len(nums)

	answer := [][]int{}

	for first := 0; first < numnsLen; first++ {

		if nums[first] > 0 {
			break
		}

		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		thrid := numnsLen - 1

		for second := first + 1; second < thrid; {

			if thrid < numnsLen-1 && nums[thrid] == nums[thrid+1] {
				thrid--
				continue
			}

			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}

			if nums[first]+nums[second]+nums[thrid] == 0 {
				answer = append(answer, []int{nums[first], nums[second], nums[thrid]})
				second++
				continue
			}

			if nums[first]+nums[second]+nums[thrid] > 0 {
				thrid--
			} else {
				second++
			}
		}

	}

	return answer
}

func main() {
	nums := []int{-1, 0, 1}
	fmt.Println(threeSum(nums))
}
