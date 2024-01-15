package main

import "fmt"

// 空间复杂度过高
// func sumSubarrayMins(arr []int) int {
// 	mod := int(1e9 + 7)

// 	n := len(arr)

// 	bp := make([][]int, n)
// 	for i := range bp {
// 		bp[i] = make([]int, n)
// 	}

// 	for i, v := range arr {
// 		bp[i][i] = v
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			bp[i][j] = min(bp[i][j-1], arr[j])
// 		}
// 	}

// 	answer := 0
// 	for i := range bp {
// 		for _, v2 := range bp[i] {
// 			answer = (answer + v2) % mod
// 		}
// 	}

// 	return answer
// }

func sumSubarrayMins(arr []int) int {
	mod := int(1e9 + 7)

	n := len(arr)
	answer := 0

	for i := 0; i < n; i++ {
		minNum := arr[i]
		for j := i; j < n; j++ {
			minNum = min(minNum, arr[j])
			answer = (answer + minNum) % mod
		}
	}
	return answer
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	arr := []int{3, 1, 2, 4}
	fmt.Println(sumSubarrayMins(arr))
}
