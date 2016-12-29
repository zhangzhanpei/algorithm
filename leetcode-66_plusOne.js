//以数组的形式给定一整数, 将该整数+1
//如[3, 2] + 1 = [3, 3], 或者[9] + 1 = [1, 0]

function plusOne(digits)
{
    var len = digits.length;
    digits[len - 1] += 1; //最后一个元素+1
    for (let i = len - 1; i > 0; i--) { //从后往前逐个元素检查, 如果有10则赋值为0, 前一个元素+1
        if (digits[i] != 10) {
            break; //如果某个元素不为10, 也就是前面的位不用+1, 直接跳出循环
        }
        digits[i] = 0;
        digits[i - 1] += 1;
    }
    if (digits[0] == 10) { //如果头元素为10, 则赋值为0, 然后数组前面插入元素1
        digits[0] = 0;
        digits.unshift(1);
    }
    return digits;
}
