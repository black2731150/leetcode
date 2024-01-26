package main

import "fmt"

var primes map[int]bool = make(map[int]bool, 100)

func init() {
	primes[1] = false
	for i := 2; i <= 100; i++ {
		primes[i] = isPrime(i)
	}
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func numPrimeArrangements(n int) int {
	mod := int(1e9 + 7)

	p := 0
	np := 0

	for i := 1; i <= n; i++ {
		if primes[i] {
			p++
		} else {
			np++
		}
	}

	pj := 1
	for i := 1; i <= p; i++ {
		pj = (pj * i) % mod
	}

	npj := 1
	for i := 1; i <= np; i++ {
		npj = (npj * i) % mod
	}
	return (pj * npj) % mod
}

func main() {
	n := 5
	fmt.Println(numPrimeArrangements(n))
}
