package main

import "fmt"

func diagonalPrime(nums [][]int) int {
	answer := 0

	n := len(nums)
	for i := 0; i < n; i++ {
		if isPirme(nums[i][i]) {
			answer = max(answer, nums[i][i])
		}

		if isPirme(nums[i][n-1-i]) {
			answer = max(answer, nums[i][n-1-i])
		}
	}

	return answer
}

func isPirme(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 10, 11},
	}
	fmt.Println(diagonalPrime(nums))
}
