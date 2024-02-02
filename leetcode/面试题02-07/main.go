package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	theMap := make(map[*ListNode]struct{})

	for headA != nil {
		theMap[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, find := theMap[headB]; find {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
//     if headA == nil || headB == nil {
//         return nil
//     }
//     pa, pb := headA, headB
//     for pa != pb {
//         if pa == nil {
//             pa = headB
//         } else {
//             pa = pa.Next
//         }
//         if pb == nil {
//             pb = headA
//         } else {
//             pb = pb.Next
//         }
//     }
//     return pa
// }

func main() {
	c := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: c}}}}}}
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: c}}}}}}

	getIntersectionNode(head1, head2)
}
