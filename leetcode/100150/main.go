package main

import "fmt"

func maximumSetSize(nums1 []int, nums2 []int) int {
	n := len(nums1)

	num1Map, num2Map := make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		num1Map[nums1[i]]++
		num2Map[nums2[i]]++
	}

	needGo1 := n / 2
	for k := range num1Map {
		for num1Map[k] > 1 {
			if needGo1 > 0 {
				num1Map[k]--
				needGo1--
			} else {
				break
			}
		}
	}

	needGo2 := n / 2
	for k := range num2Map {
		for num2Map[k] > 1 {
			if needGo2 > 0 {
				num2Map[k]--
				needGo2--
			} else {
				break
			}
		}
	}

	if needGo1 > 0 {
		for k := range num1Map {
			if _, find := num2Map[k]; find {
				delete(num1Map, k)
				needGo1--
				if needGo1 == 0 {
					break
				}
			}
		}
	}

	if needGo2 > 0 {
		for k := range num2Map {
			if _, find := num1Map[k]; find {
				delete(num2Map, k)
				needGo2--
				if needGo2 == 0 {
					break
				}
			}
		}
	}

	for k := range num1Map {
		num2Map[k]++
	}

	return len(num2Map) - needGo1 - needGo2
}

func main() {
	nums1 := []int{1, 2, 1, 2}
	nums2 := []int{1, 1, 1, 1}
	fmt.Println(maximumSetSize(nums1, nums2))
}
