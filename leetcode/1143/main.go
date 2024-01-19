package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	bp := make([][]int, len(text1)+1)
	for i := range bp {
		bp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				bp[i][j] = 1 + bp[i-1][j-1]
			} else {
				bp[i][j] = maxThree(bp[i-1][j-1], bp[i-1][j], bp[i][j-1])
			}
		}
	}

	return bp[len(text1)][len(text2)]
}

func maxTow(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxThree(x, y, z int) int {
	return maxTow(maxTow(x, y), z)
}

func main() {
	text1 := "bsbininm"
	text2 := "jmjkbkjkv"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
