package main

func multiply(A int, B int) int {
	if B == 1 {
		return A
	} else {
		if B&1 == 1 {
			return A + multiply(A<<1, B>>1)
		} else {
			return multiply(A<<1, B>>1)
		}
	}
}
