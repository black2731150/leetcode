package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	theMap := make(map[byte]int, 26)
	n := len(text)
	for i := 0; i < n; i++ {
		theMap[text[i]]++
	}

	answer := 0
	for {
		if theMap['b'] -= 1; theMap['b'] < 0 {
			break
		}

		if theMap['a'] -= 1; theMap['a'] < 0 {
			break
		}

		if theMap['l'] -= 2; theMap['l'] < 0 {
			break
		}

		if theMap['o'] -= 2; theMap['o'] < 0 {
			break
		}

		if theMap['n'] -= 1; theMap['n'] < 0 {
			break
		}

		answer++
	}

	return answer
}

func main() {
	text := "balloondosndksodms"
	fmt.Println(maxNumberOfBalloons(text))
}
