package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	w   []int
	sum int
	len int
}

func Constructor(w []int) Solution {
	n := len(w)
	prefixSums := make([]int, n)
	sum := 0
	for i, weight := range w {
		sum += weight
		prefixSums[i] = sum
	}
	return Solution{
		w:   prefixSums,
		sum: sum,
		len: n,
	}
}

func (s *Solution) PickIndex() int {
	target := rand.Intn(s.sum)
	return binarySearch(s.w, target)
}

// 辅助函数：二分搜索
func binarySearch(prefixSums []int, target int) int {
	l, r := 0, len(prefixSums)
	for l < r {
		mid := l + (r-l)/2
		if prefixSums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	w := []int{1}
	obj := Constructor(w)
	fmt.Println(obj.PickIndex())
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
