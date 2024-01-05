package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	hashMap := make(map[*ListNode]struct{}, 0)
	for head != nil {
		hashMap[head] = struct{}{}
		if head.Next != nil {
			if _, find := hashMap[head.Next]; find {
				return head.Next
			}
		}
		head = head.Next
	}
	return nil
}

func main() {
	head := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
	fmt.Println(detectCycle(head))
}
