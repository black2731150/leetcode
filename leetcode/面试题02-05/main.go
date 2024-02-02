package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	count := 0
	answer := &ListNode{}
	a := answer
	for l1 != nil && l2 != nil {
		v := l1.Val + l2.Val + count
		count = 0
		if v >= 10 {
			v = v % 10
			count = 1
		}
		answer.Next = &ListNode{Val: v}
		answer = answer.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		v := l1.Val + count
		count = 0
		if v >= 10 {
			v = v % 10
			count = 1
		}
		answer.Next = &ListNode{Val: v}
		answer = answer.Next
		l1 = l1.Next
	}

	for l2 != nil {
		v := l2.Val + count
		count = 0
		if v >= 10 {
			v = v % 10
			count = 1
		}
		answer.Next = &ListNode{Val: v}
		answer = answer.Next
		l2 = l2.Next
	}

	if count > 0 {
		answer.Next = &ListNode{Val: count}
		answer = answer.Next
	}

	return a.Next
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}}
	addTwoNumbers(l1, l2)
}
