package main

import "fmt"

func canBeIncreasing(nums []int) bool {
	lenNums := len(nums)
	for i := 1; i+1 < lenNums; i++ {
		if nums[i] > nums[i-1] {
			continue
		} else {
			nums1 := []int{}
			nums2 := []int{}

			nums1 = append(nums1, nums[:i]...)
			nums1 = append(nums1, nums[i+1:]...)

			nums2 = append(nums2, nums[:i-1]...)
			nums2 = append(nums2, nums[i:]...)
			return check(nums1) || check(nums2)
		}
	}
	return true
}

func check(nums []int) bool {
	lenNums := len(nums)
	for i := 1; i < lenNums; i++ {
		if nums[i-1] < nums[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1, 2, 10, 5, 7}
	fmt.Println(canBeIncreasing(nums))
}
