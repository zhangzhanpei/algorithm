/**
 * 给定一个单链表, 把所有奇数位置节点提到前面, 偶数位置节点提到后面
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    odd := head
    even := head.Next
    for even != nil && even.Next != nil {
        tmp := odd.Next //如 1->2->3->4, 此时 tmp = 2
        odd.Next = even.Next //1->3
        even.Next = even.Next.Next //2->4
        odd.Next.Next = tmp //3->2 此时得到 1->3->2->4
        odd = odd.Next
        even = even.Next
    }
    return head
}
