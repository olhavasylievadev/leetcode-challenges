func romanToInt(s string) int {
    roman := make(map[string]int)
    roman["I"] = 1
    roman["V"] = 5
    roman["X"] = 10
    roman["L"] = 50
    roman["C"] = 100
    roman["D"] = 500
    roman["M"] = 1000
    
    sum := 0
    
    for i, j := range s[:len(s)-1] {
        if roman[string(j)] < roman[string(s[i+1])] {
            sum -= roman[string(j)]
            continue
        }
        sum += roman[string(j)]
    }
    sum += roman[string(s[len(s)-1])]
    
    return sum
}