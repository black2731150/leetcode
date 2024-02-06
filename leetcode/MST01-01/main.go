package main

func isUnique(astr string) bool {
	x := uint(0)
	for i := 0; i < len(astr); i++ {
		ch := int(astr[i] - 'a')

		if (x>>ch)&1 == 1 {
			return false
		} else {
			x = x | (1 << ch)
		}
	}
	return true
}
