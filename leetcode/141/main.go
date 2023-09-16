package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, last := head, head

	if head == nil {
		return false
	}

	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		last = last.Next

		if last == fast {
			return true
		}
	}

	return false
}

func main() {
	listnode1 := ListNode{Val: 1}
	listnode2 := ListNode{Val: 2}
	listnode3 := ListNode{Val: 3}

	listnode1.Next = &listnode2
	listnode2.Next = &listnode3
	listnode3.Next = &listnode1

	fmt.Println(hasCycle(&listnode1))
}
