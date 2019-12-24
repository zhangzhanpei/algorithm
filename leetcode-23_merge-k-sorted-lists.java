/**
 * 给定一个有序单链表数组，合并所有单链表仍保持有序返回
 * 方案一：两两合并
 */
public ListNode mergeKLists(ListNode[] lists) {
    ListNode ret = null;
    for (ListNode head : lists) {
        ret = mergeTwoLists(ret, head);
    }
    return ret;
}

// 这里是 leetcode-21 合并两个有序链表
public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
    if (l1 == null && l2 == null) return null; // 如果两个链表都为 null, 直接返回 null
    ListNode ret = new ListNode(-1); // 定义结果链表, 只有一个头节点-1
    ListNode tmp = ret;
    while (l1 != null && l2 != null) { // 合并相同长度的部分直到其中一个单链表处理完
        if (l1.val <= l2.val) {
            tmp.next = l1;
            l1 = l1.next;
        } else {
            tmp.next = l2;
            l2 = l2.next;
        }
        tmp = tmp.next;
    }
    if (l1 == null) tmp.next = l2; // 当其中一个链表没有节点了, 直接把另一个链表跟在目标链表后
    if (l2 == null) tmp.next = l1;
    return ret.next; // 把目标链表的第二个节点当做链表的表头返回, 因为第一个节点是我们定义的-1
}

/**
 * 方案二：优先队列
 */
public ListNode mergeKLists(ListNode[] lists) {
    PriorityQueue<ListNode> queue = new PriorityQueue<ListNode>((o1, o2) -> o1.val - o2.val);
    // 逐个单链表元素计入优先队列
    for (ListNode head : lists) {
        if (head == null) {
            continue;
        }
        ListNode cur = head;
        ListNode nextNode = head.next;
        while (cur != null) {
            cur.next = null; // 这里的 next 要置为 null，否则 Memory Limit Exceeded
            queue.add(cur);
            cur = nextNode;
            if (nextNode == null) {
                break;
            }
            nextNode = nextNode.next;
        }
    }
    // 将优先队列中的链表结点取出，重新构建单链表
    ListNode tmp = new ListNode(-1);
    ListNode ret = tmp;
    while (!queue.isEmpty()) {
        tmp.next = queue.poll();
        tmp = tmp.next;
    }
    return ret.next;
}
