package main

import "fmt"

func entityParser(text string) string {
	n := len(text)
	answer := ""
	for i := 0; i < n; {
		if text[i] == '&' {
			if i+4 <= n && text[i:i+4] == "&lt;" {
				answer += "<"
				i += 4
			} else if i+4 <= n && text[i:i+4] == "&gt;" {
				answer += ">"
				i += 4
			} else if i+5 <= n && text[i:i+5] == "&amp;" {
				answer += "&"
				i += 5
			} else if i+6 <= n && text[i:i+6] == "&quot;" {
				answer += `"`
				i += 6
			} else if i+6 <= n && text[i:i+6] == "&apos;" {
				answer += `'`
				i += 6
			} else if i+7 <= n && text[i:i+7] == "&frasl;" {
				answer += `/`
				i += 7
			} else {
				answer += string(text[i])
				i++
			}
		} else {
			answer += string(text[i])
			i++
		}
	}
	return answer
}

func main() {
	text := "&amp; is an HTML entity but &ambassador; is not."
	fmt.Println(entityParser(text))
}
