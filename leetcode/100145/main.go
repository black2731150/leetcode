package main

func countCompleteSubstrings(word string, k int) int {
	sums := make([][]int, len(word)+1)
	for i := range sums {
		sums[i] = make([]int, 26)
	}

	for i, ch := range word {
		copy(sums[i+1], sums[i]) // 复制前缀和数组的当前行
		sums[i+1][ch-'a']++
	}

	answer := 0
	for i := 0; i < len(word); i++ {
		for j := i + 1; j <= len(word); j++ {
			is := true
			for x := 0; x < 26; x++ {
				if sums[j][x]-sums[i][x] != k {
					is = false
					break
				}
			}

			if is && j-i > 1 {
				for p := i; p < j-1; p++ {
					if abs(int(word[p]-'a')-int(word[p+1]-'a')) > 2 {
						is = false
						break
					}
				}
			}
			if is {
				answer++
			}
		}
	}

	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
