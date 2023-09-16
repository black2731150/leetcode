package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil {
		return nil
	}

	if head.Next.Next == nil {
		if n == 1 {
			return &ListNode{Val: head.Val, Next: nil}
		}
		if n == 2 {
			return &ListNode{Val: head.Next.Val, Next: nil}
		}
	}

	right := head
	left := head
	beforeLeft := head

	index := 0
	for right != nil && index < n {
		right = right.Next
		index++
	}

	if right != nil {
		left = left.Next
		right = right.Next
	}

	for right != nil {
		left = left.Next
		right = right.Next
		beforeLeft = beforeLeft.Next
	}

	if left.Next == nil {
		beforeLeft.Next = nil
	} else {
		*left = *left.Next
	}

	return head
}

func main() {
	list := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 5}

	list.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	answer := removeNthFromEnd(list, 1)

	for answer != nil {
		fmt.Print(answer.Val, "->")
		answer = answer.Next
	}
	fmt.Println("nil")
}
