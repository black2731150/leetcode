package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	type point struct {
		x   int
		y   int
		val int
	}

	theMap := make(map[int][]point)

	var dfs func(r *TreeNode, x, y int)
	dfs = func(r *TreeNode, x, y int) {
		if r == nil {
			return
		}

		theMap[y] = append(theMap[y], point{x: x, y: y, val: r.Val})
		if r.Left != nil {
			dfs(r.Left, x+1, y-1)
		}

		if r.Right != nil {
			dfs(r.Right, x+1, y+1)
		}
	}

	dfs(root, 0, 0)

	indexs := []int{}
	for k := range theMap {
		indexs = append(indexs, k)
	}

	sort.Ints(indexs)

	answer := [][]int{}
	for _, v := range indexs {
		nums := theMap[v]
		sort.Slice(nums, func(i, j int) bool {
			if nums[i].x == nums[j].x {
				return nums[i].val < nums[j].val
			}
			return nums[i].x < nums[j].x
		})

		answer = append(answer, []int{})
		for _, p := range nums {
			answer[len(answer)-1] = append(answer[len(answer)-1], p.val)
		}
	}

	return answer
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	verticalTraversal(root)
}
