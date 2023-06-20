/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    mergedList := &ListNode{Next: nil}
    result := mergedList
    
    for list1 != nil || list2 != nil {
        if list1 == nil {
            mergedList.Next = list2
            break
        } else if list2 == nil {
            mergedList.Next = list1
            break
        }
        if list1.Val <= list2.Val {
            mergedList.Next = list1
            list1 = list1.Next
        } else {
            mergedList.Next = list2
            list2 = list2.Next
        }
        mergedList = mergedList.Next
    }
    return result.Next
}