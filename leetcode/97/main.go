package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)

	if m+n != len(s3) {
		return false
	}

	bp := make([][]bool, n+1)
	for i := range bp {
		bp[i] = make([]bool, m+1)
	}

	bp[0][0] = true
	for i := 1; i <= n; i++ {
		bp[i][0] = bp[i-1][0] && s1[i-1] == s3[i-1]
	}

	for j := 1; j <= m; j++ {
		bp[0][j] = bp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			bp[i][j] = (bp[i-1][j] && s1[i-1] == s3[i-1+j]) || (bp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return bp[n][m]
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))
}
