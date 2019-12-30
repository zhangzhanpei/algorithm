/**
 * 给定一个二进制的单链表，返回十进制数值
 */
public int getDecimalValue(ListNode head) {
    int ret = 0;
    while (head != null) {
        ret = ret * 2 + head.val;
        head = head.next;
    }
    return ret;
}
