package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	theMap := make(map[int][]string)
	theMap[2] = []string{"a", "b", "c"}
	theMap[3] = []string{"d", "e", "f"}
	theMap[4] = []string{"g", "h", "i"}
	theMap[5] = []string{"j", "k", "l"}
	theMap[6] = []string{"m", "n", "o"}
	theMap[7] = []string{"p", "q", "r", "s"}
	theMap[8] = []string{"t", "u", "v"}
	theMap[9] = []string{"w", "x", "y", "z"}

	var answer []string
	var help func(combination string, nextDigits string)
	help = func(combination string, nextDigits string) {
		if len(nextDigits) == 0 {
			answer = append(answer, combination)
		} else {
			digit := nextDigits[0] - '0'
			letters := theMap[int(digit)]
			for _, letter := range letters {
				help(combination+letter, nextDigits[1:])
			}
		}
	}

	help("", digits)
	return answer
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
