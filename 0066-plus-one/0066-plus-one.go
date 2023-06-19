func plusOne(digits []int) []int {
    lastElement := len(digits) - 1
    
    for i := lastElement; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    
    return append([]int{1}, digits...)
}