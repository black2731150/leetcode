package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var q *ListNode

	for head != nil {
		next := head.Next
		head.Next = q
		q = head
		head = next
	}

	return q
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(reverseList(head))
}
