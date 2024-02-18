package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	answer := &ListNode{Val: 1, Next: head}

	for head != nil {
		if head == nil || head.Next == nil {
			break
		}

		n := &ListNode{}
		n = nil
		for i := 0; i < k && head != nil; i++ {
			next := head.Next
			head.Next = next
			n = head
			head = next
		}

		n = n
	}

	return answer.Next
}
