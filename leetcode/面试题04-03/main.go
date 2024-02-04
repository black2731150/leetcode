package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {

	queue := []*TreeNode{tree}

	answer := []*ListNode{}

	for len(queue) > 0 {
		cq := make([]*TreeNode, len(queue))
		copy(cq, queue)
		queue = queue[:0]

		onrList := &ListNode{}
		o := onrList
		for _, tn := range cq {
			onrList.Next = &ListNode{Val: tn.Val}
			onrList = onrList.Next

			if tn.Left != nil {
				queue = append(queue, tn.Left)
			}

			if tn.Right != nil {
				queue = append(queue, tn.Right)
			}
		}

		answer = append(answer, o.Next)
	}

	return answer

}

func main() {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	listOfDepth(tree)
}
