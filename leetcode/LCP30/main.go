package main

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	answer := old[n-1]
	*h = old[:n-1]
	return answer
}

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func magicTower(nums []int) int {
	h := minHeap{}

	n := len(nums)

	xue := 1
	last := 0
	answer := 0
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			heap.Push(&h, nums[i])
			if xue+nums[i] <= 0 {
				x := heap.Pop(&h).(int)
				xue = xue - x
				last += x
				answer++
			}
			xue += nums[i]
		} else {
			xue += nums[i]
		}
	}

	if xue+last < 0 {
		return -1
	}

	return answer
}

func main() {
	nums := []int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}
	fmt.Println(magicTower(nums))
}
