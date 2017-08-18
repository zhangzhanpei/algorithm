/**
 * 判断一个单向链表是否回文的
 */
func isPalindrome(head *ListNode) bool {
    //先使用快慢指针找到中点
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    //将后半部分链表反转
    linkedList := reverseLinkedList(slow)
    //从链表的头部和反转后的链表头部开始比较, 如果全部相等则原链表是回文的
    for head != nil && linkedList != nil {
        if head.Val != linkedList.Val {
            return false
        }
        head = head.Next
        linkedList = linkedList.Next
    }
    return true
}

//反转单链表
func reverseLinkedList(head *ListNode) *ListNode {
    var res, first, tmp *ListNode;
    res = &ListNode{Val: -1}; //创建一个新节点作为头部
    res.Next = head;
    first = res.Next;
    for first.Next != nil {
        tmp = first.Next;
        first.Next = tmp.Next;
        tmp.Next = res.Next; //这里tmp的下一个不能是first，否则tmp和first会来回交换
        res.Next = tmp;
    }
    return res.Next;
}
