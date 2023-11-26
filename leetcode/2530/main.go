package main

import (
	"fmt"
	"math"
	"sort"
)

func maxKelements(nums []int, k int) int64 {
	var answer int64 = 0

	n := len(nums)

	if n == 1 {
		answer = answer + int64(nums[0])
		for i := 1; i < k; i++ {
			answer = answer + int64(math.Ceil(float64(nums[0])/3))
		}
		return answer
	}

	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	first, second := nums[0], nums[1]
	for i := 0; i < k; i++ {
		answer = answer + int64(first)
		first = int(math.Ceil(float64(first) / 3))
		nums[0] = first

		if first > second {
			continue
		} else {
			sort.Ints(nums)
			for i2, j := 0, len(nums)-1; i2 < j; i2, j = i2+1, j-1 {
				nums[i2], nums[j] = nums[j], nums[i2]
			}

			first = nums[0]
			second = nums[1]
		}

	}

	return answer
}

func main() {
	nums := []int{1, 10, 3, 3, 3}
	k := 3
	fmt.Println(maxKelements(nums, k))
}
