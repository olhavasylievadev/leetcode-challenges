func isPalindrome(x int) bool {
    inputX := x
    reverse := 0
    
    for x>0 {
        remainder := x % 10
        reverse = (reverse * 10) + remainder
        x /= 10
    }
    
    return inputX == reverse
}