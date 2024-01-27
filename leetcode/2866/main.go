package main

import "fmt"

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	left, right := make([]int, n), make([]int, n)
	stackLeft, stackRight := []int{}, []int{}

	for i := 0; i < n; i++ {
		for len(stackLeft) > 0 && maxHeights[i] < maxHeights[stackLeft[len(stackLeft)-1]] {
			stackLeft = stackLeft[:len(stackLeft)-1]
		}

		if len(stackLeft) == 0 {
			left[i] = (i + 1) * maxHeights[i]
		} else {
			lastIndex := stackLeft[len(stackLeft)-1]
			left[i] = left[lastIndex] + (i-lastIndex)*maxHeights[i]
		}

		stackLeft = append(stackLeft, i)
	}

	answer := 0

	for i := n - 1; i >= 0; i-- {
		for len(stackRight) > 0 && maxHeights[i] < maxHeights[stackRight[len(stackRight)-1]] {
			stackRight = stackRight[:len(stackRight)-1]
		}

		if len(stackRight) == 0 {
			right[i] = (n - i) * maxHeights[i]
		} else {
			lastIndex := stackRight[len(stackRight)-1]
			right[i] = right[lastIndex] + (lastIndex-i)*maxHeights[i]
		}
		stackRight = append(stackRight, i)

		answer = max(answer, left[i]+right[i]-maxHeights[i])
	}

	return int64(answer)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	maxHights := []int{3, 2, 5, 5, 2, 3}
	fmt.Println(maximumSumOfHeights(maxHights))
}
