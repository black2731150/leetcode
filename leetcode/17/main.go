package main

import "fmt"

func letterCombinations(digits string) []string {

	answer := []string{}

	if len(digits) == 0 {
		return answer
	}

	TheMap := make(map[int][]string)
	TheMap[2] = []string{"a", "b", "c"}
	TheMap[3] = []string{"d", "e", "f"}
	TheMap[4] = []string{"g", "h", "i"}
	TheMap[5] = []string{"j", "k", "l"}
	TheMap[6] = []string{"m", "n", "o"}
	TheMap[7] = []string{"p", "q", "r", "s"}
	TheMap[6] = []string{"t", "u", "v"}
	TheMap[7] = []string{"w", "x", "y", "z"}

	var help func(have []string, i int)
	help = func(have []string, i int) {
		if i == len(digits) {
			answer = append(answer, have...)
		} else {
			for _, ch := range TheMap[int(digits[i]-'0')] {
				copyhave := make([]string, len(have))
				copy(copyhave, have)
				for i2 := range copyhave {
					copyhave[i2] += ch
				}
				help(have, i+1)
			}
		}
	}
	help(TheMap[int(digits[0]-'0')], 1)

	return answer
}

func main() {
	digit := "23"
	fmt.Println(letterCombinations(digit))
}
