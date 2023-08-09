package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {

	for _, v := range letters {
		if target < v {
			return v
		}
	}

	return letters[0]
}

func main() {
	letters := []byte{'x', 'x', 'y', 'y'}
	target := 'z'
	fmt.Println(string(nextGreatestLetter(letters, byte(target))))
}
