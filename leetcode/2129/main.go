package main

import "fmt"

func minMoves(target int, maxDoubles int) int {
	answer := 0
	for target > 1 {
		if target%2 == 0 && maxDoubles > 0 {
			target = target / 2
			maxDoubles--

		} else {
			if maxDoubles > 0 {
				target--
			} else {
				answer = answer + target - 1
				target = 1
				return answer
			}
		}
		answer++
		fmt.Println(target)
	}
	return answer
}

func main() {
	target := 656101986
	maxD := 1
	fmt.Println(minMoves(target, maxD))
}
