/**
 * 给定两个链表, 求它们的和
 * l1 = 7->2->4->3
 * l2 = 5->6->4
 * 即 7243 + 564 = 7807 = 7->8->0->7
 */
var addTwoNumbers = function(l1, l2) {
    let s1 = [], s2 = [];
    //先用两个栈存储链表的数字
    while (l1) {
        s1.push(l1.val);
        l1 = l1.next;
    }
    while (l2) {
        s2.push(l2.val);
        l2 = l2.next;
    }
    
    let sum = 0;
    newList = new ListNode(0);
    while (s1.length || s2.length) {
        if (s1.length) sum += s1.pop(); //两个栈元素出栈并相加
        if (s2.length) sum += s2.pop();
        newList.val = sum % 10 //余数设置为当前节点的值
        let head = new ListNode(parseInt(sum / 10)); //整数作为前一个节点的值, 如果还有下一趟的 sum 则会覆盖这里设置的 head 节点值
        head.next = newList;
        newList = head;
        sum = parseInt(sum / 10);
    }
    return newList.val == 0 ? newList.next : newList;
};
