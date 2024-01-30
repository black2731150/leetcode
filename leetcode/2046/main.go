package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortLinkedList(head *ListNode) *ListNode {
	h := &ListNode{Val: 0, Next: head}

	ch := h.Next
	prev := h

	for ch != nil {
		if ch.Val < 0 {
			prev.Next = ch.Next
			h.Val = ch.Val
			ch.Next = nil
			ch = prev.Next
			h = &ListNode{Val: 0, Next: h}

		} else {
			prev = ch
			ch = ch.Next
		}
	}

	return h.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: -3, Next: &ListNode{Val: -7, Next: &ListNode{Val: -9}}}}
	fmt.Println(sortLinkedList(head))
}
