//给定一个二叉树, 求左叶子的和
var sumOfLeftLeaves = function(root) {
    if (root === null) return 0;
    let list = [];
    let res = 0;
    list.push(root);

    while (list.length) {
        let node = list.shift();
        //如果左儿子没有左右节点, 说明它就是左叶子
        if (node.left && node.left.left === null && node.left.right === null) res += node.left.val;
        if (node.left) list.push(node.left);
        if (node.right) list.push(node.right);
    }
    return res;
};
