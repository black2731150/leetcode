package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	res := false
	var valid func(*TreeNode, *TreeNode) bool
	valid = func(root1 *TreeNode, root2 *TreeNode) bool {
		// 判断两个树是否相同
		if root1 == nil && root2 == nil {
			return true
		}
		if root1 == nil || root2 == nil {
			return false
		}
		if root1.Val != root2.Val {
			return false
		}
		return valid(root1.Left, root2.Left) && valid(root1.Right, root2.Right)
	}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			// 找到t1和t2根节点值相同的结点
			if root.Val == t2.Val {
				T1 := root
				T2 := t2
				// 传入到valid中 查看这个子树和t2是否相同
				res = valid(T1, T2)
			}
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(t1)
	return res
}
