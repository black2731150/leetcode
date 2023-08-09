package main

func CheckPermutation(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	nums := make([]int, 26)
	for i := range s1 {
		nums[s1[i]-'a']++
	}

	for i := range s2 {
		if nums[s2[i]-'a'] == 0 {
			return false
		} else {
			nums[s2[i]-'a']--
		}
	}
	return true
}
