package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	lenDelete := len(to_delete)
	deleteMap := make(map[int]struct{}, lenDelete)

	for _, v := range to_delete {
		deleteMap[v] = struct{}{}
	}

	var answer []*TreeNode

	var work func(node *TreeNode, isRoot bool) *TreeNode
	work = func(node *TreeNode, isRoot bool) *TreeNode {
		if node == nil {
			return nil
		}

		// Check if the current node needs to be deleted
		_, toDelete := deleteMap[node.Val]

		if isRoot && !toDelete {
			answer = append(answer, node)
		}

		node.Left = work(node.Left, toDelete)
		node.Right = work(node.Right, toDelete)

		if toDelete {
			return nil
		} else {
			return node
		}
	}

	work(root, true)

	return answer
}

func main() {
	// Example usage
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	toDelete := []int{3, 5}
	fmt.Println(delNodes(root, toDelete))
}
