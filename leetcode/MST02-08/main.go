package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	theMap := make(map[*ListNode]struct{})

	for head != nil {
		if _, find := theMap[head]; find {
			return head
		} else {
			theMap[head] = struct{}{}
			head = head.Next
		}
	}
	return nil
}

// func detectCycle(head *ListNode) *ListNode {
//     slow, fast := head, head
//     for fast != nil {
//         slow = slow.Next
//         if fast.Next == nil {
//             return nil
//         }
//         fast = fast.Next.Next
//         if fast == slow {
//             p := head
//             for p != slow {
//                 p = p.Next
//                 slow = slow.Next
//             }
//             return p
//         }
//     }
//     return nil
// }
