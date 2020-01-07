/**
 * 给定一个数字和#组成的字符串，转为相应的字母
 * 1 - 9 代表 a - i
 * 10# - 26# 代表 j - z
 */
public static String freqAlphabets(String s) {
    StringBuilder stringBuilder = new StringBuilder();
    for (int i = 0; i < s.length(); i++) {
        // 如果遇到 # 号，截取前两位字符转成数字，再转成相应的 ascii 字符
        if (i + 2 < s.length() && s.charAt(i + 2) == '#') {
            stringBuilder.append((char) (Integer.parseInt(s.substring(i, i + 2)) - 1 + 'a'));
            i += 2;
        } else {
            stringBuilder.append((char) (Integer.parseInt(s.substring(i, i + 1)) - 1 + 'a'));
        }
    }
    return stringBuilder.toString();
}
