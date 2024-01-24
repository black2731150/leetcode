package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num1) < len(num2) {
		return multiply(num2, num1)
	}

	b1, b2 := []byte(num1), []byte(num2)
	total := make([]byte, len(b1)+len(b2))
	for i := range total {
		total[i] = '0'
	}

	for i := len(b1) - 1; i >= 0; i-- {
		n1 := int(b1[i] - '0')
		for j := len(b2) - 1; j >= 0; j-- {
			n2 := int(b2[j] - '0')

			product := n1*n2 + int(total[i+j+1]-'0')
			total[i+j+1] = byte(product%10 + '0')
			total[i+j] += byte(product / 10)
		}
	}

	if total[0] == '0' {
		total = total[1:]
	}

	return string(total)
}

func main() {
	a := "123"
	b := "456"
	fmt.Println(multiply(a, b))
}
