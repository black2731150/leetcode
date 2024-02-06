package main

import "fmt"

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	answer := []byte{}
	ch := S[0]
	chNum := 0
	for i := 0; i < len(S); i++ {
		if S[i] == ch {
			chNum++
		} else {
			answer = append(answer, ch)
			answer = append(answer, numToByte(chNum)...)

			chNum = 1
			ch = S[i]
		}
	}

	answer = append(answer, ch)
	answer = append(answer, numToByte(chNum)...)

	if len(answer) >= len(S) {
		return S
	}

	return string(answer)
}

func numToByte(n int) []byte {
	answer := []byte{}
	for n > 0 {
		answer = append(answer, '0'+byte(n%10))
		n = n / 10
	}
	for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}

	return answer
}

func main() {
	S := "aabcccccaaa"
	fmt.Println(compressString(S))
}
