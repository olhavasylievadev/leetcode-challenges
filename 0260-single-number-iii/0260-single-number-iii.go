func singleNumber(nums []int) []int {
    myMap := make(map[int]int, 0)
    var unique []int
    
    for _, j := range nums {
        myMap[j]++
    }
    
    for k, l := range myMap {
        if l == 1 {
            unique = append(unique, k)
        }
    }
    
    return unique
}