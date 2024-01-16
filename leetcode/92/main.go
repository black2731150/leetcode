package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	copyHead := head

	for i := 0; i < left-1; i++ {
		head = head.Next
	}

	leftPre := head
	head = head.Next

	next := head.Next
	for i := 0; i < right-left-1; i++ {
		if head != nil && head.Next != nil {
			
		}
	}

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
