//一堆石头, 一次可取走1-3块, 谁取走最后一块则胜利.
//给定一个n块石头的堆, 自己先取, 判断是否能赢
//当石头数量是4的整数倍时, 怎么取都是输
var canWinNim = function(n) {
    if (n % 4 === 0) {
        return false;
    }
    return true;
};

canWinNim(4);
