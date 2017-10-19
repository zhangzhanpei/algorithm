/**
 * 给定一个单链表和起点m和终点n, 反转m到n部分链表
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil {
        return nil
    }
    ret := &ListNode{Val:0}
    ret.Next = head
    pre := ret
    for i := 0; i < m - 1; i++ {
        pre = pre.Next //pre为起点前一个节点
    }

    start := pre.Next //起点节点
    then := start.Next //起点下一个节点
    for i = m; i < n; i++ {
        start.Next = then.Next
        then.Next = pre.Next
        pre.Next = then
        then = start.Next
    }

    return ret.Next
}
