/**
 * 给定一个 int 单链表，对其进行排序
 */
public ListNode sortList(ListNode head) {
    if (head == null || head.next == null) {
        return head;
    }

    ListNode prev = null, slow = head, fast = head;
    // 找到单链表的中点
    while (fast != null && fast.next != null) {
        prev = slow;
        slow = slow.next;
        fast = fast.next.next;
    }

    // 这里的作用是将单链表在中点处截断(否则会无限递归下去栈溢出)
    prev.next = null;

    ListNode l1 = sortList(head);
    ListNode l2 = sortList(slow);

    return merge(l1, l2);
}

// 合并两个有序单链表
public ListNode merge(ListNode l1, ListNode l2) {
    ListNode l = new ListNode(-1);
    ListNode h = l;

    while (l1 != null && l2 != null) {
        if (l1.val <= l2.val) {
            h.next = l1;
            l1 = l1.next;
        } else {
            h.next = l2;
            l2 = l2.next;
        }
        h = h.next;
    }

    if (l1 != null) {
        h.next = l1;
    }

    if (l2 != null) {
        h.next = l2;
    }

    return l.next;
}
