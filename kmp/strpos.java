// 字串位置查找 - 朴素算法暴力匹配
public int strpos(String haystack, String needle) {
    if (haystack.length() < needle.length()) {
        return -1;
    }
    for (int i = 0; i < haystack.length() - needle.length(); i++) {
        int j = 0;
        for (; j < needle.length(); j++) {
            if (haystack.charAt(i + j) != needle.charAt(j)) {
                break;
            }
        }
        // 如果字串已全部比较完毕
        if (j >= needle.length()) {
            return i;
        }
    }
    return -1;
}
