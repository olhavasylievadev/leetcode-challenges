func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    word1s := ""
	word2s := ""

	for _, w1 := range word1 {
		word1s += w1
	}

	for _, w2 := range word2 {
		word2s += w2
	}

	return word1s == word2s
}