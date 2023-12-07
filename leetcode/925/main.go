package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	nameIndex := 0
	typedIndex := 0

	for typedIndex < len(typed) {
		if nameIndex < len(name) && name[nameIndex] == typed[typedIndex] {
			nameIndex++
			typedIndex++
		} else if typedIndex > 0 && typed[typedIndex] == typed[typedIndex-1] {
			typedIndex++
		} else {
			return false
		}

	}

	return nameIndex == len(name)
}

func main() {
	name := "alex"
	typed := "aaleex"
	fmt.Println(isLongPressedName(name, typed))
}
