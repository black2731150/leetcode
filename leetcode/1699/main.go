package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var ap, bp *ListNode
	l1 := list1
	for l1.Next != nil {
		a--
		b--
		if a == 0 {
			ap = l1
		}

		if b == 0 {
			bp = l1
			break
		}
		l1 = l1.Next
	}
	bp = bp.Next
	l2 := list2
	for l2.Next != nil {
		l2 = l2.Next
	}

	ap.Next = list2
	l2.Next = bp.Next
	return list1
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	ltst2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	a, b := 1, 1
	fmt.Println(mergeInBetween(list1, a, b, ltst2))
}
