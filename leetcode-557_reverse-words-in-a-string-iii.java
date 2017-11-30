/**
 * 给定一个字符串，反转每一个单词
 */
public String reverseWords(String s) {
    String[] strs = s.split(" "); //先切出每一个单词
    for (int i = 0; i < strs.length; i++) {
        strs[i] = new StringBuffer(strs[i]).reverse().toString(); //反转单词
    }
    return String.join(" ", strs); //再拼回字符串
}
