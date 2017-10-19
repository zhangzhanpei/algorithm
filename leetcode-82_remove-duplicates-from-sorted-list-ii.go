/**
 * 给定一个单链表, 移除有重复的节点, 只保留仅出现一次的节点
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    ret := &ListNode{Val:0}
    ret.Next = head
    h := ret
    cur := head
    for cur != nil {
        for cur.Next != nil && cur.Val == cur.Next.Val {
            cur = cur.Next
        }
        if h.Next == cur { //单独节点时
            h = h.Next //h指向单独节点
        } else { //重复节点时
            h.Next = cur.Next //h的下一个节点指向重复节点的下一个节点
        }
        cur = cur.Next
    }
    return ret.Next
}
