//单向链表
//每个节点存储有自身的值, 以及指向下一节点的指针

function ListNode(val, next = null)
{
    this.val = val;
    this.next = next;
}

var n5 = new ListNode(5);
var n4 = new ListNode(4, n5);
var n3 = new ListNode(3, n4);
var n2 = new ListNode(2, n3);
var head = new ListNode(1, n2);

while (head != null) {
    console.log(head.val);
    head = head.next;
}
