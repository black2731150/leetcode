package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	nums := []int{}

	var order func(r *TreeNode)
	order = func(r *TreeNode) {
		if r == nil {
			return
		}

		order(r.Left)
		nums = append(nums, r.Val)
		order(r.Right)
	}

	order(root)

	answer := [][]int{}
	for _, v := range queries {

		left, right := 0, len(nums)-1

		one := []int{-1, -1}
		for left <= right {
			mid := (left + right) / 2

			if nums[mid] == v {
				left = mid
				right = mid
				break
			}

			if nums[mid] > v {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		if left > -1 && left < len(nums) {
			one[1] = nums[left]
		}

		if right > -1 && right < len(nums) {
			one[0] = nums[right]
		}

		answer = append(answer, one)
	}

	return answer
}

func main() {
	root := &TreeNode{}
	q := []int{12, 1}
	closestNodes(root, q)
}
