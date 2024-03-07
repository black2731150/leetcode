package main

import "fmt"

func strongPasswordCheckerII(password string) bool {
	var lenght, haveLow, haveUp, haveDigit, haveSpe, haveLianxu bool

	lenght = len(password) >= 8

	for i := range password {
		if password[i] >= 'a' && password[i] <= 'z' {
			haveLow = true
		}

		if password[i] >= 'A' && password[i] <= 'Z' {
			haveUp = true
		}

		if password[i] >= '0' && password[i] <= '9' {
			haveDigit = true
		}

		for _, ch := range "!@#$%^&*()-+" {
			if password[i] == byte(ch) {
				haveSpe = true
			}
		}

		if i > 0 && password[i-1] == password[i] {
			haveLianxu = true
		}
	}

	return lenght && haveLow && haveUp && haveDigit && haveSpe && !haveLianxu
}

func main() {
	password := "IloveLe3tcode!"
	fmt.Println(strongPasswordCheckerII(password))
}
