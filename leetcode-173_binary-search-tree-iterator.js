/**
 * 给定一个二叉搜索树, 实现 next 方法由小到大迭代返回二叉树节点
 */
var stack = []
var BSTIterator = function(root) {
    while (root) {
        stack.push(root)
        root = root.left
    }
};

BSTIterator.prototype.hasNext = function() {
    return stack.length > 0
};

BSTIterator.prototype.next = function() {
    let top = stack.pop()
    let ret = top.val
    if (top.right) { //栈顶元素就是最小节点
        top = top.right
        while (top) {
            stack.push(top)
            top = top.left
        }
    }
    return ret
};
