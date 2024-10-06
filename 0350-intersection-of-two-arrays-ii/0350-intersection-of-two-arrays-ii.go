func intersect(nums1 []int, nums2 []int) []int {
    m := map[int]int{}
    
    intersection := make([]int, 0, len(nums2))
    
    for _, value := range nums1 {
        if _, ok := m[value]; ok {
            m[value]++
        } else {
            m[value] = 1
        }
    }
    
    for _, value := range nums2 {
        if v, ok := m[value]; ok && v > 0 {
            intersection = append(intersection, value)
            
            m[value]--
        }
    }
    
    return intersection
}