package main

import "fmt"

func minBitFlips(start int, goal int) int {
	answer := 0
	for i := 0; i < 32; i++ {
		if (start)&1^(goal)&1 == 1 {
			answer++
		}
		start = start >> 1
		goal = goal >> 1
	}

	return answer
}

//还不会位运算时的野蛮方法
// func minBitFlips(start int, goal int) int {
// 	bignum := max(start, goal)
// 	smallnnum := min(start, goal)

// 	bigbin := bin(bignum)
// 	smallbin := bin(smallnnum)

// 	lenbig := len(bigbin)
// 	lensmall := len(smallbin)

// 	for i := 0; i < lenbig-lensmall; i++ {
// 		smallbin = append(smallbin, false)
// 	}

// 	asnwer := 0
// 	for i := 0; i < lenbig; i++ {
// 		if xor(smallbin[i], bigbin[i]) {
// 			asnwer++
// 		}
// 	}

// 	return asnwer
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// func bin(x int) []bool {
// 	answer := []bool{}
// 	for x > 0 {
// 		if x%2 == 1 {
// 			answer = append(answer, true)
// 		} else {
// 			answer = append(answer, false)
// 		}
// 		x = x / 2
// 	}
// 	return answer
// }

// func xor(x, y bool) bool {
// 	return x != y
// }

func main() {
	start := 10
	goal := 7
	fmt.Println(minBitFlips(start, goal))
}
