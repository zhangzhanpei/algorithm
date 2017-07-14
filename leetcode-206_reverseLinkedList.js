//翻转单链表
var reverseList = function(head) {
    let last = null, cur = head;
    while (cur) {
        let next = cur.next;
        cur.next = last;
        last = cur;
        cur = next;
    }
    return last;
};
