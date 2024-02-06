package main

import "sort"

type StreamRank struct {
	nums      []int
	hasSorted bool
}

func Constructor() StreamRank {
	return StreamRank{
		nums:      []int{},
		hasSorted: false,
	}
}

func (s *StreamRank) Track(x int) {
	s.nums = append(s.nums, x)
	s.hasSorted = false
}

func (s *StreamRank) GetRankOfNumber(x int) int {
	if len(s.nums) == 0 {
		return 0
	}

	if !s.hasSorted {
		sort.Ints(s.nums)
		s.hasSorted = true
	}

	left, right := 0, len(s.nums)-1
	for left <= right {
		mid := (left + right) / 2

		if s.nums[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left >= len(s.nums) {
		return len(s.nums)
	}

	if s.nums[left] == x {
		return left + 1
	} else {
		return left
	}
}

/**
 * Your StreamRank object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Track(x);
 * param_2 := obj.GetRankOfNumber(x);
 */
