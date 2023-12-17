package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 || n == 1 {
		return s
	}

	start, end := 0, 0

	for i := range s {
		len1 := help(s, i, i)
		len2 := help(s, i, i+1)
		len := max(len1, len2)

		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start : end+1]

}

func help(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//时间复杂度为O(n^3)
// func longestPalindrome(s string) string {
// 	n := len(s)
// 	// fmt.Println("n= ", n)

// 	if n <= 1 {
// 		return s
// 	}

// 	var maxstring string
// 	var maxnum int
// 	for i := 0; i < n; i++ {
// 		for j := i; j < n; j++ {
// 			if s[i] == s[j] {
// 				// fmt.Println(s[i : j+1])
// 				if IsPalindrome(s[i : j+1]) {
// 					if len(s[i:j+1]) > maxnum {
// 						maxstring = s[i : j+1]
// 						maxnum = len(s[i : j+1])
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return maxstring
// }

// func IsPalindrome(s string) bool {
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] == s[len(s)-1-i] {
// 			continue
// 		} else {
// 			return false
// 		}
// 	}

// 	return true
// }

func main() {
	s := "badad"
	fmt.Println(longestPalindrome(s))
}
