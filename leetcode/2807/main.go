package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil && head.Next == nil {
		return nil
	}

	headp := head
	for head.Next != nil {
		next := head.Next
		x := head.Val
		y := next.Val
		gcd := get(x, y)
		head.Next = &ListNode{Val: gcd}
		head.Next.Next = next
		head = head.Next.Next
	}

	return headp
}

func get(x, y int) int {
	if x < y {
		return get(y, x)
	}

	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func main() {
	head := &ListNode{Val: 17, Next: &ListNode{Val: 6, Next: &ListNode{Val: 10, Next: &ListNode{Val: 3}}}}
	fmt.Println(insertGreatestCommonDivisors(head))
}
