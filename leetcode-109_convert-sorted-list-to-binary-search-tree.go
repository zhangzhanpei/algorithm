/**
 * 给定一个有序链表, 构建一个平衡二叉查找树
 * 同二叉查找算法, 使用快慢指针找到中点
 */
package main

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    return toBST(head, nil)
}

func toBST(head *ListNode, tail *ListNode) *TreeNode {
    if head == tail {
        return nil
    }
    slow := head
    fast := head
    for fast != tail && fast.Next != nil { //快慢指针, slow 即为中点
        fast = fast.Next.Next
        slow = slow.Next
    }

    node := &TreeNode{Val:slow.Val}
    node.Left = toBST(head, slow)
    node.Right = toBST(slow.Next, tail)
    return node
}
