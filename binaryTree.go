//二叉树遍历[非递归]
package main

import (
    "fmt"
    "test" //test.Tree 中定义二叉树的结构, 并初始化了一个二叉树, 根结点为Root
)

func main() {
    // preorderTraversal(&test.Root)
    // inorderTraversal(&test.Root)
    postorderTraversal(&test.Root)
}

//前序遍历
func preorderTraversal(root *test.Tree) {
    stack := []test.Tree{} //栈
    for root != nil || len(stack) > 0 {
        for root != nil {
            fmt.Println(root.Val) //输出根节点
            stack = append(stack, *root) //根节点进栈
            root = root.Left //左孩子成为根节点继续输出和进栈
        }
        if len(stack) > 0 {
            root = &stack[len(stack) - 1] //栈顶元素[最左下角的元素]出栈
            stack = stack[:len(stack) - 1]
            root = root.Right //看有没有右孩子, 如果有则成为根节点继续输出和进栈
        }
    }
}

//中序遍历
func inorderTraversal(root *test.Tree) {
    stack := []test.Tree{} //栈
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, *root) //根节点进栈
            root = root.Left //左孩子成为根节点继续进栈
        }
        if len(stack) > 0 {
            root = &stack[len(stack) - 1] //栈顶元素[最左下角的元素]出栈
            fmt.Println(root.Val) //输出栈顶元素[最左下角的元素]
            stack = stack[:len(stack) - 1]
            root = root.Right //看有没有右孩子, 如果有则成为根节点继续进栈
        }
    }
}

//后序遍历
func postorderTraversal(root *test.Tree) {
    stack := []test.Tree{} //栈
    var cur *test.Tree
    var pre *test.Tree
    stack = append(stack, *root) //根节点进栈
    for len(stack) > 0 {
        cur = &stack[len(stack) - 1] //取栈顶元素
        //如果没有左右孩子, 或左右孩子都被访问过, 则输出当前节点值
        if (cur.Left == nil && cur.Right == nil) || (pre != nil && (*pre == *cur.Left || *pre == *cur.Right)) {
            fmt.Println(cur.Val)
            stack = stack[:len(stack) - 1]
            pre = cur
        } else { //先把右孩子进栈, 再把左孩子进栈, 那么出栈顺序为左右中, 即后序遍历
            if cur.Right != nil {
                stack = append(stack, *cur.Right)
            }
            if cur.Left != nil {
                stack = append(stack, *cur.Left)
            }
        }
    }
}
