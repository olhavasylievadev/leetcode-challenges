func findTheDifference(s string, t string) byte {
    runeS := []byte(s)
    runeT := []byte(t)
    myRune := runeT[len(s)]
    
    for i, _ := range runeS {
        myRune -= runeS[i]
		myRune += runeT[i]
    }
    
    return myRune
}