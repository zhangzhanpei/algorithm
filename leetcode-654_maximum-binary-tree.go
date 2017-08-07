/**
 * 给定一个int数组生成一个二叉树
 * 根节点是数组中最大的元素
 * 左儿子是最大元素左侧的最大元素
 * 右儿子是最大元素右侧的最大元素
 * 依此类推
 */
package main

import "fmt"
import "math"

func main() {
    constructMaximumBinaryTree([]int{3,2,1,6,0,5})
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    root := buildTree(nums)
    return root
}

//递归构造二叉树
func buildTree(arr []int) *TreeNode {
    if len(arr) == 0 {
        return nil
    }
    idx, max := getMax(arr)
    node := TreeNode{Val: max}
    node.Left = buildTree(arr[:idx])
    node.Right = buildTree(arr[idx + 1:])
    return &node
}

//取数组最大元素值和下标
func getMax(arr []int) (idx, max int) {
    idx = -1
    max = math.MinInt32
    for key, val := range arr {
        if val > max {
            max = val
            idx = key
        }
    }
    return idx, max
}
