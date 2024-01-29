package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	headCopy := head

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			next := head.Next.Next
			head.Next.Next = nil
			head.Next = next
		} else {
			head = head.Next
		}
	}

	return headCopy
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}
	deleteDuplicates(head)
}
