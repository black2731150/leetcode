package main

import "fmt"

func one(n int, rotes string) int {

	answer := 0

	for i := 0; i < len(rotes)-1; i++ {
		if rotes[i] == rotes[i+1] {
			rotes = rotes[:i] + rotes[i+1:]
			n--
			i--
			answer++
		}
	}

	return answer

}

func main() {
	n := 0
	rotes := ""

	fmt.Scan(&n)
	fmt.Scan(&rotes)
	fmt.Println(one(n, rotes))
}
