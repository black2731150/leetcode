package main

import "fmt"

func permutation(S string) []string {
	bs := []byte(S)
	n := len(bs)

	answer := []string{}

	var help func(keyong []byte, yiyou []byte)
	help = func(keyong []byte, yiyou []byte) {
		if len(yiyou) == n {
			answer = append(answer, string(yiyou))
		}

		if len(keyong) == 0 {
			return
		}

		for i := 0; i < len(keyong); i++ {
			cyiyou := make([]byte, len(yiyou))
			copy(cyiyou, yiyou)
			cyiyou = append(cyiyou, keyong[i])
			ckeyong := append([]byte{}, keyong[:i]...)
			ckeyong = append(ckeyong, keyong[i+1:]...)

			help(ckeyong, cyiyou)
		}
	}

	help(bs, []byte{})

	return answer
}

func main() {
	S := "qwe"
	fmt.Println(permutation(S))
}
