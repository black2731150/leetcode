package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	answer := []int{}
	var dfs func(r *Node)
	dfs = func(r *Node) {
		if r == nil {
			return
		}

		answer = append(answer, r.Val)
		for _, n := range r.Children {
			dfs(n)
		}
	}

	dfs(root)
	return answer
}

func main() {
	root := Node{Val: 1, Children: []*Node{&Node{Val: 2}, &Node{Val: 2}}}
	preorder(&root)
}
