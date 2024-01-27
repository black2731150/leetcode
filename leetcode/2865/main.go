package main

import "fmt"

func maximumSumOfHeights(maxHeights []int) int64 {
	answer := int64(0)
	for i := range maxHeights {
		leftSum := int64(0)
		leftmin := maxHeights[i]
		for j := i - 1; j >= 0; j-- {
			if maxHeights[j] < leftmin {
				leftmin = maxHeights[j]
			}
			leftSum += int64(leftmin)
		}

		rightSum := int64(0)
		rightmin := maxHeights[i]
		for j := i + 1; j < len(maxHeights); j++ {
			if maxHeights[j] < rightmin {
				rightmin = maxHeights[j]
			}
			rightSum += int64(rightmin)
		}

		answer = max(answer, rightSum+leftSum+int64(maxHeights[i]))

	}

	return answer
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func main() {
	maxHights := []int{3, 2, 5, 5, 2, 3}
	fmt.Println(maximumSumOfHeights(maxHights))
}
