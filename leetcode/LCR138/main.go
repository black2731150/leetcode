package main

import "fmt"

func validNumber(s string) bool {

	if len(s) == 0 {
		return false
	}

	for s[0] == ' ' {
		s = s[1:]
		if len(s) == 0 {
			return false
		}
	}

	for s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
		if len(s) == 0 {
			return false
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == 'e' || s[i] == 'E' {
			return (isZhengshu(s[:i]) || isXiaoshu(s[:i])) && isZhengshu(s[i+1:])
		}
	}

	return isZhengshu(s) || isXiaoshu(s)

}

func isXiaoshu(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			continue
		} else {
			if s[i] == '.' {
				if len(s[i+1:]) == 0 && len(s[:i]) != 0 {
					return true
				} else {
					if i+1 < len(s) && (s[i+1] == '-' || s[i+1] == '+') {
						return false
					}
					return isZhengshu(s[i+1:])
				}
			} else {
				return false
			}
		}
	}
	return true
}

func isZhengshu(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "1654165"
	fmt.Println(validNumber(s))
}
