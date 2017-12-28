/**
 * 给定两个字符串, 找到一个最长的字串, 并且这个字串不能是另一个字符串的字串
 * 字符串本身就是一个字串
 */
func findLUSlength(a string, b string) int {
    //如果两个字符串相等, 即字串是另一个字符串的字串, 返回-1
    if a == b {
        return -1
    }
    if len(a) > len(b) { //返回两个字符串中长的那个, 因为长的字串不可能是另一个字符串的字串
        return len(a)
    } else {
        return len(b)
    }
}
