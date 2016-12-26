//给定一个单链表, 判断是否为一个循环链表
//快慢指针是指快指针一次移动两步，慢指针一次移动一步
//如果是一个环形单链表，那么快指针一定会追上慢指针

function hasCycle(head)
{
    if (head === null || head.next === null) {
        return false;
    }
    var slow = head;
    var fast = head;
    fast = head.next.next;
    while (true) {
        if (fast === null || fast.next === null) {
            return false;
        } else if (fast == slow || fast.next == slow) {
            return true;
        } else {
            fast = fast.next.next;
            slow = slow.next;
        }
    }
}
