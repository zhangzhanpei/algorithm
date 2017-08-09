//删除单链表中倒数第n个节点
//给出单链表1->2->3->4->5, 删除倒数第2个节点得到 1->2->3->5
//快慢指针, 一个指针先走n步, 然后两个指针一起走, 当快指针走到队尾时, 慢指针指向的就是倒数第n个节点

function removeNthNodeFromEndOfList(head, n)
{
    if (head === null || head.next === null) { //如果链表没有或只有一个节点, 都返回null
        return null;
    }
    var fast = head;
    var slow = head;
    var i = 0;
    while (i != n) { //快指针先走n步
        fast = fast.next
        if (fast === null) { //如果fast为null, 说明传入的n与链表节点数相等, 也就是删除第一个节点
            return head.next
        }
        i++;
    }
    while (fast.next) { //快慢指针一起走
        fast = fast.next
        slow = slow.next
    }
    slow.next = slow.next.next //删除节点
    return head;
}
