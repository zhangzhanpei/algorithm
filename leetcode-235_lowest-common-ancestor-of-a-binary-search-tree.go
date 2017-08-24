/**
 * 给定一个二叉查找树和两个节点, 求这两个节点的最近公共祖先
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if root.Val > p.Val && root.Val > q.Val { //如果两个节点都小于根节点, 那么公共祖先在左子树
            root = root.Left
        } else if root.Val < p.Val && root.Val < q.Val { //如果两个节点都大于根节点, 那么公共祖先在右子树
            root = root.Right
        } else { //如果其中一个小于根节点而另一个大于根节点, 那么根节点就是它们的祖先
            return root
        }
    }
}
