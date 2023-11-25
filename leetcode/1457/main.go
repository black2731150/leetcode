package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	answer := 0

	var work func(root *TreeNode, path int)
	work = func(root *TreeNode, path int) {

		if root == nil {
			return
		}

		path = path ^ (1 << root.Val)

		if root.Left == nil && root.Right == nil {
			if path&(path-1) == 0 {
				answer++
			}
		}

		work(root.Left, path)
		work(root.Right, path)
	}

	work(root, 0)

	return answer
}

//超出内存限制
// func pseudoPalindromicPaths(root *TreeNode) int {
// 	allArrengement := [][]int{}

// 	var work func(root *TreeNode, beforeNums []int, allArrengement *[][]int)
// 	work = func(root *TreeNode, beforeNums []int, allArrengement *[][]int) {
// 		if root.Left == nil && root.Right == nil {
// 			newNums := []int{}
// 			newNums = append(newNums, beforeNums...)
// 			newNums = append(newNums, root.Val)
// 			*allArrengement = append(*allArrengement, newNums)
// 			return
// 		}

// 		newNums := []int{}
// 		newNums = append(newNums, beforeNums...)
// 		newNums = append(newNums, root.Val)

// 		if root.Left != nil {
// 			work(root.Left, newNums, allArrengement)
// 		}

// 		if root.Right != nil {
// 			work(root.Right, newNums, allArrengement)
// 		}

// 	}

// 	work(root, []int{}, &allArrengement)

// 	fmt.Println(allArrengement)
// 	answer := 0
// 	for _, v := range allArrengement {
// 		theMap := make(map[int]int)
// 		for _, v2 := range v {
// 			theMap[v2]++
// 		}

// 		oodNum := 0
// 		for _, v2 := range theMap {
// 			if v2%2 == 0 {
// 				continue
// 			} else {
// 				oodNum++
// 				if oodNum == 2 {
// 					break
// 				}
// 			}
// 		}

// 		if oodNum < 2 {
// 			answer++
// 		}
// 	}

// 	return answer
// }

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}
	fmt.Println(pseudoPalindromicPaths(root))
}
