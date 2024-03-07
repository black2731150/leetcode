package main

import "fmt"

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}

	thmap := make(map[int]struct{})
	indexs := make([]int, len(primes))

	answer := []int{1}
	for len(answer) <= n {
		next := primes[0] * answer[indexs[0]]
		for i, prime := range primes {
			next = min(next, prime*answer[indexs[i]])
		}

		for i, prime := range primes {
			if next == prime*answer[indexs[i]] {
				indexs[i]++
				break
			}
		}

		if _, find := thmap[next]; !find {
			answer = append(answer, next)
			thmap[next] = struct{}{}
		}
	}

	return answer[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	n := 12
	primes := []int{2, 7, 13, 19}
	fmt.Println(nthSuperUglyNumber(n, primes))
}
