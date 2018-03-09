/**
 * 给定两个字符串A和B，判断A字符串前N个字符移到尾部能否形成B
 */
public boolean rotateString(String A, String B) {
    //A和B的长度必定相等
    //两个A拼接在一起，能找到一个子串B
    return A.length() == B.length() && (A + A).indexOf(B) > -1;
}
