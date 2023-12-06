package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	stack := []int{}
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	answer := 0
	x := 1
	for i := len(stack) - 1; i >= 0; i-- {
		answer = answer + stack[i]*x
		x = x * 2
	}

	return answer
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	fmt.Println(getDecimalValue(head))
}
