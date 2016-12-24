//合并两个有序单链表
//如1->3->5, 2->4->6合并得到1->2->3->4->5->6

function mergeTwoLists(l1, l2)
{
    if (l1 === null && l2 === null) { //如果两个链表都为null, 直接返回null
        return null;
    }
    var targetLists = new ListNode(-1); //定义目标链表, 只有一个头节点-1
    var current = targetLists; //指针指向目标链表头节点
    while (l1 !== null && l2 !== null) {
        if (l1.val <= l2.val) { //比较两个链表的节点, 小的节点跟在目标链表后面
            current.next = l1;
            l1 = l1.next;
        } else {
            current.next = l2;
            l2 = l2.next;
        }
        current = current.next;
    }
    if (l1 === null) { //当其中一个链表没有节点了, 直接把另一个链表跟在目标链表后
        current.next = l2;
    }
    if (l2 === null) {
        current.next = l1;
    }
    return targetLists.next; //把目标链表的第二个节点当做链表的表头返回, 因为第一个节点是我们定义的-1
}
