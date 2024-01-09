package main

func partition(s string) [][]string {
	answer := [][]string{}

	bp := make([][]bool, len(s))
	for i, _ := range bp {
		bp[i] = make([]bool, len(s))
	}

	for i := 1; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			bp[i][j] = (s[i] == s[j]) && bp[i-1][j+1]
			if bp[i][j] {
				answer = append(answer, s[i:j])
			}
		}
	}
}
