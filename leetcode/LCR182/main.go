package main

import "fmt"

func dynamicPassword(password string, target int) string {
	return password[target:] + password[:target]
}

func main() {
	password := "helloworld"
	target := 6
	fmt.Println(dynamicPassword(password, target))
}
