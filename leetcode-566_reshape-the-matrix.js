//转换一个矩阵的行和列数
var matrixReshape = function(nums, r, c) {
    let i = nums.length; //原来的行数
    let j = nums[0].length; //原来的列数
    if (i * j !== r * c) { //如果元素个数不一致直接返回原矩阵
        return nums;
    }
    let ret = [];
    let index = 0;
    for (let m = 0; m < r; m++) { //逐行填充
        let tmpArr = [];
        for (let n = 0; n < c; n++) {
            tmpArr.push(nums[parseInt((m * c + n) / j)][(m * c + n) % j]); //原来矩阵的元素逐个赋值到新数组
        }
        ret.push(tmpArr);
    }
    return ret;
};

console.log(matrixReshape([[1,2,3,4]], 2, 2)); // output: [[1, 2], [3, 4]]
