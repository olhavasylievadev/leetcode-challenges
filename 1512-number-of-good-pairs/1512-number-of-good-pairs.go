func numIdenticalPairs(nums []int) int {
    var myList = make([]map[int]int, 0)
    var myGoodPairs = make(map[int]int)

    for i := 0; i < len(nums)-1; i++ {
	    for j := i + 1; j < len(nums); j++ {
		    if nums[i] == nums[j] && i < j {
			    myGoodPairs[nums[i]] = nums[j]
			    myList = append(myList, myGoodPairs)
		    }
	    }
    }
    return len(myList)
}