//交换单词中的元音字符['a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U']
//两个指针从两边向中间靠拢, 交换元音字符
var reverseVowels = function(s) {
    let vowels = ['a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'];
    let chars = s.split("");
    let left = 0, right = chars.length - 1;
    while (left < right) {
        while (left < right && vowels.indexOf(chars[left]) === -1) left++;
        while (left < right && vowels.indexOf(chars[right]) === -1) right--;
        swap(chars, left, right);
        left++;
        right--;
    }
    return chars.join("");
};

var swap = function (arr, x, y) {
    let tmp = arr[x];
    arr[x] = arr[y];
    arr[y] = tmp;
};

console.log(reverseVowels('leetcode')); //output 'leotcede'
