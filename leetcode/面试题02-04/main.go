package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	h := &ListNode{Val: 0, Next: head}
	prev := h
	for head != nil {
		if head.Val < x {
			v := head.Val
			prev.Next = head.Next
			head = prev.Next
			h.Val = v
			h = &ListNode{Val: 0, Next: h}
		} else {
			head = head.Next
			prev = prev.Next
		}
	}
	return h.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}}
	x := 3
	partition(head, x)
}
