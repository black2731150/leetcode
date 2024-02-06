package main

type MyQueue struct {
	inStack  []int
	outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (m *MyQueue) Push(x int) {
	m.inStack = append(m.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (m *MyQueue) Pop() int {
	if len(m.outStack) == 0 {
		if len(m.inStack) == 0 {
			return -1
		} else {
			for i := len(m.inStack) - 1; i >= 0; i-- {
				m.outStack = append(m.outStack, m.inStack[i])
				m.inStack = m.inStack[:i]
			}
		}
	}

	answer := m.outStack[len(m.outStack)-1]
	m.outStack = m.outStack[:len(m.outStack)-1]
	return answer
}

/** Get the front element. */
func (m *MyQueue) Peek() int {
	if len(m.outStack) == 0 {
		if len(m.inStack) == 0 {
			return -1
		} else {
			return m.inStack[0]
		}
	} else {
		return m.outStack[len(m.outStack)-1]
	}
}

/** Returns whether the queue is empty. */
func (m *MyQueue) Empty() bool {
	return len(m.inStack) == 0 && len(m.outStack) == 0
}

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(3)
	obj.Pop()
	obj.Peek()
	obj.Empty()
}
