package main

import (
	"fmt"
	"math"
)

func pickGifts(gifts []int, k int) int64 {
	for ; k > 0; k-- {
		indexMax := max(gifts)
		s := math.Sqrt(float64(gifts[indexMax]))
		gifts[indexMax] = int(s)
	}
	var answer int64
	for i := 0; i < len(gifts); i++ {
		answer = answer + int64(gifts[i])
	}
	return int64(answer)
}

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 0
	}

	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}
	return max
}

func main() {
	gifts := []int{1}
	k := 53
	fmt.Println(pickGifts(gifts, k))
}
