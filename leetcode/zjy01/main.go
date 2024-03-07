package main

type Node struct {
	Val         int
	Left, Right *Node
}

func slove(root *Node, p, q int) *Node {
	if root == nil {
		return nil
	}

	if root.Val == q || root.Val == p {
		return root
	}

	left := slove(root.Left, p, q)
	right := slove(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil && right == nil {
		return left
	}

	if right != nil && left == nil {
		return right
	}

	return nil
}
