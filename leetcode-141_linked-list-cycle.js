//给定一个单链表, 判断是否有环
//快慢指针是指快指针一次移动两步，慢指针一次移动一步
//如果是一个有环单链表，那么快指针一定会追上慢指针
var hasCycle = function(head) {
    if (head === null || head.next === null) {
        return false;
    }
    var slow = head; //慢指针
    var fast = head; //快指针
    while (fast.next && fast.next.next) {
        slow = slow.next;
        fast = fast.next.next;
        if (slow === fast) { //快慢指针相遇
            return true;
        }
    }
    return false;
};
