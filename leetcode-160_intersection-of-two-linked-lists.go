/**
 * 两个单链表如果相交, 求交点, 如不相交返回null
 * 两个链表从交点往后都是一致的, 所以两个链表同时走, 如果相交之前的长度一致, 则会在交点处相遇
 * 如果相交之前的长度不一致, 当短链表走到尾部时, 让它去走长链表
 * 长链表到尾部时, 去走短链表, 这样两者走的长度相等
 * 如果两者始终没有相遇, 则没有交点 
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    p1 := headA
    p2 := headB
    if p1 == nil || p2 == nil {
        return nil
    }

    for p1 != nil && p2 != nil && p1 != p2 {
        p1 = p1.Next
        p2 = p2.Next
        if p1 == p2 {
            return p1
        }

        if p1 = nil {
            p1 = headB
        }
        if p2 = nil {
            p2 = headA
        }
    }

    //此时 p1 == nil(没交点) 或 p1 == p2(相交)
    return p1
}
