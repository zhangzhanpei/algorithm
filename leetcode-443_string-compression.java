/**
 * 给定一个char数组，将相同字符统计成数量进行压缩
 * 如 {'a','a','b','b','c','c','c'} -> {'a', '2', 'b', '2', 'c', '3'}
 */
public int compress(char[] chars) {
    int i = 0, j = 0;
    while (i < chars.length) {
        int count = 0;
        char curChar = chars[i]; //取一个字符
        while (i < chars.length && chars[i] == curChar) { //统计数量
            count++;
            i++;
        }
        chars[j++] = curChar; //保存字符
        if (count > 1) {
            for (char c : Integer.toString(count).toCharArray()) { //将数量转为字符存入数组
                chars[j++] = c;
            }
        }
    }
    return j; //返回压缩后的字符数组长度
}
