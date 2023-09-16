package main

import "math"

type MinStack struct {
	Stack []int
	Min   int
}

func Constructor() MinStack {
	return MinStack{
		Stack: make([]int, 0),
		Min:   math.MaxInt32,
	}
}

func (m *MinStack) Push(val int) {
	m.Stack = append(m.Stack, val)
	if val < m.Min {
		m.Min = val
	}
}

func (m *MinStack) Pop() {
	if len(m.Stack) > 0 {

		last := m.Stack[len(m.Stack)-1]
		m.Stack = m.Stack[:len(m.Stack)-1]

		if last == m.Min {
			if len(m.Stack) > 0 {
				m.Min = min(m.Stack)
			} else {
				m.Min = math.MaxInt32
			}
		}

	}
}

func (m *MinStack) Top() int {
	return m.Stack[len(m.Stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.Min
}

func min(nums []int) int {
	min := nums[0]
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Top()
	obj.GetMin()
	obj.Pop()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
