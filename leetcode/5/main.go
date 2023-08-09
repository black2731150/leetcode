package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	// fmt.Println("n= ", n)

	if n <= 1 {
		return s
	}

	var maxstring string
	var maxnum int
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				// fmt.Println(s[i : j+1])
				if IsPalindrome(s[i : j+1]) {
					if len(s[i:j+1]) > maxnum {
						maxstring = s[i : j+1]
						maxnum = len(s[i : j+1])
					}
				}
			}
		}
	}
	return maxstring
}

func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] == s[len(s)-1-i] {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "badad"
	fmt.Println(longestPalindrome(s))
}
