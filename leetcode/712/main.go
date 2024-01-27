package main

import "fmt"

func minimumDeleteSum(s1 string, s2 string) int {
	bp := make([][]int, len(s1)+1)
	for i := range bp {
		bp[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		bp[i][0] = bp[i-1][0] + int(s1[i-1])
	}

	for j := 1; j <= len(s2); j++ {
		bp[0][j] = bp[0][j-1] + int(s2[j-1])
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				bp[i][j] = bp[i-1][j-1]
			} else {
				bp[i][j] = min(int(s1[i-1])+bp[i-1][j], int(s2[j-1])+bp[i][j-1])
			}
		}
	}

	return bp[len(s1)][len(s2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	s1 := "ccaccjp"
	s2 := "fwosarcwge"
	fmt.Println(minimumDeleteSum(s1, s2))
}
