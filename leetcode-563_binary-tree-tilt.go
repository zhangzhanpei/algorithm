/**
 * 给出一个二叉树, 返回整个树的差值
 * 树节点的差值被定义为所有左子树节点值和所有右子树节点值之和的差的绝对值. 空节点有差值0
 * 整个树的差值被定义为所有节点的差值的总和
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(findTilt(&root))
}

var res = 0
func findTilt(root *TreeNode) int {
    tilt(root)
    return res
}

func tilt(node *TreeNode) int {
    if node == nil {
        return 0
    }
    left := tilt(node.Left)
    right := tilt(node.Right)
    res += int(math.Abs(float64(left - right)))
    return node.Val + left + right
}
