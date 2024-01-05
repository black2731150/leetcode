package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	//如果头节点是空的或者只有一个节点，则不需要排序
	if head == nil || head.Next == nil {
		return head
	}

	//通过快慢指针找到链表的中点
	frontOfSlow, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		frontOfSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	//中链表中间断开
	frontOfSlow.Next = nil

	//通过递归，归并的排序剩余的链表
	left := sortList(head)
	right := sortList(slow)

	return mergeTowLists(left, right)
}

// 用合并两个有序链表的方法去合并链表
func mergeTowLists(l1, l2 *ListNode) *ListNode {
	answer := &ListNode{}
	answerHead := answer

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			answerHead.Next = l1
			l1 = l1.Next
		} else {
			answerHead.Next = l2
			l2 = l2.Next
		}
		answerHead = answerHead.Next
	}

	if l1 != nil {
		answerHead.Next = l1
	} else if l2 != nil {
		answerHead.Next = l2
	}

	return answer.Next
}

//放进数组并进行快排，不能体现算法的本质，其本质还是归并排序，相当于再链表身上直接实现归并排序，可以实现空间复杂度为O(1)
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
