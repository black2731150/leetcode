package main

type listNode struct {
	val  int
	next *listNode
}

type MyCircularQueue struct {
	front *listNode
	rear  *listNode
	len   int
	cap   int
}

func Constructor(k int) MyCircularQueue {
	node := &listNode{}
	node.next = node

	return MyCircularQueue{
		front: node,
		rear:  node,
		len:   0,
		cap:   k,
	}
}

func (m *MyCircularQueue) EnQueue(value int) bool {
	if m.len >= m.cap {
		return false
	}

	newNode := &listNode{val: value}

	if m.len == 0 {
		m.front.next = newNode
		newNode.next = m.front
	} else {
		newNode.next = m.rear.next
		m.rear.next = newNode
	}

	m.rear = newNode
	m.len++
	return true
}

func (m *MyCircularQueue) DeQueue() bool {
	if m.len == 0 {
		return false
	}

	if m.len == 1 {
		m.rear.next = m.front
	}

	m.front.next = m.front.next.next
	m.len--
	return true
}

func (m *MyCircularQueue) Front() int {
	if m.len == 0 {
		return -1
	}
	return m.front.next.val
}

func (m *MyCircularQueue) Rear() int {
	if m.len == 0 {
		return -1
	}
	return m.rear.val
}

func (m *MyCircularQueue) IsEmpty() bool {
	return m.len == 0
}

func (m *MyCircularQueue) IsFull() bool {
	return m.len == m.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main() {
	Constructor(4)
}
