package main

import "fmt"

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n < 2 {
		return n
	}

	inc, dec := make([]int, n), make([]int, n)
	inc[0], dec[0] = 1, 1
	maxLength := 1

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			inc[i] = dec[i-1] + 1
			dec[i] = 1
		} else if arr[i] < arr[i-1] {
			dec[i] = inc[i-1] + 1
			inc[i] = 1
		} else {
			inc[i], dec[i] = 1, 1
		}
		maxLength = max(maxLength, inc[i], dec[i])
	}

	return maxLength
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func main() {
	arr := []int{4, 8, 12, 16}
	fmt.Println(maxTurbulenceSize(arr))
}
