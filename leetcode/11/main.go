package main

import "fmt"

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1

	answer := min(height[left], height[right]) * (right - left)
	for left < right {

		if height[left] > height[right] {
			area := height[right] * (right - left)
			answer = max(area, answer)
			right--
		} else {
			area := height[left] * (right - left)
			answer = max(area, answer)
			left++
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	height := []int{1, 3, 2, 5, 25, 24, 5}
	fmt.Println(maxArea(height))
}
