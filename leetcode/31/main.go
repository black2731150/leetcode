package main

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	n := len(nums)

	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	for i2, j := 0, len(nums[i+1:])-1; i2 < j; i2, j = i2+1, j-1 {
		nums[i+1:][i2], nums[i+1:][j] = nums[i+1:][j], nums[i+1:][i2]
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
}
