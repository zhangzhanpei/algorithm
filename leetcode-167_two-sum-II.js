//给定一个有序数组, 取两个元素使得和等于目标值
//从数组两头开始相加, 与目标值比较, 如果大就减小右边的数, 如果小则加大左边的数
var twoSum = function(numbers, target) {
    let i = 0, j = numbers.length - 1;
    while (i < j) {
        if (numbers[i] + numbers[j] > target) {
            j--;
        }
        if (numbers[i] + numbers[j] < target) {
            i++;
        }
        if (numbers[i] + numbers[j] == target) {
            return [i + 1, j + 1];
        }
    }
};

console.log(twoSum([2, 7, 11, 15], 9));
