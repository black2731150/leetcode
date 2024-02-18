package main

import (
	"fmt"
	"sort"
)

func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)

	if n%k != 0 {
		return false
	}

	sort.Ints(nums)

	theMap := make(map[int]int)
	for _, v := range nums {
		theMap[v]++
	}

	nk := 0

	for i := 0; i < n; i++ {
		start := nums[i]
		if n1, find1 := theMap[start]; find1 && n1 > 0 {
			for j := 0; j < k; j++ {
				if n2, find2 := theMap[start+j]; find2 && n2 > 0 {
					theMap[start+j]--
				} else {
					return false
				}
			}
			nk++
		} else {
			continue
		}
	}

	return nk*k == n
}

func main() {
	nums := []int{1, 2, 3, 3, 4, 4, 5, 6}
	k := 4
	fmt.Println(isPossibleDivide(nums, k))
}
