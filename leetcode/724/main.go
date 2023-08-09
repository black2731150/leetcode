package main

import "fmt"

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if sum(nums[:i]) == sum(nums[i+1:]) {
			return i
		}
	}

	return -1
}

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	return sum
}

// func pivotIndex(nums []int) int {
//     sum := 0
//     for _,v := range nums {
//         sum += v
//     }
//     left := 0
//     for i := 0;i < len(nums);i++ {
//         if left * 2 + nums[i] == sum {
//             return i
//         }
//         left += nums[i]
//     }
//     return -1
// }

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(pivotIndex(nums))
}
