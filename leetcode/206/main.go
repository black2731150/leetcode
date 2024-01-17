package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := &ListNode{}
	n = nil
	for head != nil {
		next := head.Next
		head.Next = n
		n = head
		head = next
	}

	return n
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	reverseList(head)
}
