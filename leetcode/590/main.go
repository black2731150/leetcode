package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	answer := []int{}
	var dfs func(r *Node)
	dfs = func(r *Node) {
		if r == nil {
			return
		}

		for _, n := range r.Children {
			dfs(n)
		}
		answer = append(answer, r.Val)
	}

	dfs(root)
	return answer
}

func main() {
	root := Node{Val: 1, Children: []*Node{{Val: 2}, {Val: 2}}}
	postorder(&root)
}
