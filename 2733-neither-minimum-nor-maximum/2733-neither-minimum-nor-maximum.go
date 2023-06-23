import "sort"

func findNonMinOrMax(nums []int) int {
    if len(nums) < 3 {
        return -1
    }
    
    sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
    return nums[1]
}