/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    myLen := len(nums)
    if myLen == 0 {
        return nil
    }
    
    mid := myLen / 2
    
    return &TreeNode {
        Val: nums[mid],
        Left: sortedArrayToBST(nums[:mid]),
        Right: sortedArrayToBST(nums[mid+1:]),
    }
}