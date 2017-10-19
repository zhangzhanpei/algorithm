/**
 * 给定两个单链表形式的数字, 求两数的和
 * 单链表表示数字时低位在前, 如 2->4->3 表示 342, 5->6->4 表示 465
 * 则 2->4->3 + 5->6->4 = 7->0->8
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sumList := &ListNode{Val: 0} //结果链表
    sumListHead := sumList
    carry := 0 //进位
    for l1 != nil || l2 != nil { //如果其中一个链表加完了, 默认用0相加
        a, b := 0, 0
        if l1 != nil {
            a = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            b = l2.Val
            l2 = l2.Next
        }

        sum := a + b + carry
        carry = sum / 10
        sumList.Next = &ListNode{Val: sum % 10}
        sumList = sumList.Next
    }
    if carry > 0 {
        sumList.Next = &ListNode{Val: carry}
    }
    return sumListHead.Next
}
