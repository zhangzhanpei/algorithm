/**
 * 给定一个二叉查找树, 求节点间的最小差值
 * 中序遍历二叉树, 后一个节点减去前一个节点求最小差值
 */
package main

import "fmt"
import "math"

func main() {
    
}

var prev *TreeNode
var min int
func getMinimumDifference(root *TreeNode) int {
    prev = nil
    min = math.MaxInt32
    return inorderTraversal(root)
}

func inorderTraversal(root *TreeNode) int {
    if root == nil {
        return min
    }
    inorderTraversal(root.Left)
    if prev != nil {
        min = int(math.Min(float64(min), float64(root.Val - prev.Val)))
    }
    prev = root
    inorderTraversal(root.Right)
    return min
}
