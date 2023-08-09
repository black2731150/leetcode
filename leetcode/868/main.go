package main

import "fmt"

func binaryGap(n int) int {
	binaryArr := []int{}
	for i := n; i > 0; i = i / 2 {
		binaryArr = append(binaryArr, i%2)
	}

	firstOne := 0
	for i := 0; i < len(binaryArr); i++ {
		if binaryArr[i] == 0 {
			continue
		} else {
			firstOne = i
			break
		}
	}

	start, end := firstOne, firstOne
	max := 0
	for end < len(binaryArr) {
		if binaryArr[start] == binaryArr[end] {
			if max < end-start {
				max = end - start
			}
			start = end
		}
		end++
	}
	return max
}

func main() {
	n := 8
	fmt.Println(binaryGap(n))
}
