/**
 * 合并两个有序单链表
 */
public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
    if (l1 == null && l2 == null) return null; //如果两个链表都为null, 直接返回null
    ListNode ret = new ListNode(-1); //定义结果链表, 只有一个头节点-1
    ListNode tmp = ret;
    while (l1 != null && l2 != null) { //合并相同长度的部分直到其中一个单链表处理完
        if (l1.val <= l2.val) {
            tmp.next = l1;
            l1 = l1.next;
        } else {
            tmp.next = l2;
            l2 = l2.next;
        }
        tmp = tmp.next;
    }
    if (l1 == null) tmp.next = l2; //当其中一个链表没有节点了, 直接把另一个链表跟在目标链表后
    if (l2 == null) tmp.next = l1;
    return ret.next; //把目标链表的第二个节点当做链表的表头返回, 因为第一个节点是我们定义的-1
}
