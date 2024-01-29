package main

import "fmt"

func flowerGame(n int, m int) int64 {
	answer := int64(0)

	nji, nou := 0, 0
	mji, mou := 0, 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			nou++
		} else {
			nji++
		}
	}

	for i := 1; i <= m; i++ {
		if i%2 == 0 {
			mou++
		} else {
			mji++
		}
	}

	answer = int64(nji)*int64(mou) + int64(nou)*int64(mji)

	return answer
}

func main() {
	m := 3
	n := 5
	fmt.Println(flowerGame(n, m))
}
