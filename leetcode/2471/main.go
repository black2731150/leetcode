package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {
	answer := 0

	qurue := []*TreeNode{root}

	for len(qurue) != 0 {
		copyQueue := make([]*TreeNode, len(qurue))
		copy(copyQueue, qurue)
		qurue = []*TreeNode{}

		nums := []int{}

		for _, tn := range copyQueue {
			nums = append(nums, tn.Val)
			if tn.Left != nil {
				qurue = append(qurue, tn.Left)
			}
			if tn.Right != nil {
				qurue = append(qurue, tn.Right)
			}
		}

		ids := make([]int, len(nums))
		for i := range ids {
			ids[i] = i
		}

		sort.Slice(ids, func(i, j int) bool { return nums[ids[i]] < nums[ids[j]] })

		answer += len(nums)
		vis := make([]bool, len(nums))

		for _, v := range ids {
			if !vis[v] {
				for ; !vis[v]; v = ids[v] {
					vis[v] = true
				}
				answer--
			}
		}
	}

	return answer
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 7}}}
	fmt.Println(minimumOperations(root))
}
