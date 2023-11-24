package main

import "fmt"

func closestPrimes(left int, right int) []int {
	answer := []int{-1, -1}

	primes := []int{}
	for i := left; i <= right; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	if len(primes) < 2 {
		return answer
	}

	min := right
	for i := 0; i < len(primes)-1; i++ {
		if x := primes[i+1] - primes[i]; x <= min {
			if x == 1 {
				answer[0] = primes[i]
				answer[1] = primes[i+1]
				return answer
			}

			if x == min {
				if primes[i] < answer[0] {
					answer[0] = primes[i]
					answer[1] = primes[i+1]
				}
			} else {
				answer[0] = primes[i]
				answer[1] = primes[i+1]
				min = x
			}
		}
	}

	return answer
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	left := 10
	right := 19
	fmt.Println(closestPrimes(left, right))
}
