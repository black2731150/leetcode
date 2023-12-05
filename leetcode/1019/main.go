package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {

	answer := []int{}

	stack := []int{head.Val}
	head = head.Next
	for head != nil {
		if len(stack) > 0 && head.Val > stack[len(stack)-1] {
			for len(stack) > 0 && head.Val > stack[len(stack)-1] {
				answer = append(answer, head.Val)
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, head.Val)
		} else {
			stack = append(stack, head.Val)
		}
		head = head.Next
	}

	if len(stack) != 0 {
		for len(stack) > 0 {
			answer = append(answer, 0)
			stack = stack[:len(stack)-1]
		}
	}

	return answer
}
