package main

type WordsFrequency struct {
	wordMap map[string]int
}

func Constructor(book []string) WordsFrequency {
	wordMap := make(map[string]int)
	for _, v := range book {
		wordMap[v]++
	}

	return WordsFrequency{
		wordMap: wordMap,
	}
}

func (w *WordsFrequency) Get(word string) int {
	return w.wordMap[word]
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
