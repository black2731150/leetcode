package main

func maxProduct(words []string) int {
	n := len(words)
	answer := 0
	theMap := make([]uint32, n)
	for i, v := range words {
		var x uint32 = 0
		for _, ch := range v {
			x = x | (1 << (ch - 'a'))
		}
		theMap[i] = x
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if theMap[i]&theMap[j] == 0 {
				answer = max(answer, len(words[i])*len(words[j]))
			}
		}
	}
	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
