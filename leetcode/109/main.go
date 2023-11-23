package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	var work func(nums []int) *TreeNode
	work = func(nums []int) *TreeNode {
		n := len(nums)
		if n == 0 {
			return nil
		}

		var answer *TreeNode = &TreeNode{Val: nums[n/2]}
		answer.Left = work(nums[:n/2])
		answer.Right = work(nums[n/2+1:])

		return answer

	}
	return work(nums)
}

func main() {
	head := &ListNode{Val: -10, Next: &ListNode{Val: -1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	fmt.Println(sortedListToBST(head))
}
