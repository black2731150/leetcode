package main

func wordBreak(s string, wordDict []string) bool {

	theWordMap := make(map[string]struct{})
	for _, v := range wordDict {
		theWordMap[v] = struct{}{}
	}

	word := 0
	for i := 0; i < len(s); i++ {

	}
}
