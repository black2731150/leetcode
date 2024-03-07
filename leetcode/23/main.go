package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	answer := &ListNode{Val: -10000000}
	for i := 0; i < len(lists); i++ {
		answer = mergeTowList(answer, lists[i])
	}

	return answer.Next
}

func mergeTowList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	d := &ListNode{}
	answer := d

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			d.Next = list1
			list1 = list1.Next
		} else {
			d.Next = list2
			list2 = list2.Next
		}

		d = d.Next
	}

	if list1 != nil {
		d.Next = list1
	}

	if list2 != nil {
		d.Next = list2
	}

	return answer.Next
}

func main() {
	list1 := &ListNode{Val: 1}
	list1Node1 := &ListNode{Val: 2}
	list1Node2 := &ListNode{Val: 4}

	list1.Next = list1Node1
	list1Node1.Next = list1Node2

	list2 := &ListNode{Val: 1}
	list2Node1 := &ListNode{Val: 3}
	list2Node2 := &ListNode{Val: 4}

	list2.Next = list2Node1
	list2Node1.Next = list2Node2

	lists := []*ListNode{list1, list2}
	answer := mergeKLists(lists)
	for answer != nil {
		fmt.Print(answer.Val, "->")
		answer = answer.Next
	}
	fmt.Println("nil")
}
