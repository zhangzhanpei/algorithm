// 删除有序链表中的重复元素, 使得每个元素只出现一次
// 给出1->1->2，return 1->2
// 给出1->1->2->3->3，return 1->2->3

function deleteDuplicates(head)
{
    if (head === null) {
        return null;
    }
    var pre = head;
    var tmp = head.next;
    while (tmp !== null) {
        if (tmp.val == pre.val) { //如果下一个节点的值与前一个节点的值相等
            pre.next = tmp.next;  //让前一个节点的指针指向tmp的下一个节点, 即跳过tmp节点
        } else {
            pre = tmp;
        }
        tmp = tmp.next;
    }
    return head;
}
