package main

import "fmt"

func minimumPartition(s string, k int) int {
	tmp := 0
	answer := 0
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		if tmp*10+digit <= k {
			tmp = tmp*10 + digit
		} else {
			answer++
			if digit > k {
				return -1
			}
			tmp = digit
		}
	}

	return answer + 1
}

func main() {
	s := "238182"
	k := 5
	fmt.Println(minimumPartition(s, k))
}
