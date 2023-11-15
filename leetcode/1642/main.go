package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hrrpQurre struct {
	sort.IntSlice
}

func (h *hrrpQurre) Push(val interface{}) {
	h.IntSlice = append(h.IntSlice, val.(int))
}

func (h *hrrpQurre) Pop() interface{} {
	answer := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return answer
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)

	diffs := &hrrpQurre{}
	heap.Init(diffs)
	for i := 1; i < n; i++ {
		if cha := heights[i] - heights[i-1]; cha > 0 {
			heap.Push(diffs, cha)
			if diffs.Len() > ladders {
				min := heap.Pop(diffs).(int)
				if bricks-min >= 0 {
					bricks -= min
				} else {
					return i - 1
				}
			}
		}
	}
	return n - 1
}

func main() {
	height := []int{2, 7, 9, 12}
	bricks := 5
	laders := 1
	fmt.Println(furthestBuilding(height, bricks, laders))
}
