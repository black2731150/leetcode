package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[0:help(len(str1), len(str2))]
}

func help(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func main() {
	s1 := "ABCABC"
	s2 := "ABC"
	fmt.Println(gcdOfStrings(s1, s2))
}
