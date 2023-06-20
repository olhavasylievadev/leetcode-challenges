func fizzBuzz(n int) []string {
    var result []string
    
    for i := 1; i <= n; i++ {
        switch {
            case i%15 == 0:
			    result = append(result, "FizzBuzz")
		    case i%3 == 0:
			    result = append(result, "Fizz")
		    case i%5 == 0:
			    result = append(result, "Buzz")
		    default:
			    myNum := strconv.Itoa(i)
			    result = append(result, myNum)
		}
    }
    return result
}
