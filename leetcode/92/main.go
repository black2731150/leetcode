package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	copyHead := head

	left = left - 1
	l, r := head, head
	for left > 0 || right > 0 {
		if left > 0 {
			left--
			l = l.Next
		}
		if right > 0 {
			right--
			r = r.Next
		}
	}

	if l.Next != nil {
		l.Next.Next = r.Next
	}
	l.Next = r

	return copyHead
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

	answer := reverseBetween(list, 2, 4)

	for answer != nil {
		fmt.Print(answer.Val, "->")
		answer = answer.Next
	}
	fmt.Println("nil")
}
