package main

type MinStack struct {
	stack []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		mins:  []int{},
	}
}

func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)
	l := len(m.mins) - 1
	for ; l >= 0 && m.mins[l] > x; l-- {
	}
	if l == -1 {
		m.mins = append([]int{x}, m.mins...)
	} else {
		m.mins = append(m.mins[:l+1], append([]int{x}, m.mins[l+1:]...)...)
	}

}

func (m *MinStack) Pop() {
	x := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if m.mins[0] == x {
		m.mins = m.mins[1:]
	}

}

func (m *MinStack) Top() int {

	if len(m.stack) > 0 {
		return m.stack[len(m.stack)-1]
	}

	return -1
}

func (m *MinStack) GetMin() int {

	if len(m.mins) > 0 {
		return m.mins[0]
	}

	return -1
}

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Pop()
	obj.Top()
	obj.GetMin()
}
