package main

import (
	"container/heap"
	"fmt"
)

type elem struct {
	index int
	cha   int
}

type minHeap []elem

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].cha > m[j].cha }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(*m)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
func (m *minHeap) Push(v interface{}) {
	*m = append(*m, v.(elem))
}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {

	n := len(reward1)

	h := minHeap{}
	heap.Init(&h)

	for i := 0; i < n; i++ {
		heap.Push(&h, elem{
			index: i,
			cha:   reward1[i] - reward2[i],
		})
	}

	onesum := 0
	for i := 0; i < k; i++ {
		x := heap.Pop(&h).(elem)
		onesum += reward1[x.index]
	}

	towsum := 0
	for h.Len() > 0 {
		x := heap.Pop(&h).(elem)
		towsum += reward2[x.index]
	}

	return onesum + towsum
}

func main() {
	reward1 := []int{1, 1, 3, 4}
	reward2 := []int{4, 4, 1, 1}
	k := 2
	fmt.Println(miceAndCheese(reward1, reward2, k))
}
