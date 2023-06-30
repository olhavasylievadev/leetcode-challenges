import "sort"

func sortedSquares(nums []int) []int {
    var temp []int
    j := 0
    
    for i := range nums {
        j = nums[i] * nums[i]
        temp = append(temp, j)
    }
    
    sort.Ints(temp)
    
    return temp
}