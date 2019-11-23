/**
 * 反转单向链表
 */
func reverseList(head *ListNode) *ListNode {
    var last *ListNode
    current := head
    for current != nil {
        next := current.Next
        current.Next = last
        last = current
        current = next
    }
    return last
}
