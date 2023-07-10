func isPowerOfThree(n int) bool {
    if n == 1{
        return true
    }
    if n % 3 != 0 || n <= 0{
        return false
    }
    num := 1
    for num <= n {
        num *= 3
        if(num == n){
            return true
        }
    }
    return false
}