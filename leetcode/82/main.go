package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nums := []int{}

	for head != nil {
		if len(nums) == 0 {
			nums = append(nums, head.Val)
			head = head.Next
			continue
		}

		if nums[len(nums)-1] == head.Val {
			val := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			for head != nil && head.Val == val {
				head = head.Next
			}
		} else {
			nums = append(nums, head.Val)
			head = head.Next
		}
	}

	fmt.Println(nums)

	if len(nums) == 0 {
		return nil
	}

	answer := &ListNode{}
	answerHead := answer
	for i := range nums {
		answer.Val = nums[i]
		if i != len(nums)-1 {
			answer.Next = &ListNode{}
			answer = answer.Next
		}
	}

	return answerHead

}

func main() {
	list := &ListNode{Val: 1}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 2}
	node4 := &ListNode{Val: 3}

	list.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	answer := deleteDuplicates(list)
	for answer != nil {
		fmt.Print(answer.Val, "->")
		answer = answer.Next
	}
	fmt.Println("nil")
}
