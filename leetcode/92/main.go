package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	copyHead := head

	nums := []int{}
	n := 0
	for copyHead != nil {
		nums = append(nums, copyHead.Val)
		n++
		copyHead = copyHead.Next
	}

	copyHead = head
	index := 0
	for copyHead != nil && index < left-1 {
		copyHead = copyHead.Next
		index++
	}

	for left <= right {
		copyHead.Val = nums[right-1]
		copyHead = copyHead.Next
		right--
	}

	return head
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
