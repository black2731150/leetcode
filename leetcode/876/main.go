package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	low, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	if fast != nil && fast.Next != nil {
		low = low.Next
	}

	return low
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	fmt.Println(middleNode(head))
}
