package main

func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] + numbers[1]
	numbers[1] = numbers[0] - numbers[1]
	numbers[0] = numbers[0] - numbers[1]
	return numbers
}

func main() {
	numbers := []int{0, 2147483647}
	swapNumbers(numbers)
}
