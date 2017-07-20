//判断一棵二叉树是否对称
var isSymmetric = function(root) {
    return root === null || isSymmetricHelp(root.left, root.right);
}

var isSymmetricHelp = function (left, right) {
    if (left === null || right === null) {
        return left === right;
    }
    if (left.val !== right.val) {
        return false;
    }
    return isSymmetricHelp(left.left, right.right) && isSymmetricHelp(left.right, right.left);
}
