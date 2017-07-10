//前序遍历一棵二叉树, 生成一个字符串
var tree2str = function(t) {
    if (t === null) {
        return '';
    }
    let str = '';
    let left = tree2str(t.left);
    let right = tree2str(t.right);
    if (left || right) {
        str += `(${left})`;
    }
    if (right) {
        str += `(${right})`;
    }
    return t.val + str;
};
