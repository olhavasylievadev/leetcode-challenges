import "strings"

func detectCapitalUse(word string) bool {
    if word == strings.ToUpper(word) || word == strings.ToLower(word)  {
        return true
    } else {
        temp := strings.ToLower(word)
        if word == strings.Title(temp) {
            return true
        }
    }
    
    return false
}