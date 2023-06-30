func reverseWords(s string) string {
    words := strings.Fields(s)
	str := ""

	for _, word := range words {
		for i:=len(word)-1; i>=0; i-- {
			str += string(word[i])
		}
		str += " "
	}

	return strings.TrimSpace(str)
}