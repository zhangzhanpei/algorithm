//分糖果
//一个偶数长度的数组中, 不同的数字代表不同类型的糖果, 把糖果平均分给哥哥和妹妹, 尽可能使妹妹的糖果种类最多
//如果种类超过数组半数, 那么妹妹可分到数组半数的种类
//如果不超过半数, 对数组去重得到所有种类, 妹妹可分到所有种类

var unique = function (arr) {
    let retArr = [];
    arr.forEach((val) => {
        if (retArr.indexOf(val) === -1) {
            retArr.push(val);
        }
    });
    return retArr;
}

var distributeCandies = function(candies) {
    return Math.min(candies.length / 2, unique(candies).length);
};

distributeCandies([1,1,2,2,3,3]);
