package main

import (
	"container/heap"
	"fmt"
)

type PointHeap []int

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[0]
	*h = old[1:n]
	return x
}

func leastInterval(tasks []byte, n int) int {

	nums := make([]int, 26)
	for _, v := range tasks {
		nums[v-'A']++
	}

	taskNums := PointHeap{}
	for _, v := range nums {
		if v > 0 {
			taskNums.Push(v)
		}
	}

	heap.Init(&taskNums)

	fmt.Println(taskNums)

	answer := 0
	for taskNums.Len() > 0 && taskNums[0] > 0 {
		tmp := []int{}

		for i := 0; i <= n; i++ {
			if taskNums.Len() > 0 {
				max := taskNums.Pop().(int)
				if max > 1 {
					tmp = append(tmp, max-1)
				}
				answer++

			} else if len(tmp) > 0 {
				answer++
			} else {
				break
			}
		}
		fmt.Println("taskNums: ", taskNums)
		fmt.Println("tmp: ", tmp)

		for _, v := range tmp {
			if v > 0 {
				taskNums.Push(v)
			}
		}

		heap.Init(&taskNums)
	}

	return answer
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}
