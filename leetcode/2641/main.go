package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Val = 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)

		sum := 0
		for i := 0; i < n; i++ {
			node := queue[i]

			if node.Left != nil {
				queue = append(queue, node.Left)
				sum += node.Left.Val
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
				sum += node.Right.Val
			}
		}

		for i := 0; i < n; i++ {
			node := queue[i]

			if node.Left != nil && node.Right != nil {
				x := node.Left.Val + node.Right.Val
				node.Left.Val = sum - x
				node.Right.Val = sum - x
			} else {
				if node.Left != nil {
					node.Left.Val = sum - node.Left.Val
				}

				if node.Right != nil {
					node.Right.Val = sum - node.Right.Val
				}
			}
		}

		queue = queue[n:]
	}

	return root
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	replaceValueInTree(root)
}
