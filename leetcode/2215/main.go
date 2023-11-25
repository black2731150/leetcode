package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map, nums2Map := make(map[int]struct{}), make(map[int]struct{})
	for _, v := range nums1 {
		if _, ok := nums1Map[v]; !ok {
			nums1Map[v] = struct{}{}
		}
	}

	for _, v := range nums2 {
		if _, ok := nums2Map[v]; !ok {
			nums2Map[v] = struct{}{}
		}
	}

	answer1 := []int{}
	answer2 := []int{}
	for k := range nums1Map {
		if _, ok := nums2Map[k]; !ok {
			answer1 = append(answer1, k)
		}
	}

	for k := range nums2Map {
		if _, ok := nums1Map[k]; !ok {
			answer2 = append(answer2, k)
		}
	}
	return [][]int{answer1, answer2}
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
}
