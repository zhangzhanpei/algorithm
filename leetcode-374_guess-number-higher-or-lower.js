//猜数字, 每次猜测会告知是大是小还是相等
//使用二分查找
var guessNumber = function (n) {
    let minNum = 1, maxNum = n;
    while (true) {
        let midNum = parseInt((maxNum - minNum) / 2) + minNum;
        let res = guess(midNum);
        if (res === 0) {
            return midNum;
        } else if (res === 1) {
            minNum = midNum + 1;
        } else {
            maxNum = midNum - 1;
        }
    }
}
