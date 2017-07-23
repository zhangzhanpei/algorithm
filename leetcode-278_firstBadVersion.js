//找到第一个 bad version, 给定的 isBadVersion() 方法可以判断是否 bad version
//二分查找
var firstBadVersion = function (n) {
    let min = 1, max = n;
    while (min < max) {
        let mid = parseInt((max - min) / 2) + min;
        if (!isBadVersion(mid)) {
            min = mid + 1;
        } else {
            max = mid;
        }
    }
    return min;
}
