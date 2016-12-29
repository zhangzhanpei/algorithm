//删除单链表中的指定元素

function removeElements(head, val)
{
    while (head !== null && head.val == val) { //如果头元素就是要删除的元素, 直接让下一元素成为头部
        head = head.next;
    }
    if (head === null) {
        return null;
    }
    tmp = head.next;
    pre = head;
    while (tmp !== null) {
        if (tmp.val == val) {
            pre.next = tmp.next;
        } else {
            pre = tmp;
        }
        tmp = tmp.next;
    }
    return head;
}
