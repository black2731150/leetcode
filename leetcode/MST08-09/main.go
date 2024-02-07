package main

import "fmt"

func generateParenthesis(n int) []string {

	answer := []string{}

	var help func(l, r int, tmp []byte)
	help = func(l, r int, tmp []byte) {
		if l == 0 && r == 0 {
			answer = append(answer, string(tmp))
			return
		}

		if l < 0 || r < 0 {
			return
		}

		if r > l {
			ntmp1 := append([]byte{}, tmp...)
			ntmp1 = append(ntmp1, '(')
			help(l-1, r, ntmp1)

			ntmp2 := append([]byte{}, tmp...)
			ntmp2 = append(ntmp2, ')')
			help(l, r-1, ntmp2)
		}

		if l == r {
			ntmp := append([]byte{}, tmp...)
			ntmp = append(ntmp, '(')
			help(l-1, r, ntmp)
		}
	}

	help(n, n, []byte{})
	return answer
}

func main() {
	n := 9
	fmt.Println(generateParenthesis(n))
}
