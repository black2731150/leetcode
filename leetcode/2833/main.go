package main

import "fmt"

func furthestDistanceFromOrigin(moves string) int {
	leftNum := 0
	rightNum := 0
	noNum := 0

	for _, ch := range moves {
		switch ch {
		case 'L':
			leftNum++
		case 'R':
			rightNum++
		case '_':
			noNum++
		}
	}

	if leftNum > rightNum {
		return leftNum - rightNum + noNum
	} else {
		return rightNum - leftNum + noNum
	}
}

func main() {
	moves := "_R__LL_"
	fmt.Println(furthestDistanceFromOrigin(moves))
}
