func findDuplicates(nums []int) []int {
    myMap := make(map[int]int, 0)
    var notUnique []int
    
    for _, j := range nums {
        myMap[j]++
    }
    
    for k, l := range myMap {
        if l != 1 {
            notUnique = append(notUnique, k)
        }
    }
    
    return notUnique
}