/**
 * 给定一个单链表, 判断是否有环
 * 快慢指针是指快指针一次移动两步，慢指针一次移动一步
 * 如果是一个有环单链表，那么快指针一定会追上慢指针
 */
public boolean hasCycle(ListNode head) {
    if (head == null || head.next == null) {
        return false;
    }

    ListNode fast = head;
    ListNode slow = head;
    while (fast.next != null && fast.next.next != null) {
        fast = fast.next.next;
        slow = slow.next;
        if (fast == slow) { // 快指针追上慢指针
            return true;
        }
    }
    return false;
}
