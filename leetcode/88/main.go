package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)
	nums3Index := 0

	num1Index, nums2Index := 0, 0
	for num1Index < m && nums2Index < n {
		if nums1[num1Index] < nums2[nums2Index] {
			nums3[nums3Index] = nums1[num1Index]
			num1Index++
		} else {
			nums3[nums3Index] = nums2[nums2Index]
			nums2Index++
		}

		nums3Index++
	}

	if num1Index < m {
		for i := num1Index; i < m; i++ {
			nums3[nums3Index] = nums1[i]
			nums3Index++
		}
	}

	if nums2Index < n {
		for i := nums2Index; i < n; i++ {
			nums3[nums3Index] = nums2[i]
			nums3Index++
		}
	}

	copy(nums1, nums3)
}

// 脑残方法
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	nums1 = nums1[:m]
// 	nums2 = nums2[:n]
// 	nums1 = append(nums1, nums2...)
// 	sort.Ints(nums1)
// }

func main() {
	nums1 := []int{1, 2, 4, 0, 0, 0}
	nums2 := []int{3, 4, 5}
	merge(nums1, len(nums1)+len(nums2), nums2, len(nums2))
}
