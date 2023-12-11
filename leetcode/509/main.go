package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums[n]
}

func main() {
	n := 10
	fmt.Println(fib(n))
}
