package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	stack := []*ListNode{}
	h := head
	for h != nil {
		next := h.Next
		stack = append(stack, h)
		h = next
	}

	n := len(stack)
	stack[n/2-1].Next = nil
	right := stack[n/2:]

	for i := range right {
		right[i].Next = nil
	}

	copyLastNode := &ListNode{}
	h = head
	for len(right) > 0 && h != nil {
		next := h.Next

		lastNode := right[len(right)-1]
		right = right[:len(right)-1]

		h.Next = lastNode
		lastNode.Next = next
		copyLastNode = lastNode

		h = next
	}

	if len(right) > 0 {
		copyLastNode.Next = right[len(right)-1]
	}
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	reorderList(head)
}
