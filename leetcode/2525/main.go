package main

import "fmt"

func categorizeBox(length int, width int, height int, mass int) string {
	answer := [2]bool{false, false}

	if length >= 10000 || width >= 10000 || height >= 10000 || length*width*height >= 1000000000 {
		answer[0] = true
	}

	if mass >= 100 {
		answer[1] = true
	}

	if answer[0] && answer[1] {
		return "Both"
	} else if !answer[0] && !answer[1] {
		return "Neither"
	} else if answer[0] && !answer[1] {
		return "Bulky"
	} else {
		return "Heavy"
	}
}

func main() {
	lenght := 1000
	width := 1000
	height := 1000
	mess := 1000

	fmt.Println(categorizeBox(lenght, width, height, mess))
}
