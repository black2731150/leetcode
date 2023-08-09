package main

import (
	"fmt"
	"sort"
)

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	i,j,n,m:=0,0,len(nums1),len(nums2)
// 	nums := make([]int, n+m)
// 	idx := 0
// 	for i < n && j < m {
// 		if nums1[i] < nums2[j] {
// 			nums[idx] =nums1[i]
// 			i++
// 		}else {
// 			nums[idx] =nums2[j]
// 			j++
// 		}
// 		idx++
// 	}

// 	for i < n {
// 		nums[idx] =nums1[i]
// 		i++
// 		idx++
// 	}

// 	for j < m {
// 		nums[idx] =nums2[j]
// 		j++
// 		idx++
// 	}

// 	mid := (n+m)/2
// 	if (n+m)%2==0 {
// 		return float64(nums[mid] + nums[mid-1])/2
// 	}else {
// 		return float64(nums[mid])
// 	}
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	n := len(nums1)
	if n%2 == 0 {
		return float64(nums1[(n/2)-1]+nums1[(n/2)]) / 2
	} else {
		return float64(nums1[n/2])
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
