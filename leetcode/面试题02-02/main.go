package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	h := head
	for i := 0; i < k; i++ {
		head = head.Next
	}

	for head != nil {
		h = h.Next
		head = head.Next
	}
	return h.Val
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}}
	k := 3
	kthToLast(head, k)
}
