package main

import "fmt"

func areaOfMaxDiagonal(dimensions [][]int) int {

	maxArea, maxL := 0, 0

	for _, juxing := range dimensions {
		chang, kuan := juxing[0], juxing[1]
		if chang*chang+kuan*kuan >= maxL {
			if chang*chang+kuan*kuan == maxL {
				if chang*kuan > maxArea {
					maxArea = chang * kuan
				}
			} else {
				maxL = chang*chang + kuan*kuan
				maxArea = chang * kuan
			}
		}
	}

	return maxArea
}

func main() {
	dimensions := [][]int{{2, 6}, {5, 1}, {3, 10}, {8, 4}}
	fmt.Println(areaOfMaxDiagonal(dimensions))
}
