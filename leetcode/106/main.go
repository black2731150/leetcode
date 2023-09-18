package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]
	postorder = postorder[:len(postorder)-1]

	node := &TreeNode{Val: val}

	inorderValIndex := -1
	for i := range inorder {
		if inorder[i] == val {
			inorderValIndex = i
		}
	}

	if inorderValIndex == -1 {
		return nil
	}

	leftInorder := inorder[:inorderValIndex]
	var rightInorder []int
	if len(inorder) == 0 {
		rightInorder = []int{}
	} else {
		rightInorder = inorder[inorderValIndex+1:]
	}

	n := len(leftInorder)

	rightPostOrder := postorder[n:]
	leftPostOrder := postorder[:n]

	node.Left = buildTree(leftInorder, leftPostOrder)
	node.Right = buildTree(rightInorder, rightPostOrder)

	return node
}

func main() {
	head := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}

	head.Left = node1
	head.Right = node2

	node2.Left = node3
	node2.Right = node4

	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	answer := buildTree(preorder, inorder)

	fmt.Println(answer)
}
