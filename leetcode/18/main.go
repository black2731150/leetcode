package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	type one struct {
		a, b, c, d int
	}

	theMap := make(map[one]struct{})
	answer := [][]int{}

	sort.Ints(nums)
	n := len(nums)

	for first := 0; first < n-3; first++ {
		for second := first + 1; second < n-2; second++ {
			left, right := second+1, n-1
			for left < right {
				t := target - nums[first] - nums[second]

				if nums[left]+nums[right] == t {
					if _, find := theMap[one{nums[first], nums[second], nums[left], nums[right]}]; !find {
						theMap[one{nums[first], nums[second], nums[left], nums[right]}] = struct{}{}
						answer = append(answer, []int{nums[first], nums[second], nums[left], nums[right]})
					}
				}

				if nums[left]+nums[right] > t {
					right--
				} else {
					left++
				}
			}

		}
	}
	return answer
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}
