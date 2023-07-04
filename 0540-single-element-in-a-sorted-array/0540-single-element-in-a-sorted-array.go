func singleNonDuplicate(nums []int) int {
    myMap := make(map[int]int, 0)
    
    for _, j := range nums {
        myMap[j]++
    }
    
    for k, l := range myMap {
        if l == 1 {
            return k
        }
    }
    
    return -1
}