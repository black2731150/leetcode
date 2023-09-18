package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := head

	n := 1
	for last.Next != nil {
		last = last.Next
		n++
	}

	fmt.Println(n)

	last.Next = head

	k = k % n
	fmt.Println(k)
	for i := 0; i < n-k-1; i++ {
		head = head.Next
		last = last.Next
	}

	last = last.Next
	last = last.Next

	head.Next = nil
	return last
}

func main() {
	head := &ListNode{Val: 0}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	// node3 := &ListNode{Val: 4}
	// node4 := &ListNode{Val: 5}

	head.Next = node1
	node1.Next = node2
	// node2.Next = node3
	// node3.Next = node4

	answer := rotateRight(head, 4)

	for answer != nil {
		fmt.Print(answer.Val, "->")
		answer = answer.Next
	}
	fmt.Println("nil")
}
