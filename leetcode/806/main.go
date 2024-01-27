package main

import "fmt"

func numberOfLines(widths []int, s string) []int {
	w := 0
	answer := 0
	for i := range s {
		theNum := widths[s[i]-'a']
		if w+theNum > 100 {
			w = theNum
			answer++
		} else {
			w += theNum
		}
	}
	return []int{answer + 1, w}
}

func main() {
	withs := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(numberOfLines(withs, s))
}
