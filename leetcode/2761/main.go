package main

import (
	"fmt"
)

// var primes map[int]struct{} = make(map[int]struct{})

// func init() {
// 	for i := 2; i <= int(1e6); i++ {
// 		if isPrime(i) {
// 			primes[i] = struct{}{}
// 		}
// 	}
// }

// func isPrime(x int) bool {
// 	if x == 1 {
// 		return false
// 	}

// 	for i := 2; i*i <= x; i++ {
// 		if x%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func findPrimePairs(n int) [][]int {
// 	if n <= 3 {
// 		return nil
// 	}

// 	type point struct {
// 		x, y int
// 	}
// 	amap := make(map[point]struct{})
// 	for x := range primes {
// 		if _, find := primes[n-x]; find {
// 			p := point{}
// 			if x <= n-x {
// 				p.x = x
// 				p.y = n - x
// 			} else {
// 				p.x = n - x
// 				p.y = x
// 			}
// 			amap[p] = struct{}{}
// 		}
// 	}

// 	answer := [][]int{}
// 	for p := range amap {
// 		answer = append(answer, []int{p.x, p.y})
// 	}

// 	sort.Slice(answer, func(i, j int) bool {
// 		return answer[i][0] < answer[j][0]
// 	})

// 	return answer
// }

const pLen int = int(1e6) + 1

var primes []bool = make([]bool, pLen)

func init() {
	for i := 1; i < pLen; i++ {
		if isPrime(i) {
			primes[i] = true
		}
	}
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true && x != 1
}

func findPrimePairs(n int) [][]int {
	if n <= 3 {
		return nil
	}

	answer := [][]int{}
	for i := 2; i <= n-i; i++ {
		if primes[i] {
			if primes[n-i] {
				answer = append(answer, []int{i, n - i})
			}
		}
	}
	return answer
}

func main() {
	n := 10
	fmt.Println(findPrimePairs(n))
}
