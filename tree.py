# 树节点定义
# class TreeNode:
#     def __init__(self, val):
#         this.val = val
#         this.left, this.right = None, None

# 中序遍历二叉树
tree=[]
def inorderTraversal(self, root):
    if root==None:
        return []
    self.inorderTraversal(root.left)
    self.tree.append(root.val)
    self.inorderTraversal(root.right)
    return self.tree

# 前序遍历
def preorderTraversal(self, root):
    if root==None:
        return []
    self.tree.append(root.val)
    self.preorderTraversal(root.left)
    self.preorderTraversal(root.right)
    return self.tree

# 后序遍历
def postorderTraversal(self, root):
    if root==None:
        return []
    self.postorderTraversal(root.left)
    self.postorderTraversal(root.right)
    self.tree.append(root.val)
    return self.tree

# 求最大深度
def maxDepth(self, root):
    if root==None:
        return 0
    leftMaxDepth=self.maxDepth(root.left)+1
    rightMaxDepth=self.maxDepth(root.right)+1
    return max(leftMaxDepth,rightMaxDepth)

# 求最小深度
def minDepth(self, root):
    if root==None:
        return 0
    if root.left==None:
        return self.minDepth(root.right)+1
    if root.right==None:
        return self.minDepth(root.left)+1
    return min(self.minDepth(root.left),self.minDepth(root.right))+1;
