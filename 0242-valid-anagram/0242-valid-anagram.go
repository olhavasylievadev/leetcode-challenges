func isAnagram(s string, t string) bool {
    myStr1 := len(s)
    myStr2 := len(t)
    
    if myStr1 != myStr2 {
        return false
    }
    
    myMap := make(map[string]int)
    
    for i := 0; i < myStr1; i++ {
		myMap[string(s[i])]++
	}

	for i := 0; i < myStr2; i++ {
		myMap[string(t[i])]--
	}

	for i := 0; i < myStr1; i++ {
		if myMap[string(s[i])] != 0 {
			return false
		}
	}

	return true
}