func addDigits(num int) int {
    if num < 9 {
        return num
    }
    
    switch num % 9 {
	case 0:
		return 9
	default:
		return num % 9
	}
}