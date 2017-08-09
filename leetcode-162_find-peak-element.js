//查找数组的峰值
//给定一数组, 相邻元素不等, 查找任一峰值
//如果后一元素比前面的大, 说明在上山, 直到后面元素比前一个小, 说明在下山, 那么前一个元素就是峰顶
function findPeakElement(nums)
{
    for (i = 1; i < nums.length; i++) {
        if (nums[i] < nums[i - 1]) {
            return i - 1;
        }
    }
    return nums.length - 1;
}

var arr = [1, 2, 3];
console.log(findPeakElement(arr));
