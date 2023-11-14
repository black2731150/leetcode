package main

import "fmt"

func countCompleteSubarrays(nums []int) int {
	n := len(nums)

	allMap := make(map[int]bool)
	for _, v := range nums {
		allMap[v] = true
	}

	annswer := 0
	for i := range nums {
		theMap := copyMap(allMap)

		for j := i; j < n; j++ {
			delete(theMap, nums[j])

			if len(theMap) == 0 {
				annswer++
			}
		}
	}
	return annswer
}

func copyMap(originalMap map[int]bool) map[int]bool {
	newMap := make(map[int]bool)
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}

func main() {
	nums := []int{1, 2, 3, 4, 3, 2}
	fmt.Println(countCompleteSubarrays(nums))
}
