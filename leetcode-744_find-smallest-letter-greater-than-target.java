/**
 * 给定一个有序char数组和一个目标字符，找到比目标字符大的第一个字符
 * char数组是环形的
 */
public char nextGreatestLetter(char[] letters, char target) {
    if (letters[letters.length - 1] <= target) { //如果最大的字符还不如目标字符大，则返回第一个字符
        return letters[0];
    }
    for (char c : letters) { //逐个查找，可优化成二分查找
        if (c > target) {
            return c;
        }
    }
    return letters[0];
}
