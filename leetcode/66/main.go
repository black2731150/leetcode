package main

func plusOne(digits []int) []int {
	count := 0
	digits[len(digits)-1]++

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + count
		if digits[i] >= 10 {
			count = 1
			digits[i] = digits[i] % 10
			continue
		} else {
			count = 0
			break
		}
	}

	if count == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
