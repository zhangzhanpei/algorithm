//翻转单链表

function reverseLinkList(l) {
    if (l === null) {
        return;
    }
    var res, first, tmp;
    res = new linkList(-1); //创建一个新节点插入到头部
    res.next = l;
    first = res.next;
    while (first.next !== null) {
        tmp = first.next;
        first.next = tmp.next;
        tmp.next = res.next; //这里tmp的下一个不能是first，否则tmp和first会来回交换
        res.next = tmp;
    }
    return res.next; //从第二个节点开始就是反转后的链表
}
