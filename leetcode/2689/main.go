package main

import (
	"fmt"
	"strings"
)

type RopeTreeNode struct {
	len   int
	val   string
	left  *RopeTreeNode
	right *RopeTreeNode
}

func getKthCharacter(root *RopeTreeNode, k int) byte {
	var builder strings.Builder
	var order func(r *RopeTreeNode)
	order = func(r *RopeTreeNode) {
		if r == nil {
			return
		}

		order(r.left)
		order(r.right)
		builder.WriteString(r.val)
	}

	order(root)

	return builder.String()[k-1]
}

func main() {
	root := &RopeTreeNode{len: 0, val: "ropetree"}
	k := 8
	fmt.Println(getKthCharacter(root, k))
}
