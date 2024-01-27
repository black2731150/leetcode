package main

import "fmt"

func minSum(nums1 []int, nums2 []int) int64 {
	sum1, zero1 := arrsum(nums1)
	sum2, zero2 := arrsum(nums2)

	if sum1 > sum2 {
		if zero2 == 0 || (sum2+int64(zero2) > sum1 && zero1 == 0) {
			return -1
		}
	}

	if sum2 > sum1 {
		if zero1 == 0 || (sum1+int64(zero1) > sum2 && zero2 == 0) {
			return -1
		}
	}

	if sum1 == sum2 {
		if (zero1 > 0 && zero2 == 0) || (zero1 == 0 && zero2 > 0) {
			return -1
		}
	}

	if sum1+int64(zero1) > sum2+int64(zero2) {
		return sum1 + int64(zero1)
	} else {
		return sum2 + int64(zero2)
	}
}

func arrsum(arr []int) (int64, int) {
	var answer int64
	var zeronum int

	for _, v := range arr {
		if v == 0 {
			zeronum++
		}
		answer += int64(v)
	}

	return answer, zeronum
}

func main() {
	nums1 := []int{0, 0, 10, 10, 12, 0, 13, 6, 0, 2, 10}
	nums2 := []int{24, 5, 12, 22}
	fmt.Println(minSum(nums1, nums2))
}
