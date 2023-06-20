import "strings"

func truncateSentence(s string, k int) string {
    myStr := strings.Split(s, " ")
    myStr = myStr[0:k]
    newStr := strings.Join(myStr, " ")
    
    return string(newStr)
}