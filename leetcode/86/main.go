package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{0, &ListNode{}}
	smallHead := small
	big := &ListNode{0, &ListNode{}}
	bigHead := big

	for ; head != nil; head = head.Next {
		if head.Val < x {
			small = small.Next
			small.Val = head.Val
			small.Next = &ListNode{}

		} else {
			big = big.Next
			big.Val = head.Val
			big.Next = &ListNode{}
		}
	}

	small.Next = nil
	big.Next = nil

	small.Next = bigHead.Next

	return smallHead.Next

}

func main() {
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node7 := &ListNode{Val: 7}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	answer := partition(head, 3)
	for answer != nil {
		fmt.Print(answer.Val, "->")
		answer = answer.Next
	}
	fmt.Println("nil")
}
