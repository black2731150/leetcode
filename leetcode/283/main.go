package main

func moveZeroes(nums []int) {
	index := 0

	for _, v := range nums {
		if v != 0 {
			nums[index] = v
			index++
		}
	}

	for index < len(nums) {
		nums[index] = 0
		index++
	}
}

func main() {
	nums := []int{1, 2, 0, 4, 5, 6, 0, 0, 0, 0}
	moveZeroes(nums)
}
