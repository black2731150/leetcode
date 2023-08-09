package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	answer := &ListNode{}
	answerHead := answer
	count := 0

	for l1 != nil || l2 != nil {

		answer.Next = &ListNode{}

		l1val := 0
		l2val := 0

		if l1 == nil {
			l1val = 0
		} else {
			l1val = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			l2val = 0
		} else {
			l2val = l2.Val
			l2 = l2.Next
		}

		temp := l1val + l2val + count
		count = 0

		if temp >= 10 {
			temp = temp - 10
			count = 1
		}
		answer.Val = temp

		answer = answer.Next

	}

	if count == 1 {
		answer.Val = 1
		count = 0
	} else {
		head := answerHead
		for head.Next.Next != nil {
			head = head.Next
		}
		head.Next = nil
	}

	return answerHead
}

func main() {
	l1 := &ListNode{Val: 2}
	l1Node1 := &ListNode{Val: 4}
	l1Node2 := &ListNode{Val: 3}
	l1.Next = l1Node1
	l1Node1.Next = l1Node2

	l2 := &ListNode{Val: 5}
	l2Node1 := &ListNode{Val: 6}
	l2Node2 := &ListNode{Val: 4}
	l2.Next = l2Node1
	l2Node1.Next = l2Node2

	fmt.Println(addTwoNumbers(l1, l2))
}
