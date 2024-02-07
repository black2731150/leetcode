package main

import "fmt"

func permutation(S string) []string {
	bs := []byte(S)
	n := len(bs)

	answer := []string{}

	theMap := make(map[string]struct{})

	var help func(keyong []byte, yiyou []byte)
	help = func(keyong []byte, yiyou []byte) {
		if len(yiyou) == n {
			theMap[string(yiyou)] = struct{}{}
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

	for k := range theMap {
		answer = append(answer, k)
	}

	return answer
}

func main() {
	S := "qwqe"
	fmt.Println(permutation(S))
}
