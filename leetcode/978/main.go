package main

import "fmt"

func maxTurbulenceSize(arr []int) int {
	n := len(arr)

	if n == 1 {
		return 1
	}

	diffNums := make([]int, n-1)
	for i := 0; i < n-1; i++ {
	}

	answer := 2
	maxLenght := 0

	right, left := 1, 0

	status := bijiao(arr[left], arr[right])

	left++
	right++

	for right < n {
		if x := bijiao(arr[left], arr[right]); x != status {
			status = x
			maxLenght++
			answer = max(answer, maxLenght)

			left++
			right++
		} else {
			maxLenght = 1
			status = x
			left++
			right++
		}
	}

	return answer
}

func bijiao(left, right int) bool {
	if left < right {
		return false
	} else {
		return true
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	arr := []int{4, 8, 12, 16}
	fmt.Println(maxTurbulenceSize(arr))
}
