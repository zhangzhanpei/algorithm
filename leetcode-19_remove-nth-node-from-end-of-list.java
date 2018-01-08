/**
 * 给定一个单链表和一个整数n，删除单链表倒数第n个节点
 * 给出单链表1->2->3->4->5, 删除倒数第2个节点得到 1->2->3->5
 * 快慢指针, 一个指针先走n步, 然后两个指针一起走, 当快指针走到队尾时, 慢指针指向的就是倒数第n个节点
 * 注：给出的n总是有效的
 */
public ListNode removeNthFromEnd(ListNode head, int n) {
    //如果没有节点或只有一个节点
    if (head == null || head.next == null) return null;
    ListNode fast = head;
    ListNode slow = head;
    int i = 0;
    //快指针先走n步
    while (i != n) {
        fast = fast.next;
        //如果fast为null, 说明传入的n与链表节点数相等, 即删除第一个节点(因为n总是有效的)
        if (fast == null) {
            return head.next;
        }
        i++;
    }
    //快慢指针一起走
    while (fast.next != null) {
        fast = fast.next;
        slow = slow.next;
    }
    //删除节点
    slow.next = slow.next.next;
    return head;
}
