//自底向上逐层返回二叉树的元素
var levelOrderBottom = function(root) {
    let list = [], ret = [];
    if (root === null) {
        return ret;
    }
    list.push(root); //list一开始先保存第一层的元素root节点
    while (list.length) {
        let len = list.length; //获取当前层的宽度
        let tmp = [];
        for (let i = 0; i < len; i++) { //遍历当前层时, 将下一层的元素压进list尾部
            let node = list.shift(); //当前层的元素逐个弹出, 当前层遍历完后剩下的元素就全是下一层的
            tmp.push(node.val); //tmp保存了一层的节点的值
            if (node.left !== null) {
                list.push(node.left);
            }
            if (node.right !== null) {
                list.push(node.right);
            }
        }
        ret.push(tmp);
    }
    return ret.reverse();
};
