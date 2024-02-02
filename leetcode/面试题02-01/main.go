package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	hashMap := make(map[int]struct{})
	prve := &ListNode{Val: 0, Next: head}
	h := prve
	for head != nil {
		if _, find := hashMap[head.Val]; find {
			prve.Next = head.Next
			head = prve.Next
		} else {
			hashMap[head.Val] = struct{}{}
			head = head.Next
			prve = prve.Next
		}
	}
	return h.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}}
	removeDuplicateNodes(head)
}
