package main

import (
	"fmt"
)

func uniqueLetterString(s string) int {
	index := [26][2]int{}
	for i := range index {
		index[i][0], index[i][1] = -1, -1
	}

	result := 0
	for i, c := range s {
		charIndex := c - 'A'
		result += (i - index[charIndex][1]) * (index[charIndex][1] - index[charIndex][0])
		index[charIndex][0], index[charIndex][1] = index[charIndex][1], i
	}

	n := len(s)
	for c := 0; c < 26; c++ {
		result += (n - index[c][1]) * (index[c][1] - index[c][0])
	}

	return result
}

//采用前缀和的方法,时间复杂度为26*n^2 过高
// func uniqueLetterString(s string) int {
// 	n := len(s)

// 	charnums := make([][]int, n+1)
// 	for i := range charnums {
// 		charnums[i] = make([]int, 26)
// 	}

// 	for i := 0; i < n; i++ {
// 		copy(charnums[i+1], charnums[i])
// 		charnums[i+1][s[i]-'A']++
// 	}

// 	answer := 0
// 	for i := 0; i <= n; i++ {
// 		for j := i + 1; j <= n; j++ {
// 			x := 0
// 			for k := 0; k < 26; k++ {
// 				if one := charnums[j][k] - charnums[i][k]; one == 1 {
// 					x++
// 				}

// 			}
// 			answer += x
// 		}
// 	}

// 	return answer
// }

func main() {
	s := "ABA"
	fmt.Println(uniqueLetterString(s))
}
