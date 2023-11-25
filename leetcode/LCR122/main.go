package main

import "fmt"

func pathEncryption(path string) string {
	n := len(path)
	answer := ""
	for i := 0; i < n; i++ {
		if path[i] == '.' {
			answer = answer + " "
		} else {
			answer = answer + string(path[i])
		}
	}
	return answer
}

func main() {
	path := "a.aef.qerf.bb"
	fmt.Println(pathEncryption(path))
}
