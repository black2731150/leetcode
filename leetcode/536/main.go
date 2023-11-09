package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func str2tree(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}

	root := &TreeNode{}

	num, s := getNum(s)
	root.Val = num

	left, s := getTreeString(s)
	right, _ := getTreeString(s)

	root.Left = str2tree(left)
	root.Right = str2tree(right)

	return root
}

func getTreeString(s string) (string, string) {
	if len(s) == 0 {
		return "", ""
	}

	if s[0] != '(' {
		return "", ""
	}

	stark := []byte{}

	var answer string
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stark = append(stark, s[i])
		}

		if s[i] == ')' {
			if stark[len(stark)-1] == '(' {
				stark = stark[:len(stark)-1]
				if len(stark) == 0 {
					answer = s[1:i]
					s = s[i+1:]
					break
				}
			}
		}
	}

	return answer, s

}

func getNum(s string) (int, string) {
	if len(s) == 0 {
		return 0, ""
	}

	end := 0
	for ; end < len(s); end++ {
		if s[end] == '(' {
			break
		}
	}

	num, _ := strconv.Atoi(s[:end])
	return num, s[end:]
}

func main() {
	s := "4(2(3)(1))(6(5))"
	fmt.Println(str2tree(s))
}
