package main

import "fmt"

func getCommon(nums1 []int, nums2 []int) int {
	answer := -1

	n1i, n2i := 0, 0

	for n1i < len(nums1) && n2i < len(nums2) {
		if nums1[n1i] == nums2[n2i] {
			return nums1[n1i]
		}

		if nums1[n1i] < nums2[n2i] {
			n1i++
		} else {
			n2i++
		}
	}

	return answer
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4}
	fmt.Println(getCommon(nums1, nums2))
}
