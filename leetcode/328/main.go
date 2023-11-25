package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	first := head
	seconde := head.Next

	copyFirst := first
	copySecond := seconde
	head = head.Next.Next
	x := 0
	for head != nil {
		next := head.Next
		head.Next = nil
		if x%2 == 0 {
			first.Next = head
			first = first.Next
		} else {
			seconde.Next = head
			seconde = seconde.Next
		}
		head = next
		x++
	}
	seconde.Next = nil

	first.Next = copySecond
	return copyFirst
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	fmt.Println(oddEvenList(head))
}
