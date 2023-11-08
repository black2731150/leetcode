package main

import "fmt"

func numTrees(n int) int {
	nums := make([]int, n+1)
	nums[0] = 1
	nums[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}
	return nums[n]
}

// func numTrees(n int) int {
//     C := 1
//     for i := 0; i < n; i++ {
//         C = C * 2 * (2 * i + 1) / (i + 2);
//     }
//     return C
// }

func main() {
	n := 17
	fmt.Println(numTrees(n))
}
