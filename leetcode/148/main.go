package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func sortList(head *ListNode) *ListNode {
// 	answer := &ListNode{Val: 0, Next: head}
// 	answerHead := answer

// 	nums := []int{}
// 	for answer.Next != nil {
// 		nums = append(nums, answer.Next.Val)
// 		answer = answer.Next
// 	}

// 	sort.Ints(nums)

// 	answer = answerHead

// 	for i := range nums {
// 		answer.Next.Val = nums[i]
// 		answer = answer.Next
// 	}

// 	return answerHead.Next
// }

func sortList(head *ListNode) *ListNode {
	answer := &ListNode{Val: 0, Next: head}
	answerHead := answer

	nums := []int{}
	for answer.Next != nil {
		nums = append(nums, answer.Next.Val)
		answer = answer.Next
	}

	sort.Ints(nums)

	answer = answerHead

	for i := range nums {
		answer.Next.Val = nums[i]
		answer = answer.Next
	}

	return answerHead.Next
}

func main() {
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 4}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	fmt.Println(sortList(head))
}
