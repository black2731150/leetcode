package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	answer := &ListNode{}
	answerHead := answer

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			answer.Val = list1.Val
			list1 = list1.Next
		} else {
			answer.Val = list2.Val
			list2 = list2.Next
		}

		answer.Next = &ListNode{}
		answer = answer.Next
	}

	if list1 == nil {
		*answer = *list2
	} else {
		*answer = *list1
	}

	return answerHead
}

func main() {
	list1 := &ListNode{Val: 1}
	list1Node1 := &ListNode{Val: 2}
	list1Node2 := &ListNode{Val: 4}

	list1.Next = list1Node1
	list1Node1.Next = list1Node2

	list2 := &ListNode{Val: 1}
	list2Node1 := &ListNode{Val: 3}
	list2Node2 := &ListNode{Val: 4}

	list2.Next = list2Node1
	list2Node1.Next = list2Node2

	answer := mergeTwoLists(list1, list2)
	for answer != nil {
		fmt.Print(answer.Val, "->")
		answer = answer.Next
	}
	fmt.Println("nil")
}
